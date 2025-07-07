package repositoryHumanresources

import (
	"context"
	"fmt"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/Jlenin5/facil_backend/internal/domain/human-resources"
	"github.com/jmoiron/sqlx"
)

type AbsenceRequestRepository struct {
	db *sqlx.DB
}

func NewAbsenceRequestRepository(db *sqlx.DB) *AbsenceRequestRepository {
	return &AbsenceRequestRepository{db: db}
}

const (
	absenceRequestsTable = "absence_requests"
	employeesTable2      = "employees"
	absenceTypesTable2   = "absence_types"
)

func (r *AbsenceRequestRepository) Create(ctx context.Context, request *humanresources.AbsenceRequest) error {
	query := fmt.Sprintf(`
		INSERT INTO %s (
			employee_id, absence_type_id, start_date, end_date, reason, status
		) VALUES (
			:employee_id, :absence_type_id, :start_date, :end_date, :reason, :status
		) RETURNING id`, absenceRequestsTable)

	rows, err := r.db.NamedQueryContext(ctx, query, request)
	if err != nil {
		return fmt.Errorf("error creating absence request: %w", err)
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.Scan(&request.Id); err != nil {
			return fmt.Errorf("error getting created ID: %w", err)
		}
	}

	return nil
}

func (r *AbsenceRequestRepository) GetAll(ctx context.Context) ([]humanresources.AbsenceRequest, error) {
	query := fmt.Sprintf(`
		SELECT
			ar.id, ar.employee_id, ar.absence_type_id, 
			ar.start_date, ar.end_date, ar.reason, ar.status, 
			ar.approved_by, ar.approved_at, ar.rejection_reason,
			e.id AS "employee.id", 
			e.names AS "employee.names",   
			e.surname AS "employee.surname", 
			e.second_surname AS "employee.second_surname", 
			e.warehouse_id AS "employee.warehouse_id",
			at.id AS "absence_type.id", 
			at.name AS "absence_type.name", 
			at.code AS "absence_type.code"
		FROM %s ar
		JOIN %s e ON ar.employee_id = e.id
		JOIN %s at ON ar.absence_type_id = at.id
		WHERE ar.deleted_at IS NULL
		ORDER BY ar.start_date DESC`, absenceRequestsTable, employeesTable2, absenceTypesTable2)

	var requests []struct {
		humanresources.AbsenceRequest
		Employee    domain.EmployeeReducedData    `db:"employee"`
		AbsenceType domain.AbsenceTypeReducedData `db:"absence_type"`
	}

	if err := r.db.SelectContext(ctx, &requests, query); err != nil {
		return nil, fmt.Errorf("error getting absence requests: %w", err)
	}

	result := make([]humanresources.AbsenceRequest, len(requests))
	for i, req := range requests {
		req.AbsenceRequest.Employee = req.Employee
		req.AbsenceRequest.AbsenceType = req.AbsenceType
		result[i] = req.AbsenceRequest
	}

	return result, nil
}

func (r *AbsenceRequestRepository) GetById(ctx context.Context, id int) (*humanresources.AbsenceRequest, error) {
	query := fmt.Sprintf(`
		SELECT
			ar.*,
			e.id AS "employee.id", 
			e.names AS "employee.names",
			-- otros campos de employee...
			at.id AS "absence_type.id",
			at.name AS "absence_type.name",
			at.code AS "absence_type.code"
		FROM %s ar
		JOIN %s e ON ar.employee_id = e.id
		JOIN %s at ON ar.absence_type_id = at.id
		WHERE ar.id = $1 AND ar.deleted_at IS NULL`, absenceRequestsTable, employeesTable2, absenceTypesTable2)

	var result struct {
		humanresources.AbsenceRequest
		Employee    domain.EmployeeReducedData    `db:"employee"`
		AbsenceType domain.AbsenceTypeReducedData `db:"absence_type"`
	}

	if err := r.db.GetContext(ctx, &result, query, id); err != nil {
		return nil, fmt.Errorf("error getting absence request by ID: %w", err)
	}

	absenceRequest := result.AbsenceRequest
	absenceRequest.Employee = result.Employee
	absenceRequest.AbsenceType = result.AbsenceType

	return &absenceRequest, nil
}

func (r *AbsenceRequestRepository) Update(ctx context.Context, request *humanresources.AbsenceRequest) error {
	query := fmt.Sprintf(`
		UPDATE %s SET
			employee_id = :employee_id,
			absence_type_id = :absence_type_id,
			start_date = :start_date,
			end_date = :end_date,
			reason = :reason,
			status = :status,
			approved_by = :approved_by,
			approved_at = :approved_at,
			rejection_reason = :rejection_reason,
			updated_at = NOW()
		WHERE id = :id AND deleted_at IS NULL`, absenceRequestsTable)

	result, err := r.db.NamedExecContext(ctx, query, request)
	if err != nil {
		return fmt.Errorf("error updating absence request: %w", err)
	}

	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return fmt.Errorf("no rows affected, record may not exist")
	}

	return nil
}

func (r *AbsenceRequestRepository) DeleteById(ctx context.Context, id int) error {
	query := fmt.Sprintf(`
		UPDATE %s 
		SET deleted_at = NOW(), updated_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL`, absenceRequestsTable)

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("error deleting absence request: %w", err)
	}

	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return fmt.Errorf("no rows affected, record may not exist")
	}

	return nil
}

func (r *AbsenceRequestRepository) Approve(ctx context.Context, id, approvedBy int) error {
	query := fmt.Sprintf(`
		UPDATE %s SET
			status = 'approved',
			approved_by = $1,
			approved_at = NOW(),
			updated_at = NOW()
		WHERE id = $2 AND deleted_at IS NULL`, absenceRequestsTable)

	return r.executeStatusUpdate(ctx, query, approvedBy, id)
}

func (r *AbsenceRequestRepository) Reject(ctx context.Context, id, approvedBy int, reason string) error {
	query := fmt.Sprintf(`
		UPDATE %s SET
			status = 'rejected',
			approved_by = $1,
			approved_at = NOW(),
			rejection_reason = $2,
			updated_at = NOW()
		WHERE id = $3 AND deleted_at IS NULL`, absenceRequestsTable)

	return r.executeStatusUpdate(ctx, query, approvedBy, id, reason)
}

func (r *AbsenceRequestRepository) executeStatusUpdate(ctx context.Context, query string, args ...interface{}) error {
	result, err := r.db.ExecContext(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("error executing status update: %w", err)
	}

	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return fmt.Errorf("no rows affected, record may not exist")
	}

	return nil
}