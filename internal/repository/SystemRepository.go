package repository

import (
	"database/sql"
	"fmt"

	"github.com/Jlenin5/facil_backend/internal/domain"

	"github.com/jmoiron/sqlx"
)

type SystemRepository struct {
	db *sqlx.DB
}

func NewSystemRepository(db *sqlx.DB) *SystemRepository {
	return &SystemRepository{db: db}
}

// Crear un configuraciones del sistema
func (r *SystemRepository) CreateSystem(system *domain.Systems) error {
	query := `
		INSERT INTO systems (
			name, description
		) VALUES (
			:name, :description
		)
	`
	_, err := r.db.NamedExec(query, system)
	if err != nil {
		fmt.Printf("Error creando la configuraciones del sistema: %v\n", err)
		return err
	}
	return nil
}

// Obtener todas las sistemas
func (r *SystemRepository) GetAllSystems() ([]domain.Systems, error) {
	var systems []domain.Systems
	query := "SELECT id, name, description FROM systems ORDER BY id DESC"
	err := r.db.Select(&systems, query)
	if err != nil {
		fmt.Printf("Error ejecutando la consulta: %v\n", err)
		return nil, err
	}
	return systems, nil
}

// Obtener un configuraciones del sistema por Id
func (r *SystemRepository) GetSystemById(systemId int) (*domain.Systems, error) {
	var system domain.Systems
	query := "SELECT id, name, description FROM systems WHERE id = $1"

	// Usar Get para mapear automáticamente los resultados
	err := r.db.Get(&system, query, systemId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Área de trabajo no encontrado
		}
		fmt.Printf("Error ejecutando la consulta: %v\n", err)
		return nil, err
	}

	return &system, nil
}

// Actualizar un configuraciones del sistema
func (r *SystemRepository) UpdateSystem(system *domain.Systems) error {
	query := `
		UPDATE systems SET
			name = :name, description = :description, updated_at = NOW()
		WHERE id = :id
	`
	_, err := r.db.NamedExec(query, system)
	if err != nil {
		return fmt.Errorf("error updating user: %w", err)
	}
	return nil
}