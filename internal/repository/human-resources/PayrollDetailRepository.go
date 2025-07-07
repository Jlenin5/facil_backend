package repositoryHumanresources

import (
	"github.com/Jlenin5/facil_backend/internal/domain/human-resources"
	"github.com/jmoiron/sqlx"
)

type PayrollDetailRepository struct {
	db *sqlx.DB
}

func NewPayrollDetailRepository(db *sqlx.DB) *PayrollDetailRepository {
	return &PayrollDetailRepository{db: db}
}

func (r *PayrollDetailRepository) Create(detail *humanresources.PayrollDetail) error {
	query := `
		INSERT INTO payroll_details (
			payroll_id, employee_id, base_salary, days_worked, hours_worked,
			overtime_hours, overtime_pay, bonuses, deductions, net_pay,
			payment_method, bank_account, status
		) VALUES (
			:payroll_id, :employee_id, :base_salary, :days_worked, :hours_worked,
			:overtime_hours, :overtime_pay, :bonuses, :deductions, :net_pay,
			:payment_method, :bank_account, :status
		)`
	_, err := r.db.NamedExec(query, detail)
	return err
}

func (r *PayrollDetailRepository) GetAll() ([]humanresources.PayrollDetail, error) {
	var payrollDetail []humanresources.PayrollDetail
	query := `SELECT * FROM payroll_details WHERE deleted_at IS NULL`
	err := r.db.Select(&payrollDetail, query)
	return payrollDetail, err
}

func (r *PayrollDetailRepository) GetByPayrollID(payrollID int) ([]humanresources.PayrollDetail, error) {
	var details []humanresources.PayrollDetail
	query := `SELECT * FROM payroll_details WHERE payroll_id = $1 AND deleted_at IS NULL`
	err := r.db.Select(&details, query, payrollID)
	return details, err
}

func (r *PayrollDetailRepository) Update(detail *humanresources.PayrollDetail) error {
	query := `
		UPDATE payroll_details SET
			base_salary = :base_salary,
			days_worked = :days_worked,
			hours_worked = :hours_worked,
			overtime_hours = :overtime_hours,
			overtime_pay = :overtime_pay,
			bonuses = :bonuses,
			deductions = :deductions,
			net_pay = :net_pay,
			payment_method = :payment_method,
			bank_account = :bank_account,
			status = :status,
			paid_at = :paid_at,
			updated_at = NOW()
		WHERE id = :id`
	_, err := r.db.NamedExec(query, detail)
	return err
}

func (r *PayrollDetailRepository) MarkAsPaid(id int) error {
	query := `UPDATE payroll_details SET status = 'paid', paid_at = NOW() WHERE id = $1`
	_, err := r.db.Exec(query, id)
	return err
}