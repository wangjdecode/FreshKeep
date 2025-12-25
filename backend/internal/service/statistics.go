package service

import (
	"context"
	"time"

	"github.com/freshkeep/backend/internal/data/repo"
)

// StatisticsService is a statistics service interface.
type StatisticsService interface {
	GetOverview(ctx context.Context, days int) (*OverviewStats, error)
	GetExpiringItems(ctx context.Context, days int) ([]*ExpiringItem, error)
	GetTrend(ctx context.Context, days int) ([]*TrendData, error)
}

type statisticsService struct {
	itemRepo repo.ItemRepo
}

// NewStatisticsService creates a new statistics service.
func NewStatisticsService(itemRepo repo.ItemRepo) StatisticsService {
	return &statisticsService{
		itemRepo: itemRepo,
	}
}

type OverviewStats struct {
	FreshCount       int32
	ExpiringSoonCount int32
	ExpiredCount     int32
	NewItemsCount    int32
}

type ExpiringItem struct {
	ID            uint64
	Name          string
	CategoryID    uint64
	CategoryName  string
	ExpiryDate    time.Time
	DaysRemaining int32
	Status        string
}

type TrendData struct {
	Date          string
	FreshCount    int32
	ExpiringCount int32
	ExpiredCount  int32
}

func (s *statisticsService) GetOverview(ctx context.Context, days int) (*OverviewStats, error) {
	now := time.Now()
	var startDate time.Time
	if days > 0 {
		startDate = now.AddDate(0, 0, -days)
	}

	filters := make(map[string]interface{})
	allItems, _, err := s.itemRepo.List(ctx, 1, 10000, filters)
	if err != nil {
		return nil, err
	}

	stats := &OverviewStats{}
	for _, item := range allItems {
		if days > 0 && item.CreatedAt.Before(startDate) {
			continue
		}

		switch item.Status {
		case "fresh":
			stats.FreshCount++
		case "expiring_soon":
			stats.ExpiringSoonCount++
		case "expired":
			stats.ExpiredCount++
		}

		if days > 0 && item.CreatedAt.After(startDate) {
			stats.NewItemsCount++
		}
	}

	return stats, nil
}

func (s *statisticsService) GetExpiringItems(ctx context.Context, days int) ([]*ExpiringItem, error) {
	threshold := time.Now().AddDate(0, 0, days)
	filters := make(map[string]interface{})
	filters["status"] = "expiring_soon"

	items, _, err := s.itemRepo.List(ctx, 1, 1000, filters)
	if err != nil {
		return nil, err
	}

	var expiringItems []*ExpiringItem
	for _, item := range items {
		if item.ExpiryDate.After(threshold) {
			continue
		}
		expiringItems = append(expiringItems, &ExpiringItem{
			ID:            item.ID,
			Name:          item.Name,
			CategoryID:    item.CategoryID,
			ExpiryDate:    item.ExpiryDate,
			DaysRemaining: int32(item.DaysRemaining()),
			Status:        item.Status,
		})
	}

	return expiringItems, nil
}

func (s *statisticsService) GetTrend(ctx context.Context, days int) ([]*TrendData, error) {
	// This is a simplified implementation
	// In production, you would aggregate data by date
	now := time.Now()
	var startDate time.Time
	if days > 0 {
		startDate = now.AddDate(0, 0, -days)
	}

	filters := make(map[string]interface{})
	items, _, err := s.itemRepo.List(ctx, 1, 10000, filters)
	if err != nil {
		return nil, err
	}

	// Group by date (simplified - in production use proper date grouping)
	dateMap := make(map[string]*TrendData)
	for _, item := range items {
		if days > 0 && item.CreatedAt.Before(startDate) {
			continue
		}

		dateKey := item.CreatedAt.Format("2006-01-02")
		if _, exists := dateMap[dateKey]; !exists {
			dateMap[dateKey] = &TrendData{
				Date: dateKey,
			}
		}

		data := dateMap[dateKey]
		switch item.Status {
		case "fresh":
			data.FreshCount++
		case "expiring_soon":
			data.ExpiringCount++
		case "expired":
			data.ExpiredCount++
		}
	}

	var trendData []*TrendData
	for _, data := range dateMap {
		trendData = append(trendData, data)
	}

	return trendData, nil
}
