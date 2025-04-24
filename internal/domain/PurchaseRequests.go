package domain

import (
	"time"
)

type PurchaseRequests struct {
	Id                     int                      `db:"id" json:"id"`
	Reference              string                   `db:"reference" json:"reference"`
	Supplier_Id            int                      `db:"supplier_id" json:"supplier_id"`
	Supplier               *Suppliers               `json:"supplier"`
	User_Id                int                      `db:"user_id" json:"user_id"`
	User                   *Users                   `json:"user"`
	Request_Date           time.Time                `db:"request_date" json:"request_date"`
	Expected_Delivery_Date NullTime                 `db:"expected_delivery_date" json:"expected_delivery_date"`
	Status                 string                   `db:"status" json:"status"`
	Priority               string                   `db:"priority" json:"priority"`
	Total_Cost             float64                  `db:"total_cost" json:"total_cost"`
	Approved_By            NullInt                  `db:"approved_by" json:"approved_by"`
	Approval_Date          NullTime                 `db:"approval_date" json:"approval_date"`
	Notes                  NullString               `db:"notes" json:"notes"`
	PurchaseRequestDetails []PurchaseRequestDetails `json:"purchase_request_details"`
	Created_at             time.Time                `db:"created_at" json:"-"`
	Updated_at             time.Time                `db:"updated_at" json:"-"`
	Deleted_at             NullTime                 `db:"deleted_at" json:"-"`
}

type PurchaseRequestDetails struct {
	Id                  int       `db:"id" json:"id"`
	Purchase_Request_Id int       `db:"purchase_request_id" json:"purchase_request_id"`
	Product_Id          NullInt   `db:"product_id" json:"product_id"`
	Product             *Products `json:"product"`
	Quantity            float64   `db:"quantity" json:"quantity"`
	Unit_Price          NullFloat `db:"unit_price" json:"unit_price"`
	Estimated_Cost      NullFloat `db:"estimated_cost" json:"estimated_cost"`
	Subtotal            NullFloat `db:"subtotal" json:"subtotal"`
	Received_Quantity   int       `db:"received_quantity" json:"received_quantity"`
	Status              string    `db:"status" json:"status"`
	Comments            string    `db:"comments" json:"comments"`
	Created_at          time.Time `db:"created_at" json:"-"`
	Updated_at          time.Time `db:"updated_at" json:"-"`
	Deleted_at          NullTime  `db:"deleted_at" json:"-"`
}