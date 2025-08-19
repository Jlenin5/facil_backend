package humanresources

import (
	"time"

	"github.com/Jlenin5/facil_backend/internal/domain"
)

type VacationBalance struct {
	Id            int                        `db:"id" json:"id"`
	EmployeeId    int                        `db:"employee_id" json:"employee_id"`
	Employee      domain.EmployeeReducedData `json:"employee"`
	Year          int                        `db:"year" json:"year"`
	TotalDays     int                        `db:"total_days" json:"total_days"`
	DaysTaken     int                        `db:"days_taken" json:"days_taken"`
	DaysRemaining int                        `db:"days_remaining" json:"days_remaining"`
	Created_at    time.Time                  `db:"created_at" json:"-"`
	Updated_at    domain.NullTime            `db:"updated_at" json:"-"`
	Deleted_at    domain.NullTime            `db:"deleted_at" json:"-"`
}