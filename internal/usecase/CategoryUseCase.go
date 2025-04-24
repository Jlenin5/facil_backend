package usecase

import (
	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/Jlenin5/facil_backend/internal/repository"
)

type CategoryUseCase struct {
	CategoryRepo *repository.CategoryRepository
}

func NewCategoryUseCase(CategoryRepo *repository.CategoryRepository) *CategoryUseCase {
	return &CategoryUseCase{CategoryRepo: CategoryRepo}
}

func (uc *CategoryUseCase) CreateCategory(category *domain.Categories) error {
	return uc.CategoryRepo.CreateCategory(category)
}

func (uc *CategoryUseCase) GetAllCategories() ([]domain.Categories, error) {
	return uc.CategoryRepo.GetAllCategories()
}

func (uc *CategoryUseCase) GetCategoryById(categoryId int) (*domain.Categories, error) {
	return uc.CategoryRepo.GetCategoryById(categoryId)
}

func (uc *CategoryUseCase) UpdateCategory(category *domain.Categories) error {
	return uc.CategoryRepo.UpdateCategory(category)
}

func (uc *CategoryUseCase) DeleteCategoryById(id int) error {
	return uc.CategoryRepo.DeleteCategoryById(id)
}

func (uc *CategoryUseCase) DeleteCategoriesByIds(ids []int) error {
	return uc.CategoryRepo.DeleteCategoriesByIds(ids)
}
