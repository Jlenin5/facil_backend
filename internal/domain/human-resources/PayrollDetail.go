package humanresources

import (
	"time"

	"github.com/Jlenin5/facil_backend/internal/domain"
)

type PayrollDetail struct {
	Id            int               `db:"id" json:"id"`
	PayrollId     int               `db:"payroll_id" json:"payrollId"`
	EmployeeId    int               `db:"employee_id" json:"employeeId"`
	BaseSalary    float64           `db:"base_salary" json:"baseSalary"`
	DaysWorked    int               `db:"days_worked" json:"daysWorked"`
	HoursWorked   float64           `db:"hours_worked" json:"hoursWorked"`
	OvertimeHours float64           `db:"overtime_hours" json:"overtimeHours,omitempty"`
	OvertimePay   float64           `db:"overtime_pay" json:"overtimePay,omitempty"`
	Bonuses       float64           `db:"bonuses" json:"bonuses,omitempty"`
	Deductions    float64           `db:"deductions" json:"deductions,omitempty"`
	NetPay        float64           `db:"net_pay" json:"netPay"`
	PaymentMethod domain.NullString `db:"payment_method" json:"paymentMethod,omitempty"`
	BankAccount   domain.NullString `db:"bank_account" json:"bankAccount,omitempty"`
	Status        string            `db:"status" json:"status"`
	PaidAt        domain.NullTime   `db:"paid_at" json:"paidAt,omitempty"`
	Created_at    time.Time         `db:"created_at" json:"-"`
	Updated_at    time.Time         `db:"updated_at" json:"-"`
	Deleted_at    domain.NullTime   `db:"deleted_at" json:"-"`
}
