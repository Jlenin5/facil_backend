package domain

import (
	"time"
)

// Summary Dashboard
type SummaryRanges struct {
	DY  string `json:"DY"`
	DT  string `json:"DT"`
	DTM string `json:"DTM"`
}

type SummaryExtra struct {
	Name  string        `json:"name"`
	Count SummaryRanges `json:"count"`
}

type SummaryData struct {
	Name  string         `json:"name"`
	Count map[string]int `json:"count"`
	Extra SummaryExtra   `json:"extra"`
}

type DashboardSummary struct {
	Ranges       []string    `json:"ranges"`
	CurrentRange string      `json:"currentRange"`
	Data         SummaryData `json:"data"`
	Detail       string      `json:"detail"`
}

// Overdue Dashboard
type OverdueExtra struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type OverdueData struct {
	Name  string       `json:"name"`
	Count int          `json:"count"`
	Extra OverdueExtra `json:"extra"`
}

type DashboardOverdue struct {
	Title  string      `json:"title"`
	Data   OverdueData `json:"data"`
	Detail string      `json:"detail"`
}

// Issues Dashboard
type IssuesExtra struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type IssuesData struct {
	Name  string      `json:"name"`
	Count int         `json:"count"`
	Extra IssuesExtra `json:"extra"`
}

type DashboardIssues struct {
	Title  string     `json:"title"`
	Data   IssuesData `json:"data"`
	Detail string     `json:"detail"`
}

// Features Dashboard
type FeaturesExtra struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type FeaturesData struct {
	Name  string        `json:"name"`
	Count int           `json:"count"`
	Extra FeaturesExtra `json:"extra"`
}

type DashboardFeatures struct {
	Title  string       `json:"title"`
	Data   FeaturesData `json:"data"`
	Detail string       `json:"detail"`
}

// GithubIssues Dashboard
type DashboardGithubIssues struct {
	Overview GithubIssuesOverview `json:"overview"`
	Ranges   GithubIssuesRanges   `json:"ranges"`
	Labels   []string             `json:"labels"`
	Series   GithubIssuesSeries   `json:"series"`
}

type GithubIssuesOverview struct {
	Sales     GithubIssuesWeek `json:"sales"`
	Purchases GithubIssuesWeek `json:"purchases"`
}

type GithubIssuesWeek struct {
	NewIssues    int64 `json:"new_issues"`
	ClosedIssues int64 `json:"closed_issues"`
	Fixed        int64 `json:"fixed"`
	WontFix      int64 `json:"wont_fix"`
	// Puedes agregar más campos si es necesario
}

type GithubIssuesRanges struct {
	Sales     string `json:"sales"`
	Purchases string `json:"purchases"`
}

type GithubIssuesLabels struct {
	Mon string `json:"mon"`
	Tue string `json:"tue"`
	Wed string `json:"wed"`
	Thu string `json:"thu"`
	Fri string `json:"fri"`
	Sat string `json:"sat"`
	Sun string `json:"sun"`
}

type GithubIssuesSeries struct {
	Sales     []GithubIssuesThisWeek `json:"sales"`
	Purchases []GithubIssuesLastWeek `json:"purchases"`
}

type GithubIssuesThisWeek struct {
	Name string  `json:"name"`
	Type string  `json:"type"`
	Data []int `json:"data"`
}

type GithubIssuesLastWeek struct {
	Name string  `json:"name"`
	Type string  `json:"type"`
	Data []int `json:"data"`
}

// Task Distribution
type TaskDistributionSeries struct {
	ThisWeek []int `json:"this-week"`
	LastWeek []int `json:"last-week"`
}

type TaskDistributionWeek struct {
	New       int `json:"new"`
	Completed int `json:"completed"`
}

type TaskDistributionOverview struct {
	ThisWeek TaskDistributionWeek `json:"this-week"`
	LastWeek TaskDistributionWeek `json:"last-week"`
}

type TaskDistributionRanges struct {
	ThisWeek string `json:"this-week"`
	LastWeek string `json:"last-week"`
}

type DashboardTaskDistribution struct {
	Ranges   TaskDistributionRanges   `json:"ranges"`
	Overview TaskDistributionOverview `json:"overview"`
	Labels   []string                 `json:"labels"`
	Series   TaskDistributionSeries   `json:"series"`
}

// Schedule Dashboard
type ScheduleSeriesTime struct {
	Title    string  `json:"title"`
	Time     string  `json:"time"`
	Location *string `json:"location"`
}

type ScheduleSeries struct {
	Today    []ScheduleSeriesTime `json:"today"`
	Tomorrow []ScheduleSeriesTime `json:"tomorrow"`
}

type ScheduleRanges struct {
	Today    string `json:"today"`
	Tomorrow string `json:"tomorrow"`
}

type DashboardSchedule struct {
	Ranges ScheduleRanges `json:"ranges"`
	Series ScheduleSeries `json:"series"`
}

// Widgets Dashboard
type DashboardWidgets struct {
	Summary          DashboardSummary          `json:"summary"`
	Overdue          DashboardSummary          `json:"overdue"`
	Issues           DashboardSummary          `json:"issues"`
	Features         DashboardSummary          `json:"features"`
	GithubIssues     DashboardGithubIssues     `json:"githubIssues"`
	TaskDistribution DashboardTaskDistribution `json:"taskDistribution"`
	Schedule         DashboardSchedule         `json:"schedule"`
	PurchaseOrders   GetDashboardProjects      `json:"purchase_orders"`
	OverdueOrders    OverdueWidget             `json:"overdue_orders"`
	RecentOrders     RecentOrdersWidget        `json:"recent_orders"`
}

type ProductFlowDay struct {
	WeekDay          string `db:"weekday"`   // Mon, Tue, …
	SalesPaid        float64  `db:"sales_paid"`
	SalesCanceled    float64  `db:"sales_canceled"`
	PurchasesPaid    float64  `db:"purchases_paid"`
	PurchasesCanceled float64  `db:"purchases_canceled"`
}

type GetDashboardProjects struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
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
