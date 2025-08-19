package humanresources

import (
	"time"

	"github.com/Jlenin5/facil_backend/internal/domain"
)

type AttendanceTypes struct {
	Id          int               `db:"id" json:"id"`
	Name        string            `db:"name" json:"name"`
	Code        domain.NullString `db:"code" json:"code"`
	Description string            `db:"description" json:"description"`
	Status      string            `db:"status" json:"status"`
	Created_at  time.Time         `db:"created_at" json:"-"`
	Updated_at  domain.NullTime   `db:"updated_at" json:"-"`
	Deleted_at  domain.NullTime   `db:"deleted_at" json:"-"`
}