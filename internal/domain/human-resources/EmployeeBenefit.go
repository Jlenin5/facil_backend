package humanresources

import (
	"time"

	"github.com/Jlenin5/facil_backend/internal/domain"
)

type EmployeeBenefit struct {
	Id          int                        `db:"id" json:"id"`
	EmployeeId  int                        `db:"employee_id" json:"employee_id"`
	Employee    domain.EmployeeReducedData `json:"employee"`
	BenefitType string                     `db:"benefit_type" json:"benefit_type"`
	Description domain.NullString          `db:"description" json:"description,omitempty"`
	Amount      domain.NullFloat           `db:"amount" json:"amount,omitempty"`
	StartDate   time.Time                  `db:"start_date" json:"start_date"`
	EndDate     domain.NullTime            `db:"end_date" json:"end_date,omitempty"`
	Status      uint8                      `db:"status" json:"status"`
	Created_at  time.Time                  `db:"created_at" json:"-"`
	Updated_at  time.Time                  `db:"updated_at" json:"-"`
	Deleted_at  domain.NullTime            `db:"deleted_at" json:"-"`
}
