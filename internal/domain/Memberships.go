package domain

import "time"

type Plans struct {
	Id                 int        `db:"id" json:"id"`
	Title              string     `db:"title" json:"title"`
	Subtitle           NullString `db:"subtitle" json:"subtitle"`
	Price              float64    `db:"price" json:"price"`
	Max_Branch_Offices int        `db:"max_branch_offices" json:"max_branch_offices"`
	Max_Warehouses     int        `db:"max_warehouses" json:"max_warehouses"`
	Max_Purchases      int        `db:"max_purchases" json:"max_purchases"`
	Max_Users          int        `db:"max_users" json:"max_users"`
	Max_Products       int        `db:"max_products" json:"max_products"`
	Max_Services       int        `db:"max_services" json:"max_services"`
	Max_Documents      int        `db:"max_documents" json:"max_documents"`
	Status             uint8      `db:"status" json:"status"`
	Created_at         time.Time  `db:"created_at" json:"-"`
	Updated_at         time.Time  `db:"updated_at" json:"-"`
	Deleted_at         NullTime   `db:"deleted_at" json:"-"`
}

type Subscriptions struct {
	Id         int                 `db:"id" json:"id"`
	Company_Id int                 `db:"company_id" json:"company_id"`
	Company    *CompanyReducedData `json:"company"`
	Plan_Id    int                 `db:"plan_id" json:"plan_id"`
	Plan       *Plans              `json:"plan"`
	Start_Date time.Time           `db:"start_date" json:"start_date"`
	End_Date   time.Time           `db:"end_date" json:"end_date"`
	Status     string              `db:"status" json:"status"`
	Created_at time.Time           `db:"created_at" json:"-"`
	Updated_at time.Time           `db:"updated_at" json:"-"`
	Deleted_at NullTime            `db:"deleted_at" json:"-"`
}