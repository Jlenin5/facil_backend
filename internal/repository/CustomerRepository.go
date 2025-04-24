package repository

import (
	"fmt"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

type CustomerRepository struct {
	db *sqlx.DB
}

func NewCustomerRepository(db *sqlx.DB) *CustomerRepository {
	return &CustomerRepository{db: db}
}

// Crear un cliente
func (r *CustomerRepository) CreateCustomer(customer *domain.Customers) error {
	query := `
		INSERT INTO customers (
			first_name, second_name, third_name, surname, second_surname, company_name, document_type, document_number, email, address, phone, status
		) VALUES (
		 	:first_name, :second_name, :third_name, :surname, :second_surname, :company_name, :document_type, :document_number, :email, :address, :phone, :status
		)
	`
	_, err := r.db.NamedExec(query, customer)
	if err != nil {
		fmt.Printf("Error creando el cliente: %v\n", err)
		return err
	}
	return nil
}

// Obtener todos los clientes
func (r *CustomerRepository) GetAllCustomers() ([]domain.Customers, error) {
	var customers []domain.Customers
	query := `
		SELECT
			id, first_name, second_name, third_name, surname, second_surname, company_name, document_type, document_number, email, address, phone, status
		FROM customers
		WHERE deleted_at IS NULL
		ORDER BY id DESC
	`
	err := r.db.Select(&customers, query)
	if err != nil {
		fmt.Printf("Error ejecutando la consulta: %v\n", err)
		return nil, err
	}
	return customers, nil
}

// Obtener un cliente por Id
func (r *CustomerRepository) GetCustomerById(customerId int) (*domain.Customers, error) {
	var customer domain.Customers
	query := `
		SELECT
			id, first_name, second_name, third_name, surname, second_surname, company_name, document_type, document_number, email, address, phone, status
		FROM customers
		WHERE id = $1 AND deleted_at IS NULL
	`
	err := r.db.Get(&customer, query, customerId)
	if err != nil {
		fmt.Printf("Error obteniendo cliente con ID %d: %v\n", customerId, err)
		return nil, err
	}
	return &customer, nil
}

// Actualizar un cliente
func (r *CustomerRepository) UpdateCustomer(customer *domain.Customers) error {
	query := `
		UPDATE customers SET
			first_name = $1, second_name = $2, third_name = $3, surname = $4, second_surname = $5, company_name = $6, document_type = $7, document_number = $8, email = $9, address = $10, phone = $11, status = $12, updated_at = NOW()
		WHERE id = $13
	`
	_, err := r.db.Exec(query, customer.First_Name, customer.Second_Name, customer.Third_Name,
		customer.Surname, customer.Second_Surname, customer.Company_Name, customer.Document_Type,
		customer.Document_Number, customer.Email, customer.Address, customer.Phone, customer.Status, customer.Id)
	return err
}

// Eliminar un cliente por Id (eliminación lógica)
func (r *CustomerRepository) DeleteCustomerById(id int) error {
	query := `UPDATE customers SET deleted_at = NOW() WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

// Eliminar múltiples clientes por Ids (eliminación lógica)
func (r *CustomerRepository) DeleteCustomersByIds(ids []int) error {
	query := `UPDATE customers SET deleted_at = NOW() WHERE id = ANY($1)`
	_, err := r.db.Exec(query, pq.Array(ids))
	return err
}
