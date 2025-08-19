package humanresources

import (
	"time"

	"github.com/Jlenin5/facil_backend/internal/domain"
)
type EmployeeIncident struct {
	Id           int                        `db:"id" json:"id"`
	EmployeeId   int                        `db:"employee_id" json:"employee_id"`
	Employee     domain.EmployeeReducedData `json:"employee"`
	IncidentType string                     `db:"incident_type" json:"incident_type"`
	IncidentDate time.Time                  `db:"incident_date" json:"incident_date"`
	Observation  domain.NullString          `db:"observation" json:"observation"`
	Discount     float64                    `db:"discount" json:"discount"`
	Total_to_pay float64                    `db:"total_to_pay" json:"total_to_pay"`
	ReportedBy   int                        `db:"reported_by" json:"reported_by"`
	Status       string                     `db:"status" json:"status"`
	ResolvedAt   domain.NullTime            `db:"resolved_at" json:"resolved_at,omitempty"`
	Created_at   time.Time                  `db:"created_at" json:"-"`
	Updated_at   domain.NullTime            `db:"updated_at" json:"-"`
	Deleted_at   domain.NullTime            `db:"deleted_at" json:"-"`
}