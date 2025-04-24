package domain

import (
	"time"
)

type PurchaseOrders struct {
	Id                   int                    `db:"id" json:"id"`
	Reference            string                 `db:"reference" json:"reference"`
	Description          string                 `db:"description" json:"description"`
	Warehouse_Id         int                    `db:"warehouse_id" json:"warehouse_id"`
	Warehouse            *Warehouses            `json:"warehouse"`
	Supplier_Id          int                    `db:"supplier_id" json:"supplier_id"`
	Supplier             *Suppliers             `json:"supplier"`
	Supplier_Document    string                 `db:"supplier_document" json:"supplier_document"`
	Exchange_Rate        float64                `db:"exchange_rate" json:"exchange_rate"`
	Discount             NullFloat              `db:"discount" json:"discount"`
	User_Id              int                    `db:"user_id" json:"user_id"`
	User                 *Users                 `json:"user"`
	Issue_Date           time.Time              `db:"issue_date" json:"issue_date"`
	Tax                  float64                `db:"tax" json:"tax"`
	Subtotal             float64                `db:"subtotal" json:"subtotal"`
	Total                float64                `db:"total" json:"total"`
	Order_Status         string                 `db:"order_status" json:"order_status"`
	Date_Approved        NullTime               `db:"date_approved" json:"date_approved"`
	Migrate_Purchase     uint8                  `db:"migrate_purchase" json:"migrate_purchase"`
	PurchaseOrderDetails []PurchaseOrderDetails `json:"purchase_order_details"`
	Created_at           time.Time              `db:"created_at" json:"-"`
	Updated_at           time.Time              `db:"updated_at" json:"-"`
	Deleted_at           NullTime               `db:"deleted_at" json:"-"`
}

type PurchaseOrderDetails struct {
	Id                int       `db:"id" json:"id"`
	Purchase_Order_Id int       `db:"purchase_order_id" json:"purchase_order_id"`
	Product_Id        NullInt   `db:"product_id" json:"product_id"`
	Quantity          float64   `db:"quantity" json:"quantity"`
	Price             float64   `db:"price" json:"price"`
	Total             float64   `db:"total" json:"total"`
	Created_at        time.Time `db:"created_at" json:"-"`
	Updated_at        time.Time `db:"updated_at" json:"-"`
	Deleted_at        NullTime  `db:"deleted_at" json:"-"`
}