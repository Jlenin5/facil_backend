package domain

import "time"

type StockControl struct {
	Id              int         `db:"id" json:"id"`
	Warehouse_Id    int         `db:"warehouse_id" json:"warehouse_id"`
	Warehouse       *Warehouses `json:"warehouse"`
	Product_Id      NullInt     `db:"product_id" json:"product_id"`
	Product         *Products   `json:"product"`
	Current_Stock   int         `db:"current_stock" json:"current_stock"`
	Current_Booking int         `db:"current_booking" json:"current_booking"`
	Min_Stock       NullInt     `db:"min_stock" json:"min_stock"`
	Max_Stock       NullInt     `db:"max_stock" json:"max_stock"`
	Created_at      time.Time   `db:"created_at" json:"-"`
	Updated_at      time.Time   `db:"updated_at" json:"-"`
	Deleted_at      NullTime    `db:"deleted_at" json:"-"`
}