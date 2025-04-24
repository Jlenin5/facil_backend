package domain

import "time"

type InventoryMovements struct {
	Id            int         `db:"id" json:"id"`
	Warehouse_Id  int         `db:"warehouse_id" json:"warehouse_id"`
	Warehouse     *Warehouses `json:"warehouse"`
	Product_Id    NullInt     `db:"product_id" json:"product_id"`
	Product       *Products   `json:"product"`
	Movement_Type string      `db:"movement_type" json:"movement_type"`
	Quantity      float64     `db:"quantity" json:"quantity"`
	Reference     NullString  `db:"reference" json:"reference"`
	User_Id       int         `db:"user_id" json:"user_id"`
	User          *Users      `json:"user"`
	Created_at    time.Time   `db:"created_at" json:"-"`
	Updated_at    time.Time   `db:"updated_at" json:"-"`
	Deleted_at    NullTime    `db:"deleted_at" json:"-"`
}