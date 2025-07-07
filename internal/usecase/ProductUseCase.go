package usecase

import (
	"fmt"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/Jlenin5/facil_backend/internal/repository"
)

type ProductUseCase struct {
	ProductRepo *repository.ProductRepository
}

func NewProductUseCase(productRepo *repository.ProductRepository) *ProductUseCase {
	return &ProductUseCase{ProductRepo: productRepo}
}

func (uc *ProductUseCase) CreateProduct(product *domain.Products) error {
	return uc.ProductRepo.CreateProduct(product)
}

func (uc *ProductUseCase) SaveProducts(products []domain.Products) error {
	if len(products) == 0 {
		return fmt.Errorf("no products to save")
	}

	// Validar que todos los productos tengan los campos requeridos
	// for _, product := range products {
	// 	if product.Name == "" {
	// 		return fmt.Errorf("invalid product data: missing required fields")
	// 	}
	// }

	// Llamar al repositorio para guardar los productos
	// if err := uc.ProductRepo.SaveProducts(products); err != nil {
	// 	return fmt.Errorf("failed to save products: %w", err)
	// }

	return uc.ProductRepo.SaveProducts(products)
}

func (uc *ProductUseCase) GetAllProducts() ([]domain.Products, error) {
	return uc.ProductRepo.GetAllProducts()
}

func (uc *ProductUseCase) GetProductById(productID int) (*domain.Products, error) {
	return uc.ProductRepo.GetProductById(productID)
}

// Obtener productos por IDs
func (uc *ProductUseCase) GetProductsByIds(ids []int) ([]domain.Products, error) {
	return uc.ProductRepo.GetProductsByIds(ids)
}

func (uc *ProductUseCase) UpdateProduct(Product *domain.Products) error {
	return uc.ProductRepo.UpdateProduct(Product)
}

func (uc *ProductUseCase) DeleteProductById(id int) error {
	return uc.ProductRepo.DeleteProductById(id)
}

func (uc *ProductUseCase) DeleteProductsByIds(ids []int) error {
	return uc.ProductRepo.DeleteProductsByIds(ids)
}

// func (uc *ProductUseCase) GetAllProductImages(productId int) ([]domain.ProductImages, error) {
// 	return uc.ProductRepo.GetAllProductImages(productId)
// }

func (uc *ProductUseCase) DeleteProductImageById(id int) error {
	return uc.ProductRepo.DeleteProductImageById(id)
}