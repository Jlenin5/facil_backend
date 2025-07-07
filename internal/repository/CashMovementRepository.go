package repository

import (
	"database/sql"
	"fmt"

	"github.com/Jlenin5/facil_backend/internal/domain"

	"github.com/jmoiron/sqlx"
)

type CashMovementRepository struct {
	db *sqlx.DB
}

func NewCashMovementRepository(db *sqlx.DB) *CashMovementRepository {
	return &CashMovementRepository{db: db}
}

// Crear un movimiento de caja
func (r *CashMovementRepository) CreateCashMovement(user *domain.CashMovements) error {
	query := `
		INSERT INTO cash_movements (
			cash_register_id, movement_type, payment_method_id, amount, description, user_id
		) VALUES (
			:cash_register_id, :movement_type, :payment_method_id, :amount, :description, :user_id
		)
	`
	_, err := r.db.NamedExec(query, user)
	if err != nil {
		fmt.Printf("Error creando el movimiento de caja: %v\n", err)
		return err
	}
	return nil
}

// Obtener todos los movimiento de cajas
func (r *CashMovementRepository) GetAllCashMovements() ([]domain.CashMovements, error) {
	var cashRegisters []domain.CashMovements
	query := `
		SELECT
			cm.id, cm.cash_register_id, cm.movement_type, cm.payment_method_id, cm.amount, cm.description, cm.user_id,
			cr.id AS "cash_register.id",
			pm.id AS "payment_method.id", pm.name AS "payment_method.name",
			u.id AS "user.id", u.employee_id AS "user.employee_id",
			eo.id AS "user.employee.id", eo.names AS "user.employee.names", eo.surname AS "user.employee.surname", eo.second_surname AS "user.employee.second_surname"
		FROM cash_movements cm
		LEFT JOIN cash_registers cr ON cm.cash_register_id=cr.id
		LEFT JOIN payment_methods pm ON cm.payment_method_id=pm.id
		LEFT JOIN users u ON cm.user_id=u.id
		LEFT JOIN employees eo ON u.employee_id=eo.id
		ORDER BY cm.id DESC
	`
	err := r.db.Select(&cashRegisters, query)
	if err != nil {
		fmt.Printf("Error ejecutando la consulta: %v\n", err)
		return nil, err
	}

	return cashRegisters, nil
}

// Obtener un movimiento de caja por Id
func (r *CashMovementRepository) GetCashMovementById(cashRegisterId int) (*domain.CashMovements, error) {
	var cashRegister domain.CashMovements
	query := `
		SELECT
			cm.id, cm.cash_register_id, cm.movement_type, cm.payment_method_id, cm.amount, cm.description, cm.user_id,
			cr.id AS "cash_register.id",
			pm.id AS "payment_method.id", pm.name AS "payment_method.name",
			u.id AS "user.id", u.employee_id AS "user.employee_id",
			eo.id AS "user.employee.id", eo.first_name AS "user.employee.names", eo.surname AS "user.employee.surname", eo.second_surname AS "user.employee.second_surname"
		FROM cash_movements cm
		LEFT JOIN cash_registers cr ON cm.cash_register_id=cr.id
		LEFT JOIN payment_methods pm ON cm.payment_method_id=pm.id
		LEFT JOIN users u ON cm.user_id=u.id
		LEFT JOIN employees eo ON u.employee_id=eo.id
		WHERE cm.id = $1
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

// Actualizar un movimiento de caja
func (r *CashMovementRepository) UpdateCashMovement(user *domain.CashMovements) error {
	query := `
		UPDATE cash_movements SET 
			cash_register_id = :cash_register_id,
			movement_type = :movement_type,
			payment_method_id = :payment_method_id,
			amount = :amount,
			description = :description,
			user_id = :user_id,
			updated_at = NOW()
		WHERE id = :id AND deleted_at IS NULL
	`
	_, err := r.db.NamedExec(query, user)
	if err != nil {
		return fmt.Errorf("error updating user: %w", err)
	}
	return nil
}