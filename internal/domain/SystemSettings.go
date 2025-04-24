package domain

import "time"

type Systems struct {
	Id          int        `db:"id" json:"id"`
	Name        string     `db:"name" json:"name"`
	Description NullString `db:"description" json:"description"`
	Created_at  time.Time  `db:"created_at" json:"-"`
	Updated_at  time.Time  `db:"updated_at" json:"-"`
}

type Keys struct {
	Id           int                `db:"id" json:"id"`
	System_Id    NullInt            `db:"system_id" json:"system_id"`
	System       *SystemReducedData `json:"system"`
	Name         string             `db:"name" json:"name"`
	Description  NullString         `db:"description" json:"description"`
	Config_Key   string             `db:"config_key" json:"config_key"`
	Config_Value string             `db:"config_value" json:"config_value"`
	Created_By   NullInt            `db:"created_by" json:"created_by"`
	Updated_By   NullInt            `db:"updated_by" json:"updated_by"`
	Status       uint8              `db:"status" json:"status"`
	Created_at   time.Time          `db:"created_at" json:"-"`
	Updated_at   time.Time          `db:"updated_at" json:"-"`
}