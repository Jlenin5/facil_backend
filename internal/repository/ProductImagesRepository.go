package repository

import (
	"fmt"

	"github.com/Jlenin5/facil_backend/internal/domain"

	"github.com/jmoiron/sqlx"
)

type ProductImageRepository struct {
	db *sqlx.DB
}

func NewProductImageRepository(db *sqlx.DB) *ProductImageRepository {
	return &ProductImageRepository{db: db}
}

// Obtener todas las imágenes de productos
func (r *ProductImageRepository) GetAllProductImagess(productId int) ([]domain.ProductImages, error) {
	var images []domain.ProductImages

	query := `
		SELECT id, product_id, url, featured
		FROM product_images
		WHERE product_id = $1
	`

	err := r.db.Select(&images, query, productId)
	if err != nil {
		fmt.Printf("Error al obtener imágenes del producto: %v\n", err)
		return nil, err
	}

	return images, nil
}

func (r *ProductImageRepository) DeleteProductImageByIds(imageId int) error {
	query := `
		DELETE FROM product_images
		WHERE id = $1
	`

	_, err := r.db.Exec(query, imageId)
	if err != nil {
		fmt.Printf("Error al eliminar la imagen del producto: %v\n", err)
		return err
	}

	return nil
}
