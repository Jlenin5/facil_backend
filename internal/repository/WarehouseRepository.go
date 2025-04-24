package repository

import (
	"database/sql"
	"fmt"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

type WarehouseRepository struct {
	db *sqlx.DB
}

func NewWarehouseRepository(db *sqlx.DB) *WarehouseRepository {
	return &WarehouseRepository{db: db}
}

// Crear un almacén
func (r *WarehouseRepository) CreateWarehouse(warehouse *domain.Warehouses) error {
	query := "INSERT INTO warehouses (branch_office_id, name, description, address, status) VALUES (:branch_office_id, :name, :description, :address, :status)"
	_, err := r.db.NamedExec(query, warehouse)
	if err != nil {
		fmt.Printf("Error creando el almacén: %v\n", err)
		return err
	}
	return nil
}

// Obtener todos los almacenes
func (r *WarehouseRepository) GetAllWarehouses() ([]domain.Warehouses, error) {
	var warehouses []domain.Warehouses
	query := "SELECT id, branch_office_id, name, description, address, status FROM warehouses WHERE deleted_at IS NULL"
	err := r.db.Select(&warehouses, query)
	if err != nil {
		fmt.Printf("Error ejecutando la consulta: %v\n", err)
		return nil, err
	}

	// Iterar sobre los almacenes
	for i := range warehouses {
		// Obtener sucursal del almacén
		var branchOffice domain.BranchOffices
		err = r.db.Get(&branchOffice, `
			SELECT id, company_id, name, description, address, phone, status
			FROM branch_offices
			WHERE id = $1 AND deleted_at IS NULL
		`, warehouses[i].Branch_Office_Id)
		if err != nil && err != sql.ErrNoRows { // Manejar casos donde no hay sucursal asociada
			fmt.Printf("Error obteniendo sucursal para el almacén %d: %v\n", warehouses[i].Id, err)
			return nil, err
		}

		// Obtener empresa de la sucursal (si existe)
		if branchOffice.Company_Id > 0 {
			var company domain.Companies
			err = r.db.Get(&company, `
				SELECT id, name, ruc, email, phone, web_site, address, status
				FROM companies 
				WHERE id = $1 AND deleted_at IS NULL
			`, branchOffice.Company_Id)

			if err != nil && err != sql.ErrNoRows {
				fmt.Printf("Error obteniendo empresa %d para la sucursal %d: %v\n", branchOffice.Company_Id, branchOffice.Id, err)
				return nil, err
			}

			if err == nil {
				branchOffice.Company = &company
			}
		}

		warehouses[i].Branch_Office = &branchOffice
	}

	return warehouses, nil
}

// Obtener un almacén por Id
func (r *WarehouseRepository) GetWarehouseById(warehouseId int) (*domain.Warehouses, error) {
	query := `SELECT id, branch_office_id, name, description, address, status FROM warehouses WHERE id = $1 AND deleted_at IS NULL`
	row := r.db.QueryRow(query, warehouseId)

	var warehouse domain.Warehouses
	if err := row.Scan(&warehouse.Id, &warehouse.Branch_Office_Id, &warehouse.Name, &warehouse.Description, &warehouse.Address, &warehouse.Status); err != nil {
		return nil, err
	}

	// Obtener sucursal del almacén
	var branchOffice domain.BranchOffices
	err := r.db.Get(&branchOffice, `
		SELECT id, company_id, name, description, address, phone, status
		FROM branch_offices
		WHERE id = $1 AND deleted_at IS NULL
	`, warehouse.Branch_Office_Id)
	if err != nil && err != sql.ErrNoRows { // Manejar casos donde no hay sucursal asociada
		fmt.Printf("Error obteniendo sucursal para el almacén %d: %v\n", warehouseId, err)
		return nil, err
	}

	// Obtener empresa de la sucursal (si existe)
	if branchOffice.Company_Id > 0 {
		var company domain.Companies
		err = r.db.Get(&company, `
			SELECT id, name, ruc, email, phone, web_site, address, status
			FROM companies 
			WHERE id = $1 AND deleted_at IS NULL
		`, branchOffice.Company_Id)

		if err != nil && err != sql.ErrNoRows {
			fmt.Printf("Error obteniendo empresa %d para la sucursal %d: %v\n", branchOffice.Company_Id, branchOffice.Id, err)
			return nil, err
		}

		if err == nil {
			branchOffice.Company = &company
		}
	}

	warehouse.Branch_Office = &branchOffice

	return &warehouse, nil
}

// Actualizar un almacén
func (r *WarehouseRepository) UpdateWarehouse(warehouse *domain.Warehouses) error {
	query := `
		UPDATE warehouses SET 
			branch_office_id = $1, name = $2, description = $3, address = $4, status = $5, updated_at = NOW()
		WHERE id = $6 AND deleted_at IS NULL
	`
	_, err := r.db.Exec(query, warehouse.Branch_Office_Id, warehouse.Name, warehouse.Description, warehouse.Address, warehouse.Status, warehouse.Id)
	if err != nil {
		fmt.Printf("Error actualizando el almacén con Id %d: %v\n", warehouse.Id, err)
		return err
	}
	return nil
}

// Eliminar un almacén por Id (eliminación lógica)
func (r *WarehouseRepository) DeleteWarehouseById(id int) error {
	query := `
		UPDATE warehouses
		SET deleted_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
	`
	_, err := r.db.Exec(query, id)
	if err != nil {
		fmt.Printf("Error eliminando el almacén con Id %d: %v\n", id, err)
		return err
	}
	return nil
}

// Eliminar múltiples almacenes por Ids (eliminación lógica)
func (r *WarehouseRepository) DeleteWarehousesByIds(ids []int) error {
	query := `UPDATE warehouses SET deleted_at = NOW() WHERE id = ANY($1)`
	_, err := r.db.Exec(query, pq.Array(ids))
	return err
}
