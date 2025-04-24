package domain

import (
	"time"
)

type Customers struct {
	Id                 int              `db:"id" json:"id"`
	First_Name         NullString       `db:"first_name" json:"first_name"`
	Second_Name        NullString       `db:"second_name" json:"second_name"`
	Third_Name         NullString       `db:"third_name" json:"third_name"`
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