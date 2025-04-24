package repository

import (
	"fmt"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

type SupplierRepository struct {
	db *sqlx.DB
}

func NewSupplierRepository(db *sqlx.DB) *SupplierRepository {
	return &SupplierRepository{db: db}
}

// Crear un proveedor
func (r *SupplierRepository) CreateSupplier(supplier *domain.Suppliers) error {
	query := `
		INSERT INTO suppliers (
			name, ruc, email, phone, web_site, address, status
		) VALUES (
		 	:name, :ruc, :email, :phone, :web_site, :address, :status
		)
	`
	_, err := r.db.NamedExec(query, supplier)
	if err != nil {
		fmt.Printf("Error creando el proveedor: %v\n", err)
		return err
	}
	return nil
}

// Obtener todos los proveedores
func (r *SupplierRepository) GetAllSuppliers() ([]domain.Suppliers, error) {
	var suppliers []domain.Suppliers
	query := "SELECT id, name, ruc, email, phone, web_site, address, status FROM suppliers WHERE deleted_at IS NULL"
	err := r.db.Select(&suppliers, query)
	if err != nil {
		fmt.Printf("Error ejecutando la consulta: %v\n", err)
		return nil, err
	}
	return suppliers, nil
}

// Obtener un proveedor por Id
func (r *SupplierRepository) GetSupplierById(supplierId int) (*domain.Suppliers, error) {
	var supplier domain.Suppliers
	query := "SELECT id, name, ruc, email, phone, web_site, address, status FROM suppliers WHERE id = $1 AND deleted_at IS NULL"
	err := r.db.Get(&supplier, query, supplierId)
	if err != nil {
		fmt.Printf("Error obteniendo proveedor con ID %d: %v\n", supplierId, err)
		return nil, err
	}
	return &supplier, nil
}

// Actualizar un proveedor
func (r *SupplierRepository) UpdateSupplier(supplier *domain.Suppliers) error {
	query := `
		UPDATE suppliers SET
			name = $1, ruc = $2, email = $3, phone = $4, web_site = $5, address = $6, status = $7, updated_at = NOW()
		WHERE id = $8
	`
	_, err := r.db.Exec(query, supplier.Name, supplier.Ruc, supplier.Email, supplier.Phone, supplier.Web_Site, supplier.Address, supplier.Status, supplier.Id)
	return err
}

// Eliminar un proveedor por Id (eliminación lógica)
func (r *SupplierRepository) DeleteSupplierById(id int) error {
	query := `UPDATE suppliers SET deleted_at = NOW() WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

// Eliminar múltiples proveedores por Ids (eliminación lógica)
func (r *SupplierRepository) DeleteSuppliersByIds(ids []int) error {
	query := `UPDATE suppliers SET deleted_at = NOW() WHERE id = ANY($1)`
	_, err := r.db.Exec(query, pq.Array(ids))
	return err
}
