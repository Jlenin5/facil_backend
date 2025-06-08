package usecase

import (
	"fmt"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/Jlenin5/facil_backend/internal/repository"
)

type ProjectUseCase struct {
	ProjecRepo *repository.ProjectRepository
}

func NewProjectUseCase(ProjecRepo *repository.ProjectRepository) *ProjectUseCase {
	return &ProjectUseCase{ProjecRepo: ProjecRepo}
}

// func (uc *ProjectUseCase) DashboardSummary() ([]domain.DashboardSummary, error) {
// 	return uc.ProjecRepo.DashboardSummary()
// }

func flowsTotals(flows []domain.ProductFlowDay, field string) float64 {
	var total float64
	for _, f := range flows {
		switch field {
		case "sales_paid":
			total += f.SalesPaid
		case "sales_canceled":
			total += f.SalesCanceled
		case "purchases_paid":
			total += f.PurchasesPaid
		case "purchases_canceled":
			total += f.PurchasesCanceled
		}
	}
	return total
}

func (uc *ProjectUseCase) GetDashboardWidgets() (*domain.DashboardWidgets, error) {
	saleStatusCounts, err := uc.ProjecRepo.GetSaleStatusSummary()
	if err != nil {
		return nil, err
	}

	purchaseStatusCounts, err := uc.ProjecRepo.GetPurchaseStatusSummary()
	if err != nil {
		return nil, err
	}

	customerStatusCounts, err := uc.ProjecRepo.GetCustomerStatusSummary()
	if err != nil {
		return nil, err
	}

	supplierStatusCounts, err := uc.ProjecRepo.GetSupplierStatusSummary()
	if err != nil {
		return nil, err
	}

	flows, err := uc.ProjecRepo.GetProductsBySalesAndPurchases()
	if err != nil {
		return nil, err
	}

	var (
		salesSeriesPaid         = make([]int, 7)
		salesSeriesCanceled     = make([]int, 7)
		purchasesSeriesPaid     = make([]int, 7)
		purchasesSeriesCanceled = make([]int, 7)
	)

	location := func(s string) *string { return &s }

	// rellenamos las series según el día
	for _, f := range flows {
		switch f.WeekDay {
		case "Mon":
			salesSeriesPaid[0], salesSeriesCanceled[0] = int(f.SalesPaid), int(f.SalesCanceled)
			purchasesSeriesPaid[0], purchasesSeriesCanceled[0] = int(f.PurchasesPaid), int(f.PurchasesCanceled)
		case "Tue":
			salesSeriesPaid[1], salesSeriesCanceled[1] = int(f.SalesPaid), int(f.SalesCanceled)
			purchasesSeriesPaid[1], purchasesSeriesCanceled[1] = int(f.PurchasesPaid), int(f.PurchasesCanceled)
		case "Wed":
			salesSeriesPaid[2], salesSeriesCanceled[2] = int(f.SalesPaid), int(f.SalesCanceled)
			purchasesSeriesPaid[2], purchasesSeriesCanceled[2] = int(f.PurchasesPaid), int(f.PurchasesCanceled)
		case "Thu":
			salesSeriesPaid[3], salesSeriesCanceled[3] = int(f.SalesPaid), int(f.SalesCanceled)
			purchasesSeriesPaid[3], purchasesSeriesCanceled[3] = int(f.PurchasesPaid), int(f.PurchasesCanceled)
		case "Fri":
			salesSeriesPaid[4], salesSeriesCanceled[4] = int(f.SalesPaid), int(f.SalesCanceled)
			purchasesSeriesPaid[4], purchasesSeriesCanceled[4] = int(f.PurchasesPaid), int(f.PurchasesCanceled)
		case "Sat":
			salesSeriesPaid[5], salesSeriesCanceled[5] = int(f.SalesPaid), int(f.SalesCanceled)
			purchasesSeriesPaid[5], purchasesSeriesCanceled[5] = int(f.PurchasesPaid), int(f.PurchasesCanceled)
		case "Sun":
			salesSeriesPaid[6], salesSeriesCanceled[6] = int(f.SalesPaid), int(f.SalesCanceled)
			purchasesSeriesPaid[6], purchasesSeriesCanceled[6] = int(f.PurchasesPaid), int(f.PurchasesCanceled)
		}
	}

	response := &domain.DashboardWidgets{
		Summary: domain.DashboardSummary{
			Ranges:       []string{"issued", "paid", "unpaid", "pending", "canceled"},
			CurrentRange: "paid",
			Data: domain.SummaryData{
				Name:  "Sales",
				Count: saleStatusCounts,
				Extra: domain.SummaryExtra{
					Name: "Canceled & Pending",
					Count: domain.SummaryRanges{
						DY:  fmt.Sprintf("%d", saleStatusCounts["canceled"]),
						DT:  fmt.Sprintf("%d", saleStatusCounts["pending"]),
						DTM: "-",
					},
				},
			},
		},
		Overdue: domain.DashboardSummary{
			Ranges:       []string{"draft", "received", "partial", "paid", "unpaid", "canceled"},
			CurrentRange: "paid",
			Data: domain.SummaryData{
				Name:  "Purchases",
				Count: purchaseStatusCounts,
				Extra: domain.SummaryExtra{
					Name: "Canceled & Partial",
					Count: domain.SummaryRanges{
						DY:  fmt.Sprintf("%d", saleStatusCounts["canceled"]),
						DT:  fmt.Sprintf("%d", saleStatusCounts["partial"]),
						DTM: "-",
					},
				},
			},
		},
		Issues: domain.DashboardSummary{
			Ranges:       []string{"dni", "ruc", "ce"},
			CurrentRange: "dni",
			Data: domain.SummaryData{
				Name:  "Customers",
				Count: customerStatusCounts,
				Extra: domain.SummaryExtra{
					Name: "Active & Inactive",
					Count: domain.SummaryRanges{
						DY:  fmt.Sprintf("%d", customerStatusCounts["active"]),
						DT:  fmt.Sprintf("%d", customerStatusCounts["inactive"]),
						DTM: "-",
					},
				},
			},
		},
		Features: domain.DashboardSummary{
			Ranges:       []string{"active", "inactive"},
			CurrentRange: "active",
			Data: domain.SummaryData{
				Name:  "Suppliers",
				Count: supplierStatusCounts,
				Extra: domain.SummaryExtra{
					Name: "Active & Inactive",
					Count: domain.SummaryRanges{
						DY:  fmt.Sprintf("%d", supplierStatusCounts["active"]),
						DT:  fmt.Sprintf("%d", supplierStatusCounts["inactive"]),
						DTM: "-",
					},
				},
			},
		},
		GithubIssues: domain.DashboardGithubIssues{
			Overview: domain.GithubIssuesOverview{
				Sales: domain.GithubIssuesWeek{
					NewIssues:    int64(flowsTotals(flows, "sales_paid")),     // total vendidos
					ClosedIssues: int64(flowsTotals(flows, "sales_canceled")), // total cancelados
					Fixed:        int64(flowsTotals(flows, "sales_paid")),     // pagados
					WontFix:      int64(flowsTotals(flows, "sales_canceled")), // cancelados
					// los demás campos si los necesitas…
				},
				Purchases: domain.GithubIssuesWeek{
					NewIssues:    int64(flowsTotals(flows, "purchases_paid")),
					ClosedIssues: int64(flowsTotals(flows, "purchases_canceled")),
					Fixed:        int64(flowsTotals(flows, "purchases_paid")),
					WontFix:      int64(flowsTotals(flows, "purchases_canceled")),
				},
			},
			Ranges: domain.GithubIssuesRanges{
				Sales:     "Sales",
				Purchases: "Purchases",
			},
			Labels: []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"},
			Series: domain.GithubIssuesSeries{
				Sales: []domain.GithubIssuesThisWeek{
					{
						Name: "Sales (paid)",
						Type: "line",
						Data: salesSeriesPaid,
					},
					{
						Name: "Sales (canceled)",
						Type: "column",
						Data: salesSeriesCanceled,
					},
				},
				Purchases: []domain.GithubIssuesLastWeek{
					{
						Name: "Purchases (received)",
						Type: "line",
						Data: purchasesSeriesPaid,
					},
					{
						Name: "Purchases (canceled)",
						Type: "column",
						Data: purchasesSeriesCanceled,
					},
				},
			},
		},
		TaskDistribution: domain.DashboardTaskDistribution{
			Ranges: domain.TaskDistributionRanges{
				ThisWeek: "This Week",
				LastWeek: "Last Week",
			},
			Overview: domain.TaskDistributionOverview{
				ThisWeek: domain.TaskDistributionWeek{
					New:       594,
					Completed: 287,
				},
				LastWeek: domain.TaskDistributionWeek{
					New:       526,
					Completed: 260,
				},
			},
			Labels: []string{"API", "Backend", "Frontend", "Issues"},
			Series: domain.TaskDistributionSeries{
				ThisWeek: []int{15, 20, 38, 27},
				LastWeek: []int{19, 16, 42, 23},
			},
		},
		Schedule: domain.DashboardSchedule{
			Ranges: domain.ScheduleRanges{
				Today:    "Today",
				Tomorrow: "Tomorrow",
			},
			Series: domain.ScheduleSeries{
				Today: []domain.ScheduleSeriesTime{
					{Title: "Group Meeting", Time: "in 32 minutes", Location: location("Conference room 1B")},
					{Title: "Coffee Break", Time: "10:30 AM"},
					{Title: "Public Beta Release", Time: "11:00 AM"},
					{Title: "Lunch", Time: "12:10 PM"},
					{Title: "Dinner with David", Time: "05:30 PM", Location: location("Magnolia")},
					{Title: "Jane's Birthday Party", Time: "07:30 PM", Location: location("Home")},
					{Title: "Overseer's Retirement Party", Time: "09:30 PM", Location: location("Overseer's room")},
				},
				Tomorrow: []domain.ScheduleSeriesTime{
					{Title: "Marketing Meeting", Time: "09:00 AM", Location: location("Conference room 1A")},
					{Title: "Public Announcement", Time: "11:00 AM"},
					{Title: "Lunch", Time: "12:10 PM"},
					{Title: "Meeting with Beta Testers", Time: "03:00 PM", Location: location("Conference room 2C")},
					{Title: "Live Stream", Time: "05:30 PM"},
					{Title: "Release Party", Time: "07:30 PM", Location: location("CEO's house")},
					{Title: "CEO's Private Party", Time: "09:30 PM", Location: location("CEO's Penthouse")},
				},
			},
		},
	}

	return response, nil
}

func (uc *ProjectUseCase) GetDashboardProjects() (*[]domain.GetDashboardProjects, error) {
	response := &[]domain.GetDashboardProjects{
		{
			Id:   1,
			Name: "ACME Corp. Backend App",
		},
		{
			Id:   2,
			Name: "ACME Corp. Frontend App",
		},
		{
			Id:   3,
			Name: "Creapond",
		},
		{
			Id:   4,
			Name: "Withinpixels",
		},
	}

	return response, nil
}
