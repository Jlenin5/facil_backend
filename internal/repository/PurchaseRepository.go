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
func (r *PurchaseRepository) CreatePurchase(purchase *domain.Purchases, details []domain.PurchaseDetails) error {
	tx, err := r.db.Beginx()
	if err != nil {
		fmt.Printf("Error iniciando transacción: %v\n", err)
		return err
	}

	// Descomponer la estructura purchase en sus campos individuales
	purchaseMap := map[string]interface{}{
		"reference":           purchase.Reference,
		"invoice_number":      purchase.Invoice_Number,
		"supplier_id":         purchase.Supplier_Id,
		"warehouse_id":        purchase.Warehouse_Id,
		"currency_id":         purchase.Currency_Id,
		"exchange_rate":       purchase.Exchange_Rate,
		"purchase_status":     purchase.Purchase_Status,
		"purchase_order_id":   purchase.Purchase_Order_Id,
		"issue_date":          purchase.Issue_Date,
		"received_date":       purchase.Received_Date,
		"payment_date":        purchase.Payment_Date,
		"discount":            purchase.Discount,
		"subtotal":            purchase.Subtotal,
		"tax":                 purchase.Tax,
		"total":               purchase.Total,
		"total_paid":          purchase.Total_Paid,
		"change":              purchase.Change,
		"payment_method_id":   purchase.Payment_Method_Id,
		"created_by":          purchase.Created_By,
		"document_attachment": purchase.Document_Attachment,
		"notes":               purchase.Notes,
	}

	// Insertar la compra
	purchaseQuery := `
		INSERT INTO purchases (
			reference, invoice_number, supplier_id, warehouse_id, currency_id, exchange_rate, purchase_status, purchase_order_id, issue_date, received_date, payment_date, discount, subtotal, tax, total, total_paid, change, payment_method_id, created_by, document_attachment, notes
		) VALUES (
			:reference, :invoice_number, :supplier_id, :warehouse_id, :currency_id, :exchange_rate, :purchase_status, :purchase_order_id, :issue_date, :received_date, :payment_date, :discount, :subtotal, :tax, :total, :total_paid, :change, :payment_method_id, :created_by, :document_attachment, :notes
		) RETURNING id
	`

	var purchaseId int
	stmt, err := tx.PrepareNamed(purchaseQuery) // Preparar la consulta nombrada
	if err != nil {
		tx.Rollback()
		fmt.Printf("Error preparando la consulta: %v\n", err)
		return err
	}
	defer stmt.Close()

	err = stmt.Get(&purchaseId, purchaseMap) // Ejecutar la consulta y obtener el ID
	if err != nil {
		tx.Rollback()
		fmt.Printf("Error insertando la compra: %v\n", err)
		return err
	}

	// Insertar los detalles de la compra
	for _, detail := range details {
		detail.Purchase_Id = purchaseId
		detailQuery := `
			INSERT INTO purchase_order_details (
				purchase_order_id, product_id, quantity, discount_method, discount, price, subtotal, total
			) VALUES (
				:purchase_order_id, :product_id, :quantity, :discount_method, :discount, :price, :subtotal, :total
			)
		`
		_, err = tx.NamedExec(detailQuery, detail)
		if err != nil {
			fmt.Printf("Error creando el detalle de la compra: %v\n", err)
			tx.Rollback()
			return err
		}

		if purchase.Purchase_Status == "received" {
			stockUpdateQuery := `
				UPDATE stock_control
				SET current_stock = current_stock + $1, updated_at = NOW()
				WHERE product_id = $2 AND warehouse_id = $3
			`
			_, err := tx.Exec(stockUpdateQuery, detail.Quantity, detail.Product_Id, purchase.Warehouse_Id)
			if err != nil {
				fmt.Printf("Error actualizando stock_control para producto %v: %v\n", detail.Product_Id, err)
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
func (r *PurchaseRepository) GetPurchaseById(purchaseId int) (*domain.Purchases, error) {
	var purchase domain.Purchases
	query := querySelectPurchase("p.id = $1 AND p.deleted_at IS NULL")
	if err := r.db.Get(&purchase, query, purchaseId); err != nil {
		return nil, fmt.Errorf("error obteniendo compra con ID %d: %w", purchaseId, err)
	}

	if err := r.fetchPurchaseDetails(&purchase); err != nil {
		return nil, fmt.Errorf("error obteniendo detalles de la compra %d: %w", purchaseId, err)
	}

	return &purchase, nil
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
func (r *PurchaseRepository) UpdatePurchase(purchase *domain.Purchases, details []domain.PurchaseDetails) error {
	tx, err := r.db.Beginx()
	if err != nil {
		fmt.Printf("Error iniciando transacción: %v\n", err)
		return err
	}

	// Actualizar la compra
	purchaseQuery := `
		UPDATE purchases
		SET
			reference =           :reference,
			invoice_number =      :invoice_number,
			supplier_id =         :supplier_id,
			warehouse_id =        :warehouse_id,
			currency_id =         :currency_id,
			exchange_rate =       :exchange_rate,
			purchase_status =     :purchase_status,
			purchase_order_id =   :purchase_order_id,
			issue_date =          :issue_date,
			received_date =       :received_date,
			payment_date =        :payment_date,
			discount =            :discount,
			subtotal =            :subtotal,
			tax =                 :tax,
			total =               :total,
			total_paid =          :total_paid,
			change =              :change,
			payment_method_id =   :payment_method_id,
			created_by =          :created_by,
			document_attachment = :document_attachment,
			notes =               :notes,
			updated_at = NOW()
		WHERE id = :id
	`
	_, err = tx.NamedExec(purchaseQuery, purchase)
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
			detail.Purchase_Id = purchase.Id
			detailQuery := `
				INSERT INTO purchase_order_details (
					purchase_order_id, product_id, quantity, discount_method, discount, price, subtotal, total
				) VALUES (
					:purchase_order_id, :product_id, :quantity, :discount_method, :discount, :price, :subtotal, :total
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
				UPDATE purchase_details
				SET
					purchase_id = :purchase_id,
					product_id = :product_id,
					quantity = :quantity,
					discount_method = :discount_method,
					discount = :discount,
					price = :price,
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

		if purchase.Purchase_Status == "received" {
			stockUpdateQuery := `
				UPDATE stock_control
				SET current_stock = current_stock + $1, updated_at = NOW()
				WHERE product_id = $2 AND warehouse_id = $3
			`
			_, err := tx.Exec(stockUpdateQuery, detail.Quantity, detail.Product_Id, purchase.Warehouse_Id)
			if err != nil {
				fmt.Printf("Error actualizando stock_control para producto %v: %v\n", detail.Product_Id, err)
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

// Obtener detalles de una compra
func (r *PurchaseRepository) fetchPurchaseDetails(purchase *domain.Purchases) error {
	var details []domain.PurchaseDetails
	query := `
		SELECT
			pd.id, pd.purchase_id, pd.product_id, pd.quantity, pd.price, pd.discount, pd.subtotal, pd.total,
			p.id AS "product.id", p.name AS "product.name", p.featured_pcf AS "product.featured_pcf", p.cost AS "product.cost"
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
			p.id, p.reference, p.invoice_number, p.supplier_id, p.warehouse_id, p.currency_id, p.exchange_rate, p.purchase_status, p.purchase_order_id, p.issue_date, p.received_date, p.payment_date, p.discount, p.subtotal, p.tax, p.total, p.total_paid, p.change, p.payment_method_id, p.created_by, p.document_attachment, p.notes,
			c.id AS "supplier.id", c.name AS "supplier.name", c.ruc AS "supplier.ruc", c.email AS "supplier.email", c.phone AS "supplier.phone",
			cu.id AS "currency.id", cu.name AS "currency.name", cu.code AS "currency.code", cu.symbol AS "currency.symbol",
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

// Obtener la última referencia registrada
func (r *PurchaseRepository) GetLastPurchaseReference() (string, error) {
	var lastReference string

	query := `SELECT reference FROM purchases ORDER BY id DESC LIMIT 1`
	err := r.db.QueryRow(query).Scan(&lastReference)
	if err != nil {
		// Si no hay registros, devuelve cadena vacía
		if err == sql.ErrNoRows {
			return "", nil
		}
		return "", err
	}

	return lastReference, nil
}