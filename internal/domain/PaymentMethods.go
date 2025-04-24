package domain

import (
	"time"
)

type PaymentMethods struct {
	Id          int        `db:"id" json:"id"`
	Name        string     `db:"name" json:"name"`
	Description NullString `db:"description" json:"description"`
	Status      uint8      `db:"status" json:"status"`
	Created_at  time.Time  `db:"created_at" json:"-"`
	Updated_at  time.Time  `db:"updated_at" json:"-"`
	Deleted_at  NullTime   `db:"deleted_at" json:"-"`
}