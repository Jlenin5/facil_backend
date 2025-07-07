package humanresources

import (
	"time"

	"github.com/Jlenin5/facil_backend/internal/domain"
)

type Payroll struct {
	Id          int               `db:"id" json:"id"`
	Reference   string            `db:"reference" json:"reference"`
	PeriodStart time.Time         `db:"period_start" json:"periodStart"`
	PeriodEnd   time.Time         `db:"period_end" json:"periodEnd"`
	PaymentDate time.Time         `db:"payment_date" json:"paymentDate"`
	Status      string            `db:"status" json:"status"`
	Notes       domain.NullString `db:"notes" json:"notes,omitempty"`
	CreatedBy   int               `db:"created_by" json:"createdBy"`
	ApprovedBy  domain.NullInt    `db:"approved_by" json:"approvedBy,omitempty"`
	ApprovedAt  domain.NullTime   `db:"approved_at" json:"approvedAt,omitempty"`
	Created_at  time.Time         `db:"created_at" json:"-"`
	Updated_at  time.Time         `db:"updated_at" json:"-"`
	Deleted_at  domain.NullTime   `db:"deleted_at" json:"-"`
}
