package domain

import (
	"time"
)

type Services struct {
	Id           int       `db:"id" json:"id"`
	Name         string    `db:"name" json:"name"`
	Description  string    `db:"description" json:"description"`
	Price        float64   `db:"price" json:"price"`
	Cost         float64   `db:"cost" json:"cost"`
	Duration     float64   `db:"duration" json:"duration"`
	Requirements float64   `db:"requirements" json:"requirements"`
	Rating       NullFloat `db:"rating" json:"rating"`
	Status       uint8     `db:"status" json:"status"`
	Created_at   time.Time `db:"created_at" json:"-"`
	Updated_at   time.Time `db:"updated_at" json:"-"`
	Deleted_at   NullTime  `db:"deleted_at" json:"-"`
}