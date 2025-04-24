package repository

import (
	"database/sql"
	"fmt"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

type CashRegisterRepository struct {
	db *sqlx.DB
}

func NewCashRegisterRepository(db *sqlx.DB) *CashRegisterRepository {
	return &CashRegisterRepository{db: db}
}

// Crear un registro de caja
func (r *CashRegisterRepository) CreateCashRegister(user *domain.CashRegisters) error {
	query := `
		INSERT INTO cash_registers (
			warehouse_id, user_open_id, user_close_id, opening_date, closing_date, initial_amount, closing_amount, difference, status
		) VALUES (
			:warehouse_id, :user_open_id, :user_close_id, :opening_date, :closing_date, :initial_amount, :closing_amount, :difference, :status
		)
	`
	_, err := r.db.NamedExec(query, user)
	if err != nil {
		fmt.Printf("Error creando el registro de caja: %v\n", err)
		return err
	}
	return nil
}

// Obtener todos los registro de cajas
func (r *CashRegisterRepository) GetAllCashRegisters() ([]domain.CashRegisters, error) {
	var cashRegisters []domain.CashRegisters
	query := `
		SELECT
			cr.id, cr.warehouse_id, cr.user_open_id, cr.user_close_id, cr.opening_date, cr.closing_date, cr.initial_amount, cr.closing_amount, cr.difference, cr.status,
			w.id AS "warehouse.id", w.name AS "warehouse.name",
			uo.id AS "user_open.id", uo.employee_id AS "user_open.employee_id",
			eo.id AS "user_open.employee.id", eo.first_name AS "user_open.employee.first_name", eo.second_name AS "user_open.employee.second_name", eo.third_name AS "user_open.employee.third_name", eo.surname AS "user_open.employee.surname", eo.second_surname AS "user_open.employee.second_surname",
			COALESCE(uc.id, 0) AS "user_close.id", COALESCE(uc.employee_id, 0) AS "user_close.employee_id",
			COALESCE(ec.id, 0) AS "user_close.employee.id", COALESCE(ec.first_name, '') AS "user_close.employee.first_name", COALESCE(ec.second_name, '') AS "user_close.employee.second_name", COALESCE(ec.third_name, '') AS "user_close.employee.third_name", COALESCE(ec.surname, '') AS "user_close.employee.surname", COALESCE(ec.second_surname, '') AS "user_close.employee.second_surname"
		FROM cash_registers cr
		LEFT JOIN warehouses w ON cr.warehouse_id=w.id
		LEFT JOIN users uo ON cr.user_open_id=uo.id
		LEFT JOIN employees eo ON uo.employee_id=eo.id
		LEFT JOIN users uc ON cr.user_close_id=uc.id
		LEFT JOIN employees ec ON uc.employee_id=ec.id
		WHERE cr.deleted_at IS NULL
		ORDER BY cr.id DESC
	`
	err := r.db.Select(&cashRegisters, query)
	if err != nil {
		fmt.Printf("Error ejecutando la consulta: %v\n", err)
		return nil, err
	}

	return cashRegisters, nil
}

// Obtener un registro de caja por Id
func (r *CashRegisterRepository) GetCashRegisterById(cashRegisterId int) (*domain.CashRegisters, error) {
	var cashRegister domain.CashRegisters
	query := `
		SELECT
			cr.id, cr.warehouse_id, cr.user_open_id, cr.user_close_id, cr.opening_date, cr.closing_date, cr.initial_amount, cr.closing_amount, cr.difference, cr.status,
			w.id AS "warehouse.id", w.name AS "warehouse.name",
			uo.id AS "user_open.id", uo.employee_id AS "user_open.employee_id",
			eo.id AS "user_open.employee.id", eo.first_name AS "user_open.employee.first_name", eo.second_name AS "user_open.employee.second_name", eo.third_name AS "user_open.employee.third_name", eo.surname AS "user_open.employee.surname", eo.second_surname AS "user_open.employee.second_surname",
			COALESCE(uc.id, 0) AS "user_close.id", COALESCE(uc.employee_id, 0) AS "user_close.employee_id",
			COALESCE(ec.id, 0) AS "user_close.employee.id", COALESCE(ec.first_name, '') AS "user_close.employee.first_name", COALESCE(ec.second_name, '') AS "user_close.employee.second_name", COALESCE(ec.third_name, '') AS "user_close.employee.third_name", COALESCE(ec.surname, '') AS "user_close.employee.surname", COALESCE(ec.second_surname, '') AS "user_close.employee.second_surname"
		FROM cash_registers cr
		LEFT JOIN warehouses w ON cr.warehouse_id=w.id
		LEFT JOIN users uo ON cr.user_open_id=uo.id
		LEFT JOIN employees eo ON uo.employee_id=eo.id
		LEFT JOIN users uc ON cr.user_close_id=uc.id
		LEFT JOIN employees ec ON uc.employee_id=ec.id
		WHERE cr.id = $1 AND cr.deleted_at IS NULL
	`

	// Usar Get para mapear automáticamente los resultados
	err := r.db.Get(&cashRegister, query, cashRegisterId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // Registro de caja no encontrado
		}
		fmt.Printf("Error ejecutando la consulta: %v\n", err)
		return nil, err
	}

	return &cashRegister, nil
}

// Actualizar un registro de caja
func (r *CashRegisterRepository) UpdateCashRegister(user *domain.CashRegisters) error {
	query := `
		UPDATE cash_registers SET 
			warehouse_id = :warehouse_id,
			user_open_id = :user_open_id,
			user_close_id = :user_close_id,
			opening_date = :opening_date,
			closing_date = :closing_date,
			initial_amount = :initial_amount,
			closing_amount = :closing_amount,
			difference = :difference,
			status = :status,
			updated_at = NOW()
		WHERE id = :id AND deleted_at IS NULL
	`
	_, err := r.db.NamedExec(query, user)
	if err != nil {
		return fmt.Errorf("error updating user: %w", err)
	}
	return nil
}

// Eliminar un registro de caja por Id (eliminación lógica)
func (r *CashRegisterRepository) DeleteCashRegisterById(id int) error {
	query := `UPDATE cash_registers SET deleted_at = NOW() WHERE id = $1 AND deleted_at IS NULL`
	_, err := r.db.Exec(query, id)
	if err != nil {
		fmt.Printf("Error eliminando el registro de caja con Id %d: %v\n", id, err)
		return err
	}
	return nil
}

// Eliminar múltiples registro de cajas por Ids (eliminación lógica)
func (r *CashRegisterRepository) DeleteCashRegistersByIds(ids []int) error {
	query := `UPDATE cash_registers SET deleted_at = NOW() WHERE id = ANY($1)`
	_, err := r.db.Exec(query, pq.Array(ids))
	return err
}