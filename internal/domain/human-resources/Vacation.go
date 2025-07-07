package humanresources

import (
	"time"

	"github.com/Jlenin5/facil_backend/internal/domain"
)

type Vacation struct {
	Id              int                        `db:"id" json:"id"`
	EmployeeId      int                        `db:"employee_id" json:"employee_id"`
	Employee        domain.EmployeeReducedData `json:"employee"`
	StartDate       time.Time                  `db:"start_date" json:"start_date"`
	EndDate         time.Time                  `db:"end_date" json:"end_date"`
	DaysTaken       int                        `db:"days_taken" json:"days_taken"`
	Status          string                     `db:"status" json:"status"`
	ApprovedBy      domain.NullInt             `db:"approved_by" json:"approved_by,omitempty"`
	ApprovedAt      domain.NullTime            `db:"approved_at" json:"approved_at,omitempty"`
	RejectionReason domain.NullString          `db:"rejection_reason" json:"rejection_reason,omitempty"`
	Created_at      time.Time                  `db:"created_at" json:"-"`
	Updated_at      time.Time                  `db:"updated_at" json:"-"`
	Deleted_at      domain.NullTime            `db:"deleted_at" json:"-"`
}
