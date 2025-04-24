package repository

import (
	"database/sql"
	"fmt"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type SaleRepository struct {
	db *sqlx.DB
}

func NewSaleRepository(db *sqlx.DB) *SaleRepository {
	return &SaleRepository{db: db}
}

// Crear una venta junto con sus detalles
func (r *SaleRepository) CreateSale(sale *domain.Sales, details []domain.SaleDetails) error {
	tx, err := r.db.Beginx()
	if err != nil {
		fmt.Printf("Error iniciando transacción: %v\n", err)
		return err
	}

	// Descomponer la estructura sale en sus campos individuales
	saleMap := map[string]interface{}{
		"document_type":      sale.Document_Type,
		"series":             sale.Series,
		"number":             sale.Number,
		"bill":               sale.Bill,
		"warehouse_id":       sale.Warehouse_Id,
		"customer_id":        sale.Customer_Id,
		"currency_id":        sale.Currency_Id,
		"user_id":            sale.User_Id,
		"issue_date":         sale.Issue_Date,
		"exchange_rate":      sale.Exchange_Rate,
		"discount":           sale.Discount,
		"subtotal":           sale.Subtotal,
		"total":              sale.Total,
		"total_paid":         sale.Total_Paid,
		"change":             sale.Change,
		"sale_status":        sale.Sale_Status,
		"payment_method_id":  sale.Payment_Method_Id,
		"sale_order_id":      sale.Sale_Order_Id,
		"tax_identification": sale.Tax_Identification,
		"retention":          sale.Retention,
		"perception":         sale.Perception,
	}

	// Insertar la venta
	saleQuery := `
		INSERT INTO sales (
			document_type, series, number, bill, warehouse_id, customer_id, currency_id, user_id, issue_date, exchange_rate, discount, subtotal, total, total_paid, change, sale_status, payment_method_id, sale_order_id, tax_identification, retention, perception
		) VALUES (
			:document_type, :series, :number, :bill, :warehouse_id, :customer_id, :currency_id, :user_id, :issue_date, :exchange_rate, :discount, :subtotal, :total, :total_paid, :change, :sale_status, :payment_method_id, :sale_order_id, :tax_identification, :retention, :perception
		) RETURNING id
	`

	var saleId int
	stmt, err := tx.PrepareNamed(saleQuery) // Preparar la consulta nombrada
	if err != nil {
		tx.Rollback()
		fmt.Printf("Error preparando la consulta: %v\n", err)
		return err
	}
	defer stmt.Close()

	err = stmt.Get(&saleId, saleMap) // Ejecutar la consulta y obtener el ID
	if err != nil {
		tx.Rollback()
		fmt.Printf("Error insertando la venta: %v\n", err)
		return err
	}

	// Insertar los detalles de la venta
	for _, detail := range details {
		detail.Sale_Id = saleId
		detailQuery := `
			INSERT INTO sale_details (
				product_name, sale_id, product_id, quantity, price, discount_method, discount, subtotal, total
			) VALUES (
				:product_name, :sale_id, :product_id, :quantity, :price, :discount_method, :discount, :subtotal, :total
			)
		`
		_, err = tx.NamedExec(detailQuery, detail)
		if err != nil {
			fmt.Printf("Error creando el detalle de la venta: %v\n", err)
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		fmt.Printf("Error confirmando transacción: %v\n", err)
		return err
	}

	return nil
}

// Obtener todas las ventas
func (r *SaleRepository) GetAllSales() ([]domain.Sales, error) {
	var sales []domain.Sales
	query := querySelectSale("s.deleted_at IS NULL")
	if err := r.db.Select(&sales, query); err != nil {
		return nil, fmt.Errorf("error obteniendo ventas: %w", err)
	}

	// Obtener detalles de cada venta
	for i := range sales {
		if err := r.fetchSaleDetails(&sales[i]); err != nil {
			return nil, fmt.Errorf("error obteniendo detalles de la venta %d: %w", sales[i].Id, err)
		}
	}

	return sales, nil
}

// Obtener una venta por ID junto con sus detalles
func (r *SaleRepository) GetSaleById(saleId int) (*domain.Sales, error) {
	var sale domain.Sales
	query := querySelectSale("s.id = $1 AND s.deleted_at IS NULL")
	if err := r.db.Get(&sale, query, saleId); err != nil {
		return nil, fmt.Errorf("error obteniendo venta con ID %d: %w", saleId, err)
	}

	if err := r.fetchSaleDetails(&sale); err != nil {
		return nil, fmt.Errorf("error obteniendo detalles de la venta %d: %w", saleId, err)
	}

	return &sale, nil
}

// Obtener una venta por IDs
func (r *SaleRepository) GetSalesByIds(ids []int) ([]domain.Sales, error) {
	var sales []domain.Sales

	query := querySelectSale("s.id = ANY($1)")

	err := r.db.Select(&sales, query, pq.Array(ids))
	if err != nil {
		return nil, err
	}

	return sales, nil
}

// Actualizar una venta
func (r *SaleRepository) UpdateSale(sale *domain.Sales, details []domain.SaleDetails) error {
	tx, err := r.db.Beginx()
	if err != nil {
		fmt.Printf("Error iniciando transacción: %v\n", err)
		return err
	}

	// Actualizar la venta
	saleQuery := `
		UPDATE sales
		SET
			document_type = :document_type,
			series = :series,
			number = :number,
			bill = :bill,
			warehouse_id = :warehouse_id,
			customer_id = :customer_id,
			currency_id = :currency_id,
			user_id = :user_id,
			issue_date = :issue_date,
			exchange_rate = :exchange_rate,
			discount = :discount,
			subtotal = :subtotal,
			total = :total,
			total_paid = :total_paid,
			change = :change,
			sale_status = :sale_status,
			payment_method_id = :payment_method_id,
			sale_order_id = :sale_order_id,
			tax_identification = :tax_identification,
			retention = :retention,
			perception = :perception,
			updated_at = NOW()
		WHERE id = :id
	`
	_, err = tx.NamedExec(saleQuery, sale)
	if err != nil {
		fmt.Printf("Error actualizando la venta: %v\n", err)
		tx.Rollback()
		return err
	}

	// Actualizar o insertar los detalles de la venta
	for _, detail := range details {
		cadena := fmt.Sprintf("%d", detail.Id)
		longitud := len(cadena)
		if longitud == 13 {
			// Insertar nuevo detalle
			detail.Sale_Id = sale.Id
			detailQuery := `
				INSERT INTO sale_details (
					product_name, sale_id, product_id, quantity, price, discount_method, discount, subtotal, total
				) VALUES (
					:product_name, :sale_id, :product_id, :quantity, :price, :discount_method, :discount, :subtotal, :total
				)
			`
			_, err = tx.NamedExec(detailQuery, detail)
			if err != nil {
				fmt.Printf("Error insertando detalle: %v\n", err)
				tx.Rollback()
				return err
			}
		} else {
			// Actualizar detalle existente
			detailQuery := `
				UPDATE sale_details
				SET
					product_name = :product_name,
					product_id = :product_id,
					quantity = :quantity,
					price = :price,
					discount_method = :discount_method,
					discount = :discount,
					subtotal = :subtotal,
					total = :total,
					updated_at = NOW()
				WHERE id = :id AND deleted_at IS NULL
			`
			_, err = tx.NamedExec(detailQuery, detail)
			if err != nil {
				fmt.Printf("Error actualizando detalle: %v\n", err)
				tx.Rollback()
				return err
			}
		}
	}

	err = tx.Commit()
	if err != nil {
		fmt.Printf("Error confirmando transacción: %v\n", err)
		return err
	}

	return nil
}

// Obtener venta por documento
func (r *SaleRepository) GetSaleByBill(bill string) (*domain.Sales, error) {
	var sale domain.Sales

	// Escribir la consulta SQL manualmente
	query := querySelectSale("s.bill = $1 AND s.deleted_at IS NULL")

	// Ejecutar la consulta y escanear el resultado en la estructura sale
	err := r.db.Get(&sale, query, bill)
	if err != nil {
		return nil, err
	}

	if err := r.fetchSaleDetails(&sale); err != nil {
		return nil, fmt.Errorf("error obteniendo detalles de la venta %d: %w", sale.Id, err)
	}

	return &sale, nil
}

// Obtener la última referencia registrada
func (r *SaleRepository) GetLastSaleBill(document_type string) (string, error) {
	var lastReference string

	query := querySelectSale("s.document_type=$1 ORDER BY s.id DESC LIMIT 1")
	err := r.db.Get(&lastReference, query, document_type)
	if err != nil {
		// Si no hay registros, devuelve cadena vacía
		if err == sql.ErrNoRows {
			return "", nil
		}
		return "", err
	}

	return lastReference, nil
}

// Obtener detalles de una venta
func (r *SaleRepository) fetchSaleDetails(sale *domain.Sales) error {
	var details []domain.SaleDetails
	query := `
		SELECT
			sd.id, sd.product_name, sd.sale_id, sd.product_id, sd.quantity, sd.price, sd.discount_method, sd.discount, sd.subtotal, sd.total,
			p.id AS "product.id", p.name AS "product.name", p.price AS "product.price", p.cost AS "product.cost"
		FROM sale_details sd
		LEFT JOIN products p ON sd.product_id=p.id
		WHERE sd.sale_id = $1
	`

	if err := r.db.Select(&details, query, sale.Id); err != nil {
		return fmt.Errorf("error obteniendo detalles de la venta: %w", err)
	}

	sale.SaleDetails = details
	return nil
}

// Obtener la consulta SQL
func querySelectSale(whereClause string) string {
	query := fmt.Sprintf(`
		SELECT
			s.id, s.document_type, s.series, s.number, s.bill, s.warehouse_id, s.customer_id, s.currency_id, s.user_id, s.issue_date, s.exchange_rate, s.discount, s.subtotal, s.total, s.total_paid, s.change, s.sale_status, s.payment_method_id, s.sale_order_id, s.tax_identification, s.retention, s.perception,
			c.id AS "customer.id", c.first_name AS "customer.first_name", c.second_name AS "customer.second_name", c.third_name AS "customer.third_name", c.surname AS "customer.surname", c.second_surname AS "customer.second_surname", c.company_name AS "customer.company_name", c.document_number AS "customer.document_number",
			u.id AS "user.id", u.employee_id AS "user.employee_id",
			cu.id AS "currency.id", cu.name AS "currency.name", cu.code AS "currency.code", cu.symbol AS "currency.symbol",
			e.id AS "user.employee.id", e.first_name AS "user.employee.first_name", e.second_name AS "user.employee.second_name", e.third_name AS "user.employee.third_name", e.surname AS "user.employee.surname", e.second_surname AS "user.employee.second_surname",
			w.id AS "warehouse.id", w.branch_office_id AS "warehouse.branch_office_id", w.name AS "warehouse.name",
			bo.id AS "warehouse.branch_office.id", bo.company_id AS "warehouse.branch_office.company_id", bo.name AS "warehouse.branch_office.name",
			co.id AS "warehouse.branch_office.company.id", co.name AS "warehouse.branch_office.company.name", co.ruc AS "warehouse.branch_office.company.ruc", co.email AS "warehouse.branch_office.company.email", co.phone AS "warehouse.branch_office.company.phone", co.address AS "warehouse.branch_office.company.address",
			COALESCE(so.id, 0) AS "sale_order.id", COALESCE(so.reference, '') AS "sale_order.reference",
			pm.id AS "payment_method.id", pm.name AS "payment_method.name"
		FROM sales s
		LEFT JOIN customers c ON s.customer_id=c.id
		LEFT JOIN currencies cu ON s.currency_id=cu.id
		LEFT JOIN warehouses w ON s.warehouse_id=w.id
		LEFT JOIN branch_offices bo ON w.branch_office_id=bo.id
		LEFT JOIN companies co ON bo.company_id=co.id
		LEFT JOIN users u ON s.user_id=u.id
		LEFT JOIN employees e ON u.employee_id=e.id
		LEFT JOIN sale_orders so ON s.sale_order_id=so.id
		LEFT JOIN payment_methods pm ON s.payment_method_id=pm.id
		WHERE %s
		ORDER BY s.id DESC
	`, whereClause)
	
	return query
}