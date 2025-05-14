package repository

import (
	"database/sql"
	"fmt"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

type StockControlRepository struct {
	db *sqlx.DB
}

func NewStockControlRepository(db *sqlx.DB) *StockControlRepository {
	return &StockControlRepository{db: db}
}

// Crear una marca
func (r *StockControlRepository) CreateStockControl(brand *domain.StockControl) error {
	query := `
		INSERT INTO stock_control (
			warehouse_id, product_id, current_stock, current_booking, min_stock, max_stock
		) VALUES (
		 	:warehouse_id, :product_id, :current_stock, :current_booking, :min_stock, :max_stock
		)
	`
	_, err := r.db.NamedExec(query, brand)
	if err != nil {
		fmt.Printf("Error creando la marca: %v\n", err)
		return err
	}
	return nil
}

// Obtener todos los controles de stock
func (r *StockControlRepository) GetAllStockControl() ([]domain.StockControl, error) {
	var stockControl []domain.StockControl
	query := `
		SELECT
			sc.id, sc.warehouse_id, sc.product_id, sc.current_stock, sc.current_booking, sc.min_stock, sc.max_stock
		FROM stock_control sc
		WHERE sc.deleted_at IS NULL
		ORDER BY sc.id DESC
	`
	err := r.db.Select(&stockControl, query)
	if err != nil {
		fmt.Printf("Error ejecutando la consulta: %v\n", err)
		return nil, err
	}

	// Iterar sobre el control de stock
	for i := range stockControl {
		// Obtener almacén del control de stock
		var customer domain.Warehouses
		err = r.db.Get(&customer, `
			SELECT id, branch_office_id, name, description, address, status
			FROM warehouses
			WHERE id = $1 AND deleted_at IS NULL
		`, stockControl[i].Warehouse_Id)
		if err != nil && err != sql.ErrNoRows { // Manejar casos donde no hay almacén asociado
			fmt.Printf("Error obteniendo el almacén para del control de stock %d: %v\n", stockControl[i].Id, err)
			return nil, err
		}
		stockControl[i].Warehouse = &customer

		// Obtener producto del control de stock
		var product domain.Products
		err = r.db.Get(&product, `
			SELECT *
			FROM products
			WHERE id = $1 AND deleted_at IS NULL
		`, stockControl[i].Product_Id)
		if err != nil && err != sql.ErrNoRows { // Manejar casos donde no hay producto asociado
			fmt.Printf("Error obteniendo el producto para del control de stock %d: %v\n", stockControl[i].Id, err)
			return nil, err
		}
		stockControl[i].Product = &product
	}

	return stockControl, nil
}

// Obtener una marca por Id
func (r *StockControlRepository) GetStockControlById(id int) (*domain.StockControl, error) {
	query := `SELECT id, warehouse_id, product_id, current_stock, current_booking, min_stock, max_stock FROM stock_control WHERE id = $1 AND deleted_at IS NULL`
	row := r.db.QueryRow(query, id)

	var stockControl domain.StockControl
	if err := row.Scan(&stockControl.Id, &stockControl.Warehouse_Id, &stockControl.Product_Id, &stockControl.Current_Stock, &stockControl.Current_Booking, &stockControl.Min_Stock, &stockControl.Max_Stock); err != nil {
		return nil, err
	}

	// Obtener almacén del control de stock
	var customer domain.Warehouses
	err := r.db.Get(&customer, `
		SELECT id, branch_office_id, name, description, address, status
		FROM warehouses
		WHERE id = $1 AND deleted_at IS NULL
	`, stockControl.Warehouse_Id)
	if err != nil && err != sql.ErrNoRows { // Manejar casos donde no hay almacén asociado
		fmt.Printf("Error obteniendo el almacén para del control de stock %d: %v\n", stockControl.Id, err)
		return nil, err
	}
	stockControl.Warehouse = &customer

	// Obtener producto del control de stock
	var product domain.Products
	err = r.db.Get(&product, `
		SELECT *
		FROM products
		WHERE id = $1 AND deleted_at IS NULL
	`, stockControl.Product_Id)
	if err != nil && err != sql.ErrNoRows { // Manejar casos donde no hay producto asociado
		fmt.Printf("Error obteniendo el producto para del control de stock %d: %v\n", stockControl.Id, err)
		return nil, err
	}
	stockControl.Product = &product

	return &stockControl, nil
}

// Actualizar una marca
func (r *StockControlRepository) UpdateStockControl(stockControl *domain.StockControl) error {
	query := `
		UPDATE stock_control SET 
			warehouse_id = $1, product_id = $2, current_stock = $3, current_booking = $4, min_stock = $5, max_stock = $6, updated_at = NOW()
		WHERE id = $7 AND deleted_at IS NULL
	`
	_, err := r.db.Exec(query, stockControl.Warehouse_Id, stockControl.Product_Id, stockControl.Current_Stock, stockControl.Current_Booking, stockControl.Min_Stock, stockControl.Max_Stock, stockControl.Id)
	if err != nil {
		fmt.Printf("Error actualizando control de stock con Id %d: %v\n", stockControl.Id, err)
		return err
	}
	return nil
}

// Eliminar una marca por Id (eliminación lógica)
func (r *StockControlRepository) DeleteStockControlById(id int) error {
	query := `
		UPDATE stock_control
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
func (r *StockControlRepository) DeleteStockControlByIds(ids []int) error {
	query := `UPDATE stock_control SET deleted_at = NOW() WHERE id = ANY($1)`
	_, err := r.db.Exec(query, pq.Array(ids))
	return err
}
