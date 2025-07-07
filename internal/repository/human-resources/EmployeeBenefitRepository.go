package repositoryHumanresources

import (
	"context"
	"fmt"

	"github.com/Jlenin5/facil_backend/internal/domain"
	humanresources "github.com/Jlenin5/facil_backend/internal/domain/human-resources"
	"github.com/jmoiron/sqlx"
)

type EmployeeBenefitRepository struct {
	db *sqlx.DB
}

func NewEmployeeBenefitRepository(db *sqlx.DB) *EmployeeBenefitRepository {
	return &EmployeeBenefitRepository{db: db}
}

const (
	employeeBenefitsTable = "employee_benefits"
	employeesTable3       = "employees"
)

func (r *EmployeeBenefitRepository) Create(ctx context.Context, benefit *humanresources.EmployeeBenefit) error {
	query := fmt.Sprintf(`
		INSERT INTO %s (
			employee_id, benefit_type, description, amount, 
			start_date, end_date, status
		) VALUES (
			:employee_id, :benefit_type, :description, :amount, 
			:start_date, :end_date, :status
		) RETURNING id, created_at, updated_at`, employeeBenefitsTable)

	rows, err := r.db.NamedQueryContext(ctx, query, benefit)
	if err != nil {
		return fmt.Errorf("error creating employee benefit: %w", err)
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.Scan(
			&benefit.Id,
			&benefit.Created_at,
			&benefit.Updated_at,
		); err != nil {
			return fmt.Errorf("error getting created benefit data: %w", err)
		}
	}

	return nil
}

func (r *EmployeeBenefitRepository) GetAll(ctx context.Context) ([]humanresources.EmployeeBenefit, error) {
	query := fmt.Sprintf(`
		SELECT
			eb.id, eb.employee_id, eb.benefit_type, 
			eb.description, eb.amount, eb.start_date, 
			eb.end_date, eb.status,
			e.id AS "employee.id",
			e.names AS "employee.names",
			e.surname AS "employee.surname",
			e.second_surname AS "employee.second_surname",
			e.warehouse_id AS "employee.warehouse_id"
		FROM %s eb
		JOIN %s e ON eb.employee_id = e.id
		WHERE eb.deleted_at IS NULL
		ORDER BY eb.start_date DESC`, employeeBenefitsTable, employeesTable3)

	var results []struct {
		humanresources.EmployeeBenefit
		Employee domain.EmployeeReducedData `db:"employee"`
	}

	if err := r.db.SelectContext(ctx, &results, query); err != nil {
		return nil, fmt.Errorf("error getting employee benefits: %w", err)
	}

	benefits := make([]humanresources.EmployeeBenefit, len(results))
	for i, res := range results {
		res.EmployeeBenefit.Employee = res.Employee
		benefits[i] = res.EmployeeBenefit
	}

	return benefits, nil
}

func (r *EmployeeBenefitRepository) GetById(ctx context.Context, id int) (*humanresources.EmployeeBenefit, error) {
	query := fmt.Sprintf(`
		SELECT
			eb.*,
			e.id AS "employee.id",
			e.names AS "employee.names",
			e.surname AS "employee.surname",
			e.second_surname AS "employee.second_surname",
			e.warehouse_id AS "employee.warehouse_id"
		FROM %s eb
		JOIN %s e ON eb.employee_id = e.id
		WHERE eb.id = $1 AND eb.deleted_at IS NULL`, employeeBenefitsTable, employeesTable3)

	var result struct {
		humanresources.EmployeeBenefit
		Employee domain.EmployeeReducedData `db:"employee"`
	}

	if err := r.db.GetContext(ctx, &result, query, id); err != nil {
		return nil, fmt.Errorf("error getting employee benefit by ID: %w", err)
	}

	benefit := result.EmployeeBenefit
	benefit.Employee = result.Employee

	return &benefit, nil
}

func (r *EmployeeBenefitRepository) Update(ctx context.Context, benefit *humanresources.EmployeeBenefit) error {
	query := fmt.Sprintf(`
		UPDATE %s SET
			benefit_type = :benefit_type,
			description = :description,
			amount = :amount,
			start_date = :start_date,
			end_date = :end_date,
			status = :status,
			updated_at = NOW()
		WHERE id = :id AND deleted_at IS NULL`, employeeBenefitsTable)

	result, err := r.db.NamedExecContext(ctx, query, benefit)
	if err != nil {
		return fmt.Errorf("error updating employee benefit: %w", err)
	}

	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return fmt.Errorf("no rows affected, record may not exist")
	}

	return nil
}

func (r *EmployeeBenefitRepository) Deactivate(ctx context.Context, id int) error {
	query := fmt.Sprintf(`
		UPDATE %s 
		SET status = false, updated_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL`, employeeBenefitsTable)

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("error deactivating employee benefit: %w", err)
	}

	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return fmt.Errorf("no rows affected, record may not exist")
	}

	return nil
}

func (r *EmployeeBenefitRepository) DeleteById(ctx context.Context, id int) error {
	query := fmt.Sprintf(`
		UPDATE %s 
		SET deleted_at = NOW(), updated_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL`, employeeBenefitsTable)

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("error deleting employee benefit: %w", err)
	}

	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return fmt.Errorf("no rows affected, record may not exist")
	}

	return nil
}
