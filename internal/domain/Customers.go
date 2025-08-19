package domain

import (
	"time"
)

type Customers struct {
	Id                 int              `db:"id" json:"id"`
	Names              NullString       `db:"names" json:"names"`
	Surname            NullString       `db:"surname" json:"surname"`
	Second_Surname     NullString       `db:"second_surname" json:"second_surname"`
	Company_Name       NullString       `db:"company_name" json:"company_name"`
	Document_Type      string           `db:"document_type" json:"document_type"`
	Document_Number    string           `db:"document_number" json:"document_number"`
	Email              string           `db:"email" json:"email"`
	Address            NullString       `db:"address" json:"address"`
	Phone              NullString       `db:"phone" json:"phone"`
	Status             uint8            `db:"status" json:"status"`
	Created_at         time.Time        `db:"created_at" json:"-"`
	Updated_at         time.Time        `db:"updated_at" json:"-"`
	Deleted_at         NullTime         `db:"deleted_at" json:"-"`
}