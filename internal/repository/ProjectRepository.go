package repository

import (
	"fmt"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/jmoiron/sqlx"
)

type ProjectRepository struct {
	db *sqlx.DB
}

func NewProjectRepository(db *sqlx.DB) *ProjectRepository {
	return &ProjectRepository{db: db}
}

// Resumen del dashboard
func (r *ProjectRepository) DashboardSummary() ([]domain.DashboardSummary, error) {
	var summary domain.DashboardSummary

	// Obtener total de ventas
	err := r.db.QueryRow("SELECT COALESCE(SUM(total), 0) FROM sales WHERE sale_status = 'paid'").Scan(&summary.TotalSales)
	if err != nil {
		return nil, fmt.Errorf("error al obtener el total de ventas: %v", err)
	}

	// Obtener total de compras
	err = r.db.QueryRow("SELECT COALESCE(SUM(total), 0) FROM purchase_orders WHERE order_status = 'paid'").Scan(&summary.TotalPurchases)
	if err != nil {
		return nil, fmt.Errorf("error al obtener el total de compras: %v", err)
	}

	// Obtener productos con stock bajo
	err = r.db.QueryRow("SELECT COUNT(*) FROM stock_control WHERE current_stock < min_stock").Scan(&summary.LowStockProducts)
	if err != nil {
		return nil, fmt.Errorf("error al obtener productos con stock bajo: %v", err)
	}

	// Obtener clientes activos
	err = r.db.QueryRow("SELECT COUNT(*) FROM customers WHERE status = B'1'").Scan(&summary.ActiveCustomers)
	if err != nil {
		return nil, fmt.Errorf("error al obtener clientes activos: %v", err)
	}

	// Obtener empleados activos
	err = r.db.QueryRow("SELECT COUNT(*) FROM employees WHERE status = B'1'").Scan(&summary.ActiveEmployees)
	if err != nil {
		return nil, fmt.Errorf("error al obtener empleados activos: %v", err)
	}

	return []domain.DashboardSummary{summary}, nil
}

// Obtener resumen de órdenes de venta
func (r *ProjectRepository) GetSaleOrders() (map[string]int, error) {
	query := `
		SELECT 
			COUNT(*) FILTER (WHERE issue_date >= NOW() - INTERVAL '1 day') AS dy,
			COUNT(*) FILTER (WHERE issue_date >= NOW()::date) AS dt,
			COUNT(*) FILTER (WHERE issue_date >= NOW()::date + INTERVAL '1 day') AS dtm
		FROM sale_orders
		WHERE deleted_at IS NULL
	`
	var counts struct {
		DY  int `db:"dy"`
		DT  int `db:"dt"`
		DTM int `db:"dtm"`
	}
	err := r.db.Get(&counts, query)
	if err != nil {
		return nil, fmt.Errorf("error obteniendo resumen de órdenes de venta: %v", err)
	}

	return map[string]int{
		"DY":  counts.DY,
		"DT":  counts.DT,
		"DTM": counts.DTM,
	}, nil
}

// Obtener resumen de órdenes de compra
func (r *ProjectRepository) GetPurchaseOrders() (map[string]int, error) {
	query := `
		SELECT 
			COUNT(*) FILTER (WHERE issue_date >= NOW() - INTERVAL '1 day') AS dy,
			COUNT(*) FILTER (WHERE issue_date >= NOW()::date) AS dt,
			COUNT(*) FILTER (WHERE issue_date >= NOW()::date + INTERVAL '1 day') AS dtm
		FROM purchase_orders
		WHERE deleted_at IS NULL
	`
	var counts struct {
		DY  int `db:"dy"`
		DT  int `db:"dt"`
		DTM int `db:"dtm"`
	}
	err := r.db.Get(&counts, query)
	if err != nil {
		return nil, fmt.Errorf("error obteniendo resumen de órdenes de compra: %v", err)
	}

	return map[string]int{
		"DY":  counts.DY,
		"DT":  counts.DT,
		"DTM": counts.DTM,
	}, nil
}

// Obtener órdenes recientes
func (r *ProjectRepository) GetRecentOrders() ([]domain.OrderInfo, error) {
	query := `
		SELECT id, reference, customer_id, issue_date, total, order_status
		FROM sale_orders
		WHERE deleted_at IS NULL
		ORDER BY issue_date DESC
		LIMIT 5
	`
	var orders []domain.OrderInfo
	err := r.db.Select(&orders, query)
	if err != nil {
		return nil, fmt.Errorf("error obteniendo órdenes recientes: %v", err)
	}
	return orders, nil
}
