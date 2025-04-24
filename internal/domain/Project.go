package domain

import (
	"time"
)

type DashboardSummary struct {
	TotalSales       float64 `json:"total_sales"`
	TotalPurchases   float64 `json:"total_purchases"`
	LowStockProducts int     `json:"low_stock_products"`
	ActiveCustomers  int     `json:"active_customers"`
	ActiveEmployees  int     `json:"active_employees"`
}




type ProjectResponse struct {
	SaleOrders     Widget             `json:"sale_orders"`
	PurchaseOrders Widget             `json:"purchase_orders"`
	OverdueOrders  OverdueWidget      `json:"overdue_orders"`
	RecentOrders   RecentOrdersWidget `json:"recent_orders"`
}

type Widget struct {
	Status       map[string]string `db:"order_status" json:"order_status"`
	CurrentRange string            `json:"currentRange"`
	Data         WidgetData        `json:"data"`
	Detail       string            `json:"detail"`
}

type OverdueWidget struct {
	Title  string     `json:"title"`
	Data   WidgetData `json:"data"`
	Detail string     `json:"detail"`
}

type RecentOrdersWidget struct {
	Title  string      `json:"title"`
	Orders []OrderInfo `json:"orders"`
}

type WidgetData struct {
	Name  string         `json:"name"`
	Count map[string]int `json:"count"`
	Extra ExtraData      `json:"extra"`
}

type ExtraData struct {
	Name  string         `json:"name"`
	Count map[string]int `json:"count"`
}

type OrderInfo struct {
	ID          int       `db:"id" json:"id"`
	Reference   string    `db:"reference" json:"reference"`
	CustomerID  int       `db:"customer_id" json:"customer_id"`
	IssueDate   time.Time `db:"issue_date" json:"issue_date"`
	Total       float64   `db:"total" json:"total"`
	OrderStatus string    `db:"order_status" json:"order_status"`
}