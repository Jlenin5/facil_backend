package repository

import (
	"database/sql"
	"fmt"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type PurchaseRepository struct {
	db *sqlx.DB
}

func NewPurchaseRepository(db *sqlx.DB) *PurchaseRepository {
	return &PurchaseRepository{db: db}
}

// Crear una compra junto con sus detalles
func (r *PurchaseRepository) CreatePurchase(sale *domain.Purchases, details []domain.PurchaseDetails) error {
	tx, err := r.db.Beginx()
	if err != nil {
		fmt.Printf("Error iniciando transacción: %v\n", err)
		return err
	}

	// Descomponer la estructura sale en sus campos individuales
	saleMap := map[string]interface{}{
		"reference":      sale.Reference,
		"warehouse_id":       sale.Warehouse_Id,
		"currency_id":        sale.Currency_Id,
		"issue_date":         sale.Issue_Date,
		"exchange_rate":      sale.Exchange_Rate,
		"discount":           sale.Discount,
		"subtotal":           sale.Subtotal,
		"total":              sale.Total,
		"total_paid":         sale.Total_Paid,
		"change":             sale.Change,
		"sale_status":        sale.Purchase_Status,
		"payment_method_id":  sale.Payment_Method_Id,
		"sale_order_id":      sale.Purchase_Order_Id,
	}

	// Insertar la compra
	saleQuery := `
		INSERT INTO purchases (
			reference, invoice_number, number, bill, warehouse_id, customer_id, currency_id, user_id, issue_date, exchange_rate, discount, subtotal, total, total_paid, change, sale_status, payment_method_id, sale_order_id, tax_identification, retention, perception
		) VALUES (
			:reference, :invoice_number, :number, :bill, :warehouse_id, :customer_id, :currency_id, :user_id, :issue_date, :exchange_rate, :discount, :subtotal, :total, :total_paid, :change, :sale_status, :payment_method_id, :sale_order_id, :tax_identification, :retention, :perception
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
		fmt.Printf("Error insertando la compra: %v\n", err)
		return err
	}

	// Insertar los detalles de la compra
	for _, detail := range details {
		detail.Purchase_Id = saleId
		detailQuery := `
			INSERT INTO sale_details (
				product_name, sale_id, product_id, quantity, price, discount_method, discount, subtotal, total
			) VALUES (
				:product_name, :sale_id, :product_id, :quantity, :price, :discount_method, :discount, :subtotal, :total
			)
		`
		_, err = tx.NamedExec(detailQuery, detail)
		if err != nil {
			fmt.Printf("Error creando el detalle de la compra: %v\n", err)
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

// Obtener todas las compras
func (r *PurchaseRepository) GetAllPurchases() ([]domain.Purchases, error) {
	var purchases []domain.Purchases
	query := querySelectPurchase("p.deleted_at IS NULL")
	if err := r.db.Select(&purchases, query); err != nil {
		return nil, fmt.Errorf("error obteniendo compras: %w", err)
	}

	// Obtener detalles de cada compra
	for i := range purchases {
		if err := r.fetchPurchaseDetails(&purchases[i]); err != nil {
			return nil, fmt.Errorf("error obteniendo detalles de la compra %d: %w", purchases[i].Id, err)
		}
	}

	return purchases, nil
}

// Obtener una compra por ID junto con sus detalles
func (r *PurchaseRepository) GetPurchaseById(saleId int) (*domain.Purchases, error) {
	var sale domain.Purchases
	query := querySelectPurchase("p.id = $1 AND p.deleted_at IS NULL")
	if err := r.db.Get(&sale, query, saleId); err != nil {
		return nil, fmt.Errorf("error obteniendo compra con ID %d: %w", saleId, err)
	}

	if err := r.fetchPurchaseDetails(&sale); err != nil {
		return nil, fmt.Errorf("error obteniendo detalles de la compra %d: %w", saleId, err)
	}

	return &sale, nil
}

// Obtener una compra por IDs
func (r *PurchaseRepository) GetPurchasesByIds(ids []int) ([]domain.Purchases, error) {
	var purchases []domain.Purchases

	query := querySelectPurchase("p.id = ANY($1)")

	err := r.db.Select(&purchases, query, pq.Array(ids))
	if err != nil {
		return nil, err
	}

	return purchases, nil
}

// Actualizar una compra
func (r *PurchaseRepository) UpdatePurchase(sale *domain.Purchases, details []domain.PurchaseDetails) error {
	tx, err := r.db.Beginx()
	if err != nil {
		fmt.Printf("Error iniciando transacción: %v\n", err)
		return err
	}

	// Actualizar la compra
	saleQuery := `
		UPDATE purchases
		SET
			reference = :reference,
			invoice_number = :invoice_number,
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
		fmt.Printf("Error actualizando la compra: %v\n", err)
		tx.Rollback()
		return err
	}

	// Actualizar o insertar los detalles de la compra
	for _, detail := range details {
		cadena := fmt.Sprintf("%d", detail.Id)
		longitud := len(cadena)
		if longitud == 13 {
			// Insertar nuevo detalle
			detail.Purchase_Id = sale.Id
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

// Obtener compra por documento
func (r *PurchaseRepository) GetPurchaseByBill(bill string) (*domain.Purchases, error) {
	var sale domain.Purchases

	// Escribir la consulta SQL manualmente
	query := querySelectPurchase("p.bill = $1 AND p.deleted_at IS NULL")

	// Ejecutar la consulta y escanear el resultado en la estructura sale
	err := r.db.Get(&sale, query, bill)
	if err != nil {
		return nil, err
	}

	if err := r.fetchPurchaseDetails(&sale); err != nil {
		return nil, fmt.Errorf("error obteniendo detalles de la compra %d: %w", sale.Id, err)
	}

	return &sale, nil
}

// Obtener la última referencia registrada
func (r *PurchaseRepository) GetLastPurchaseBill(reference string) (string, error) {
	var lastReference string

	query := querySelectPurchase("p.reference=$1 ORDER BY p.id DESC LIMIT 1")
	err := r.db.Get(&lastReference, query, reference)
	if err != nil {
		// Si no hay registros, devuelve cadena vacía
		if err == sql.ErrNoRows {
			return "", nil
		}
		return "", err
	}

	return lastReference, nil
}

// Obtener detalles de una compra
func (r *PurchaseRepository) fetchPurchaseDetails(purchase *domain.Purchases) error {
	var details []domain.PurchaseDetails
	query := `
		SELECT
			pd.id, pd.purchase_id, pd.product_id, pd.quantity, pd.price, pd.discount, pd.subtotal, pd.total,
			p.id AS "product.id", p.name AS "product.name", p.price AS "product.price", p.cost AS "product.cost"
		FROM purchase_details pd
		LEFT JOIN products p ON pd.product_id=p.id
		WHERE pd.purchase_id = $1
	`

	if err := r.db.Select(&details, query, purchase.Id); err != nil {
		return fmt.Errorf("error obteniendo detalles de la compra: %w", err)
	}

	purchase.PurchaseDetails = details
	return nil
}

// Obtener la consulta SQL
func querySelectPurchase(whereClause string) string {
	query := fmt.Sprintf(`
		SELECT
			p.id, p.reference, p.invoice_number, p.warehouse_id, p.supplier_id, p.currency_id, p.issue_date, p.exchange_rate, p.discount, p.subtotal, p.total, p.total_paid, p.payment_method_id,
			c.id AS "supplier.id", c.name AS "supplier.name", c.ruc AS "supplier.ruc", c.email AS "supplier.email", c.phone AS "supplier.phone",
			w.id AS "warehouse.id", w.branch_office_id AS "warehouse.branch_office_id", w.name AS "warehouse.name",
			bo.id AS "warehouse.branch_office.id", bo.company_id AS "warehouse.branch_office.company_id", bo.name AS "warehouse.branch_office.name",
			co.id AS "warehouse.branch_office.company.id", co.name AS "warehouse.branch_office.company.name", co.ruc AS "warehouse.branch_office.company.ruc", co.email AS "warehouse.branch_office.company.email", co.phone AS "warehouse.branch_office.company.phone", co.address AS "warehouse.branch_office.company.address",
			COALESCE(po.id, 0) AS "purchase_order.id", COALESCE(po.reference, '') AS "purchase_order.reference",
			COALESCE(pm.id, 0) AS "payment_method.id", COALESCE(pm.name, '') AS "payment_method.name"
		FROM purchases p
		LEFT JOIN suppliers c ON p.supplier_id=c.id
		LEFT JOIN currencies cu ON p.currency_id=cu.id
		LEFT JOIN warehouses w ON p.warehouse_id=w.id
		LEFT JOIN branch_offices bo ON w.branch_office_id=bo.id
		LEFT JOIN companies co ON bo.company_id=co.id
		LEFT JOIN purchase_orders po ON p.purchase_order_id=po.id
		LEFT JOIN payment_methods pm ON p.payment_method_id=pm.id
		WHERE %s
		ORDER BY p.id DESC
	`, whereClause)
	
	return query
}