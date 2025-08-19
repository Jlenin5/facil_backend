package humanresources

import (
	"time"

	"github.com/Jlenin5/facil_backend/internal/domain"
)

type AbsenceType struct {
	Id               int               `db:"id" json:"id"`
	Name             string            `db:"name" json:"name"`
	Code             string            `db:"code" json:"code"`
	Description      domain.NullString `db:"description" json:"description,omitempty"`
	RequiresApproval uint8             `db:"requires_approval" json:"requires_approval"`
	IsPaid           uint8             `db:"is_paid" json:"is_paid"`
	DeductsVacation  uint8             `db:"deducts_vacation" json:"deducts_vacation"`
	Status           uint8             `db:"status" json:"status"`
	Created_at       time.Time         `db:"created_at" json:"created_at"`
	Updated_at       domain.NullTime   `db:"updated_at" json:"updated_at"`
	Deleted_at       domain.NullTime   `db:"deleted_at" json:"-"`
}