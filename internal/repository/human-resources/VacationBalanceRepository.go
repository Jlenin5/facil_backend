package repositoryHumanresources

import (
	"context"
	"fmt"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/Jlenin5/facil_backend/internal/domain/human-resources"
	"github.com/jmoiron/sqlx"
)

type VacationBalanceRepository struct {
	db *sqlx.DB
}

func NewVacationBalanceRepository(db *sqlx.DB) *VacationBalanceRepository {
	return &VacationBalanceRepository{db: db}
}

const (
	vacationBalancesTable = "vacation_balances"
	employeesTable8        = "employees"
)

func (r *VacationBalanceRepository) Create(ctx context.Context, balance *humanresources.VacationBalance) error {
	query := fmt.Sprintf(`
		INSERT INTO %s (
			employee_id, year, total_days, days_taken, days_remaining
		) VALUES (
			:employee_id, :year, :total_days, :days_taken, :days_remaining
		) RETURNING id`, vacationBalancesTable)

	rows, err := r.db.NamedQueryContext(ctx, query, balance)
	if err != nil {
		return fmt.Errorf("error creating vacation balance: %w", err)
	}
	defer rows.Close()

	if rows.Next() {
		if err := rows.Scan(&balance.Id); err != nil {
			return fmt.Errorf("error getting created ID: %w", err)
		}
	}

	return nil
}

func (r *VacationBalanceRepository) GetAll(ctx context.Context) ([]humanresources.VacationBalance, error) {
	query := fmt.Sprintf(`
		SELECT
			vb.id, vb.employee_id, vb.year, vb.total_days, 
			vb.days_taken, vb.days_remaining, vb.created_at, vb.updated_at,
			e.id AS "employee.id", 
			e.names AS "employee.names", 
			e.surname AS "employee.surname", 
			e.second_surname AS "employee.second_surname", 
			e.warehouse_id AS "employee.warehouse_id"
		FROM %s vb
		JOIN %s e ON vb.employee_id = e.id
		WHERE vb.deleted_at IS NULL
		ORDER BY vb.year DESC`, vacationBalancesTable, employeesTable8)

	var balances []struct {
		humanresources.VacationBalance
		Employee domain.EmployeeReducedData `db:"employee"`
	}

	if err := r.db.SelectContext(ctx, &balances, query); err != nil {
		return nil, fmt.Errorf("error getting vacation balances: %w", err)
	}

	result := make([]humanresources.VacationBalance, len(balances))
	for i, bal := range balances {
		bal.VacationBalance.Employee = bal.Employee
		result[i] = bal.VacationBalance
	}

	return result, nil
}

func (r *VacationBalanceRepository) GetById(ctx context.Context, id int) (*humanresources.VacationBalance, error) {
	query := fmt.Sprintf(`
		SELECT
			vb.id, vb.employee_id, vb.year, vb.total_days, 
			vb.days_taken, vb.days_remaining, vb.created_at, vb.updated_at,
			e.id AS "employee.id", 
			e.names AS "employee.names", 
			e.surname AS "employee.surname", 
			e.second_surname AS "employee.second_surname", 
			e.warehouse_id AS "employee.warehouse_id"
		FROM %s vb
		JOIN %s e ON vb.employee_id = e.id
		WHERE vb.id = $1 AND vb.deleted_at IS NULL`, vacationBalancesTable, employeesTable8)

	var result struct {
		humanresources.VacationBalance
		Employee domain.EmployeeReducedData `db:"employee"`
	}

	if err := r.db.GetContext(ctx, &result, query, id); err != nil {
		return nil, fmt.Errorf("error getting vacation balance by ID: %w", err)
	}

	vacationBalance := result.VacationBalance
	vacationBalance.Employee = result.Employee

	return &vacationBalance, nil
}

func (r *VacationBalanceRepository) Update(ctx context.Context, balance *humanresources.VacationBalance) error {
	query := fmt.Sprintf(`
		UPDATE %s SET
			total_days = :total_days,
			days_taken = :days_taken,
			days_remaining = :days_remaining,
			updated_at = NOW()
		WHERE id = :id AND deleted_at IS NULL`, vacationBalancesTable)

	result, err := r.db.NamedExecContext(ctx, query, balance)
	if err != nil {
		return fmt.Errorf("error updating vacation balance: %w", err)
	}

	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return fmt.Errorf("no rows affected, record may not exist")
	}

	return nil
}

func (r *VacationBalanceRepository) DeleteById(ctx context.Context, id int) error {
	query := fmt.Sprintf(`
		UPDATE %s 
		SET deleted_at = NOW(), updated_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL`, vacationBalancesTable)

	result, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return fmt.Errorf("error deleting vacation balance: %w", err)
	}

	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return fmt.Errorf("no rows affected, record may not exist")
	}

	return nil
}

func (r *VacationBalanceRepository) GetByEmployeeId(ctx context.Context, employeeID int) ([]humanresources.VacationBalance, error) {
	query := fmt.Sprintf(`
		SELECT
			vb.id, vb.employee_id, vb.year, vb.total_days, 
			vb.days_taken, vb.days_remaining, vb.created_at, vb.updated_at,
			e.id AS "employee.id", 
			e.names AS "employee.names", 
			e.surname AS "employee.surname", 
			e.second_surname AS "employee.second_surname", 
			e.warehouse_id AS "employee.warehouse_id"
		FROM %s vb
		JOIN %s e ON vb.employee_id = e.id
		WHERE vb.employee_id = $1 AND vb.deleted_at IS NULL
		ORDER BY vb.year DESC`, vacationBalancesTable, employeesTable8)

	var balances []struct {
		humanresources.VacationBalance
		Employee domain.EmployeeReducedData `db:"employee"`
	}

	if err := r.db.SelectContext(ctx, &balances, query, employeeID); err != nil {
		return nil, fmt.Errorf("error getting vacation balances by employee ID: %w", err)
	}

	result := make([]humanresources.VacationBalance, len(balances))
	for i, bal := range balances {
		bal.VacationBalance.Employee = bal.Employee
		result[i] = bal.VacationBalance
	}

	return result, nil
}

func (r *VacationBalanceRepository) GetByEmployeeAndYear(ctx context.Context, employeeID, year int) (*humanresources.VacationBalance, error) {
	query := fmt.Sprintf(`
		SELECT
			vb.id, vb.employee_id, vb.year, vb.total_days, 
			vb.days_taken, vb.days_remaining, vb.created_at, vb.updated_at,
			e.id AS "employee.id", 
			e.names AS "employee.names", 
			e.surname AS "employee.surname", 
			e.second_surname AS "employee.second_surname", 
			e.warehouse_id AS "employee.warehouse_id"
		FROM %s vb
		JOIN %s e ON vb.employee_id = e.id
		WHERE vb.employee_id = $1 AND vb.year = $2 AND vb.deleted_at IS NULL`, vacationBalancesTable, employeesTable8)

	var result struct {
		humanresources.VacationBalance
		Employee domain.EmployeeReducedData `db:"employee"`
	}

	if err := r.db.GetContext(ctx, &result, query, employeeID, year); err != nil {
		return nil, fmt.Errorf("error getting vacation balance by employee and year: %w", err)
	}

	vacationBalance := result.VacationBalance
	vacationBalance.Employee = result.Employee

	return &vacationBalance, nil
}