package service

import (
	"context"
	"time"

	"github.com/freshkeep/backend/internal/data/model"
	"github.com/freshkeep/backend/internal/data/repo"
)

// ItemService is a item service interface.
type ItemService interface {
	CreateItem(ctx context.Context, req *CreateItemRequest) (*model.Item, error)
	GetItem(ctx context.Context, id uint64) (*model.Item, error)
	ListItems(ctx context.Context, req *ListItemsRequest) ([]*model.Item, int64, error)
	UpdateItem(ctx context.Context, id uint64, req *UpdateItemRequest) (*model.Item, error)
	DeleteItem(ctx context.Context, id uint64) error
}

type itemService struct {
	itemRepo     repo.ItemRepo
	categoryRepo repo.CategoryRepo
	reminderRepo repo.ReminderRepo
}

// NewItemService creates a new item service.
func NewItemService(itemRepo repo.ItemRepo, categoryRepo repo.CategoryRepo, reminderRepo repo.ReminderRepo) ItemService {
	return &itemService{
		itemRepo:     itemRepo,
		categoryRepo: categoryRepo,
		reminderRepo: reminderRepo,
	}
}

// NewItemServiceWithRepos creates a new item service with all repos
func NewItemServiceWithRepos(itemRepo repo.ItemRepo, categoryRepo repo.CategoryRepo, reminderRepo repo.ReminderRepo) ItemService {
	return NewItemService(itemRepo, categoryRepo, reminderRepo)
}

type CreateItemRequest struct {
	Name              string
	CategoryID        uint64
	ExpiryDate        time.Time
	PurchaseDate      *time.Time
	Quantity          int
	ImagePath         string
	Barcode           string
	Notes             string
	ReminderDaysBefore int
}

type ListItemsRequest struct {
	Page       int
	PageSize   int
	CategoryID *uint64
	Status     string
	Search     string
}

type UpdateItemRequest struct {
	Name              string
	CategoryID        uint64
	ExpiryDate        time.Time
	PurchaseDate      *time.Time
	Quantity          int
	ImagePath         string
	Barcode           string
	Notes             string
	ReminderDaysBefore int
}

func (s *itemService) CreateItem(ctx context.Context, req *CreateItemRequest) (*model.Item, error) {
	item := &model.Item{
		Name:         req.Name,
		CategoryID:   req.CategoryID,
		ExpiryDate:   req.ExpiryDate,
		PurchaseDate: req.PurchaseDate,
		Quantity:     req.Quantity,
		ImagePath:    req.ImagePath,
		Barcode:      req.Barcode,
		Notes:        req.Notes,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	item.CalculateStatus()

	createdItem, err := s.itemRepo.Create(ctx, item)
	if err != nil {
		return nil, err
	}

	// Create reminder if specified
	if req.ReminderDaysBefore > 0 {
		reminder := &model.Reminder{
			ItemID:      createdItem.ID,
			DaysBefore:  req.ReminderDaysBefore,
			Enabled:     true,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		}
		_, err = s.reminderRepo.Create(ctx, reminder)
		if err != nil {
			// Log error but don't fail the item creation
		}
	}

	return createdItem, nil
}

func (s *itemService) GetItem(ctx context.Context, id uint64) (*model.Item, error) {
	return s.itemRepo.Get(ctx, id)
}

func (s *itemService) ListItems(ctx context.Context, req *ListItemsRequest) ([]*model.Item, int64, error) {
	filters := make(map[string]interface{})
	if req.CategoryID != nil {
		filters["category_id"] = *req.CategoryID
	}
	if req.Status != "" {
		filters["status"] = req.Status
	}
	if req.Search != "" {
		filters["search"] = req.Search
	}

	page := req.Page
	if page < 1 {
		page = 1
	}
	pageSize := req.PageSize
	if pageSize < 1 {
		pageSize = 20
	}

	return s.itemRepo.List(ctx, page, pageSize, filters)
}

func (s *itemService) UpdateItem(ctx context.Context, id uint64, req *UpdateItemRequest) (*model.Item, error) {
	item, err := s.itemRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	item.Name = req.Name
	item.CategoryID = req.CategoryID
	item.ExpiryDate = req.ExpiryDate
	item.PurchaseDate = req.PurchaseDate
	item.Quantity = req.Quantity
	item.ImagePath = req.ImagePath
	item.Barcode = req.Barcode
	item.Notes = req.Notes
	item.UpdatedAt = time.Now()
	item.CalculateStatus()

	err = s.itemRepo.Update(ctx, item)
	if err != nil {
		return nil, err
	}

	// Update reminder
	if req.ReminderDaysBefore > 0 {
		reminder, err := s.reminderRepo.GetByItemID(ctx, id)
		if err == nil {
			reminder.DaysBefore = req.ReminderDaysBefore
			reminder.UpdatedAt = time.Now()
			s.reminderRepo.Update(ctx, reminder)
		} else {
			// Create new reminder
			reminder := &model.Reminder{
				ItemID:     id,
				DaysBefore: req.ReminderDaysBefore,
				Enabled:    true,
				CreatedAt:  time.Now(),
				UpdatedAt:  time.Now(),
			}
			s.reminderRepo.Create(ctx, reminder)
		}
	}

	return item, nil
}

func (s *itemService) DeleteItem(ctx context.Context, id uint64) error {
	return s.itemRepo.Delete(ctx, id)
}
