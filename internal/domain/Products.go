package domain

import (
	"encoding/json"
	"time"
)

type Products struct {
	Id               int                   `db:"id" json:"id"`
	Name             string                `db:"name" json:"name"`
	Brand_Id         NullInt               `db:"brand_id" json:"brand_id"`
	Brand            *BrandReducedData     `json:"brand"`
	Handle           NullString            `db:"handle" json:"handle"`
	Description      NullString            `db:"description" json:"description"`
	Tags             json.RawMessage       `db:"tags" json:"tags"`
	Categories       []CategoryReducedData `json:"categories"`
	FeaturedImage    NullString            `db:"featured_image" json:"featured_image"`
	Images           json.RawMessage       `db:"images" json:"images"`
	Prices_cf        json.RawMessage       `db:"prices_cf" json:"prices_cf"`
	Prices_sf        json.RawMessage       `db:"prices_sf" json:"prices_sf"`
	Prices_box       json.RawMessage       `db:"prices_box" json:"prices_box"`
	Cost             json.RawMessage       `db:"cost" json:"cost"`
	TaxRate          NullFloat             `db:"tax_rate" json:"tax_rate"`
	Quantity         float64               `db:"quantity" json:"quantity"`
	SKU              NullString            `db:"sku" json:"sku"`
	Width            float64               `db:"width" json:"width"`
	Height           float64               `db:"height" json:"height"`
	Depth            float64               `db:"depth" json:"depth"`
	Liters           float64               `db:"liters" json:"liters"`
	Weight           float64               `db:"weight" json:"weight"`
	Barcode          NullString            `db:"barcode" json:"barcode"`
	Rating           float64               `db:"rating" json:"rating"`
	ExtraShippingFee float64               `db:"extra_shipping_fee" json:"extra_shipping_fee"`
	Status           uint8                 `db:"status" json:"status"`
	Created_By       int                   `db:"created_by" json:"created_by"`
	Updated_By       NullInt               `db:"updated_by" json:"updated_by"`
	Stock            int                   `json:"stock"`
	Booking          []Booking             `json:"booking"`
	Created_at       time.Time             `db:"created_at" json:"-"`
	Updated_at       time.Time             `db:"updated_at" json:"-"`
	Deleted_at       NullTime              `db:"deleted_at" json:"-"`
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

// Unit of Measurement structure
type UnitOfMeasurement struct {
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

// ProductUnitsOfMeasurement relationship table
type ProductUnitOfMeasurement struct {
	Id                  int `db:"id" json:"id"`
	ProductId           int `db:"product_id" json:"productId"`
	UnitOfMeasurementId int `db:"unit_of_measurement_id" json:"unitOfMeasurementId"`
}

type Stock struct {
	Warehouse_Id NullInt `json:"warehouse_id"`
	Total        float64 `json:"total"`
}

type Booking struct {
	Warehouse_Id NullInt `json:"warehouse_id"`
	Total        float64 `json:"total"`
}