package repository

import (
	"fmt"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

type CategoryRepository struct {
	db *sqlx.DB
}

func NewCategoryRepository(db *sqlx.DB) *CategoryRepository {
	return &CategoryRepository{db: db}
}

// Crear una categoría
func (r *CategoryRepository) CreateCategory(category *domain.Categories) error {
	query := "INSERT INTO categories (name, description, status) VALUES (:name, :description, :status)"
	_, err := r.db.NamedExec(query, category)
	if err != nil {
		fmt.Printf("Error creando la categoría: %v\n", err)
		return err
	}
	return nil
}

// Obtener todas las categorías
func (r *CategoryRepository) GetAllCategories() ([]domain.Categories, error) {
	var categories []domain.Categories
	query := "SELECT id, name, description, status FROM categories WHERE deleted_at IS NULL"
	err := r.db.Select(&categories, query)
	if err != nil {
		fmt.Printf("Error ejecutando la consulta: %v\n", err)
		return nil, err
	}
	return categories, nil
}

// Obtener una categoría por Id
func (r *CategoryRepository) GetCategoryById(id int) (*domain.Categories, error) {
	query := `SELECT id, name, description, status FROM categories WHERE id = $1 AND deleted_at IS NULL`
	row := r.db.QueryRow(query, id)

	var category domain.Categories
	if err := row.Scan(&category.Id, &category.Name, &category.Description, &category.Status); err != nil {
		return nil, err
	}
	return &category, nil
}

// Actualizar una categoría
func (r *CategoryRepository) UpdateCategory(category *domain.Categories) error {
	query := `
		UPDATE categories SET 
			name = $1, description = $2, status = $3, updated_at = NOW()
		WHERE id = $4 AND deleted_at IS NULL
	`
	_, err := r.db.Exec(query, category.Name, category.Description, category.Status, category.Id)
	if err != nil {
		fmt.Printf("Error actualizando la categoría con Id %d: %v\n", category.Id, err)
		return err
	}
	return nil
}

// Eliminar una categoría por Id (eliminación lógica)
func (r *CategoryRepository) DeleteCategoryById(id int) error {
	query := `
		UPDATE categories
		SET deleted_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
	`
	_, err := r.db.Exec(query, id)
	if err != nil {
		fmt.Printf("Error eliminando la categoría con Id %d: %v\n", id, err)
		return err
	}
	return nil
}

// Eliminar múltiples categorías por Ids (eliminación lógica)
func (r *CategoryRepository) DeleteCategoriesByIds(ids []int) error {
	query := `UPDATE categories SET deleted_at = NOW() WHERE id = ANY($1)`
	_, err := r.db.Exec(query, pq.Array(ids))
	return err
}
