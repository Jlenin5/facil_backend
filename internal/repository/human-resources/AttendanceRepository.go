package repositoryHumanresources

import (
	"context"
	"fmt"
	"time"

	"github.com/Jlenin5/facil_backend/internal/domain"
	humanresources "github.com/Jlenin5/facil_backend/internal/domain/human-resources"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type AttendanceRepository struct {
	db *sqlx.DB
}

func NewAttendanceRepository(db *sqlx.DB) *AttendanceRepository {
	return &AttendanceRepository{db: db}
}

const (
	attendancesTable     = "attendances"
	employeesTable       = "employees"
	attendanceTypesTable = "attendance_types"
)

func (r *AttendanceRepository) Create(ctx context.Context, attendance *humanresources.Attendance) error {
	query := fmt.Sprintf(`
		INSERT INTO %s (
			employee_id, attendance_type_id, date, check_in, check_out, 
			worked_hours, late_minutes, early_departure_minutes, notes, status
		) VALUES (
			:employee_id, :attendance_type_id, :date, :check_in, :check_out,
			:worked_hours, :late_minutes, :early_departure_minutes, :notes, :status
		) RETURNING id, created_at, updated_at`, attendancesTable)

	rows, err := r.db.NamedQueryContext(ctx, query, attendance)
	if err != nil {
		return fmt.Errorf("error creating attendance: %w", err)
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.Scan(
			&attendance.Id,
			&attendance.Created_at,
			&attendance.Updated_at,
		); err != nil {
			return fmt.Errorf("error getting created attendance data: %w", err)
		}
	}

	return nil
}

func (r *AttendanceRepository) GetAll(ctx context.Context) ([]humanresources.Attendance, error) {
	query := fmt.Sprintf(`
		SELECT
			a.id, a.employee_id, a.attendance_type_id, 
			a.date, a.check_in, a.check_out, 
			a.worked_hours, a.late_minutes, a.early_departure_minutes, 
			a.notes, a.status, a.approved_by, a.approved_at,
			at.id AS "attendance_type.id", 
			at.name AS "attendance_type.name", 
			at.code AS "attendance_type.code",
			e.id AS "employee.id", 
			e.names AS "employee.names", 
			e.surname AS "employee.surname", 
			e.second_surname AS "employee.second_surname", 
			e.warehouse_id AS "employee.warehouse_id"
		FROM %s a
		JOIN %s at ON a.attendance_type_id = at.id
		JOIN %s e ON a.employee_id = e.id
		WHERE a.deleted_at IS NULL
		ORDER BY a.date DESC, a.check_in DESC`,
		attendancesTable, attendanceTypesTable, employeesTable)

	var results []struct {
		humanresources.Attendance
		Employee       domain.EmployeeReducedData       `db:"employee"`
		AttendanceType domain.AttendanceTypeReducedData `db:"attendance_type"`
	}

	if err := r.db.SelectContext(ctx, &results, query); err != nil {
		return nil, fmt.Errorf("error getting attendances: %w", err)
	}

	attendances := make([]humanresources.Attendance, len(results))
	for i, res := range results {
		res.Attendance.Employee = res.Employee
		res.Attendance.AttendanceType = res.AttendanceType
		attendances[i] = res.Attendance
	}

	return attendances, nil
}

func (r *AttendanceRepository) GetById(ctx context.Context, id int) (*humanresources.Attendance, error) {
	query := fmt.Sprintf(`
		SELECT
			a.*,
			e.id AS "employee.id",
			e.names AS "employee.names
			e.surname AS "employee.surname",
			e.second_surname AS "employee.second_surname",
			e.warehouse_id AS "employee.warehouse_id",
			at.id AS "attendance_type.id",
			at.name AS "attendance_type.name",
			at.code AS "attendance_type.code"
		FROM %s a
		JOIN %s e ON a.employee_id = e.id
		JOIN %s at ON a.attendance_type_id = at.id
		WHERE a.id = $1 AND a.deleted_at IS NULL`, attendancesTable, employeesTable, attendanceTypesTable)

	var result struct {
		humanresources.Attendance
		Employee       domain.EmployeeReducedData       `db:"employee"`
		AttendanceType domain.AttendanceTypeReducedData `db:"attendance_type"`
	}

	if err := r.db.GetContext(ctx, &result, query, id); err != nil {
		return nil, fmt.Errorf("error getting attendance by ID: %w", err)
	}

	attendance := result.Attendance
	attendance.Employee = result.Employee
	attendance.AttendanceType = result.AttendanceType

	return &attendance, nil
}

func (r *AttendanceRepository) Update(ctx context.Context, attendance *humanresources.Attendance) error {
	query := fmt.Sprintf(`
		UPDATE %s SET
			employee_id = :employee_id,
			attendance_type_id = :attendance_type_id,
			date = :date,
			check_in = :check_in,
			check_out = :check_out,
			worked_hours = :worked_hours,
			late_minutes = :late_minutes,
			early_departure_minutes = :early_departure_minutes,
			notes = :notes,
			status = :status,
			approved_by = :approved_by,
			approved_at = :approved_at,
			updated_at = NOW()
		WHERE id = :id AND deleted_at IS NULL`, attendancesTable)

	result, err := r.db.NamedExecContext(ctx, query, attendance)
	if err != nil {
		return fmt.Errorf("error updating attendance: %w", err)
	}

	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return fmt.Errorf("no rows affected, record may not exist")
	}

	return nil
}

func (r *AttendanceRepository) DeleteById(ctx context.Context, id int) error {
	query := fmt.Sprintf(`
		UPDATE %s 
		SET deleted_at = NOW(), updated_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL`, attendancesTable)

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("error deleting attendance: %w", err)
	}

	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return fmt.Errorf("no rows affected, record may not exist")
	}

	return nil
}

func (r *AttendanceRepository) DeleteByIds(ctx context.Context, ids []int) error {
	if len(ids) == 0 {
		return fmt.Errorf("no IDs provided for deletion")
	}

	query := fmt.Sprintf(`
		UPDATE %s 
		SET deleted_at = NOW(), updated_at = NOW()
		WHERE id = ANY($1) AND deleted_at IS NULL`, attendancesTable)

	result, err := r.db.ExecContext(ctx, query, pq.Array(ids))
	if err != nil {
		return fmt.Errorf("error deleting multiple attendances: %w", err)
	}

	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return fmt.Errorf("no rows affected, records may not exist")
	}

	return nil
}

func (r *AttendanceRepository) GetByEmployeeAndDateRange(ctx context.Context, employeeID int, startDate, endDate time.Time) ([]humanresources.Attendance, error) {
	query := fmt.Sprintf(`
		SELECT * FROM %s 
		WHERE employee_id = $1 
		AND date BETWEEN $2 AND $3 
		AND deleted_at IS NULL
		ORDER BY date, check_in`, attendancesTable)

	var attendances []humanresources.Attendance
	if err := r.db.SelectContext(ctx, &attendances, query, employeeID, startDate, endDate); err != nil {
		return nil, fmt.Errorf("error getting attendances by employee and date range: %w", err)
	}

	return attendances, nil
}
