package humanresources

import (
	"time"

	"github.com/Jlenin5/facil_backend/internal/domain"
)

type Attendance struct {
	Id                    int                              `db:"id" json:"id"`
	EmployeeId            int                              `db:"employee_id" json:"employee_id"`
	Employee              domain.EmployeeReducedData       `json:"employee"`
	AttendanceTypeId      int                              `db:"attendance_type_id" json:"attendance_type_id"`
	AttendanceType        domain.AttendanceTypeReducedData `db:"attendance_type" json:"attendance_type"`
	Date                  time.Time                        `db:"date" json:"date"`
	CheckIn               domain.NullTime                  `db:"check_in" json:"check_in,omitempty"`
	CheckOut              domain.NullTime                  `db:"check_out" json:"check_out,omitempty"`
	WorkedHours           float64                          `db:"worked_hours" json:"worked_hours,omitempty"`
	LateMinutes           int                              `db:"late_minutes" json:"late_minutes,omitempty"`
	EarlyDepartureMinutes int                              `db:"early_departure_minutes" json:"early_departure_minutes,omitempty"`
	Notes                 domain.NullString                `db:"notes" json:"notes,omitempty"`
	Status                string                           `db:"status" json:"status"`
	ApprovedBy            domain.NullInt                   `db:"approved_by" json:"approved_by,omitempty"`
	ApprovedAt            domain.NullTime                  `db:"approved_at" json:"approved_at,omitempty"`
	Created_at            time.Time                        `db:"created_at" json:"-"`
	Updated_at            time.Time                        `db:"updated_at" json:"-"`
	Deleted_at            domain.NullTime                  `db:"deleted_at" json:"-"`
}
