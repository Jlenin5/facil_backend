package domain

import (
	"time"
)

type Sales struct {
	Id                 int                   `db:"id" json:"id"`
	Document_Type      string                `db:"document_type" json:"document_type"`
	Series             NullString            `db:"series" json:"series"`
	Number             NullInt               `db:"number" json:"number"`
	Bill               NullString            `db:"bill" json:"bill"`
	Issue_Date         time.Time             `db:"issue_date" json:"issue_date"`
	Warehouse_Id       int                   `db:"warehouse_id" json:"warehouse_id"`
	Warehouse          *WarehouseReducedData `json:"warehouse"`
	Customer_Id        int                   `db:"customer_id" json:"customer_id"`
	Customer           *CustomerReducedData  `json:"customer"`
	Currency_Id        int                   `db:"currency_id" json:"currency_id"`
	Currency           *CurrencyReducedData  `json:"currency"`
	User_Id            int                   `db:"user_id" json:"user_id"`
	User               *UserReducedData      `json:"user"`
	Exchange_Rate      float64               `db:"exchange_rate" json:"exchange_rate"`
	Discount           NullFloat             `db:"discount" json:"discount"`
	Subtotal           float64               `db:"subtotal" json:"subtotal"`
	Total              float64               `db:"total" json:"total"`
	Total_Paid         float64               `db:"total_paid" json:"total_paid"`
	Change             float64               `db:"change" json:"change"`
	Sale_Status        string                `db:"sale_status" json:"sale_status"`
	Payment_Method_Id  int                   `db:"payment_method_id" json:"payment_method_id"`
	Payment_Method     *PaymentMethods       `json:"payment_method"`
	Sale_Order_Id      NullInt               `db:"sale_order_id" json:"sale_order_id"`
	Sale_Order         *SaleOrderReducedData `json:"sale_order"`
	Tax_Identification NullString            `db:"tax_identification" json:"tax_identification"`
	Retention          NullFloat             `db:"retention" json:"retention"`
	Perception         NullFloat             `db:"perception" json:"perception"`
	SaleDetails        []SaleDetails         `json:"sale_details"`
	Created_at         time.Time             `db:"created_at" json:"-"`
	Updated_at         time.Time             `db:"updated_at" json:"-"`
	Deleted_at         NullTime              `db:"deleted_at" json:"-"`
}

type SaleDetails struct {
	Id              int                 `db:"id" json:"id"`
	Product_Name    string              `db:"product_name" json:"product_name"`
	Sale_Id         int                 `db:"sale_id" json:"sale_id"`
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
