package domain

import (
	"time"
)

type CashRegisters struct {
	Id             int                   `db:"id" json:"id"`
	Warehouse_Id   int                   `db:"warehouse_id" json:"warehouse_id"`
	Warehouse      *WarehouseReducedData `json:"warehouse"`
	User_Open_Id   int                   `db:"user_open_id" json:"user_open_id"`
	User_Open      *UserReducedData      `json:"user_open"`
	User_Close_Id  NullInt                   `db:"user_close_id" json:"user_close_id"`
	User_Close     *UserReducedData      `json:"user_close"`
	Opening_Date   time.Time             `db:"opening_date" json:"opening_date"`
	Closing_Date   NullTime              `db:"closing_date" json:"closing_date"`
	Initial_Amount float64               `db:"initial_amount" json:"initial_amount"`
	Closing_Amount NullFloat             `db:"closing_amount" json:"closing_amount"`
	Difference     NullFloat             `db:"difference" json:"difference"`
	Status         string                `db:"status" json:"status"`
	Created_at     time.Time             `db:"created_at" json:"-"`
	Updated_at     time.Time             `db:"updated_at" json:"-"`
	Deleted_at     NullTime              `db:"deleted_at" json:"-"`
}

type CashMovements struct {
	Id                int              `db:"id" json:"id"`
	Cash_Register_Id  int              `db:"cash_register_id" json:"cash_register_id"`
	Cash_Register     *CashRegisters   `json:"cash_register"`
	Movement_Type     string           `db:"movement_type" json:"movement_type"`
	Payment_Method_Id int              `db:"payment_method_id" json:"payment_method_id"`
	Payment_Method    *PaymentMethods  `json:"payment_method"`
	Amount            float64          `db:"amount" json:"amount"`
	Description       string           `db:"description" json:"description"`
	User_Id           int              `db:"user_id" json:"user_id"`
	User              *UserReducedData `json:"user"`
	Created_at        time.Time        `db:"created_at" json:"-"`
	Updated_at        time.Time        `db:"updated_at" json:"-"`
}