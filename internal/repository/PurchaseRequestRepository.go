package repository

import (
	"database/sql"
	"fmt"
	"reflect"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/jmoiron/sqlx"
)

type PurchaseRequestRepository struct {
	db *sqlx.DB
}

func NewPurchaseRequestRepository(db *sqlx.DB) *PurchaseRequestRepository {
	return &PurchaseRequestRepository{db: db}
}

// Crear una orden de venta junto con sus detalles
func (r *PurchaseRequestRepository) CreatePurchaseRequest(order *domain.PurchaseRequests, details []domain.PurchaseRequestDetails) error {
	tx, err := r.db.Beginx()
	if err != nil {
		fmt.Printf("Error iniciando transacción: %v\n", err)
		return err
	}

	// Insertar la orden de venta
	orderQuery := `
		INSERT INTO purchase_requests (
			reference, warehouse_id, customer_id, user_id, date, tax, discount, subtotal, total, payment_id, order_status
		) VALUES (
			:reference, :warehouse_id, :customer_id, :user_id, :date, :tax, :discount, :subtotal, :total, :payment_id, :order_status
		) RETURNING id
	`
	var orderId int
	err = tx.QueryRowx(orderQuery, order).Scan(&orderId)
	if err != nil {
		fmt.Printf("Error creando la orden de venta: %v\n", err)
		tx.Rollback()
		return err
	}

	// Insertar los detalles de la orden
	for _, detail := range details {
		detail.Id = 0
		detail.Purchase_Request_Id = orderId
		detailQuery := `
			INSERT INTO purchase_request_details (
				product_name, purchase_request_id, product_id, quantity, igv, discount_method, discount, price, subtotal, total
			) VALUES (
				:product_name, :purchase_request_id, :product_id, :quantity, :igv, :discount_method, :discount, :price, :subtotal, :total
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

// Obtener todas las órdenes de venta
func (r *PurchaseRequestRepository) GetAllPurchaseRequests() ([]domain.PurchaseRequests, error) {
	var purchaseRequests []domain.PurchaseRequests
	query := "SELECT * FROM purchase_requests WHERE deleted_at IS NULL"
	err := r.db.Select(&purchaseRequests, query)
	if err != nil {
		fmt.Printf("Error obteniendo órdenes de venta: %v\n", err)
		return nil, err
	}

	// Iterar sobre las órdenes de venta y agregar los detalles de cada orden
	for i := range purchaseRequests {
		// Obtener usuario de la orden de venta
		var user domain.Users
		err = r.db.Get(&user, `
			SELECT id, role_id, display_name, photo_url, email, status FROM users WHERE id = $1 AND deleted_at IS NULL
		`, purchaseRequests[i].User_Id)
		if err != nil && err != sql.ErrNoRows { // Manejar casos donde no hay usuario asociado
			fmt.Printf("Error obteniendo el usuario para la orden de venta %d: %v\n", purchaseRequests[i].Id, err)
			return nil, err
		}
		purchaseRequests[i].User = &user

		var purchaseRequestDetails []domain.PurchaseRequestDetails
		// Consulta para obtener los detalles de cada orden de venta
		detailsQuery := `
			SELECT * FROM purchase_request_details WHERE purchase_request_id = $1
		`

		err := r.db.Select(&purchaseRequestDetails, detailsQuery, purchaseRequests[i].Id)
		if err != nil {
			fmt.Printf("Error obteniendo detalles de la orden de venta %d: %v\n", purchaseRequests[i].Id, err)
			return nil, err
		}

		// Obtener productos de los detalles
		for i, detail := range purchaseRequestDetails {
			if detail.Product_Id.Valid { // Si hay un producto asociado
				var product domain.Products
				err = r.db.Get(&product, `
					SELECT * FROM products WHERE id = $1 AND deleted_at IS NULL
				`, detail.Product_Id.Int)
				if err != nil && err != sql.ErrNoRows {
					fmt.Printf("Error obteniendo producto %d para la orden de venta %d: %v\n", *detail.Product_Id.Int, purchaseRequests[i].Id, err)
					return nil, err
				}
				if err == nil { // Si se encontró el producto, asignarlo
					purchaseRequestDetails[i].Product = &product
				}
			}
		}

		// Asignar los detalles a la orden correspondiente
		purchaseRequests[i].PurchaseRequestDetails = purchaseRequestDetails
	}

	return purchaseRequests, nil
}

// Obtener una orden de venta por ID junto con sus detalles
func (r *PurchaseRequestRepository) GetPurchaseRequestById(orderId int) (*domain.PurchaseRequests, error) {
	var purchaseRequests domain.PurchaseRequests
	query := "SELECT * FROM purchase_requests WHERE id = $1 AND deleted_at IS NULL"
	err := r.db.Get(&purchaseRequests, query, orderId)
	if err != nil {
		fmt.Printf("Error obteniendo orden de venta con ID %d: %v\n", orderId, err)
		return nil, err
	}

	// Obtener usuario de la orden de venta
	var user domain.Users
	err = r.db.Get(&user, `
		SELECT id, role_id, display_name, employee_id, photo_url, email, status FROM users WHERE id = $1 AND deleted_at IS NULL
	`, purchaseRequests.User_Id)
	if err != nil && err != sql.ErrNoRows { // Manejar casos donde no hay usuario asociado
		fmt.Printf("Error obteniendo usuario para la orden de venta %d: %v\n", orderId, err)
		return nil, err
	}
	purchaseRequests.User = &user

	// Obtener emploeado del usuario (si existe)
	if user.Employee_Id.Valid {
		var employee domain.Employees
		err = r.db.Get(&employee, `
			SELECT * FROM employees WHERE id = $1 AND deleted_at IS NULL
		`, user.Employee_Id)

		if err != nil && err != sql.ErrNoRows {
			return nil, err
		}

		if err == nil {
			user.Employee = &employee
		}
	}

	var purchaseRequestDetails []domain.PurchaseRequestDetails
	detailQuery := `
		SELECT * FROM purchase_request_details WHERE purchase_request_id = $1 AND deleted_at IS NULL
	`
	err = r.db.Select(&purchaseRequestDetails, detailQuery, orderId)
	if err != nil {
		fmt.Printf("Error obteniendo detalles para la orden de venta %d: %v\n", orderId, err)
		return nil, err
	}

	// Obtener productos de los detalles
	for i, detail := range purchaseRequestDetails {
		if detail.Product_Id.Valid { // Si hay un producto asociado
			var product domain.Products
			err = r.db.Get(&product, `
				SELECT * FROM products WHERE id = $1 AND deleted_at IS NULL
			`, detail.Product_Id.Int)
			if err != nil && err != sql.ErrNoRows {
				fmt.Printf("Error obteniendo producto %d para la orden de venta %d: %v\n", *detail.Product_Id.Int, orderId, err)
				return nil, err
			}
			if err == nil { // Si se encontró el producto, asignarlo
				purchaseRequestDetails[i].Product = &product
			}
		}
	}

	purchaseRequests.PurchaseRequestDetails = purchaseRequestDetails

	return &purchaseRequests, nil
}

// Actualizar una orden de venta
func (r *PurchaseRequestRepository) UpdatePurchaseRequest(order *domain.PurchaseRequests, details []domain.PurchaseRequestDetails) error {
	tx, err := r.db.Beginx()
	if err != nil {
		fmt.Printf("Error iniciando transacción: %v\n", err)
		return err
	}

	// Actualizar la orden de venta
	orderQuery := `
		UPDATE purchase_requests
		SET
			reference = :reference,
			warehouse_id = :warehouse_id,
			customer_id = :customer_id,
			user_id = :user_id,
			tax = :tax,
			discount = :discount,
			subtotal = :subtotal,
			total = :total,
			payment_id = :payment_id,
			date = :date,
			date_approved = :date_approved,
			updated_at = NOW()
		WHERE id = :id
	`
	_, err = tx.NamedExec(orderQuery, order)
	if err != nil {
		fmt.Printf("Error actualizando la orden de venta: %v\n", err)
		tx.Rollback()
		return err
	}

	// Actualizar o insertar los detalles de la orden
	for _, detail := range details {
		if detail.Id == 0 || reflect.TypeOf(detail.Id).Kind() == reflect.String {
			detail.Id = 0
			// Insertar nuevo detalle
			detail.Purchase_Request_Id = order.Id
			detailQuery := `
				INSERT INTO purchase_request_details (
					product_name, purchase_request_id, product_id, quantity, igv, discount_method, discount, price, subtotal, total
				) VALUES (
					:product_name, :purchase_request_id, :product_id, :quantity, :igv, :discount_method, :discount, :price, :subtotal, :total
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
				UPDATE purchase_request_details
				SET
					product_name = :product_name,
					product_id = :product_id,
					quantity = :quantity,
					igv = :igv,
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

// Eliminar (suavemente) una orden de venta y sus detalles
func (r *PurchaseRequestRepository) DeletePurchaseRequest(orderID int) error {
	tx, err := r.db.Beginx()
	if err != nil {
		fmt.Printf("Error iniciando transacción: %v\n", err)
		return err
	}

	orderQuery := "UPDATE purchase_requests SET deleted_at = NOW() WHERE id = $1 AND deleted_at IS NULL"
	_, err = tx.Exec(orderQuery, orderID)
	if err != nil {
		fmt.Printf("Error eliminando orden de venta: %v\n", err)
		tx.Rollback()
		return err
	}

	detailQuery := "UPDATE purchase_request_details SET deleted_at = NOW() WHERE purchase_request_id = $1 AND deleted_at IS NULL"
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

// Obtener la última referencia registrada
func (r *PurchaseRequestRepository) GetLastPurchaseRequestReference() (string, error) {
	var lastReference string

	query := `SELECT reference FROM purchase_requests ORDER BY id DESC LIMIT 1`
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
