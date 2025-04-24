package repository

import (
	"fmt"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

type PaymentMethodRepository struct {
	db *sqlx.DB
}

func NewPaymentMethodRepository(db *sqlx.DB) *PaymentMethodRepository {
	return &PaymentMethodRepository{db: db}
}

// Crear un método de pago
func (r *PaymentMethodRepository) CreatePaymentMethod(paymentMethod *domain.PaymentMethods) error {
	query := `
		INSERT INTO payment_methods (
			name, description, status
		) VALUES (
		 	:name, :description, :status
		)
	`
	_, err := r.db.NamedExec(query, paymentMethod)
	if err != nil {
		fmt.Printf("Error creando la método de pago: %v\n", err)
		return err
	}
	return nil
}

// Obtener todas las métodos de pago
func (r *PaymentMethodRepository) GetAllPaymentMethods() ([]domain.PaymentMethods, error) {
	var paymentMethods []domain.PaymentMethods
	query := "SELECT id, name, description, status FROM payment_methods WHERE deleted_at IS NULL"
	err := r.db.Select(&paymentMethods, query)
	if err != nil {
		fmt.Printf("Error ejecutando la consulta: %v\n", err)
		return nil, err
	}
	return paymentMethods, nil
}

// Obtener un método de pago por Id
func (r *PaymentMethodRepository) GetPaymentMethodById(id int) (*domain.PaymentMethods, error) {
	query := `SELECT id, name, description, status FROM payment_methods WHERE id = $1 AND deleted_at IS NULL`
	row := r.db.QueryRow(query, id)

	var paymentMethod domain.PaymentMethods
	if err := row.Scan(&paymentMethod.Id, &paymentMethod.Name, &paymentMethod.Description, &paymentMethod.Status); err != nil {
		return nil, err
	}
	return &paymentMethod, nil
}

// Actualizar un método de pago
func (r *PaymentMethodRepository) UpdatePaymentMethod(paymentMethod *domain.PaymentMethods) error {
	query := `
		UPDATE payment_methods SET 
			name = $1, description = $2, status = $5, updated_at = NOW()
		WHERE id = $6 AND deleted_at IS NULL
	`
	_, err := r.db.Exec(query, paymentMethod.Name, paymentMethod.Description, paymentMethod.Status, paymentMethod.Id)
	if err != nil {
		fmt.Printf("Error actualizando la método de pago con Id %d: %v\n", paymentMethod.Id, err)
		return err
	}
	return nil
}

// Eliminar un método de pago por Id (eliminación lógica)
func (r *PaymentMethodRepository) DeletePaymentMethodById(id int) error {
	query := `
		UPDATE payment_methods
		SET deleted_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
	`
	_, err := r.db.Exec(query, id)
	if err != nil {
		fmt.Printf("Error eliminando la método de pago con Id %d: %v\n", id, err)
		return err
	}
	return nil
}

// Eliminar múltiples métodos de pago por Ids (eliminación lógica)
func (r *PaymentMethodRepository) DeletePaymentMethodsByIds(ids []int) error {
	query := `UPDATE payment_methods SET deleted_at = NOW() WHERE id = ANY($1)`
	_, err := r.db.Exec(query, pq.Array(ids))
	return err
}
