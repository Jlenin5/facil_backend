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
/*unc (r *ProjectRepository) DashboardSummary() ([]domain.DashboardSummary, error) {
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
}*/

// Obtener resumen de ventas
func (r *ProjectRepository) GetSaleStatusSummary() (map[string]int, error) {
	query := `
		SELECT 
			sale_status, 
			COUNT(*) AS count
		FROM sales
		WHERE deleted_at IS NULL
		GROUP BY sale_status
	`

	rows, err := r.db.Queryx(query)
	if err != nil {
		return nil, fmt.Errorf("error obteniendo resumen de estados de venta: %v", err)
	}
	defer rows.Close()

	statusCounts := make(map[string]int)

	for rows.Next() {
		var status string
		var count int
		if err := rows.Scan(&status, &count); err != nil {
			return nil, fmt.Errorf("error escaneando fila: %v", err)
		}
		statusCounts[status] = count
	}

	return statusCounts, nil
}

// Obtener resumen de compras
func (r *ProjectRepository) GetPurchaseStatusSummary() (map[string]int, error) {
	query := `
		SELECT 
			purchase_status, 
			COUNT(*) AS count
		FROM purchases
		WHERE deleted_at IS NULL
		GROUP BY purchase_status
	`
	rows, err := r.db.Queryx(query)
	if err != nil {
		return nil, fmt.Errorf("error obteniendo resumen de estados de compra: %v", err)
	}
	defer rows.Close()

	statusCounts := make(map[string]int)

	for rows.Next() {
		var status string
		var count int
		if err := rows.Scan(&status, &count); err != nil {
			return nil, fmt.Errorf("error escaneando fila: %v", err)
		}
		statusCounts[status] = count
	}

	return statusCounts, nil
}

// Obtener resumen de clientes
func (r *ProjectRepository) GetCustomerStatusSummary() (map[string]int, error) {
	query := `
		SELECT 
			document_type, 
			COUNT(*) AS count
		FROM customers
		WHERE deleted_at IS NULL
		GROUP BY document_type
	`
	rows, err := r.db.Queryx(query)
	if err != nil {
		return nil, fmt.Errorf("error obteniendo resumen de tipo de documento: %v", err)
	}
	defer rows.Close()

	statusCounts := make(map[string]int)

	for rows.Next() {
		var status string
		var count int
		if err := rows.Scan(&status, &count); err != nil {
			return nil, fmt.Errorf("error escaneando fila: %v", err)
		}
		statusCounts[status] = count
	}

	return statusCounts, nil
}

// Obtener resumen de suppliers
func (r *ProjectRepository) GetSupplierStatusSummary() (map[string]int, error) {
	query := `
		SELECT 
			CASE 
				WHEN status = B'1' THEN 'active'
				WHEN status = B'0' THEN 'inactive'
				ELSE 'unknown'
			END AS status_label,
			COUNT(*) AS count
		FROM suppliers
		WHERE deleted_at IS NULL
		GROUP BY status_label
	`
	rows, err := r.db.Queryx(query)
	if err != nil {
		return nil, fmt.Errorf("error obteniendo resumen de tipo de documento: %v", err)
	}
	defer rows.Close()

	statusCounts := make(map[string]int)

	for rows.Next() {
		var status string
		var count int
		if err := rows.Scan(&status, &count); err != nil {
			return nil, fmt.Errorf("error escaneando fila: %v", err)
		}
		statusCounts[status] = count
	}

	return statusCounts, nil
}

func (r *ProjectRepository) GetProductsBySalesAndPurchases() ([]domain.ProductFlowDay, error) {
	query := `
		WITH sales_agg AS (
    SELECT
        TO_CHAR(s.issue_date, 'Dy') AS weekday,
        SUM(sd.quantity) FILTER (WHERE s.sale_status = 'paid')     AS sales_paid,
        SUM(sd.quantity) FILTER (WHERE s.sale_status = 'canceled') AS sales_canceled
    FROM sales s
    JOIN sale_details sd ON sd.sale_id = s.id
    WHERE s.deleted_at IS NULL
      AND sd.deleted_at IS NULL
    GROUP BY weekday
),
purchases_agg AS (
    SELECT
        TO_CHAR(p.issue_date, 'Dy') AS weekday,
        SUM(pd.quantity) FILTER (WHERE p.purchase_status IN ('paid', 'received')) AS purchases_paid,
        SUM(pd.quantity) FILTER (WHERE p.purchase_status = 'canceled')            AS purchases_canceled
    FROM purchases p
    JOIN purchase_details pd ON pd.purchase_id = p.id
    WHERE p.deleted_at IS NULL
      AND pd.deleted_at IS NULL
    GROUP BY weekday
)
SELECT
    COALESCE(s.weekday, p.weekday)           AS weekday,
    COALESCE(s.sales_paid, 0)                AS sales_paid,
    COALESCE(s.sales_canceled, 0)            AS sales_canceled,
    COALESCE(p.purchases_paid, 0)            AS purchases_paid,
    COALESCE(p.purchases_canceled, 0)        AS purchases_canceled
FROM sales_agg s
FULL OUTER JOIN purchases_agg p ON p.weekday = s.weekday
ORDER BY
    CASE COALESCE(s.weekday, p.weekday)
        WHEN 'Mon' THEN 1
        WHEN 'Tue' THEN 2
        WHEN 'Wed' THEN 3
        WHEN 'Thu' THEN 4
        WHEN 'Fri' THEN 5
        WHEN 'Sat' THEN 6
        ELSE 7
    END;
	`

	var flows []domain.ProductFlowDay
	if err := r.db.Select(&flows, query); err != nil {
		return nil, fmt.Errorf("error obteniendo flujo de productos: %v", err)
	}
	return flows, nil
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
