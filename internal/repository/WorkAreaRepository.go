package repository

import (
	"database/sql"
	"fmt"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

type WorkAreaRepository struct {
	db *sqlx.DB
}

func NewWorkAreaRepository(db *sqlx.DB) *WorkAreaRepository {
	return &WorkAreaRepository{db: db}
}

// Crear un área de trabajo
func (r *WorkAreaRepository) CreateWorkArea(workArea *domain.WorkAreas) error {
	query := "INSERT INTO work_areas (name, description, status) VALUES (:name, :description, :status)"
	_, err := r.db.NamedExec(query, workArea)
	if err != nil {
		fmt.Printf("Error creando la área de trabajo: %v\n", err)
		return err
	}
	return nil
}

// Obtener todas las áreas de trabajo
func (r *WorkAreaRepository) GetAllWorkAreas() ([]domain.WorkAreas, error) {
	var workAreas []domain.WorkAreas
	query := "SELECT id, name, description, status FROM work_areas WHERE deleted_at IS NULL ORDER BY id DESC"
	err := r.db.Select(&workAreas, query)
	if err != nil {
		fmt.Printf("Error ejecutando la consulta: %v\n", err)
		return nil, err
	}
	return workAreas, nil
}

// Obtener un área de trabajo por Id
func (r *WorkAreaRepository) GetWorkAreaById(workAreaId int) (*domain.WorkAreas, error) {
	var workArea domain.WorkAreas
	query := "SELECT id, name, description, status FROM work_areas WHERE id = $1 AND deleted_at IS NULL"

	// Usar Get para mapear automáticamente los resultados
	err := r.db.Get(&workArea, query, workAreaId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Área de trabajo no encontrado
		}
		fmt.Printf("Error ejecutando la consulta: %v\n", err)
		return nil, err
	}

	return &workArea, nil
}

// Actualizar un área de trabajo
func (r *WorkAreaRepository) UpdateWorkArea(workArea *domain.WorkAreas) error {
	query := `
		UPDATE work_areas SET
			name = :name, description = :description, status = :status, updated_at = NOW()
		WHERE id = :id AND deleted_at IS NULL
	`
	_, err := r.db.NamedExec(query, workArea)
	if err != nil {
		return fmt.Errorf("error updating user: %w", err)
	}
	return nil
}

// Eliminar un área de trabajo por Id (eliminación lógica)
func (r *WorkAreaRepository) DeleteWorkAreaById(id int) error {
	query := `
		UPDATE work_areas
		SET deleted_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
	`
	_, err := r.db.Exec(query, id)
	if err != nil {
		fmt.Printf("Error eliminando la área de trabajo con Id %d: %v\n", id, err)
		return err
	}
	return nil
}

// Eliminar múltiples áreas de trabajo por Ids (eliminación lógica)
func (r *WorkAreaRepository) DeleteWorkAreasByIds(ids []int) error {
	query := `UPDATE work_areas SET deleted_at = NOW() WHERE id = ANY($1)`
	_, err := r.db.Exec(query, pq.Array(ids))
	return err
}
