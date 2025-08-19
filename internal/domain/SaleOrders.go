package domain

import (
	"time"
)

type SaleOrders struct {
	Id                 int                   `db:"id" json:"id"`
	Reference          string                `db:"reference" json:"reference"`
	Warehouse_Id       int                   `db:"warehouse_id" json:"warehouse_id"`
	Warehouse          *WarehouseReducedData `json:"warehouse"`
	Customer_Id        int                   `db:"customer_id" json:"customer_id"`
	Customer           *CustomerReducedData  `db:"customer" json:"customer"`
	Currency_Id        int                   `db:"currency_id" json:"currency_id"`
	Currency           *CurrencyReducedData  `json:"currency"`
	User_Id            int                   `db:"user_id" json:"user_id"`
	User               *UserReducedData      `json:"user"`
	Issue_Date         time.Time             `db:"issue_date" json:"issue_date"`
	Exchange_Rate      float64               `db:"exchange_rate" json:"exchange_rate"`
	Discount           NullFloat             `db:"discount" json:"discount"`
	Subtotal           float64               `db:"subtotal" json:"subtotal"`
	Total              float64               `db:"total" json:"total"`
	Order_Status       string                `db:"order_status" json:"order_status"`
	Date_Approved      NullTime              `db:"date_approved" json:"date_approved"`
	Migrate_Sale_Order uint8                 `db:"migrate_sale_order" json:"migrate_sale_order"`
	Quote_Id           NullInt               `db:"quote_id" json:"quote_id"`
	Quote              *QuoteReducedData     `json:"quote"`
	SaleOrderDetails   []SaleOrderDetails    `json:"sale_order_details"`
	Created_at         time.Time             `db:"created_at" json:"-"`
	Updated_at         time.Time             `db:"updated_at" json:"-"`
	Deleted_at         NullTime              `db:"deleted_at" json:"-"`
}

type SaleOrderDetails struct {
	Id              int                 `db:"id" json:"id"`
	Product_Name    string              `db:"product_name" json:"product_name"`
	Sale_Order_Id   int                 `db:"sale_order_id" json:"sale_order_id"`
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
