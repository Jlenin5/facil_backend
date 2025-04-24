package domain

import (
	"time"
)

type Quotes struct {
	Id              int                   `db:"id" json:"id"`
	Reference       string                `db:"reference" json:"reference"`
	Warehouse_Id    int                   `db:"warehouse_id" json:"warehouse_id"`
	Warehouse       *WarehouseReducedData `json:"warehouse"`
	Customer_Id     int                   `db:"customer_id" json:"customer_id"`
	Customer        *CustomerReducedData  `json:"customer"`
	Currency_Id     int                   `db:"currency_id" json:"currency_id"`
	Currency        *CurrencyReducedData  `json:"currency"`
	User_Id         int                   `db:"user_id" json:"user_id"`
	User            *UserReducedData      `json:"user"`
	Issue_Date      time.Time             `db:"issue_date" json:"issue_date"`
	Exchange_Rate   float64               `db:"exchange_rate" json:"exchange_rate"`
	Expiration_Date NullTime              `db:"expiration_date" json:"expiration_date"`
	Approved_By     NullInt               `db:"approved_by" json:"approved_by"`
	Approved_At     NullTime              `db:"approved_at" json:"approved_at"`
	Canceled_By     NullInt               `db:"canceled_by" json:"canceled_by"`
	Canceled_At     NullTime              `db:"canceled_at" json:"canceled_at"`
	Discount        float64               `db:"discount" json:"discount"`
	Subtotal        float64               `db:"subtotal" json:"subtotal"`
	Total           float64               `db:"total" json:"total"`
	Quote_Status    string                `db:"quote_status" json:"quote_status"`
	Migrate_Quote   uint8                 `db:"migrate_quote" json:"migrate_quote"`
	QuoteDetails    []QuoteDetails        `json:"quote_details"`
	Created_at      time.Time             `db:"created_at" json:"-"`
	Updated_at      time.Time             `db:"updated_at" json:"-"`
	Deleted_at      NullTime              `db:"deleted_at" json:"-"`
}

type QuoteDetails struct {
	Id              int                 `db:"id" json:"id"`
	Product_Name    NullString          `db:"product_name" json:"product_name"`
	Quote_Id        int                 `db:"quote_id" json:"quote_id"`
	Product_Id      NullInt             `db:"product_id" json:"product_id"`
	Product         *ProductReducedData `json:"product"`
	Quantity        float64             `db:"quantity" json:"quantity"`
	Discount_Method uint8               `db:"discount_method" json:"discount_method"`
	Discount        float64             `db:"discount" json:"discount"`
	Price           float64             `db:"price" json:"price"`
	Subtotal        float64             `db:"subtotal" json:"subtotal"`
	Total           float64             `db:"total" json:"total"`
	Created_at      time.Time           `db:"created_at" json:"-"`
	Updated_at      time.Time           `db:"updated_at" json:"-"`
	Deleted_at      NullTime            `db:"deleted_at" json:"-"`
}