package usecase

import (
	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/Jlenin5/facil_backend/internal/repository"
)

type ProjectUseCase struct {
	ProjecRepo *repository.ProjectRepository
}

func NewProjectUseCase(ProjecRepo *repository.ProjectRepository) *ProjectUseCase {
	return &ProjectUseCase{ProjecRepo: ProjecRepo}
}

func (uc *ProjectUseCase) DashboardSummary() ([]domain.DashboardSummary, error) {
	return uc.ProjecRepo.DashboardSummary()
}

func (uc *ProjectUseCase) GetProjectData() (*domain.ProjectResponse, error) {
	saleOrders, err := uc.ProjecRepo.GetSaleOrders()
	if err != nil {
		return nil, err
	}

	purchaseOrders, err := uc.ProjecRepo.GetPurchaseOrders()
	if err != nil {
		return nil, err
	}

	recentOrders, err := uc.ProjecRepo.GetRecentOrders()
	if err != nil {
		return nil, err
	}

	response := &domain.ProjectResponse{
		SaleOrders: domain.Widget{
			Status: map[string]string{
				"DY":  "Yesterday",
				"DT":  "Today",
				"DTM": "Tomorrow",
			},
			CurrentRange: "DT",
			Data: domain.WidgetData{
				Name:  "Sale Orders",
				Count: saleOrders,
				Extra: domain.ExtraData{
					Name: "Completed",
					Count: map[string]int{
						"DY":  5,
						"DT":  7,
						"DTM": 0,
					},
				},
			},
			Detail: "You can show some detailed information about sale orders in here.",
		},
		PurchaseOrders: domain.Widget{
			Status: map[string]string{
				"DY":  "Yesterday",
				"DT":  "Today",
				"DTM": "Tomorrow",
			},
			CurrentRange: "DT",
			Data: domain.WidgetData{
				Name:  "Purchase Orders",
				Count: purchaseOrders,
				Extra: domain.ExtraData{
					Name: "Completed",
					Count: map[string]int{
						"DY":  3,
						"DT":  4,
						"DTM": 0,
					},
				},
			},
			Detail: "You can show some detailed information about purchase orders in here.",
		},
		OverdueOrders: domain.OverdueWidget{
			Title: "Overdue Orders",
			Data: domain.WidgetData{
				Name: "Orders",
				Count: map[string]int{
					"total": 4,
				},
				Extra: domain.ExtraData{
					Name: "Yesterday's overdue",
					Count: map[string]int{
						"total": 2,
					},
				},
			},
			Detail: "You can show some detailed information about overdue orders in here.",
		},
		RecentOrders: domain.RecentOrdersWidget{
			Title:  "Recent Orders",
			Orders: recentOrders,
		},
	}

	return response, nil
}
