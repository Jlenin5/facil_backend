package humanresources

import (
	"time"

	"github.com/Jlenin5/facil_backend/internal/domain"
)

type OvertimeRequest struct {
	Id              int                        `db:"id" json:"id"`
	EmployeeId      int                        `db:"employee_id" json:"employee_id"`
	Employee        domain.EmployeeReducedData `json:"employee"`
	Date            time.Time                  `db:"date" json:"date"`
	StartTime       string                     `db:"start_time" json:"start_time"`
	EndTime         string                     `db:"end_time" json:"end_time"`
	Hours           float64                    `db:"hours" json:"hours"`
	Reason          string                     `db:"reason" json:"reason"`
	Status          string                     `db:"status" json:"status"`
	ApprovedBy      domain.NullInt             `db:"approved_by" json:"approved_by,omitempty"`
	ApprovedAt      domain.NullTime            `db:"approved_at" json:"approved_at,omitempty"`
	RejectionReason domain.NullString          `db:"rejection_reason" json:"rejection_reason,omitempty"`
	Created_at      time.Time                  `db:"created_at" json:"created_at"`
	Updated_at      domain.NullTime            `db:"updated_at" json:"updated_at"`
	Deleted_at      domain.NullTime            `db:"deleted_at" json:"-"`
}
