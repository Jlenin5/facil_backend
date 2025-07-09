package domain

type EmployeeReducedData struct {
	Id             int        `json:"id"`
	First_Name     string     `json:"first_name"`
	Second_Name    NullString `json:"second_name"`
	Third_Name     NullString `json:"third_name"`
	Surname        NullString `json:"surname"`
	Second_Surname NullString `json:"second_surname"`
	Warehouse_Id   int        `json:"warehouse_id"`
}

type UserReducedData struct {
	Id          int                  `json:"id"`
	Employee_Id NullInt              `json:"employee_id"`
	Employee    *EmployeeReducedData `json:"employee"`
}

type CustomerReducedData struct {
	Id              int        `json:"id"`
	First_Name      NullString `json:"first_name"`
	Second_Name     NullString `json:"second_name"`
	Third_Name      NullString `json:"third_name"`
	Surname         NullString `json:"surname"`
	Second_Surname  NullString `json:"second_surname"`
	Company_Name    NullString `json:"company_name"`
	Document_Number string     `json:"document_number"`
	Address         NullString `db:"address" json:"address"`
	Phone           NullString `json:"phone"`
}

type CurrencyReducedData struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Code   string `json:"code"`
	Symbol string `json:"symbol"`
}

type CompanyReducedData struct {
	Id      int        `json:"id"`
	Name    string     `json:"name"`
	Ruc     string     `json:"ruc"`
	Email   string     `json:"email"`
	Phone   NullString `json:"phone"`
	Address string     `json:"address"`
}

type BranchOfficeReducedData struct {
	Id         int                 `json:"id"`
	Company_Id int                 `json:"company_id"`
	Company    *CompanyReducedData `json:"company"`
	Name       string              `json:"name"`
}

type WarehouseReducedData struct {
	Id               int                      `json:"id"`
	Branch_Office_Id int                      `json:"branch_office_id"`
	Branch_Office    *BranchOfficeReducedData `json:"branch_office"`
	Name             string                   `json:"name"`
}

type ProductReducedData struct {
	Id    int     `json:"id"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
	Cost  float64 `json:"cost"`
}

type BrandReducedData struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type CategoryReducedData struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type UnitOfMeasurementReduceData struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Shortcut    string `json:"shortcut"`
}

type SaleOrderReducedData struct {
	Id        int    `json:"id"`
	Reference string `json:"reference"`
}

type PurchaseOrderReducedData struct {
	Id        int    `json:"id"`
	Reference string `json:"reference"`
}

type QuoteReducedData struct {
	Id        int    `json:"id"`
	Reference string `json:"reference"`
}

type PlanReducedData struct {
	Id       int        `json:"id"`
	Title    string     `json:"title"`
	Subtitle NullString `json:"subtitle"`
	Price    float64    `json:"price"`
}

type WorkAreaReducedData struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type JobPositionReducedData struct {
	Id           int                  `json:"id"`
	Work_Area_Id int                  `json:"work_area_id"`
	Work_Area    *WorkAreaReducedData `json:"work_area"`
	Name         string               `json:"name"`
}

type SystemReducedData struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type AbsenceTypeReducedData struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

type AttendanceTypeReducedData struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Code string `json:"code"`
}

type ScheduleReducedData struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
