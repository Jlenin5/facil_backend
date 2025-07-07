package repositoryHumanresources

import (
	"context"
	"fmt"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/Jlenin5/facil_backend/internal/domain/human-resources"
	"github.com/jmoiron/sqlx"
)

type VacationRepository struct {
	db *sqlx.DB
}

func NewVacationRepository(db *sqlx.DB) *VacationRepository {
	return &VacationRepository{db: db}
}

const (
	vacationsTable  = "vacations"
	employeesTable9 = "employees"
)

func (r *VacationRepository) Create(ctx context.Context, vacation *humanresources.Vacation) error {
	query := fmt.Sprintf(`
		INSERT INTO %s (
			employee_id, start_date, end_date, days_taken, status
		) VALUES (
			:employee_id, :start_date, :end_date, :days_taken, :status
		) RETURNING id`, vacationsTable)

	rows, err := r.db.NamedQueryContext(ctx, query, vacation)
	if err != nil {
		return fmt.Errorf("error creating vacation: %w", err)
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.Scan(&vacation.Id); err != nil {
			return fmt.Errorf("error getting created ID: %w", err)
		}
	}

	return nil
}

func (r *VacationRepository) GetAll(ctx context.Context) ([]humanresources.Vacation, error) {
	query := fmt.Sprintf(`
		SELECT
			v.id, v.employee_id, v.start_date, v.end_date, 
			v.days_taken, v.status, v.approved_by, v.approved_at, 
			v.rejection_reason, v.created_at, v.updated_at,
			e.id AS "employee.id", 
			e.names AS "employee.names",
			e.surname AS "employee.surname", 
			e.second_surname AS "employee.second_surname", 
			e.warehouse_id AS "employee.warehouse_id"
		FROM %s v
		JOIN %s e ON v.employee_id = e.id
		WHERE v.deleted_at IS NULL
		ORDER BY v.start_date DESC`, vacationsTable, employeesTable9)

	var vacations []struct {
		humanresources.Vacation
		Employee domain.EmployeeReducedData `db:"employee"`
	}

	if err := r.db.SelectContext(ctx, &vacations, query); err != nil {
		return nil, fmt.Errorf("error getting vacations: %w", err)
	}

	result := make([]humanresources.Vacation, len(vacations))
	for i, vac := range vacations {
		vac.Vacation.Employee = vac.Employee
		result[i] = vac.Vacation
	}

	return result, nil
}

func (r *VacationRepository) GetById(ctx context.Context, id int) (*humanresources.Vacation, error) {
	query := fmt.Sprintf(`
		SELECT
			v.id, v.employee_id, v.start_date, v.end_date, 
			v.days_taken, v.status, v.approved_by, v.approved_at, 
			v.rejection_reason, v.created_at, v.updated_at,
			e.id AS "employee.id", 
			e.names AS "employee.names",
			e.surname AS "employee.surname", 
			e.second_surname AS "employee.second_surname", 
			e.warehouse_id AS "employee.warehouse_id"
		FROM %s v
		JOIN %s e ON v.employee_id = e.id
		WHERE v.id = $1 AND v.deleted_at IS NULL`, vacationsTable, employeesTable9)

	var result struct {
		humanresources.Vacation
		Employee domain.EmployeeReducedData `db:"employee"`
	}

	if err := r.db.GetContext(ctx, &result, query, id); err != nil {
		return nil, fmt.Errorf("error getting vacation by ID: %w", err)
	}

	vacation := result.Vacation
	vacation.Employee = result.Employee

	return &vacation, nil
}

func (r *VacationRepository) Update(ctx context.Context, vacation *humanresources.Vacation) error {
	query := fmt.Sprintf(`
		UPDATE %s SET
			employee_id = :employee_id,
			start_date = :start_date,
			end_date = :end_date,
			days_taken = :days_taken,
			status = :status,
			approved_by = :approved_by,
			approved_at = :approved_at,
			rejection_reason = :rejection_reason,
			updated_at = NOW()
		WHERE id = :id AND deleted_at IS NULL`, vacationsTable)

	result, err := r.db.NamedExecContext(ctx, query, vacation)
	if err != nil {
		return fmt.Errorf("error updating vacation: %w", err)
	}

	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return fmt.Errorf("no rows affected, record may not exist")
	}

	return nil
}

func (r *VacationRepository) DeleteById(ctx context.Context, id int) error {
	query := fmt.Sprintf(`
		UPDATE %s 
		SET deleted_at = NOW(), updated_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL`, vacationsTable)

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("error deleting vacation: %w", err)
	}

	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return fmt.Errorf("no rows affected, record may not exist")
	}

	return nil
}

func (r *VacationRepository) Approve(ctx context.Context, id, approvedBy int) error {
	query := fmt.Sprintf(`
		UPDATE %s SET
			status = 'approved',
			approved_by = $1,
			approved_at = NOW(),
			updated_at = NOW()
		WHERE id = $2 AND deleted_at IS NULL`, vacationsTable)

	result, err := r.db.ExecContext(ctx, query, approvedBy, id)
	if err != nil {
		return fmt.Errorf("error approving vacation: %w", err)
	}

	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return fmt.Errorf("no rows affected, record may not exist")
	}

	return nil
}

func (r *VacationRepository) Reject(ctx context.Context, id, approvedBy int, reason string) error {
	query := fmt.Sprintf(`
		UPDATE %s SET
			status = 'rejected',
			approved_by = $1,
			approved_at = NOW(),
			rejection_reason = $2,
			updated_at = NOW()
		WHERE id = $3 AND deleted_at IS NULL`, vacationsTable)

	result, err := r.db.ExecContext(ctx, query, approvedBy, reason, id)
	if err != nil {
		return fmt.Errorf("error rejecting vacation: %w", err)
	}

	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return fmt.Errorf("no rows affected, record may not exist")
	}

	return nil
}