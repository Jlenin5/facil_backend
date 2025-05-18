package domain

import (
	"time"
)

type Currencies struct {
	Id          int        `db:"id" json:"id"`
	Name        string     `db:"name" json:"name"`
	Description NullString `db:"description" json:"description"`
	Code        string     `db:"code" json:"code"`
	Symbol      string     `db:"symbol" json:"symbol"`
	Status      uint8      `db:"status" json:"status"`
	Created_at  time.Time  `db:"created_at" json:"-"`
	Updated_at  time.Time  `db:"updated_at" json:"-"`
	Deleted_at  NullTime   `db:"deleted_at" json:"-"`
}

type ExchangeRates struct {
	Id                 int                  `db:"id" json:"id"`
	Base_Currency_Id   int                  `db:"base_currency_id" json:"base_currency_id"`
	Base_Currency      *CurrencyReducedData `json:"base_currency"`
	Target_Currency_Id int                  `db:"target_currency_id" json:"target_currency_id"`
	Target_Currency    *CurrencyReducedData `json:"target_currency"`
	Exchange_Rate      float64              `db:"exchange_rate" json:"exchange_rate"`
	Created_By         int                  `db:"created_by" json:"created_by"`
	Updated_By         NullInt              `db:"updated_by" json:"updated_by"`
	Created_at         time.Time            `db:"created_at" json:"created_at"`
	Updated_at         time.Time            `db:"updated_at" json:"-"`
	Deleted_at         NullTime             `db:"deleted_at" json:"-"`
}