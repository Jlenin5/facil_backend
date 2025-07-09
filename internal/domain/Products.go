package domain

import (
	"encoding/json"
	"time"
)

type Products struct {
	Id                  int                          `db:"id" json:"id"`
	Name                string                       `db:"name" json:"name"`
	Brand_Id            NullInt                      `db:"brand_id" json:"brand_id"`
	Brand               *BrandReducedData            `json:"brand"`
	Handle              NullString                   `db:"handle" json:"handle"`
	Description         NullString                   `db:"description" json:"description"`
	Tags                json.RawMessage              `db:"tags" json:"tags"`
	Categories          []CategoryReducedData        `json:"categories"`
	UnitOfMeasurementId NullInt                      `db:"unit_of_measurement_id" json:"unit_of_measurement_id"`
	UnitOfMeasurement   *UnitOfMeasurementReduceData `db:"unit_of_measurement" json:"unit_of_measurement"`
	FeaturedImage       NullString                   `db:"featured_image" json:"featured_image"`
	Images              json.RawMessage              `db:"images" json:"images"`
	Prices_cf           json.RawMessage              `db:"prices_cf" json:"prices_cf"`
	Prices_sf           json.RawMessage              `db:"prices_sf" json:"prices_sf"`
	Prices_box          json.RawMessage              `db:"prices_box" json:"prices_box"`
	Quantity_In_Box     NullInt                      `db:"quantity_in_box" json:"quantity_in_box"`
	Featured_Pcf        NullFloat                    `db:"featured_pcf" json:"featured_pcf"`
	Featured_Psf        NullFloat                    `db:"featured_psf" json:"featured_psf"`
	Featured_Pbox       NullFloat                    `db:"featured_pbox" json:"featured_pbox"`
	Cost                float64                      `db:"cost" json:"cost"`
	TaxRate             NullFloat                    `db:"tax_rate" json:"tax_rate"`
	Quantity            float64                      `db:"quantity" json:"quantity"`
	SKU                 NullString                   `db:"sku" json:"sku"`
	Width               float64                      `db:"width" json:"width"`
	Height              float64                      `db:"height" json:"height"`
	Depth               float64                      `db:"depth" json:"depth"`
	Liters              float64                      `db:"liters" json:"liters"`
	Weight              float64                      `db:"weight" json:"weight"`
	Barcode             NullString                   `db:"barcode" json:"barcode"`
	Rating              float64                      `db:"rating" json:"rating"`
	ExtraShippingFee    float64                      `db:"extra_shipping_fee" json:"extra_shipping_fee"`
	Status              uint8                        `db:"status" json:"status"`
	Created_By          int                          `db:"created_by" json:"created_by"`
	Updated_By          NullInt                      `db:"updated_by" json:"updated_by"`
	Stock               int                          `json:"stock"`
	Booking             []Booking                    `json:"booking"`
	Created_at          time.Time                    `db:"created_at" json:"created_at"`
	Updated_at          NullTime                     `db:"updated_at" json:"updated_at"`
	Deleted_at          NullTime                     `db:"deleted_at" json:"-"`
}

type Brands struct {
	Id          int       `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Description string    `db:"description" json:"description"`
	Logo_url    string    `db:"logo_url" json:"logo_url"`
	Website_url string    `db:"website_url" json:"website_url"`
	Status      uint8     `db:"status" json:"status"`
	Created_By  int       `db:"created_by" json:"created_by"`
	Updated_By  NullInt   `db:"updated_by" json:"updated_by"`
	Created_at  time.Time `db:"created_at" json:"-"`
	Updated_at  time.Time `db:"updated_at" json:"-"`
	Deleted_at  NullTime  `db:"deleted_at" json:"-"`
}

type Categories struct {
	Id          int       `db:"id" json:"id"`
	Name        string    `db:"name" json:"name"`
	Description string    `db:"description" json:"description"`
	Status      uint8     `db:"status" json:"status"`
	Created_By  int       `db:"created_by" json:"created_by"`
	Updated_By  NullInt   `db:"updated_by" json:"updated_by"`
	Created_at  time.Time `db:"created_at" json:"-"`
	Updated_at  time.Time `db:"updated_at" json:"-"`
	Deleted_at  NullTime  `db:"deleted_at" json:"-"`
}

type ProductCategories struct {
	Id          int `db:"id" json:"id"`
	Product_Id  int `db:"product_id" json:"product_id"`
	Category_Id int `db:"category_id" json:"category_id"`
}

type UnitsOfMeasurement struct {
	Id          int        `db:"id" json:"id"`
	Name        string     `db:"name" json:"name"`
	Description string     `db:"description" json:"description"`
	Shortcut    NullString `db:"shortcut" json:"shortcut"`
	Status      uint8      `db:"status" json:"status"`
	Created_By  int        `db:"created_by" json:"created_by"`
	Updated_By  NullInt    `db:"updated_by" json:"updated_by"`
	Created_at  time.Time  `db:"created_at" json:"-"`
	Updated_at  time.Time  `db:"updated_at" json:"-"`
	Deleted_at  NullTime   `db:"deleted_at" json:"-"`
}

type Stock struct {
	Warehouse_Id NullInt `json:"warehouse_id"`
	Total        float64 `json:"total"`
}

type Booking struct {
	Warehouse_Id NullInt `json:"warehouse_id"`
	Total        float64 `json:"total"`
}
type PriceProducts struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

type PriceType string
type PriceMargins struct {
	Price_1 float64 `json:"price_1"`
	Price_2 float64 `json:"price_2"`
	Price_3 float64 `json:"price_3"`
	Price_4 float64 `json:"price_4"`
	Price_5 float64 `json:"price_5"`
	Price_6 float64 `json:"price_6"`
	Price_7 float64 `json:"price_7"`
	Price_8 float64 `json:"price_8"`
	Price_9 float64 `json:"price_9"`
	Price_10 float64 `json:"price_10"`
	Price_11 float64 `json:"price_11"`
	Price_12 float64 `json:"price_12"`
	Price_13 float64 `json:"price_13"`
	Price_14 float64 `json:"price_14"`
	Price_15 float64 `json:"price_15"`
	Price_16 float64 `json:"price_16"`
	Price_17 float64 `json:"price_17"`
	Price_18 float64 `json:"price_18"`
	Price_19 float64 `json:"price_19"`
	Price_20 float64 `json:"price_20"`
}