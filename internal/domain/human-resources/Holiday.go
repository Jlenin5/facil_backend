package humanresources

import (
	"time"

	"github.com/Jlenin5/facil_backend/internal/domain"
)

type Holiday struct {
	Id          int               `db:"id" json:"id"`
	Name        string            `db:"name" json:"name"`
	Description domain.NullString `db:"description" json:"description,omitempty"`
	Date        time.Time         `db:"date" json:"date"`
	Recurring   uint8             `db:"recurring" json:"recurring"`
	Status      uint8             `db:"status" json:"status"`
	Created_at  time.Time         `db:"created_at" json:"-"`
	Updated_at  time.Time         `db:"updated_at" json:"-"`
	Deleted_at  domain.NullTime   `db:"deleted_at" json:"-"`
}
