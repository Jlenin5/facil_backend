package humanresources

import (
	"time"

	"github.com/Jlenin5/facil_backend/internal/domain"
)

type ScheduleDetail struct {
	Id           int             `db:"id" json:"id"`
	ScheduleId   int             `db:"schedule_id" json:"scheduleId"`
	DayOfWeek    int             `db:"day_of_week" json:"dayOfWeek"`
	StartTime    string          `db:"start_time" json:"startTime"`
	EndTime      string          `db:"end_time" json:"endTime"`
	IsWorkingDay uint8           `db:"is_working_day" json:"isWorkingDay"`
	Created_at   time.Time       `db:"created_at" json:"-"`
	Updated_at   time.Time       `db:"updated_at" json:"-"`
	Deleted_at   domain.NullTime `db:"deleted_at" json:"-"`
}
