package repository

import (
	"fmt"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

type ExchangeRateRepository struct {
	db *sqlx.DB
}

func NewExchangeRateRepository(db *sqlx.DB) *ExchangeRateRepository {
	return &ExchangeRateRepository{db: db}
}

// Crear una tipo de cambio
func (r *ExchangeRateRepository) CreateExchangeRate(exchangeRate *domain.ExchangeRates) error {
	query := `
		INSERT INTO exchange_rates (
			base_currency_id, target_currency_id, exchange_rate, created_by, updated_by
		) VALUES (
		 	:base_currency_id, :target_currency_id, :exchange_rate, :created_by, :updated_by
		)
	`
	_, err := r.db.NamedExec(query, exchangeRate)
	if err != nil {
		fmt.Printf("Error creando la tipo de cambio: %v\n", err)
		return err
	}
	return nil
}

// Obtener todas los tipos de cambios
func (r *ExchangeRateRepository) GetAllExchangeRates() ([]domain.ExchangeRates, error) {
	var exchangeRates []domain.ExchangeRates
	query := `
		SELECT
			er.id, er.base_currency_id, er.target_currency_id, er.exchange_rate, er.created_by, er.updated_by,
			c.id AS "base_currency.id", c.name AS "base_currency.name", c.code AS "base_currency.code", c.symbol AS "base_currency.symbol",
			cu.id AS "target_currency.id", cu.name AS "target_currency.name", cu.code AS "target_currency.code", cu.symbol AS "target_currency.symbol"
		FROM exchange_rates er
		LEFT JOIN currencies c ON er.base_currency_id=c.id
		LEFT JOIN currencies cu ON er.target_currency_id=cu.id
		WHERE er.deleted_at IS NULL
		ORDER BY er.id DESC
	`
	err := r.db.Select(&exchangeRates, query)
	if err != nil {
		fmt.Printf("Error ejecutando la consulta: %v\n", err)
		return nil, err
	}
	return exchangeRates, nil
}

// Obtener un tipo de cambio por Id
func (r *ExchangeRateRepository) GetExchangeRateById(exchangeRateId int) (*domain.ExchangeRates, error) {
	var exchangeRates domain.ExchangeRates
	query := `
		SELECT
			er.id, er.base_currency_id, er.target_currency_id, er.exchange_rate, er.created_by, er.updated_by,
			c.id AS "base_currency.id", c.name AS "base_currency.name", c.code AS "base_currency.code", c.symbol AS "base_currency.symbol",
			cu.id AS "target_currency.id", cu.name AS "target_currency.name", cu.code AS "target_currency.code", cu.symbol AS "target_currency.symbol",
			u.id AS "created_by.id", u.employee_id AS "created_by.employee_id",
			e.id AS "created_by.employee.id", e.first_name AS "created_by.employee.first_name", e.second_name AS "created_by.employee.second_name", e.third_name AS "created_by.employee.third_name", e.surname AS "created_by.employee.surname", e.second_surname AS "created_by.employee.second_surname",
			COALESCE(us.id, 0) AS "updated_by.id", us.employee_id AS "updated_by.employee_id",
			COALESCE(em.id, 0) AS "updated_by.employee.id", COALESCE(em.first_name, '') AS "updated_by.employee.first_name", COALESCE(em.second_name, '') AS "updated_by.employee.second_name", COALESCE(em.third_name, '') AS "updated_by.employee.third_name", COALESCE(em.surname, '') AS "updated_by.employee.surname", COALESCE(em.second_surname, '') AS "updated_by.employee.second_surname"
		FROM exchange_rates er
		LEFT JOIN currencies c ON er.base_currency_id=c.id
		LEFT JOIN currencies cu ON er.target_currency_id=cu.id
		LEFT JOIN users u ON er.created_by=u.id
		LEFT JOIN employees e ON u.employee_id=e.id
		LEFT JOIN users us ON er.updated_by=us.id
		LEFT JOIN employees em ON us.employee_id=em.id
		WHERE er.id = $1 AND er.deleted_at IS NULL
		ORDER BY er.id DESC
	`
	err := r.db.Get(&exchangeRates, query, exchangeRateId)
	if err != nil {
		fmt.Printf("Error obteniendo tipo de cambio con Id %d: %v\n", exchangeRateId, err)
		return nil, err
	}

	return &exchangeRates, nil
}

// Actualizar un tipo de cambio
func (r *ExchangeRateRepository) UpdateExchangeRate(exchangeRate *domain.ExchangeRates) error {
	tx, err := r.db.Beginx()
	if err != nil {
		fmt.Printf("Error iniciando transacción: %v\n", err)
		return err
	}
	exchangeRateQuery := `
		UPDATE exchange_rates
		SET
			base_currency_id = :base_currency_id,
			target_currency_id = :target_currency_id,
			exchange_rate = :exchange_rate,
			created_by = :created_by,
			updated_by = :updated_by,
			updated_at = NOW()
		WHERE id = :id
	`
	_, err = tx.NamedExec(exchangeRateQuery, exchangeRate)
	if err != nil {
		fmt.Printf("Error actualizando el tipo de cambio: %v\n", err)
		tx.Rollback()
		return err
	}
	return nil
}

// Eliminar un tipo de cambio por Id (eliminación lógica)
func (r *ExchangeRateRepository) DeleteExchangeRateById(id int) error {
	query := `
		UPDATE exchange_rates
		SET deleted_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
	`
	_, err := r.db.Exec(query, id)
	if err != nil {
		fmt.Printf("Error eliminando el tipo de cambio con Id %d: %v\n", id, err)
		return err
	}
	return nil
}

// Eliminar múltiples tipos de cambios por Ids (eliminación lógica)
func (r *ExchangeRateRepository) DeleteExchangeRatesByIds(ids []int) error {
	query := `UPDATE exchange_rates SET deleted_at = NOW() WHERE id = ANY($1)`
	_, err := r.db.Exec(query, pq.Array(ids))
	return err
}
