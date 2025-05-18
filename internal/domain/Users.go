package domain

import "time"
import "encoding/json"

type Users struct {
	Id           int                 `db:"id" json:"id"`
	Company_Id   NullInt             `db:"company_id" json:"company_id"`
	Company      *CompanyReducedData `json:"company"`
	Password     string              `db:"password" json:"password"`
	Role_Id      NullInt             `db:"role_id" json:"role_id"`
	Role         *Roles              `json:"role"`
	Username     string              `db:"username" json:"username"`
	Employee_Id  NullInt             `db:"employee_id" json:"employee_id"`
	Employee     *Employees          `json:"employee"`
	Avatar       NullString          `db:"avatar" json:"avatar"`
	Email        string              `db:"email" json:"email"`
	Settings     json.RawMessage         `db:"settings" json:"settings"` // Para manejar JSONB
	Shortcuts    json.RawMessage         `db:"shortcuts" json:"shortcuts"`
	Status       uint8               `db:"status" json:"status"`
	Subscription *Subscriptions      `json:"subscription,omitempty"`
	Created_at   time.Time           `db:"created_at" json:"-"`
	Updated_at   time.Time           `db:"updated_at" json:"-"`
	Deleted_at   NullTime            `db:"deleted_at" json:"-"`
}

type Roles struct {
	Id    int    `db:"id" json:"id"`
	Name  string `db:"name" json:"name"`
	Description  string `db:"description" json:"description"`
}

type Permiissions struct {
	Id    int    `db:"id" json:"id"`
	Name  string `db:"name" json:"name"`
	Description  string `db:"description" json:"description"`
}