package domain

import (
	"time"
)

type Taxes struct {
	Id          int       `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Description string    `db:"description" json:"description"`
	Rate        float64   `db:"rate" json:"rate"`
	Tax_type    uint8     `db:"tax_type" json:"tax_type"`
	Status      uint8     `db:"status" json:"status"`
	Created_at  time.Time `db:"created_at" json:"-"`
	Updated_at  time.Time `db:"updated_at" json:"-"`
	Deleted_at  NullTime  `db:"deleted_at" json:"-"`
}