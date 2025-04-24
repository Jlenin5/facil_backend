package domain

import (
	"time"
)

type Companies struct {
	Id         int              `db:"id" json:"id"`
	Name       string           `db:"name" json:"name"`
	Logo       NullString       `db:"logo" json:"logo"`
	Ruc        string           `db:"ruc" json:"ruc"`
	Email      string           `db:"email" json:"email"`
	Phone      NullString       `db:"phone" json:"phone"`
	Web_Site   NullString       `db:"web_site" json:"web_site"`
	Address    string           `db:"address" json:"address"`
	Status     uint8            `db:"status" json:"status"`
	Plan       *PlanReducedData `json:"plan"`
	Created_at time.Time        `db:"created_at" json:"-"`
	Updated_at time.Time        `db:"updated_at" json:"-"`
	Deleted_at NullTime         `db:"deleted_at" json:"-"`
}

type BranchOffices struct {
	Id          int        `db:"id" json:"id"`
	Company_Id  int        `db:"company_id" json:"company_id"`
	Company     *Companies `json:"company"`
	Name        string     `db:"name" json:"name"`
	Description NullString `db:"description" json:"description"`
	Address     string     `db:"address" json:"address"`
	Phone       NullString `db:"phone" json:"phone"`
	Status      uint8      `db:"status" json:"status"`
	Created_at  time.Time  `db:"created_at" json:"-"`
	Updated_at  time.Time  `db:"updated_at" json:"-"`
	Deleted_at  NullTime   `db:"deleted_at" json:"-"`
}

type Warehouses struct {
	Id               int            `db:"id" json:"id"`
	Company_Id       NullInt        `db:"company_id" json:"company_id"`
	Branch_Office_Id int            `db:"branch_office_id" json:"branch_office_id"`
	Branch_Office    *BranchOffices `json:"branch_office"`
	Name             string         `db:"name" json:"name"`
	Description      NullString     `db:"description" json:"description"`
	Address          string         `db:"address" json:"address"`
	Phone            NullString     `db:"phone" json:"phone"`
	Status           uint8          `db:"status" json:"status"`
	Created_at       time.Time      `db:"created_at" json:"-"`
	Updated_at       time.Time      `db:"updated_at" json:"-"`
	Deleted_at       NullTime       `db:"deleted_at" json:"-"`
}