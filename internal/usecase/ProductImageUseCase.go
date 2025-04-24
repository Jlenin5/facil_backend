package usecase

import (
	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/Jlenin5/facil_backend/internal/repository"
)

type ProductImageUseCase struct {
	ProductImageRepo *repository.ProductImageRepository
}

func NewProductImageUseCase(ProductImageRepo *repository.ProductImageRepository) *ProductImageUseCase {
	return &ProductImageUseCase{ProductImageRepo: ProductImageRepo}
}

func (uc *ProductImageUseCase) GetAllProductImagess(productId int) ([]domain.ProductImages, error) {
	return uc.ProductImageRepo.GetAllProductImagess(productId)
}

func (uc *ProductImageUseCase) DeleteProductImageByIds(id int) error {
	return uc.ProductImageRepo.DeleteProductImageByIds(id)
}
