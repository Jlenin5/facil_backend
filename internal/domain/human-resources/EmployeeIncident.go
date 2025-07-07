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
	Description  string                     `db:"description" json:"description"`
	Severity     string                     `db:"severity" json:"severity"`
	ActionTaken  domain.NullString          `db:"action_taken" json:"action_taken,omitempty"`
	ReportedBy   int                        `db:"reported_by" json:"reported_by"`
	Status       string                     `db:"status" json:"status"`
	ResolvedAt   domain.NullTime            `db:"resolved_at" json:"resolved_at,omitempty"`
	Created_at   time.Time                  `db:"created_at" json:"-"`
	Updated_at   time.Time                  `db:"updated_at" json:"-"`
	Deleted_at   domain.NullTime            `db:"deleted_at" json:"-"`
}
