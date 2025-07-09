package repository

import (
	"fmt"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

type UnitOfMeasurementRepository struct {
	db *sqlx.DB
}

func NewUnitOfMeasurementRepository(db *sqlx.DB) *UnitOfMeasurementRepository {
	return &UnitOfMeasurementRepository{db: db}
}

// Crear una marca
func (r *UnitOfMeasurementRepository) CreateUnitOfMeasurement(uom *domain.UnitsOfMeasurement) error {
	query := `
		INSERT INTO units_of_measurement (
			name, shortcut, description, status, created_by, updated_by
		) VALUES (
		 	:name, :shortcut, :description, :status, :created_by, :updated_by
		)
	`
	_, err := r.db.NamedExec(query, uom)
	if err != nil {
		fmt.Printf("Error creando la marca: %v\n", err)
		return err
	}
	return nil
}

// Obtener todas las marcas
func (r *UnitOfMeasurementRepository) GetAllUnitsOfMeasurement() ([]domain.UnitsOfMeasurement, error) {
	var uom []domain.UnitsOfMeasurement
	query := "SELECT id, name, shortcut, description, status, created_by, updated_by FROM units_of_measurement WHERE deleted_at IS NULL"
	err := r.db.Select(&uom, query)
	if err != nil {
		fmt.Printf("Error ejecutando la consulta: %v\n", err)
		return nil, err
	}
	return uom, nil
}

// Obtener una marca por Id
func (r *UnitOfMeasurementRepository) GetUnitOfMeasurementById(id int) (*domain.UnitsOfMeasurement, error) {
	query := `SELECT id, name, shortcut, description, status, created_by, updated_by FROM units_of_measurement WHERE id = $1 AND deleted_at IS NULL`
	row := r.db.QueryRow(query, id)

	var uom domain.UnitsOfMeasurement
	if err := row.Scan(&uom.Id, &uom.Name, &uom.Shortcut, &uom.Description, &uom.Status, &uom.Created_By, &uom.Updated_By); err != nil {
		return nil, err
	}
	return &uom, nil
}

// Actualizar una marca
func (r *UnitOfMeasurementRepository) UpdateUnitOfMeasurement(uom *domain.UnitsOfMeasurement) error {
	query := `
		UPDATE units_of_measurement SET 
			name = $1, shortcut = $2, description = $3, status = $4, created_by = $5, updated_by = $6, updated_at = NOW()
		WHERE id = $7 AND deleted_at IS NULL
	`
	_, err := r.db.Exec(query, uom.Name, uom.Shortcut, uom.Description, uom.Status, uom.Created_By, uom.Updated_By, uom.Id)
	if err != nil {
		fmt.Printf("Error actualizando la marca con Id %d: %v\n", uom.Id, err)
		return err
	}
	return nil
}

// Eliminar una marca por Id (eliminación lógica)
func (r *UnitOfMeasurementRepository) DeleteUnitOfMeasurementById(id int) error {
	query := `
		UPDATE units_of_measurement
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
func (r *UnitOfMeasurementRepository) DeleteUnitsOfMeasurementByIds(ids []int) error {
	query := `UPDATE units_of_measurement SET deleted_at = NOW() WHERE id = ANY($1)`
	_, err := r.db.Exec(query, pq.Array(ids))
	return err
}
