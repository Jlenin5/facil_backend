package domain

import (
	"time"
)

type OpportunityTracking struct {
	Id               int        `db:"id" json:"id"`
	Customer_Id      int        `db:"customer_id" json:"customer_id"`
	Customer         *Customers `json:"customer"`
	User_Id          int        `db:"user_id" json:"user_id"`
	User             *Users     `json:"user"`
	Title            string     `db:"title" json:"title"`
	Description      NullString `db:"description" json:"description"`
	Status           string     `db:"status" json:"status"`
	Expected_Revenue NullFloat  `db:"expected_revenue" json:"expected_revenue"`
	Probability      NullInt    `db:"probability" json:"probability"`
	Created_at       time.Time  `db:"created_at" json:"-"`
	Updated_at       time.Time  `db:"updated_at" json:"-"`
	Deleted_at       NullTime   `db:"deleted_at" json:"-"`
}