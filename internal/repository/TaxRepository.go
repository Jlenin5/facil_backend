package repository

import (
	"fmt"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

type TaxRepository struct {
	db *sqlx.DB
}

func NewTaxRepository(db *sqlx.DB) *TaxRepository {
	return &TaxRepository{db: db}
}

// Crear un impuesto
func (r *TaxRepository) CreateTax(tax *domain.Taxes) error {
	query := `
		INSERT INTO taxes (
			name, description, rate, tax_type, status
		) VALUES (
		 	:name, :description, :rate, :tax_type, :status
		)
	`
	_, err := r.db.NamedExec(query, tax)
	if err != nil {
		fmt.Printf("Error creando el impuesto: %v\n", err)
		return err
	}
	return nil
}

// Obtener todas las impuestos
func (r *TaxRepository) GetAllTaxes() ([]domain.Taxes, error) {
	var taxes []domain.Taxes
	query := "SELECT id, name, description, rate, tax_type, status FROM taxes WHERE deleted_at IS NULL"
	err := r.db.Select(&taxes, query)
	if err != nil {
		fmt.Printf("Error ejecutando la consulta: %v\n", err)
		return nil, err
	}
	return taxes, nil
}

// Obtener un impuesto por Id
func (r *TaxRepository) GetTaxById(id int) (*domain.Taxes, error) {
	query := `SELECT id, name, description, rate, tax_type, status FROM taxes WHERE id = $1 AND deleted_at IS NULL`
	row := r.db.QueryRow(query, id)

	var tax domain.Taxes
	if err := row.Scan(&tax.Id, &tax.Name, &tax.Description, &tax.Rate, &tax.Tax_type, &tax.Status); err != nil {
		return nil, err
	}
	return &tax, nil
}

// Actualizar un impuesto
func (r *TaxRepository) UpdateTax(tax *domain.Taxes) error {
	query := `
		UPDATE taxes SET 
			name = $1, description = $2, rate = $3, tax_type = $4, status = $5, updated_at = NOW()
		WHERE id = $6 AND deleted_at IS NULL
	`
	_, err := r.db.Exec(query, tax.Name, tax.Description, tax.Rate, tax.Tax_type, tax.Status, tax.Id)
	if err != nil {
		fmt.Printf("Error actualizando el impuesto con Id %d: %v\n", tax.Id, err)
		return err
	}
	return nil
}

// Eliminar un impuesto por Id (eliminación lógica)
func (r *TaxRepository) DeleteTaxById(id int) error {
	query := `
		UPDATE taxes
		SET deleted_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
	`
	_, err := r.db.Exec(query, id)
	if err != nil {
		fmt.Printf("Error eliminando el impuesto con Id %d: %v\n", id, err)
		return err
	}
	return nil
}

// Eliminar múltiples impuestos por Ids (eliminación lógica)
func (r *TaxRepository) DeleteTaxesByIds(ids []int) error {
	query := `UPDATE taxes SET deleted_at = NOW() WHERE id = ANY($1)`
	_, err := r.db.Exec(query, pq.Array(ids))
	return err
}
