package repository

import (
	"database/sql"
	"fmt"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/jmoiron/sqlx"
)

type PurchaseOrderRepository struct {
	db *sqlx.DB
}

func NewPurchaseOrderRepository(db *sqlx.DB) *PurchaseOrderRepository {
	return &PurchaseOrderRepository{db: db}
}

// Crear una orden de compra junto con sus detalles
func (r *PurchaseOrderRepository) CreatePurchaseOrder(order *domain.PurchaseOrders, details []domain.PurchaseOrderDetails) error {
	tx, err := r.db.Beginx()
	if err != nil {
		fmt.Printf("Error iniciando transacción: %v\n", err)
		return err
	}

	// Descomponer la estructura sale en sus campos individuales
	orderMap := map[string]interface{}{
		"reference":           order.Reference,
		"warehouse_id":        order.Warehouse_Id,
		"supplier_id":         order.Supplier_Id,
		"currency_id":         order.Currency_Id,
		"exchange_rate":       order.Exchange_Rate,
		"discount":            order.Discount,
		"issue_date":          order.Issue_Date,
		"tax":                 order.Tax,
		"subtotal":            order.Subtotal,
		"total":               order.Total,
		"document_attachment": order.Document_Attachment,
		"order_status":        order.Order_Status,
		"approval_date":       order.Approval_Date,
		"created_by":          order.Created_By,
		"approved_by":         order.Approved_By,
		"migrate_purchase":    order.Migrate_Purchase,
		"notes":               order.Notes,
	}

	// Insertar la orden de compra
	orderQuery := `
		INSERT INTO purchase_orders (
			reference, warehouse_id, supplier_id, currency_id, exchange_rate, discount, issue_date, tax, subtotal, total, document_attachment, order_status, approval_date, created_by, approved_by, migrate_purchase, notes
		) VALUES (
			:reference, :warehouse_id, :supplier_id, :currency_id, :exchange_rate, :discount, :issue_date, :tax, :subtotal, :total, :document_attachment, :order_status, :approval_date, :created_by, :approved_by, :migrate_purchase, :notes
		) RETURNING id
	`

	var orderId int
	stmt, err := tx.PrepareNamed(orderQuery) // Preparar la consulta nombrada
	if err != nil {
		tx.Rollback()
		fmt.Printf("Error preparando la consulta: %v\n", err)
		return err
	}
	defer stmt.Close()

	err = stmt.Get(&orderId, orderMap) // Ejecutar la consulta y obtener el ID
	if err != nil {
		tx.Rollback()
		fmt.Printf("Error insertando la orden de compra: %v\n", err)
		return err
	}

	// Insertar los detalles de la orden
	for _, detail := range details {
		detail.Purchase_Order_Id = orderId
		detailQuery := `
			INSERT INTO purchase_order_details (
				purchase_order_id, product_id, quantity, discount_method, discount, price, subtotal, total
			) VALUES (
				:purchase_order_id, :product_id, :quantity, :discount_method, :discount, :price, :subtotal, :total
			)
		`
		_, err = tx.NamedExec(detailQuery, detail)
		if err != nil {
			fmt.Printf("Error creando el detalle de la orden: %v\n", err)
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

// Obtener todas las órdenes de compra
func (r *PurchaseOrderRepository) GetAllPurchaseOrders() ([]domain.PurchaseOrders, error) {
	var purchaseOrders []domain.PurchaseOrders
	query := querySelectPurchaseOrder("po.deleted_at IS NULL")
	if err := r.db.Select(&purchaseOrders, query); err != nil {
		return nil, fmt.Errorf("error obteniendo órdenes de compra: %w", err)
	}

	// Iterar sobre las órdenes de compra y agregar los detalles de cada orden
	for i := range purchaseOrders {
		if err := r.fetchPurchaseOrderDetails(&purchaseOrders[i]); err != nil {
			return nil, fmt.Errorf("error obteniendo detalles de la orden de compra %d: %w", purchaseOrders[i].Id, err)
		}
	}

	return purchaseOrders, nil
}

// Obtener una orden de compra por ID junto con sus detalles
func (r *PurchaseOrderRepository) GetPurchaseOrderByID(orderId int) (*domain.PurchaseOrders, error) {
	var purchaseOrder domain.PurchaseOrders
	query := querySelectPurchaseOrder("po.id = $1 AND po.deleted_at IS NULL")
	if err := r.db.Get(&purchaseOrder, query, orderId); err != nil {
		return nil, fmt.Errorf("error obteniendo compra con ID %d: %w", orderId, err)
	}

	if err := r.fetchPurchaseOrderDetails(&purchaseOrder); err != nil {
		return nil, fmt.Errorf("error obteniendo detalles de la orden de compra %d: %w", orderId, err)
	}

	return &purchaseOrder, nil
}

// Actualizar una orden de compra
func (r *PurchaseOrderRepository) UpdatePurchaseOrder(order *domain.PurchaseOrders, details []domain.PurchaseOrderDetails) error {
	tx, err := r.db.Beginx()
	if err != nil {
		fmt.Printf("Error iniciando transacción: %v\n", err)
		return err
	}

	// Actualizar la orden de compra
	orderQuery := `
		UPDATE purchase_orders
		SET
			reference           = :reference,
			warehouse_id        = :warehouse_id,
			supplier_id         = :supplier_id,
			currency_id         = :currency_id,
			exchange_rate       = :exchange_rate,
			discount            = :discount,
			issue_date          = :issue_date,
			tax                 = :tax,
			subtotal            = :subtotal,
			total               = :total,
			document_attachment = :document_attachment,
			order_status        = :order_status,
			approval_date       = :approval_date,
			created_by          = :created_by,
			approved_by         = :approved_by,
			migrate_purchase    = :migrate_purchase,
			notes               = :notes,
			updated_at = NOW()
		WHERE id = :id AND deleted_at IS NULL
	`
	_, err = tx.NamedExec(orderQuery, order)
	if err != nil {
		fmt.Printf("Error actualizando la orden de compra: %v\n", err)
		tx.Rollback()
		return err
	}

	// Actualizar o insertar los detalles de la orden
	for _, detail := range details {
		cadena := fmt.Sprintf("%d", detail.Id)
		longitud := len(cadena)
		if longitud == 13 {
			// Insertar nuevo detalle
			detail.Purchase_Order_Id = order.Id
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
				UPDATE purchase_order_details
				SET
					purchase_order_id = :purchase_order_id,
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
	}

	err = tx.Commit()
	if err != nil {
		fmt.Printf("Error confirmando transacción: %v\n", err)
		return err
	}

	return nil
}

// Obtener detalles de una orden de compra
func (r *PurchaseOrderRepository) fetchPurchaseOrderDetails(purchaseOrder *domain.PurchaseOrders) error {
	var details []domain.PurchaseOrderDetails
	query := `
		SELECT
			pod.id, pod.purchase_order_id, pod.product_id, pod.quantity, pod.price, pod.discount, pod.subtotal, pod.total,
			p.id AS "product.id", p.name AS "product.name", p.featured_pcf AS "product.featured_pcf", p.cost AS "product.cost"
		FROM purchase_order_details pod
		LEFT JOIN products p ON pod.product_id=p.id
		WHERE pod.purchase_order_id = $1
	`

	if err := r.db.Select(&details, query, purchaseOrder.Id); err != nil {
		return fmt.Errorf("error obteniendo detalles de la compra: %w", err)
	}

	purchaseOrder.PurchaseOrderDetails = details
	return nil
}

// Obtener la consulta SQL
func querySelectPurchaseOrder(whereClause string) string {
	query := fmt.Sprintf(`
		SELECT
			po.id, po.reference, po.warehouse_id, po.supplier_id, po.currency_id, po.exchange_rate, po.discount, po.issue_date, po.tax, po.subtotal, po.total, po.document_attachment, po.order_status, po.approval_date, po.created_by, po.approved_by, po.migrate_purchase, po.notes,
			s.id AS "supplier.id", s.name AS "supplier.name", s.ruc AS "supplier.ruc", s.email AS "supplier.email", s.phone AS "supplier.phone",
			cu.id AS "currency.id", cu.name AS "currency.name", cu.code AS "currency.code", cu.symbol AS "currency.symbol",
			w.id AS "warehouse.id", w.branch_office_id AS "warehouse.branch_office_id", w.name AS "warehouse.name",
			bo.id AS "warehouse.branch_office.id", bo.company_id AS "warehouse.branch_office.company_id", bo.name AS "warehouse.branch_office.name",
			co.id AS "warehouse.branch_office.company.id", co.name AS "warehouse.branch_office.company.name", co.ruc AS "warehouse.branch_office.company.ruc", co.email AS "warehouse.branch_office.company.email", co.phone AS "warehouse.branch_office.company.phone", co.address AS "warehouse.branch_office.company.address"
		FROM purchase_orders po
		LEFT JOIN suppliers s ON po.supplier_id=s.id
		LEFT JOIN currencies cu ON po.currency_id=cu.id
		LEFT JOIN warehouses w ON po.warehouse_id=w.id
		LEFT JOIN branch_offices bo ON w.branch_office_id=bo.id
		LEFT JOIN companies co ON bo.company_id=co.id
		WHERE %s
		ORDER BY po.id DESC
	`, whereClause)
	
	return query
}

// Obtener la última referencia registrada
func (r *PurchaseOrderRepository) GetLastPurchaseOrderReference() (string, error) {
	var lastReference string

	query := `SELECT reference FROM purchase_orders ORDER BY id DESC LIMIT 1`
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