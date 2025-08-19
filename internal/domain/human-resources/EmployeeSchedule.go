package humanresources

import (
	"time"

	"github.com/Jlenin5/facil_backend/internal/domain"
)

type EmployeeSchedule struct {
	Id            int                        `db:"id" json:"id"`
	EmployeeId    int                        `db:"employee_id" json:"employee_id"`
	Employee      domain.EmployeeReducedData `json:"employee"`
	ScheduleId    int                        `db:"schedule_id" json:"schedule_id"`
	Schedule      domain.ScheduleReducedData `db:"work_schedule" json:"work_schedule"`
	EffectiveDate time.Time                  `db:"effective_date" json:"effective_date"`
	EndDate       domain.NullTime            `db:"end_date" json:"end_date,omitempty"`
	Created_at    time.Time                  `db:"created_at" json:"-"`
	Updated_at    domain.NullTime            `db:"updated_at" json:"-"`
	Deleted_at    domain.NullTime            `db:"deleted_at" json:"-"`
}
