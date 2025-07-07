package repositoryHumanresources

import (
	humanresources "github.com/Jlenin5/facil_backend/internal/domain/human-resources"
	"github.com/jmoiron/sqlx"
)

type PayrollRepository struct {
	db *sqlx.DB
}

func NewPayrollRepository(db *sqlx.DB) *PayrollRepository {
	return &PayrollRepository{db: db}
}

func (r *PayrollRepository) Create(payroll *humanresources.Payroll) error {
	query := `
		INSERT INTO payrolls (
			reference, period_start, period_end, payment_date, status, notes, created_by
		) VALUES (
			:reference, :period_start, :period_end, :payment_date, :status, :notes, :created_by
		) RETURNING id`

	rows, err := r.db.NamedQuery(query, payroll)
	if err != nil {
		return err
	}
	defer rows.Close()

	if rows.Next() {
		err = rows.Scan(&payroll.Id)
	}
	return err
}

func (r *PayrollRepository) GetAll() ([]humanresources.Payroll, error) {
	var payrolls []humanresources.Payroll
	query := `SELECT * FROM payrolls WHERE deleted_at IS NULL ORDER BY period_start DESC`
	err := r.db.Select(&payrolls, query)
	return payrolls, err
}

func (r *PayrollRepository) GetById(id int) (*humanresources.Payroll, error) {
	var payroll humanresources.Payroll
	query := `SELECT * FROM payrolls WHERE id = $1 AND deleted_at IS NULL`
	err := r.db.Get(&payroll, query, id)
	return &payroll, err
}

func (r *PayrollRepository) Update(payroll *humanresources.Payroll) error {
	query := `
		UPDATE payrolls SET
			reference = :reference,
			period_start = :period_start,
			period_end = :period_end,
			payment_date = :payment_date,
			status = :status,
			notes = :notes,
			approved_by = :approved_by,
			approved_at = :approved_at,
			updated_at = NOW()
		WHERE id = :id`
	_, err := r.db.NamedExec(query, payroll)
	return err
}

func (r *PayrollRepository) ChangeStatus(id int, status string) error {
	query := `UPDATE payrolls SET status = $1, updated_at = NOW() WHERE id = $2`
	_, err := r.db.Exec(query, status, id)
	return err
}

func (r *PayrollRepository) Approve(id, approvedBy int) error {
	query := `
		UPDATE payrolls SET
			status = 'approved',
			approved_by = $1,
			approved_at = NOW(),
			updated_at = NOW()
		WHERE id = $2`
	_, err := r.db.Exec(query, approvedBy, id)
	return err
}
