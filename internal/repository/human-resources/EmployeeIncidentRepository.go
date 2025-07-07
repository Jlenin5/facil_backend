package repositoryHumanresources

import (
	"context"
	"fmt"

	"github.com/Jlenin5/facil_backend/internal/domain"
	humanresources "github.com/Jlenin5/facil_backend/internal/domain/human-resources"
	"github.com/jmoiron/sqlx"
)

type EmployeeIncidentRepository struct {
	db *sqlx.DB
}

func NewEmployeeIncidentRepository(db *sqlx.DB) *EmployeeIncidentRepository {
	return &EmployeeIncidentRepository{db: db}
}

const (
	employeeIncidentsTable = "employee_incidents"
	employeesTable4        = "employees"
)

func (r *EmployeeIncidentRepository) Create(ctx context.Context, incident *humanresources.EmployeeIncident) error {
	query := fmt.Sprintf(`
		INSERT INTO %s (
			employee_id, incident_type, incident_date, description,
			severity, action_taken, reported_by, status
		) VALUES (
			:employee_id, :incident_type, :incident_date, :description,
			:severity, :action_taken, :reported_by, :status
		) RETURNING id, created_at, updated_at`, employeeIncidentsTable)

	rows, err := r.db.NamedQueryContext(ctx, query, incident)
	if err != nil {
		return fmt.Errorf("error creating employee incident: %w", err)
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.Scan(
			&incident.Id,
			&incident.Created_at,
			&incident.Updated_at,
		); err != nil {
			return fmt.Errorf("error getting created incident data: %w", err)
		}
	}

	return nil
}

func (r *EmployeeIncidentRepository) GetAll(ctx context.Context) ([]humanresources.EmployeeIncident, error) {
	query := fmt.Sprintf(`
		SELECT
			ei.id, ei.employee_id, ei.incident_type, ei.incident_date,
			ei.description, ei.severity, ei.action_taken, ei.reported_by,
			ei.resolved_at, ei.status,
			e.id AS "employee.id",
			e.names AS "employee.names",
			e.surname AS "employee.surname",
			e.second_surname AS "employee.second_surname",
			e.warehouse_id AS "employee.warehouse_id"
		FROM %s ei
		JOIN %s e ON ei.employee_id = e.id
		WHERE ei.deleted_at IS NULL
		ORDER BY ei.incident_date DESC`, employeeIncidentsTable, employeesTable4)

	var results []struct {
		humanresources.EmployeeIncident
		Employee domain.EmployeeReducedData `db:"employee"`
	}

	if err := r.db.SelectContext(ctx, &results, query); err != nil {
		return nil, fmt.Errorf("error getting employee incidents: %w", err)
	}

	incidents := make([]humanresources.EmployeeIncident, len(results))
	for i, res := range results {
		res.EmployeeIncident.Employee = res.Employee
		incidents[i] = res.EmployeeIncident
	}

	return incidents, nil
}

func (r *EmployeeIncidentRepository) GetById(ctx context.Context, id int) (*humanresources.EmployeeIncident, error) {
	query := fmt.Sprintf(`
		SELECT
			ei.*,
			e.id AS "employee.id",
			e.names AS "employee.names",
			e.surname AS "employee.surname",
			e.second_surname AS "employee.second_surname",
			e.warehouse_id AS "employee.warehouse_id"
		FROM %s ei
		JOIN %s e ON ei.employee_id = e.id
		WHERE ei.id = $1 AND ei.deleted_at IS NULL`, employeeIncidentsTable, employeesTable4)

	var result struct {
		humanresources.EmployeeIncident
		Employee domain.EmployeeReducedData `db:"employee"`
	}

	if err := r.db.GetContext(ctx, &result, query, id); err != nil {
		return nil, fmt.Errorf("error getting employee incident by ID: %w", err)
	}

	incident := result.EmployeeIncident
	incident.Employee = result.Employee

	return &incident, nil
}

func (r *EmployeeIncidentRepository) Update(ctx context.Context, incident *humanresources.EmployeeIncident) error {
	query := fmt.Sprintf(`
		UPDATE %s SET
			incident_type = :incident_type,
			description = :description,
			severity = :severity,
			action_taken = :action_taken,
			status = :status,
			resolved_at = :resolved_at,
			updated_at = NOW()
		WHERE id = :id AND deleted_at IS NULL`, employeeIncidentsTable)

	result, err := r.db.NamedExecContext(ctx, query, incident)
	if err != nil {
		return fmt.Errorf("error updating employee incident: %w", err)
	}

	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return fmt.Errorf("no rows affected, record may not exist")
	}

	return nil
}

func (r *EmployeeIncidentRepository) GetByStatus(ctx context.Context, status string) ([]humanresources.EmployeeIncident, error) {
	query := fmt.Sprintf(`
		SELECT
			ei.*,
			e.id AS "employee.id",
			e.names AS "employee.names",
			e.surname AS "employee.surname",
			e.second_surname AS "employee.second_surname",
			e.warehouse_id AS "employee.warehouse_id"
		FROM %s ei
		JOIN %s e ON ei.employee_id = e.id
		WHERE ei.status = $1 AND ei.deleted_at IS NULL
		ORDER BY ei.incident_date DESC`, employeeIncidentsTable, employeesTable4)

	var results []struct {
		humanresources.EmployeeIncident
		Employee domain.EmployeeReducedData `db:"employee"`
	}

	if err := r.db.SelectContext(ctx, &results, query, status); err != nil {
		return nil, fmt.Errorf("error getting employee incidents by status: %w", err)
	}

	incidents := make([]humanresources.EmployeeIncident, len(results))
	for i, res := range results {
		res.EmployeeIncident.Employee = res.Employee
		incidents[i] = res.EmployeeIncident
	}

	return incidents, nil
}

func (r *EmployeeIncidentRepository) DeleteById(ctx context.Context, id int) error {
	query := fmt.Sprintf(`
		UPDATE %s 
		SET deleted_at = NOW(), updated_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL`, employeeIncidentsTable)

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("error deleting employee incident: %w", err)
	}

	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return fmt.Errorf("no rows affected, record may not exist")
	}

	return nil
}

func (r *EmployeeIncidentRepository) Resolve(ctx context.Context, id int, actionTaken string) error {
	query := fmt.Sprintf(`
		UPDATE %s SET
			status = 'resolved',
			action_taken = $1,
			resolved_at = NOW(),
			updated_at = NOW()
		WHERE id = $2 AND deleted_at IS NULL`, employeeIncidentsTable)

	result, err := r.db.ExecContext(ctx, query, actionTaken, id)
	if err != nil {
		return fmt.Errorf("error resolving employee incident: %w", err)
	}

	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return fmt.Errorf("no rows affected, record may not exist")
	}

	return nil
}

func (r *EmployeeIncidentRepository) Close(ctx context.Context, id int) error {
	query := fmt.Sprintf(`
		UPDATE %s SET
			status = 'closed',
			updated_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL`, employeeIncidentsTable)

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("error closing employee incident: %w", err)
	}

	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return fmt.Errorf("no rows affected, record may not exist")
	}

	return nil
}
