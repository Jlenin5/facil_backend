package repository

import (
	"database/sql"
	"fmt"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/lib/pq"

	"github.com/jmoiron/sqlx"
)

type EmployeeRepository struct {
	db *sqlx.DB
}

func NewEmployeeRepository(db *sqlx.DB) *EmployeeRepository {
	return &EmployeeRepository{db: db}
}

// Crear un empleado
func (r *EmployeeRepository) CreateEmployee(employee *domain.Employees) error {
	query := `
		INSERT INTO employees (
			first_name, second_name, third_name, surname, second_surname, photo, warehouse_id, document_type, document_number, birth_date, gender, email, phone, address, hire_date, job_position_id, salary, status
		) VALUES (
		 	:first_name, :second_name, :third_name, :surname, :second_surname, :photo, :warehouse_id, :document_type, :document_number, :birth_date, :gender, :email, :phone, :address, :hire_date, :job_position_id, :salary, :status
		)
	`
	_, err := r.db.NamedExec(query, employee)
	if err != nil {
		fmt.Printf("Error creando el empleado: %v\n", err)
		return err
	}
	return nil
}

// Obtener todos los empleados
func (r *EmployeeRepository) GetAllEmployees() ([]domain.Employees, error) {
	var employees []domain.Employees
	query := `
		SELECT 
			e.id, e.names, e.surname, e.second_surname, e.photo, 
			e.warehouse_id, e.document_type, e.document_number, e.birth_date, e.gender, 
			e.email, e.phone, e.address, e.hire_date, e.job_position_id, e.salary, e.status,
			w.id AS "warehouse.id", w.name AS "warehouse.name", w.address AS "warehouse.address", w.status AS "warehouse.status", w.branch_office_id AS "warehouse.branch_office_id",
			b.id AS "warehouse.branch_office.id", b.name AS "warehouse.branch_office.name", b.address AS "warehouse.branch_office.address", b.phone AS "warehouse.branch_office.phone", b.company_id AS "warehouse.branch_office.company_id",
			c.id AS "warehouse.branch_office.company.id", c.name AS "warehouse.branch_office.company.name", c.ruc AS "warehouse.branch_office.company.ruc", c.phone AS "warehouse.branch_office.company.phone", c.address AS "warehouse.branch_office.company.address",
			jp.id AS "job_position.id", jp.work_area_id AS "job_position.work_area_id", jp.name AS "job_position.name",
			wa.id AS "job_position.work_area.id", wa.name AS "job_position.work_area.name"
		FROM employees e
		LEFT JOIN warehouses w ON e.warehouse_id = w.id AND w.deleted_at IS NULL
		LEFT JOIN branch_offices b ON w.branch_office_id = b.id AND b.deleted_at IS NULL
		LEFT JOIN companies c ON b.company_id = c.id AND c.deleted_at IS NULL
		LEFT JOIN job_positions jp ON e.job_position_id = jp.id AND jp.deleted_at IS NULL
		LEFT JOIN work_areas wa ON jp.work_area_id = wa.id AND wa.deleted_at IS NULL
		WHERE e.deleted_at IS NULL
		ORDER BY e.id DESC
	`
	err := r.db.Select(&employees, query)
	if err != nil {
		fmt.Printf("Error ejecutando la consulta: %v\n", err)
		return nil, err
	}

	// Iterar sobre los empleados
	for i := range employees {
		// Obtener almacén del empleadoo
		var warehouse domain.Warehouses
		err = r.db.Get(&warehouse, `
			SELECT id, name, address, status FROM warehouses WHERE id = $1 AND deleted_at IS NULL
		`, employees[i].Warehouse_Id)
		if err != nil && err != sql.ErrNoRows { // Manejar casos donde no hay almacén asociado
			fmt.Printf("Error obteniendo el almacén para la orden de venta %d: %v\n", employees[i].Id, err)
			return nil, err
		}
		employees[i].Warehouse = &warehouse
	}
	return employees, nil
}

// Obtener un empleado por Id
func (r *EmployeeRepository) GetEmployeeById(employeeId int) (*domain.Employees, error) {
	var employee domain.Employees

	query := `
		SELECT 
			e.id, e.names, e.surname, e.second_surname, e.photo, 
			e.warehouse_id, e.document_type, e.document_number, e.birth_date, e.gender, 
			e.email, e.phone, e.address, e.hire_date, e.job_position_id, e.salary, e.status,
			w.id AS "warehouse.id", w.name AS "warehouse.name", w.address AS "warehouse.address", w.status AS "warehouse.status", w.branch_office_id AS "warehouse.branch_office_id",
			b.id AS "warehouse.branch_office.id", b.name AS "warehouse.branch_office.name", b.address AS "warehouse.branch_office.address", b.phone AS "warehouse.branch_office.phone", b.company_id AS "warehouse.branch_office.company_id",
			c.id AS "warehouse.branch_office.company.id", c.name AS "warehouse.branch_office.company.name", c.ruc AS "warehouse.branch_office.company.ruc", c.phone AS "warehouse.branch_office.company.phone", c.address AS "warehouse.branch_office.company.address",
			jp.id AS "job_position.id", jp.work_area_id AS "job_position.work_area_id", jp.name AS "job_position.name",
			wa.id AS "job_position.work_area.id", wa.name AS "job_position.work_area.name"
		FROM employees e
		LEFT JOIN warehouses w ON e.warehouse_id = w.id AND w.deleted_at IS NULL
		LEFT JOIN branch_offices b ON w.branch_office_id = b.id AND b.deleted_at IS NULL
		LEFT JOIN companies c ON b.company_id = c.id AND c.deleted_at IS NULL
		LEFT JOIN job_positions jp ON e.job_position_id = jp.id AND jp.deleted_at IS NULL
		LEFT JOIN work_areas wa ON jp.work_area_id = wa.id AND wa.deleted_at IS NULL
		WHERE e.id = $1 AND e.deleted_at IS NULL
	`

	err := r.db.Get(&employee, query, employeeId)
	if err != nil {
		fmt.Printf("Error obteniendo empleado con ID %d: %v\n", employeeId, err)
		return nil, err
	}

	return &employee, nil
}

// Actualizar un empleado
func (r *EmployeeRepository) UpdateEmployee(employee *domain.Employees) error {
	query := `
		UPDATE employees SET
			first_name = $first_name, second_name = $second_name, third_name = $third_name,
			surname = $surname, second_surname = $second_surname, photo = $photo,
			warehouse_id = $warehouse_id, document_type = $document_type,
			document_number = $document_number, birth_date = $birth_date, gender = $gender,
			email = $email, phone = $phone, address = $address, hire_date = $hire_date, job_position_id = $job_position_id,
			position = $position, salary = $salary, status = $status, updated_at = NOW()
		WHERE id = :id
	`
	_, err := r.db.NamedExec(query, employee)
	return err
}

// Eliminar un empleado por Id (eliminación lógica)
func (r *EmployeeRepository) DeleteEmployeeById(id int) error {
	query := `UPDATE employees SET deleted_at = NOW() WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}

// Eliminar múltiples empleados por Ids (eliminación lógica)
func (r *EmployeeRepository) DeleteEmployeesByIds(ids []int) error {
	query := `UPDATE employees SET deleted_at = NOW() WHERE id = ANY($1)`
	_, err := r.db.Exec(query, pq.Array(ids))
	return err
}
