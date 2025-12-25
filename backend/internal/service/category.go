package service

import (
	"context"

	"github.com/freshkeep/backend/internal/data/model"
	"github.com/freshkeep/backend/internal/data/repo"
)

// CategoryService is a category service interface.
type CategoryService interface {
	ListCategories(ctx context.Context) ([]*model.Category, error)
	CreateCategory(ctx context.Context, req *CreateCategoryRequest) (*model.Category, error)
	UpdateCategory(ctx context.Context, id uint64, req *UpdateCategoryRequest) (*model.Category, error)
	DeleteCategory(ctx context.Context, id uint64) error
}

type categoryService struct {
	categoryRepo repo.CategoryRepo
}

// NewCategoryService creates a new category service.
func NewCategoryService(categoryRepo repo.CategoryRepo) CategoryService {
	return &categoryService{
		categoryRepo: categoryRepo,
	}
}

type CreateCategoryRequest struct {
	Name  string
	Icon  string
	Color string
}

type UpdateCategoryRequest struct {
	Name  string
	Icon  string
	Color string
}

func (s *categoryService) ListCategories(ctx context.Context) ([]*model.Category, error) {
	return s.categoryRepo.List(ctx)
}

func (s *categoryService) CreateCategory(ctx context.Context, req *CreateCategoryRequest) (*model.Category, error) {
	category := &model.Category{
		Name:  req.Name,
		Icon:  req.Icon,
		Color: req.Color,
	}
	return s.categoryRepo.Create(ctx, category)
}

func (s *categoryService) UpdateCategory(ctx context.Context, id uint64, req *UpdateCategoryRequest) (*model.Category, error) {
	category, err := s.categoryRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	category.Name = req.Name
	category.Icon = req.Icon
	category.Color = req.Color

	err = s.categoryRepo.Update(ctx, category)
	if err != nil {
		return nil, err
	}

	return category, nil
}

func (s *categoryService) DeleteCategory(ctx context.Context, id uint64) error {
	return s.categoryRepo.Delete(ctx, id)
}
