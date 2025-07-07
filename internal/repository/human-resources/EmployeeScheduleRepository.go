package repositoryHumanresources

import (
	"context"
	"fmt"
	"time"

	"github.com/Jlenin5/facil_backend/internal/domain"
	humanresources "github.com/Jlenin5/facil_backend/internal/domain/human-resources"
	"github.com/jmoiron/sqlx"
)

type EmployeeScheduleRepository struct {
	db *sqlx.DB
}

func NewEmployeeScheduleRepository(db *sqlx.DB) *EmployeeScheduleRepository {
	return &EmployeeScheduleRepository{db: db}
}

const (
	employeeSchedulesTable = "employee_schedules"
	workSchedulesTable2    = "work_schedules"
	employeesTable5        = "employees"
)

func (r *EmployeeScheduleRepository) Create(ctx context.Context, incident *humanresources.EmployeeSchedule) error {
	query := fmt.Sprintf(`
		INSERT INTO %s (
			employee_id, schedule_id, schedule_id, end_date
		) VALUES (
			:employee_id, :schedule_id, :schedule_id, :end_date
		) RETURNING id, created_at, updated_at`, employeeSchedulesTable)

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

func (r *EmployeeScheduleRepository) GetAll(ctx context.Context) ([]humanresources.EmployeeSchedule, error) {
	query := fmt.Sprintf(`
		SELECT
			es.id, es.employee_id, es.schedule_id, 
			es.effective_date, es.end_date,
			ws.id AS "work_schedule.id", 
			ws.name AS "work_schedule.name",
			e.id AS "employee.id", 
			e.names AS "employee.names",
			e.surname AS "employee.surname", 
			e.second_surname AS "employee.second_surname", 
			e.warehouse_id AS "employee.warehouse_id"
		FROM %s es
		JOIN %s ws ON es.schedule_id = ws.id
		JOIN %s e ON es.employee_id = e.id
		WHERE es.deleted_at IS NULL
		ORDER BY es.effective_date DESC`,
		employeeSchedulesTable, workSchedulesTable2, employeesTable)

	var results []struct {
		humanresources.EmployeeSchedule
		Schedule domain.ScheduleReducedData `db:"work_schedule"`
		Employee domain.EmployeeReducedData `db:"employee"`
	}

	if err := r.db.SelectContext(ctx, &results, query); err != nil {
		return nil, fmt.Errorf("error getting employee schedules: %w", err)
	}

	schedules := make([]humanresources.EmployeeSchedule, len(results))
	for i, res := range results {
		res.EmployeeSchedule.Schedule = res.Schedule
		res.EmployeeSchedule.Employee = res.Employee
		schedules[i] = res.EmployeeSchedule
	}

	return schedules, nil
}

func (r *EmployeeScheduleRepository) Assign(ctx context.Context, schedule *humanresources.EmployeeSchedule) error {
	// Finalizar el horario actual si existe
	endCurrentQuery := fmt.Sprintf(`
		UPDATE %s 
		SET end_date = $1, updated_at = NOW() 
		WHERE employee_id = $2 AND end_date IS NULL
		AND deleted_at IS NULL`, employeeSchedulesTable)

	result, err := r.db.ExecContext(ctx, endCurrentQuery,
		schedule.EffectiveDate.AddDate(0, 0, -1), schedule.EmployeeId)
	if err != nil {
		return fmt.Errorf("error ending current schedule: %w", err)
	}

	// Verificar si se actualizó algún registro
	if rowsAffected, _ := result.RowsAffected(); rowsAffected > 0 {
		// Esperar un breve momento para evitar condiciones de carrera
		time.Sleep(100 * time.Millisecond)
	}

	// Asignar nuevo horario
	assignQuery := fmt.Sprintf(`
		INSERT INTO %s (
			employee_id, schedule_id, effective_date
		) VALUES (
			:employee_id, :schedule_id, :effective_date
		) RETURNING id, created_at, updated_at`, employeeSchedulesTable)

	rows, err := r.db.NamedQueryContext(ctx, assignQuery, schedule)
	if err != nil {
		return fmt.Errorf("error assigning new schedule: %w", err)
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.Scan(
			&schedule.Id,
			&schedule.Created_at,
			&schedule.Updated_at,
		); err != nil {
			return fmt.Errorf("error getting assigned schedule data: %w", err)
		}
	}

	return nil
}

func (r *EmployeeScheduleRepository) GetById(ctx context.Context, id int) (*humanresources.EmployeeSchedule, error) {
	query := fmt.Sprintf(`
		SELECT
			ei.*,
			e.id AS "employee.id",
			e.names AS "employee.names"
			e.third_name AS "employee.third_name",
			e.surname AS "employee.surname",
			e.second_surname AS "employee.second_surname",
			e.warehouse_id AS "employee.warehouse_id"
		FROM %s ei
		JOIN %s e ON ei.employee_id = e.id
		WHERE ei.id = $1 AND ei.deleted_at IS NULL`, employeeSchedulesTable, employeesTable5)

	var result struct {
		humanresources.EmployeeSchedule
		Employee domain.EmployeeReducedData `db:"employee"`
	}

	if err := r.db.GetContext(ctx, &result, query, id); err != nil {
		return nil, fmt.Errorf("error getting employee schedule by ID: %w", err)
	}

	schedule := result.EmployeeSchedule
	schedule.Employee = result.Employee

	return &schedule, nil
}

func (r *EmployeeScheduleRepository) Update(ctx context.Context, incident *humanresources.EmployeeSchedule) error {
	query := fmt.Sprintf(`
		UPDATE %s SET
			schedule_id = :schedule_id,
			effective_date = :effective_date,
			end_date = :end_date,
			updated_at = NOW()
		WHERE id = :id AND deleted_at IS NULL`, employeeSchedulesTable)

	result, err := r.db.NamedExecContext(ctx, query, incident)
	if err != nil {
		return fmt.Errorf("error updating employee incident: %w", err)
	}

	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return fmt.Errorf("no rows affected, record may not exist")
	}

	return nil
}

