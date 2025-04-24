package repository

import (
	"fmt"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

type CurrencyRepository struct {
	db *sqlx.DB
}

func NewCurrencyRepository(db *sqlx.DB) *CurrencyRepository {
	return &CurrencyRepository{db: db}
}

// Crear una moneda
func (r *CurrencyRepository) CreateCurrency(currency *domain.Currencies) error {
	query := "INSERT INTO currencies (name, description, status) VALUES (:name, :description, :status)"
	_, err := r.db.NamedExec(query, currency)
	if err != nil {
		fmt.Printf("Error creando la moneda: %v\n", err)
		return err
	}
	return nil
}

// Obtener todas las monedas
func (r *CurrencyRepository) GetAllCurrencies() ([]domain.Currencies, error) {
	var currencies []domain.Currencies
	query := "SELECT id, name, description, code, symbol, status FROM currencies WHERE deleted_at IS NULL"
	err := r.db.Select(&currencies, query)
	if err != nil {
		fmt.Printf("Error ejecutando la consulta: %v\n", err)
		return nil, err
	}
	return currencies, nil
}

// Obtener una moneda por Id
func (r *CurrencyRepository) GetCurrencyById(id int) (*domain.Currencies, error) {
	query := `SELECT id, name, description, code, symbol, status FROM currencies WHERE id = $1 AND deleted_at IS NULL`
	row := r.db.QueryRow(query, id)

	var currency domain.Currencies
	if err := row.Scan(&currency.Id, &currency.Name, &currency.Description, &currency.Code, &currency.Symbol, &currency.Status); err != nil {
		return nil, err
	}
	return &currency, nil
}

// Actualizar una moneda
func (r *CurrencyRepository) UpdateCurrency(currency *domain.Currencies) error {
	query := `
		UPDATE currencies SET 
			name = $1, description = $2, status = $3, updated_at = NOW()
		WHERE id = $4 AND deleted_at IS NULL
	`
	_, err := r.db.Exec(query, currency.Name, currency.Description, currency.Status, currency.Id)
	if err != nil {
		fmt.Printf("Error actualizando la moneda con Id %d: %v\n", currency.Id, err)
		return err
	}
	return nil
}

// Eliminar una moneda por Id (eliminación lógica)
func (r *CurrencyRepository) DeleteCurrencyById(id int) error {
	query := `
		UPDATE currencies
		SET deleted_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
	`
	_, err := r.db.Exec(query, id)
	if err != nil {
		fmt.Printf("Error eliminando la moneda con Id %d: %v\n", id, err)
		return err
	}
	return nil
}

// Eliminar múltiples monedas por Ids (eliminación lógica)
func (r *CurrencyRepository) DeleteCurrenciesByIds(ids []int) error {
	query := `UPDATE currencies SET deleted_at = NOW() WHERE id = ANY($1)`
	_, err := r.db.Exec(query, pq.Array(ids))
	return err
}
