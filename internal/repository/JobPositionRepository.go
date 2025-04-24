package repository

import (
	"database/sql"
	"fmt"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

type JobPositionRepository struct {
	db *sqlx.DB
}

func NewJobPositionRepository(db *sqlx.DB) *JobPositionRepository {
	return &JobPositionRepository{db: db}
}

// Crear una posición laboral
func (r *JobPositionRepository) CreateJobPosition(jobPosition *domain.JobPositions) error {
	query := `
		INSERT INTO job_positions (
			work_area_id, name, description, status
		) VALUES (
			:work_area_id, :name, :description, :status
		)
	`
	_, err := r.db.NamedExec(query, jobPosition)
	if err != nil {
		fmt.Printf("Error creando la posición laboral: %v\n", err)
		return err
	}
	return nil
}

// Obtener todas las áreas de trabajo
func (r *JobPositionRepository) GetAllJobPositions() ([]domain.JobPositions, error) {
	var jobPositions []domain.JobPositions
	query := `
		SELECT
			jp.id, jp.work_area_id, jp.name, jp.description, jp.status,
			wa.id AS "work_area.id", wa.name AS "work_area.name"
		FROM job_positions jp
		LEFT JOIN work_areas wa ON jp.work_area_id=wa.id
		WHERE jp.deleted_at IS NULL
		ORDER BY jp.id DESC
	`
	err := r.db.Select(&jobPositions, query)
	if err != nil {
		fmt.Printf("Error ejecutando la consulta: %v\n", err)
		return nil, err
	}
	return jobPositions, nil
}

// Obtener una posición laboral por Id
func (r *JobPositionRepository) GetJobPositionById(jobPositionId int) (*domain.JobPositions, error) {
	var jobPosition domain.JobPositions
	query := `
		SELECT
			jp.id, jp.work_area_id, jp.name, jp.description, jp.status,
			wa.id AS "work_area.id", wa.name AS "work_area.name"
		FROM job_positions jp
		LEFT JOIN work_areas wa ON jp.work_area_id=wa.id
		WHERE jp.id = $1 AND jp.deleted_at IS NULL
	`

	// Usar Get para mapear automáticamente los resultados
	err := r.db.Get(&jobPosition, query, jobPositionId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Área de trabajo no encontrado
		}
		fmt.Printf("Error ejecutando la consulta: %v\n", err)
		return nil, err
	}

	return &jobPosition, nil
}

// Actualizar una posición laboral
func (r *JobPositionRepository) UpdateJobPosition(jobPosition *domain.JobPositions) error {
	query := `
		UPDATE job_positions SET
			work_area_id = :work_area_id, name = :name, description = :description, status = :status, updated_at = NOW()
		WHERE id = :id AND deleted_at IS NULL
	`
	_, err := r.db.NamedExec(query, jobPosition)
	if err != nil {
		return fmt.Errorf("error updating user: %w", err)
	}
	return nil
}

// Eliminar una posición laboral por Id (eliminación lógica)
func (r *JobPositionRepository) DeleteJobPositionById(id int) error {
	query := `
		UPDATE job_positions
		SET deleted_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
	`
	_, err := r.db.Exec(query, id)
	if err != nil {
		fmt.Printf("Error eliminando la posición laboral con Id %d: %v\n", id, err)
		return err
	}
	return nil
}

// Eliminar múltiples áreas de trabajo por Ids (eliminación lógica)
func (r *JobPositionRepository) DeleteJobPositionsByIds(ids []int) error {
	query := `UPDATE job_positions SET deleted_at = NOW() WHERE id = ANY($1)`
	_, err := r.db.Exec(query, pq.Array(ids))
	return err
}
