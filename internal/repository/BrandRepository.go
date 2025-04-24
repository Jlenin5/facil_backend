package repository

import (
	"fmt"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

type BrandRepository struct {
	db *sqlx.DB
}

func NewBrandRepository(db *sqlx.DB) *BrandRepository {
	return &BrandRepository{db: db}
}

// Crear una marca
func (r *BrandRepository) CreateBrand(brand *domain.Brands) error {
	query := `
		INSERT INTO brands (
			name, description, logo_url, website_url, status
		) VALUES (
		 	:name, :description, :logo_url, :website_url, :status
		)
	`
	_, err := r.db.NamedExec(query, brand)
	if err != nil {
		fmt.Printf("Error creando la marca: %v\n", err)
		return err
	}
	return nil
}

// Obtener todas las marcas
func (r *BrandRepository) GetAllBrands() ([]domain.Brands, error) {
	var brands []domain.Brands
	query := "SELECT id, name, description, logo_url, website_url, status FROM brands WHERE deleted_at IS NULL"
	err := r.db.Select(&brands, query)
	if err != nil {
		fmt.Printf("Error ejecutando la consulta: %v\n", err)
		return nil, err
	}
	return brands, nil
}

// Obtener una marca por Id
func (r *BrandRepository) GetBrandById(id int) (*domain.Brands, error) {
	query := `SELECT id, name, description, logo_url, website_url, status FROM brands WHERE id = $1 AND deleted_at IS NULL`
	row := r.db.QueryRow(query, id)

	var brand domain.Brands
	if err := row.Scan(&brand.Id, &brand.Name, &brand.Description, &brand.Logo_url, &brand.Website_url, &brand.Status); err != nil {
		return nil, err
	}
	return &brand, nil
}

// Actualizar una marca
func (r *BrandRepository) UpdateBrand(brand *domain.Brands) error {
	query := `
		UPDATE brands SET 
			name = $1, description = $2, logo_url = $3, website_url = $4, status = $5, updated_at = NOW()
		WHERE id = $6 AND deleted_at IS NULL
	`
	_, err := r.db.Exec(query, brand.Name, brand.Description, brand.Logo_url, brand.Website_url, brand.Status, brand.Id)
	if err != nil {
		fmt.Printf("Error actualizando la marca con Id %d: %v\n", brand.Id, err)
		return err
	}
	return nil
}

// Eliminar una marca por Id (eliminación lógica)
func (r *BrandRepository) DeleteBrandById(id int) error {
	query := `
		UPDATE brands
		SET deleted_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
	`
	_, err := r.db.Exec(query, id)
	if err != nil {
		fmt.Printf("Error eliminando la marca con Id %d: %v\n", id, err)
		return err
	}
	return nil
}

// Eliminar múltiples marcas por Ids (eliminación lógica)
func (r *BrandRepository) DeleteBrandsByIds(ids []int) error {
	query := `UPDATE brands SET deleted_at = NOW() WHERE id = ANY($1)`
	_, err := r.db.Exec(query, pq.Array(ids))
	return err
}
