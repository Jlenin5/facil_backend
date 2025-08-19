package repositoryHumanresources

import (
	"context"
	"fmt"
	"time"

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
			employee_id, incident_type, incident_date, observation, discount, total_to_pay, reported_by, status
		) VALUES (
			:employee_id, :incident_type, :incident_date, :observation, :discount, :total_to_pay, :reported_by, :status
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

	// Obtener las fechas de inicio y fin de la semana actual
	now := time.Now()
	weekday := now.Weekday()

	// Ajustar al lunes de esta semana
	var startOfWeek time.Time
	if weekday == time.Sunday {
		// Si es domingo, el lunes fue hace 6 días
		startOfWeek = now.AddDate(0, 0, -6)
	} else {
		// Restar los días desde el lunes
		startOfWeek = now.AddDate(0, 0, -int(weekday-time.Monday))
	}

	// El domingo de esta semana es 6 días después del lunes
	endOfWeek := startOfWeek.AddDate(0, 0, 6)
	
	// Formatear fechas para SQL
	startDate := startOfWeek.Format("2006-01-02")
	endDate := endOfWeek.Format("2006-01-02")

	query := fmt.Sprintf(`
		SELECT
			ei.id, ei.employee_id, ei.incident_type, ei.incident_date, ei.observation, ei.discount, ei.total_to_pay, ei.reported_by, ei.resolved_at, ei.status,
			e.id AS "employee.id",
			e.names AS "employee.names",
			COALESCE(e.surname, '') AS "employee.surname",
			COALESCE(e.second_surname, '') AS "employee.second_surname"
		FROM %s ei
		INNER JOIN %s e ON ei.employee_id = e.id
		WHERE ei.deleted_at IS NULL
		AND ei.incident_date BETWEEN '%s 00:00:00' AND '%s 23:59:59'
		ORDER BY ei.incident_date DESC`,
		employeeIncidentsTable, employeesTable4, startDate, endDate)

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
			ei.id, ei.employee_id, ei.incident_type, ei.incident_date, ei.observation, ei.discount, ei.total_to_pay, ei.reported_by, ei.resolved_at, ei.status,
			e.id AS "employee.id",
			e.names AS "employee.names",
			COALESCE(e.surname, '') AS "employee.surname",
			COALESCE(e.second_surname, '') AS "employee.second_surname"
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
			observation = :observation,
			discount = :discount,
			total_to_pay = :total_to_pay,
			resolved_at = :resolved_at,
			status = :status,
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
			COALESCE(e.surname, '') AS "employee.surname",
			COALESCE(e.second_surname, '') AS "employee.second_surname"
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