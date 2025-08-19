package humanresources

import (
	"time"

	"github.com/Jlenin5/facil_backend/internal/domain"
)

type WorkSchedule struct {
	Id          int               `db:"id" json:"id"`
	Name        string            `db:"name" json:"name"`
	Description domain.NullString `db:"description" json:"description,omitempty"`
	IsDefault   uint8             `db:"is_default" json:"is_default"`
	Status      uint8             `db:"status" json:"status"`
	Created_at  time.Time         `db:"created_at" json:"-"`
	Updated_at  domain.NullTime   `db:"updated_at" json:"-"`
	Deleted_at  domain.NullTime   `db:"deleted_at" json:"-"`
}