package repository

import (
	"database/sql"
	"fmt"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

type OpportunityTrackingRepository struct {
	db *sqlx.DB
}

func NewOpportunityTrackingRepository(db *sqlx.DB) *OpportunityTrackingRepository {
	return &OpportunityTrackingRepository{db: db}
}

// Crear un empleado
func (r *OpportunityTrackingRepository) CreateOpportunityTracking(opportunityTracking *domain.OpportunityTracking) error {
	query := `
		INSERT INTO opportunity_tracking (
			customer_id, user_id, title, description, status, expected_revenue, probability
		) VALUES (
		 	:customer_id, :user_id, :title, :description, :status, :expected_revenue, :probability
		)
	`
	_, err := r.db.NamedExec(query, opportunityTracking)
	if err != nil {
		fmt.Printf("Error creando el empleado: %v\n", err)
		return err
	}
	return nil
}

// Obtener todos los seguimiento de oportunidades
func (r *OpportunityTrackingRepository) GetAllOpportunityTracking() ([]domain.OpportunityTracking, error) {
	var opportunityTracking []domain.OpportunityTracking
	query := "SELECT id, customer_id, user_id, title, description, status, expected_revenue, probability FROM opportunity_tracking WHERE deleted_at IS NULL"
	err := r.db.Select(&opportunityTracking, query)
	if err != nil {
		fmt.Printf("Error ejecutando la consulta: %v\n", err)
		return nil, err
	}

	// Iterar sobre los seguimiento de oportunidades
	for i := range opportunityTracking {
		// Obtener cliente de la orden de venta
		var customer domain.Customers
		err = r.db.Get(&customer, `
			SELECT id, first_name, second_name, third_name, surname, second_surname, company_name, document_type, document_number, email, address, phone, status
			FROM customers
			WHERE id = $1 AND deleted_at IS NULL
		`, opportunityTracking[i].Customer_Id)
		if err != nil && err != sql.ErrNoRows { // Manejar casos donde no hay cliente asociado
			fmt.Printf("Error obteniendo el cliente para la orden de venta %d: %v\n", opportunityTracking[i].Id, err)
			return nil, err
		}
		opportunityTracking[i].Customer = &customer

		// Obtener usuario de la orden de venta
		var user domain.Users
		err = r.db.Get(&user, `
			SELECT id, role_id, username, employee_id, avatar, email, status FROM users WHERE id = $1 AND deleted_at IS NULL
		`, opportunityTracking[i].User_Id)
		if err != nil && err != sql.ErrNoRows { // Manejar casos donde no hay usuario asociado
			fmt.Printf("Error obteniendo el usuario para la orden de venta %d: %v\n", opportunityTracking[i].Id, err)
			return nil, err
		}
		opportunityTracking[i].User = &user

		// Obtener empleado del usuario (si existe)
		if user.Employee_Id.Valid {
			var company domain.Employees
			err = r.db.Get(&company, `
				SELECT id, first_name, second_name, third_name, surname, second_surname, photo, warehouse_id, document_type, document_number, birth_date, gender, email, phone, address, hire_date, position, salary, status
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
	return opportunityTracking, nil
}

// Obtener un empleado por Id
func (r *OpportunityTrackingRepository) GetOpportunityTrackingById(opportunityTrackingId int) (*domain.OpportunityTracking, error) {
	var opportunityTracking domain.OpportunityTracking
	query := `
		SELECT 
			id, customer_id, user_id, title, description, status, expected_revenue, probability
		FROM opportunity_tracking
		WHERE id = $1 AND deleted_at IS NULL
	`
	err := r.db.Get(&opportunityTracking, query, opportunityTrackingId)
	if err != nil {
		fmt.Printf("Error obteniendo empleado con ID %d: %v\n", opportunityTrackingId, err)
		return nil, err
	}

	// Obtener cliente de la orden de venta
	var customer domain.Customers
	err = r.db.Get(&customer, `
		SELECT c.id, c.first_name, c.second_name, c.third_name, c.surname, c.second_surname, c.company_name, c.document_type, c.document_number, c.email, c.address, c.phone, c.status
		FROM customers c
		WHERE c.id = $1 AND c.deleted_at IS NULL
	`, opportunityTracking.Customer_Id)
	if err != nil && err != sql.ErrNoRows { // Manejar casos donde no hay cliente asociado
		fmt.Printf("Error obteniendo cliente para la orden de venta %d: %v\n", opportunityTrackingId, err)
		return nil, err
	}
	opportunityTracking.Customer = &customer

	// Obtener usuario de la orden de venta
	var user domain.Users
	err = r.db.Get(&user, `
		SELECT id, role_id, username, employee_id, avatar, email, status FROM users WHERE id = $1 AND deleted_at IS NULL
	`, opportunityTracking.User_Id)
	if err != nil && err != sql.ErrNoRows { // Manejar casos donde no hay usuario asociado
		fmt.Printf("Error obteniendo usuario para la orden de venta %d: %v\n", opportunityTrackingId, err)
		return nil, err
	}
	opportunityTracking.User = &user

	return &opportunityTracking, nil
}

// Actualizar un empleado
func (r *OpportunityTrackingRepository) UpdateOpportunityTracking(opportunityTracking *domain.OpportunityTracking) error {
	query := `
		UPDATE opportunity_tracking SET
			customer_id = $1, user_id = $2, title = $3, description = $4, status = $5, expected_revenue = $6, probability = $7, updated_at = NOW()
		WHERE id = $8
	`
	_, err := r.db.Exec(query, opportunityTracking.Customer_Id, opportunityTracking.User_Id, opportunityTracking.Title, opportunityTracking.Description, opportunityTracking.Status, opportunityTracking.Expected_Revenue, opportunityTracking.Probability, opportunityTracking.Id)
	return err
}

// Eliminar un empleado por Id (eliminación lógica)
func (r *OpportunityTrackingRepository) DeleteOpportunityTrackingById(id int) error {
	query := `UPDATE opportunity_tracking SET deleted_at = NOW() WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

// Eliminar múltiples seguimiento de oportunidades por Ids (eliminación lógica)
func (r *OpportunityTrackingRepository) DeleteOpportunityTrackingByIds(ids []int) error {
	query := `UPDATE opportunity_tracking SET deleted_at = NOW() WHERE id = ANY($1)`
	_, err := r.db.Exec(query, pq.Array(ids))
	return err
}