func (r *EmployeeScheduleRepository) DeleteById(ctx context.Context, id int) error {
	query := fmt.Sprintf(`
		UPDATE %s 
		SET deleted_at = NOW(), updated_at = NOW() 
		WHERE id = $1 AND deleted_at IS NULL`, employeeSchedulesTable)

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("error deleting employee schedule: %w", err)
	}

	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return fmt.Errorf("no rows affected, record may not exist")
	}

	return nil
}

func (r *EmployeeScheduleRepository) GetCurrentByEmployeeID(ctx context.Context, employeeID int) (*humanresources.EmployeeSchedule, error) {
	query := fmt.Sprintf(`
		SELECT
			es.*,
			ws.id AS "work_schedule.id",
			ws.name AS "work_schedule.name",
			e.id AS "employee.id",
			e.names AS "employee.names"
			e.third_name AS "employee.third_name",
			e.surname AS "employee.surname",
			e.second_surname AS "employee.second_surname",
			e.warehouse_id AS "employee.warehouse_id"
		FROM %s es
		JOIN %s ws ON es.schedule_id = ws.id
		JOIN %s e ON es.employee_id = e.id
		WHERE es.employee_id = $1 
		AND (es.end_date IS NULL OR es.end_date >= CURRENT_DATE)
		AND es.deleted_at IS NULL
		ORDER BY es.effective_date DESC
		LIMIT 1`,
		employeeSchedulesTable, workSchedulesTable2, employeesTable5)

	var result struct {
		humanresources.EmployeeSchedule
		Schedule domain.ScheduleReducedData `db:"work_schedule"`
		Employee domain.EmployeeReducedData `db:"employee"`
	}

	if err := r.db.GetContext(ctx, &result, query, employeeID); err != nil {
		return nil, fmt.Errorf("error getting current employee schedule: %w", err)
	}

	schedule := result.EmployeeSchedule
	schedule.Schedule = result.Schedule
	schedule.Employee = result.Employee

	return &schedule, nil
}

func (r *EmployeeScheduleRepository) GetHistoryByEmployeeID(ctx context.Context, employeeID int) ([]humanresources.EmployeeSchedule, error) {
	query := fmt.Sprintf(`
		SELECT
			es.*,
			ws.id AS "work_schedule.id",
			ws.name AS "work_schedule.name",
			e.id AS "employee.id",
			e.names AS "employee.names"
			e.third_name AS "employee.third_name",
			e.surname AS "employee.surname",
			e.second_surname AS "employee.second_surname",
			e.warehouse_id AS "employee.warehouse_id"
		FROM %s es
		JOIN %s ws ON es.schedule_id = ws.id
		JOIN %s e ON es.employee_id = e.id
		WHERE es.employee_id = $1 AND es.deleted_at IS NULL
		ORDER BY es.effective_date DESC`,
		employeeSchedulesTable, workSchedulesTable2, employeesTable5)

	var results []struct {
		humanresources.EmployeeSchedule
		Schedule domain.ScheduleReducedData `db:"work_schedule"`
		Employee domain.EmployeeReducedData `db:"employee"`
	}

	if err := r.db.SelectContext(ctx, &results, query, employeeID); err != nil {
		return nil, fmt.Errorf("error getting employee schedule history: %w", err)
	}

	schedules := make([]humanresources.EmployeeSchedule, len(results))
	for i, res := range results {
		res.EmployeeSchedule.Schedule = res.Schedule
		res.EmployeeSchedule.Employee = res.Employee
		schedules[i] = res.EmployeeSchedule
	}

	return schedules, nil
}

func (r *EmployeeScheduleRepository) EndCurrentSchedule(ctx context.Context, employeeID int, endDate time.Time) error {
	query := fmt.Sprintf(`
		UPDATE %s 
		SET end_date = $1, updated_at = NOW() 
		WHERE employee_id = $2 AND end_date IS NULL
		AND deleted_at IS NULL`, employeeSchedulesTable)

	result, err := r.db.ExecContext(ctx, query, endDate, employeeID)
	if err != nil {
		return fmt.Errorf("error ending current schedule: %w", err)
	}

	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return fmt.Errorf("no rows affected, record may not exist")
	}

	return nil
}
