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

	// Insertar la orden de compra
	orderQuery := `
		INSERT INTO purchase_orders (
			reference, description, warehouse_id, supplier_id, supplier_document, exchange_rate, discount, user_id, issue_date, tax, subtotal, total, order_status, date_approved, migrate_purchase
		) VALUES (
			:reference, :description, :warehouse_id, :supplier_id, :supplier_document, :exchange_rate, :discount, :user_id, :date, :tax, :subtotal, :total, :order_status, :date_approved, :migrate_purchase
		) RETURNING id
	`
	var orderId int
	err = tx.QueryRowx(orderQuery, order).Scan(&orderId)
	if err != nil {
		fmt.Printf("Error creando la orden de compra: %v\n", err)
		tx.Rollback()
		return err
	}

	// Insertar los detalles de la orden
	for _, detail := range details {
		detail.Purchase_Order_Id = orderId
		detailQuery := `
			INSERT INTO purchase_order_details (
				 purchase_order_id, product_id, quantity, price, total
			) VALUES (
				:purchase_order_id, :product_id, :quantity, :price, :total
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
	query := "SELECT id, reference, description, warehouse_id, supplier_id, supplier_document, exchange_rate, discount, user_id, issue_date, tax, subtotal, total, order_status, date_approved, migrate_purchase FROM purchase_orders WHERE deleted_at IS NULL"
	err := r.db.Select(&purchaseOrders, query)
	if err != nil {
		fmt.Printf("Error obteniendo órdenes de compra: %v\n", err)
		return nil, err
	}

	// Iterar sobre las órdenes de compra y agregar los detalles de cada orden
	for i := range purchaseOrders {
		// Obtener almacén de la orden de compra
		var warehouse domain.Warehouses
		err = r.db.Get(&warehouse, `
			SELECT w.id, w.name, w.address, w.status
			FROM warehouses w
			WHERE w.id = $1 AND w.deleted_at IS NULL
		`, purchaseOrders[i].Warehouse_Id)
		if err != nil && err != sql.ErrNoRows { // Manejar casos donde no hay almacén asociado
			fmt.Printf("Error obteniendo el almacén para la orden de compra %d: %v\n", purchaseOrders[i].Id, err)
			return nil, err
		}
		purchaseOrders[i].Warehouse = &warehouse

		// Obtener proveedor de la orden de compra
		var supplier domain.Suppliers
		err = r.db.Get(&supplier, `
			SELECT s.id, s.name, s.email, s.address, s.phone, s.status
			FROM suppliers s
			WHERE s.id = $1 AND s.deleted_at IS NULL
		`, purchaseOrders[i].Supplier_Id)
		if err != nil && err != sql.ErrNoRows { // Manejar casos donde no hay proveedor asociado
			fmt.Printf("Error obteniendo el proveedor para la orden de compra %d: %v\n", purchaseOrders[i].Id, err)
			return nil, err
		}
		purchaseOrders[i].Supplier = &supplier

		// Obtener usuario de la orden de compra
		var user domain.Users
		err = r.db.Get(&user, `
			SELECT u.id, u.role_id, u.username, u.avatar, u.email, u.status
			FROM users u
			WHERE u.id = $1 AND u.deleted_at IS NULL
		`, purchaseOrders[i].User_Id)
		if err != nil && err != sql.ErrNoRows { // Manejar casos donde no hay usuario asociado
			fmt.Printf("Error obteniendo el usuario para la orden de compra %d: %v\n", purchaseOrders[i].Id, err)
			return nil, err
		}
		purchaseOrders[i].User = &user

		var purchaseOrderDetails []domain.PurchaseOrderDetails
		// Consulta para obtener los detalles de cada orden de compra
		detailsQuery := `
			SELECT pod.id, pod.purchase_order_id, pod.product_id, pod.quantity, pod.price, pod.total FROM purchase_order_details pod WHERE pod.purchase_order_id = $1 AND pod.deleted_at IS NULL
		`

		err := r.db.Select(&purchaseOrderDetails, detailsQuery, purchaseOrders[i].Id)
		if err != nil {
			fmt.Printf("Error obteniendo detalles de la orden de compra %d: %v\n", purchaseOrders[i].Id, err)
			return nil, err
		}

		// Asignar los detalles a la orden correspondiente
		purchaseOrders[i].PurchaseOrderDetails = purchaseOrderDetails
	}

	return purchaseOrders, nil
}

// Obtener una orden de compra por ID junto con sus detalles
func (r *PurchaseOrderRepository) GetPurchaseOrderByID(orderId int) (*domain.PurchaseOrders, error) {
	var purchaseOrders domain.PurchaseOrders
	query := "SELECT id, reference, description, warehouse_id, supplier_id, supplier_document, exchange_rate, discount, user_id, issue_date, tax, subtotal, total, order_status, date_approved, migrate_purchase FROM purchase_orders WHERE id = $1 AND deleted_at IS NULL"
	err := r.db.Get(&purchaseOrders, query, orderId)
	if err != nil {
		fmt.Printf("Error obteniendo orden de compra con ID %d: %v\n", orderId, err)
		return nil, err
	}

	var purchaseOrderDetails []domain.PurchaseOrderDetails
	detailQuery := `
		SELECT pod.id, pod.purchase_order_id, pod.product_id, pod.quantity, pod.price, pod.total FROM purchase_order_details pod WHERE pod.purchase_order_id = $1 AND pod.deleted_at IS NULL
	`
	err = r.db.Select(&purchaseOrderDetails, detailQuery, orderId)
	if err != nil {
		fmt.Printf("Error obteniendo detalles para la orden de compra %d: %v\n", orderId, err)
		return nil, err
	}

	purchaseOrders.PurchaseOrderDetails = purchaseOrderDetails

	return &purchaseOrders, nil
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
			reference = :reference,
			description = :description,
			warehouse_id = :warehouse_id,
			supplier_id = :supplier_id,
			supplier_document = :supplier_document,
			exchange_rate = :exchange_rate,
			discount = :discount,
			user_id = :user_id,
			date = :date,
			tax = :tax,
			subtotal = :subtotal,
			total = :total,
			order_status = :order_status,
			date_approved = :date_approved,
			migrate_purchase = :migrate_purchase,
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
		if detail.Id == 0 {
			// Insertar nuevo detalle
			detail.Purchase_Order_Id = order.Id
			detailQuery := `
				INSERT INTO sale_order_details (
					 purchase_order_id, product_id, quantity, price, total
				) VALUES (
					:purchase_order_id, :product_id, :quantity, :price, :total
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
				UPDATE sale_order_details
				SET
					product_id = :product_id,
					quantity = :quantity,
					price = :price,
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

// Eliminar (suavemente) una orden de compra y sus detalles
func (r *PurchaseOrderRepository) DeletePurchaseOrder(orderID int) error {
	tx, err := r.db.Beginx()
	if err != nil {
		fmt.Printf("Error iniciando transacción: %v\n", err)
		return err
	}

	orderQuery := "UPDATE purchase_orders SET deleted_at = NOW() WHERE id = $1 AND deleted_at IS NULL"
	_, err = tx.Exec(orderQuery, orderID)
	if err != nil {
		fmt.Printf("Error eliminando orden de compra: %v\n", err)
		tx.Rollback()
		return err
	}

	detailQuery := "UPDATE purchase_order_details SET deleted_at = NOW() WHERE purchase_order_id = $1 AND deleted_at IS NULL"
	_, err = tx.Exec(detailQuery, orderID)
	if err != nil {
		fmt.Printf("Error eliminando detalles de la orden: %v\n", err)
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	if err != nil {
		fmt.Printf("Error confirmando transacción: %v\n", err)
		return err
	}

	return nil
}
