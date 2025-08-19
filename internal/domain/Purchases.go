package domain

import (
	"time"
)

type Purchases struct {
	Id                  int                       `db:"id" json:"id"`
	Reference           string                    `db:"reference" json:"reference"`
	Invoice_Number      string                    `db:"invoice_number" json:"invoice_number"`
	Supplier_Id         int                       `db:"supplier_id" json:"supplier_id"`
	Supplier            *Suppliers                `json:"supplier"`
	Warehouse_Id        int                       `db:"warehouse_id" json:"warehouse_id"`
	Warehouse           *Warehouses               `json:"warehouse"`
	Currency_Id         int                       `db:"currency_id" json:"currency_id"`
	Currency            *CurrencyReducedData      `json:"currency"`
	Exchange_Rate       float64                   `db:"exchange_rate" json:"exchange_rate"`
	Purchase_Status     string                    `db:"purchase_status" json:"purchase_status"`
	Purchase_Order_Id   NullInt                   `db:"purchase_order_id" json:"purchase_order_id"`
	Purchase_Order      *PurchaseOrderReducedData `json:"purchase_order"`
	Issue_Date          time.Time                 `db:"issue_date" json:"issue_date"`
	Received_Date       NullTime                  `db:"received_date" json:"received_date"`
	Payment_Date        NullTime                  `db:"payment_date" json:"payment_date"`
	Discount            float64                   `db:"discount" json:"discount"`
	Subtotal            float64                   `db:"subtotal" json:"subtotal"`
	Tax                 float64                   `db:"tax" json:"tax"`
	Total               float64                   `db:"total" json:"total"`
	Total_Paid          float64                   `db:"total_paid" json:"total_paid"`
	Change              float64                   `db:"change" json:"change"`
	Payment_Method_Id   NullInt                   `db:"payment_method_id" json:"payment_method_id"`
	Payment_Method      *PaymentMethods           `json:"payment_method"`
	Created_By          int                       `db:"created_by" json:"created_by"`
	Document_Attachment NullString                `db:"document_attachment" json:"document_attachment"`
	Notes               NullString                `db:"notes" json:"notes"`
	PurchaseDetails     []PurchaseDetails         `json:"purchase_details"`
	Created_at          time.Time                 `db:"created_at" json:"-"`
	Updated_at          NullTime                  `db:"updated_at" json:"-"`
	Deleted_at          NullTime                  `db:"deleted_at" json:"-"`
}

type PurchaseDetails struct {
	Id              int                 `db:"id" json:"id"`
	Purchase_Id     int                 `db:"purchase_id" json:"purchase_id"`
	Product_Id      NullInt             `db:"product_id" json:"product_id"`
	Product         *ProductReducedData `json:"product"`
	Quantity        float64             `db:"quantity" json:"quantity"`
	Discount_Method uint8               `db:"discount_method" json:"discount_method"`
	Discount        NullFloat           `db:"discount" json:"discount"`
	Price           float64             `db:"price" json:"price"`
	Subtotal        float64             `db:"subtotal" json:"subtotal"`
	Total           float64             `db:"total" json:"total"`
	Created_at      time.Time           `db:"created_at" json:"-"`
	Updated_at      NullTime     `db:"updated_at" json:"-"`
	Deleted_at      NullTime            `db:"deleted_at" json:"-"`
}
