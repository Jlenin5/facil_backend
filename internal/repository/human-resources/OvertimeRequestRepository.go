package repositoryHumanresources

import (
	"context"
	"fmt"

	"github.com/Jlenin5/facil_backend/internal/domain/human-resources"
	"github.com/jmoiron/sqlx"
)

type OvertimeRequestRepository struct {
	db *sqlx.DB
}

func NewOvertimeRequestRepository(db *sqlx.DB) *OvertimeRequestRepository {
	return &OvertimeRequestRepository{db: db}
}

const (
	overtimeRequestsTable = "overtime_requests"
	employeesTable6       = "employees"
)

func (r *OvertimeRequestRepository) Create(ctx context.Context, request *humanresources.OvertimeRequest) error {
	query := fmt.Sprintf(`
		INSERT INTO %s (
			employee_id, date, start_time, end_time, reason, status
		) VALUES (
			:employee_id, :date, :start_time, :end_time, :reason, :status
		) RETURNING id`, overtimeRequestsTable)

	rows, err := r.db.NamedQueryContext(ctx, query, request)
	if err != nil {
		return fmt.Errorf("error creating overtime request: %w", err)
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.Scan(&request.Id); err != nil {
			return fmt.Errorf("error getting created ID: %w", err)
		}
	}

	return nil
}

func (r *OvertimeRequestRepository) GetAll(ctx context.Context) ([]humanresources.OvertimeRequest, error) {
	query := fmt.Sprintf(`
		SELECT
			ors.id, ors.employee_id, ors.date, ors.start_time, ors.end_time, 
			ors.hours, ors.reason, ors.status, ors.approved_by, ors.approved_at, 
			ors.rejection_reason, ors.created_at, ors.updated_at,
			e.id AS "employee.id", 
			e.names AS "employee.names", 
			e.surname AS "employee.surname", 
			e.second_surname AS "employee.second_surname", 
			e.warehouse_id AS "employee.warehouse_id"
		FROM %s ors
		INNER JOIN %s e ON ors.employee_id = e.id
		WHERE ors.deleted_at IS NULL
		ORDER BY ors.date DESC`, overtimeRequestsTable, employeesTable6)

	var requests []humanresources.OvertimeRequest
	if err := r.db.SelectContext(ctx, &requests, query); err != nil {
		return nil, fmt.Errorf("error getting all overtime requests: %w", err)
	}

	return requests, nil
}

func (r *OvertimeRequestRepository) GetById(ctx context.Context, id int) (*humanresources.OvertimeRequest, error) {
	query := fmt.Sprintf(`
		SELECT
			ors.id, ors.employee_id, ors.date, ors.start_time, ors.end_time, 
			ors.hours, ors.reason, ors.status, ors.approved_by, ors.approved_at, 
			ors.rejection_reason, ors.created_at, ors.updated_at,
			e.id AS "employee.id", 
			e.names AS "employee.names", 
			e.surname AS "employee.surname", 
			e.second_surname AS "employee.second_surname", 
			e.warehouse_id AS "employee.warehouse_id"
		FROM %s ors
		INNER JOIN %s e ON ors.employee_id = e.id
		WHERE ors.id = $1 AND ors.deleted_at IS NULL`, overtimeRequestsTable, employeesTable6)

	var request humanresources.OvertimeRequest
	if err := r.db.GetContext(ctx, &request, query, id); err != nil {
		return nil, fmt.Errorf("error getting overtime request by ID: %w", err)
	}

	return &request, nil
}

func (r *OvertimeRequestRepository) Update(ctx context.Context, request *humanresources.OvertimeRequest) error {
	query := fmt.Sprintf(`
		UPDATE %s SET
			date = :date,
			start_time = :start_time,
			end_time = :end_time,
			reason = :reason,
			status = :status,
			approved_by = :approved_by,
			approved_at = :approved_at,
			rejection_reason = :rejection_reason,
			updated_at = NOW()
		WHERE id = :id AND deleted_at IS NULL`, overtimeRequestsTable)

	result, err := r.db.NamedExecContext(ctx, query, request)
	if err != nil {
		return fmt.Errorf("error updating overtime request: %w", err)
	}

	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return fmt.Errorf("no rows affected, record may not exist")
	}

	return nil
}

func (r *OvertimeRequestRepository) DeleteById(ctx context.Context, id int) error {
	query := fmt.Sprintf(`
		UPDATE %s 
		SET deleted_at = NOW(), updated_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL`, overtimeRequestsTable)

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("error deleting overtime request: %w", err)
	}

	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return fmt.Errorf("no rows affected, record may not exist")
	}

	return nil
}

func (r *OvertimeRequestRepository) Approve(ctx context.Context, id, approvedBy int) error {
	query := fmt.Sprintf(`
		UPDATE %s SET
			status = 'approved',
			approved_by = $1,
			approved_at = NOW(),
			updated_at = NOW()
		WHERE id = $2 AND deleted_at IS NULL`, overtimeRequestsTable)

	result, err := r.db.ExecContext(ctx, query, approvedBy, id)
	if err != nil {
		return fmt.Errorf("error approving overtime request: %w", err)
	}

	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return fmt.Errorf("no rows affected, record may not exist")
	}

	return nil
}

func (r *OvertimeRequestRepository) Reject(ctx context.Context, id, approvedBy int, reason string) error {
	query := fmt.Sprintf(`
		UPDATE %s SET
			status = 'rejected',
			approved_by = $1,
			approved_at = NOW(),
			rejection_reason = $2,
			updated_at = NOW()
		WHERE id = $3 AND deleted_at IS NULL`, overtimeRequestsTable)

	result, err := r.db.ExecContext(ctx, query, approvedBy, reason, id)
	if err != nil {
		return fmt.Errorf("error rejecting overtime request: %w", err)
	}

	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return fmt.Errorf("no rows affected, record may not exist")
	}

	return nil
}