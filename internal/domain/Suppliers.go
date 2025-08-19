package domain

import (
	"time"
)

type Suppliers struct {
	Id         int        `db:"id" json:"id"`
	Name       string     `db:"name" json:"name"`
	Ruc        string     `db:"ruc" json:"ruc"`
	Email      NullString `db:"email" json:"email"`
	Phone      NullString `db:"phone" json:"phone"`
	Web_Site   NullString `db:"web_site" json:"web_site"`
	Address    NullString `db:"address" json:"address"`
	Status     uint8      `db:"status" json:"status"`
	Created_at time.Time  `db:"created_at" json:"created_at"`
	Updated_at NullTime   `db:"updated_at" json:"updated_at"`
	Deleted_at NullTime   `db:"deleted_at" json:"-"`
}