package domain

import "time"

type Employees struct {
	Id              int                     `db:"id" json:"id"`
	Names           string                  `db:"names" json:"names"`
	Surname         NullString              `db:"surname" json:"surname"`
	Second_Surname  NullString              `db:"second_surname" json:"second_surname"`
	Photo           NullString              `db:"photo" json:"photo"`
	Document_Type   string                  `db:"document_type" json:"document_type"`
	Document_Number string                  `db:"document_number" json:"document_number"`
	Birth_Date      time.Time               `db:"birth_date" json:"birth_date"`
	Gender          string                  `db:"gender" json:"gender"`
	Email           NullString              `db:"email" json:"email"`
	Phone           NullString              `db:"phone" json:"phone"`
	Warehouse_Id    int                     `db:"warehouse_id" json:"warehouse_id"`
	Warehouse       *Warehouses             `json:"warehouse"`
	Address         NullString              `db:"address" json:"address"`
	Hire_Date       time.Time               `db:"hire_date" json:"hire_date"`
	Job_Position_Id NullInt                 `db:"job_position_id" json:"job_position_id"`
	Job_Position    *JobPositionReducedData `json:"job_position"`
	Salary          float64                 `db:"salary" json:"salary"`
	Status          uint8                   `db:"status" json:"status"`
	Created_at      time.Time               `db:"created_at" json:"-"`
	Updated_at      NullTime                `db:"updated_at" json:"-"`
	Deleted_at      NullTime                `db:"deleted_at" json:"-"`
}

type WorkAreas struct {
	Id          int        `db:"id" json:"id"`
	Name        string     `db:"name" json:"name"`
	Description NullString `db:"description" json:"description"`
	Status      uint8      `db:"status" json:"status"`
	Created_at  time.Time  `db:"created_at" json:"-"`
	Updated_at  time.Time  `db:"updated_at" json:"-"`
	Deleted_at  NullTime   `db:"deleted_at" json:"-"`
}

type JobPositions struct {
	Id           int                  `db:"id" json:"id"`
	Work_Area_Id int                  `db:"work_area_id" json:"work_area_id"`
	Work_Area    *WorkAreaReducedData `json:"work_area"`
	Name         string               `db:"name" json:"name"`
	Description  NullString           `db:"description" json:"description"`
	Status       uint8                `db:"status" json:"status"`
	Created_at   time.Time            `db:"created_at" json:"-"`
	Updated_at   time.Time            `db:"updated_at" json:"-"`
	Deleted_at   NullTime             `db:"deleted_at" json:"-"`
}
