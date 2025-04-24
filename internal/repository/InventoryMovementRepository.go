package repository

import (
	"database/sql"
	"fmt"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

type InventoryMovementRepository struct {
	db *sqlx.DB
}

func NewInventoryMovementRepository(db *sqlx.DB) *InventoryMovementRepository {
	return &InventoryMovementRepository{db: db}
}

// Crear una marca
func (r *InventoryMovementRepository) CreateInventoryMovement(inventoryMovement *domain.InventoryMovements) error {
	query := `
		INSERT INTO inventory_movements (
			warehouse_id, product_id, movement_type, quantity, reference, user_id
		) VALUES (
		 	:warehouse_id, :product_id, :movement_type, :quantity, :reference, :user_id
		)
	`
	_, err := r.db.NamedExec(query, inventoryMovement)
	if err != nil {
		fmt.Printf("Error creando la marca: %v\n", err)
		return err
	}
	return nil
}

// Obtener todas las marcas
func (r *InventoryMovementRepository) GetAllInventoryMovements() ([]domain.InventoryMovements, error) {
	var inventoryMovement []domain.InventoryMovements
	query := "SELECT id, warehouse_id, product_id, movement_type, quantity, reference, user_id FROM inventory_movements WHERE deleted_at IS NULL"
	err := r.db.Select(&inventoryMovement, query)
	if err != nil {
		fmt.Printf("Error ejecutando la consulta: %v\n", err)
		return nil, err
	}

	// Iterar sobre el control de stock
	for i := range inventoryMovement {
		// Obtener almacén del control de stock
		var customer domain.Warehouses
		err = r.db.Get(&customer, `
			SELECT id, branch_office_id, name, description, address, status
			FROM warehouses
			WHERE id = $1 AND deleted_at IS NULL
		`, inventoryMovement[i].Warehouse_Id)
		if err != nil && err != sql.ErrNoRows { // Manejar casos donde no hay almacén asociado
			fmt.Printf("Error obteniendo el almacén para del control de stock %d: %v\n", inventoryMovement[i].Id, err)
			return nil, err
		}
		inventoryMovement[i].Warehouse = &customer

		// Obtener producto del control de stock
		var product domain.Products
		err = r.db.Get(&product, `
			SELECT *
			FROM products
			WHERE id = $1 AND deleted_at IS NULL
		`, inventoryMovement[i].Product_Id)
		if err != nil && err != sql.ErrNoRows { // Manejar casos donde no hay producto asociado
			fmt.Printf("Error obteniendo el producto para del control de stock %d: %v\n", inventoryMovement[i].Id, err)
			return nil, err
		}
		inventoryMovement[i].Product = &product

		// Obtener usuario de la orden de venta
		var user domain.Users
		err = r.db.Get(&user, `
			SELECT id, role_id, display_name, employee_id, photo_url, email, status FROM users WHERE id = $1 AND deleted_at IS NULL
		`, inventoryMovement[i].User_Id)
		if err != nil && err != sql.ErrNoRows { // Manejar casos donde no hay usuario asociado
			fmt.Printf("Error obteniendo el usuario para la orden de venta %d: %v\n", inventoryMovement[i].Id, err)
			return nil, err
		}
		inventoryMovement[i].User = &user

		// Obtener empleado del usuario (si existe)
		if user.Employee_Id.Valid {
			var company domain.Employees
			err = r.db.Get(&company, `
				SELECT id, first_name, second_name, third_name, surname, second_surname, photo_url, warehouse_id, document_type, document_number, birth_date, gender, email, phone, address, hire_date, position, salary, status
				FROM employees 
				WHERE id = $1 AND deleted_at IS NULL
			`, user.Employee_Id)

			if err != nil && err != sql.ErrNoRows {
				return nil, err
			}

			if err == nil {
				user.Employee = &company
			}
		}
	}

	return inventoryMovement, nil
}

// Obtener una marca por Id
func (r *InventoryMovementRepository) GetInventoryMovementById(id int) (*domain.InventoryMovements, error) {
	query := `SELECT id, warehouse_id, product_id, movement_type, quantity, reference, user_id FROM inventory_movements WHERE id = $1 AND deleted_at IS NULL`
	row := r.db.QueryRow(query, id)

	var inventoryMovement domain.InventoryMovements
	if err := row.Scan(&inventoryMovement.Id, &inventoryMovement.Warehouse_Id, &inventoryMovement.Product_Id, &inventoryMovement.Movement_Type, &inventoryMovement.Quantity, &inventoryMovement.Reference, &inventoryMovement.User_Id); err != nil {
		return nil, err
	}

	// Obtener almacén del control de stock
	var customer domain.Warehouses
	err := r.db.Get(&customer, `
		SELECT id, branch_office_id, name, description, address, status
		FROM warehouses
		WHERE id = $1 AND deleted_at IS NULL
	`, inventoryMovement.Warehouse_Id)
	if err != nil && err != sql.ErrNoRows { // Manejar casos donde no hay almacén asociado
		fmt.Printf("Error obteniendo el almacén para del control de stock %d: %v\n", inventoryMovement.Id, err)
		return nil, err
	}
	inventoryMovement.Warehouse = &customer

	// Obtener producto del control de stock
	var product domain.Products
	err = r.db.Get(&product, `
		SELECT *
		FROM products
		WHERE id = $1 AND deleted_at IS NULL
	`, inventoryMovement.Product_Id)
	if err != nil && err != sql.ErrNoRows { // Manejar casos donde no hay producto asociado
		fmt.Printf("Error obteniendo el producto para del control de stock %d: %v\n", inventoryMovement.Id, err)
		return nil, err
	}
	inventoryMovement.Product = &product

	// Obtener usuario de la orden de venta
	var user domain.Users
	err = r.db.Get(&user, `
		SELECT id, role_id, display_name, employee_id, photo_url, email, status FROM users WHERE id = $1 AND deleted_at IS NULL
	`, inventoryMovement.User_Id)
	if err != nil && err != sql.ErrNoRows { // Manejar casos donde no hay usuario asociado
		fmt.Printf("Error obteniendo el usuario para la orden de venta %d: %v\n", inventoryMovement.Id, err)
		return nil, err
	}
	inventoryMovement.User = &user

	// Obtener empleado del usuario (si existe)
	if user.Employee_Id.Valid {
		var company domain.Employees
		err = r.db.Get(&company, `
			SELECT id, first_name, second_name, third_name, surname, second_surname, photo_url, warehouse_id, document_type, document_number, birth_date, gender, email, phone, address, hire_date, position, salary, status
			FROM employees 
			WHERE id = $1 AND deleted_at IS NULL
		`, user.Employee_Id)

		if err != nil && err != sql.ErrNoRows {
			return nil, err
		}

		if err == nil {
			user.Employee = &company
		}
	}

	return &inventoryMovement, nil
}

// Actualizar una marca
func (r *InventoryMovementRepository) UpdateInventoryMovement(inventoryMovement *domain.InventoryMovements) error {
	query := `
		UPDATE inventory_movements SET 
			warehouse_id = $1, product_id = $2, movement_type = $3, quantity = $4, reference = $5, user_id = $6, updated_at = NOW()
		WHERE id = $7 AND deleted_at IS NULL
	`
	_, err := r.db.Exec(query, inventoryMovement.Warehouse_Id, inventoryMovement.Product_Id, inventoryMovement.Movement_Type, inventoryMovement.Quantity, inventoryMovement.Reference, inventoryMovement.User_Id, inventoryMovement.Id)
	if err != nil {
		fmt.Printf("Error actualizando la marca con Id %d: %v\n", inventoryMovement.Id, err)
		return err
	}
	return nil
}

// Eliminar una marca por Id (eliminación lógica)
func (r *InventoryMovementRepository) DeleteInventoryMovementById(id int) error {
	query := `
		UPDATE inventory_movements
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
func (r *InventoryMovementRepository) DeleteInventoryMovementsByIds(ids []int) error {
	query := `UPDATE inventory_movements SET deleted_at = NOW() WHERE id = ANY($1)`
	_, err := r.db.Exec(query, pq.Array(ids))
	return err
}
