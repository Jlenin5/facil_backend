package usecase

import (
	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/Jlenin5/facil_backend/internal/repository"
)

type BrandUseCase struct {
	BrandRepo *repository.BrandRepository
}

func NewBrandUseCase(BrandRepo *repository.BrandRepository) *BrandUseCase {
	return &BrandUseCase{BrandRepo: BrandRepo}
}

func (uc *BrandUseCase) CreateBrand(brand *domain.Brands) error {
	return uc.BrandRepo.CreateBrand(brand)
}

func (uc *BrandUseCase) GetAllBrands() ([]domain.Brands, error) {
	return uc.BrandRepo.GetAllBrands()
}

func (uc *BrandUseCase) GetBrandById(brandId int) (*domain.Brands, error) {
	return uc.BrandRepo.GetBrandById(brandId)
}

func (uc *BrandUseCase) UpdateBrand(Brand *domain.Brands) error {
	return uc.BrandRepo.UpdateBrand(Brand)
}

func (uc *BrandUseCase) DeleteBrandById(id int) error {
	return uc.BrandRepo.DeleteBrandById(id)
}

func (uc *BrandUseCase) DeleteBrandsByIds(ids []int) error {
	return uc.BrandRepo.DeleteBrandsByIds(ids)
}
