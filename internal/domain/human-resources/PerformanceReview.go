package humanresources

import (
	"time"

	"github.com/Jlenin5/facil_backend/internal/domain"
)

type PerformanceReview struct {
	Id                  int                        `db:"id" json:"id"`
	EmployeeId          int                        `db:"employee_id" json:"employee_id"`
	Employee            domain.EmployeeReducedData `json:"employee"`
	ReviewerId          int                        `db:"reviewer_id" json:"reviewer_id"`
	ReviewDate          time.Time                  `db:"review_date" json:"review_date"`
	NextReviewDate      domain.NullTime            `db:"next_review_date" json:"next_review_date,omitempty"`
	PerformanceScore    domain.NullInt             `db:"performance_score" json:"performance_score,omitempty"`
	Strengths           domain.NullString          `db:"strengths" json:"strengths,omitempty"`
	AreasForImprovement domain.NullString          `db:"areas_for_improvement" json:"areas_for_improvement,omitempty"`
	Comments            domain.NullString          `db:"comments" json:"comments,omitempty"`
	Status              string                     `db:"status" json:"status"`
	AcknowledgedAt      domain.NullTime            `db:"acknowledged_at" json:"acknowledged_at,omitempty"`
	Created_at          time.Time                  `db:"created_at" json:"-"`
	Updated_at          time.Time                  `db:"updated_at" json:"-"`
	Deleted_at          domain.NullTime            `db:"deleted_at" json:"-"`
}
