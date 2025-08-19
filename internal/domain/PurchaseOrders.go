package domain

import (
	"time"
)

type PurchaseOrders struct {
	Id                   int                    `db:"id" json:"id"`
	Reference            string                 `db:"reference" json:"reference"`
	Warehouse_Id         int                    `db:"warehouse_id" json:"warehouse_id"`
	Warehouse            *Warehouses            `json:"warehouse"`
	Supplier_Id          int                    `db:"supplier_id" json:"supplier_id"`
	Supplier             *Suppliers             `json:"supplier"`
	Currency_Id          int                    `db:"currency_id" json:"currency_id"`
	Currency             *CurrencyReducedData   `json:"currency"`
	Exchange_Rate        float64                `db:"exchange_rate" json:"exchange_rate"`
	Discount             NullFloat              `db:"discount" json:"discount"`
	Issue_Date           time.Time              `db:"issue_date" json:"issue_date"`
	Tax                  float64                `db:"tax" json:"tax"`
	Subtotal             float64                `db:"subtotal" json:"subtotal"`
	Total                float64                `db:"total" json:"total"`
	Document_Attachment  NullString             `db:"document_attachment" json:"document_attachment"`
	Order_Status         string                 `db:"order_status" json:"order_status"`
	Approval_Date        NullTime               `db:"approval_date" json:"approval_date"`
	Created_By           int                    `db:"created_by" json:"created_by"`
	Approved_By          NullInt                `db:"approved_by" json:"approved_by"`
	Migrate_Purchase     uint8                  `db:"migrate_purchase" json:"migrate_purchase"`
	PurchaseOrderDetails []PurchaseOrderDetails `json:"purchase_order_details"`
	Notes                NullString             `db:"notes" json:"notes"`
	Created_at           time.Time              `db:"created_at" json:"created_at"`
	Updated_at           NullTime               `db:"updated_at" json:"updated_at"`
	Deleted_at           NullTime               `db:"deleted_at" json:"-"`
}

type PurchaseOrderDetails struct {
	Id                int                 `db:"id" json:"id"`
	Purchase_Order_Id int                 `db:"purchase_order_id" json:"purchase_order_id"`
	Product_Id        NullInt             `db:"product_id" json:"product_id"`
	Product           *ProductReducedData `json:"product"`
	Quantity          float64             `db:"quantity" json:"quantity"`
	Price             float64             `db:"price" json:"price"`
	Discount_Method   uint8               `db:"discount_method" json:"discount_method"`
	Discount          NullFloat           `db:"discount" json:"discount"`
	Subtotal          float64             `db:"subtotal" json:"subtotal"`
	Total             float64             `db:"total" json:"total"`
	Created_at        time.Time           `db:"created_at" json:"-"`
	Updated_at        time.Time           `db:"updated_at" json:"-"`
	Deleted_at        NullTime            `db:"deleted_at" json:"-"`
}