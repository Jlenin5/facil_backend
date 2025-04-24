package domain

type Account struct {
	Id             int            `json:"id"`
	Name           string         `json:"name"`
	Username       string         `json:"username"`
	Password       string         `json:"password"`
	Company_Name   string         `json:"company_name"`
	Avatar         string         `json:"avatar"`
	Email          string         `json:"email"`
	Phone          NullString     `json:"phone"`
	Employee_Id    NullInt        `json:"employee_id"`
	Subscription   *Subscriptions `json:"subscription,omitempty"`
}

type ChangePassword struct {
	Id               int    `json:"id"`
	Current_Password string `json:"current_password"`
	New_Password     string `json:"new_password"`
}

type PlanBilling struct {
	Id   int `json:"id"`
	Plan_Id int `json:"plan_id"`
	Plan Plans `json:"plan"`
}