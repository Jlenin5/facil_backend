package repository

import (
	"database/sql"
	"fmt"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/jmoiron/sqlx"
)

type SaleOrderRepository struct {
	db *sqlx.DB
}

func NewSaleOrderRepository(db *sqlx.DB) *SaleOrderRepository {
	return &SaleOrderRepository{db: db}
}

// Crear una orden de venta junto con sus detalles
func (r *SaleOrderRepository) CreateSaleOrder(order *domain.SaleOrders, details []domain.SaleOrderDetails) error {
	tx, err := r.db.Beginx()
	if err != nil {
		fmt.Printf("Error iniciando transacción: %v\n", err)
		return err
	}

	// Descomponer la estructura sale en sus campos individuales
	orderMap := map[string]interface{}{
		"reference":          order.Reference,
		"warehouse_id":       order.Warehouse_Id,
		"customer_id":        order.Customer_Id,
		"currency_id":        order.Currency_Id,
		"user_id":            order.User_Id,
		"issue_date":         order.Issue_Date,
		"exchange_rate":      order.Exchange_Rate,
		"discount":           order.Discount,
		"subtotal":           order.Subtotal,
		"total":              order.Total,
		"order_status":       order.Order_Status,
		"date_approved":      order.Date_Approved,
		"migrate_sale_order": order.Migrate_Sale_Order,
		"quote_id":           order.Quote_Id,
	}

	// Insertar la orden de venta
	orderQuery := `
		INSERT INTO sale_orders (
			reference, warehouse_id, customer_id, currency_id, user_id, issue_date, exchange_rate, discount, subtotal, total, order_status, date_approved, migrate_sale_order, quote_id
		) VALUES (
			:reference, :warehouse_id, :customer_id, :currency_id, :user_id, :issue_date, :exchange_rate, :discount, :subtotal, :total, :order_status, :date_approved, :migrate_sale_order, :quote_id
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
		fmt.Printf("Error insertando la orden de venta: %v\n", err)
		return err
	}

	// Insertar los detalles de la orden
	for _, detail := range details {
		detail.Id = 0
		detail.Sale_Order_Id = orderId
		detailQuery := `
			INSERT INTO sale_order_details (
				product_name, sale_order_id, product_id, quantity, discount_method, discount, price, subtotal, total
			) VALUES (
				:product_name, :sale_order_id, :product_id, :quantity, :discount_method, :discount, :price, :subtotal, :total
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
func (r *SaleOrderRepository) GetAllSaleOrders() ([]domain.SaleOrders, error) {
	var saleOrders []domain.SaleOrders
	query := `
		SELECT
			so.id, so.reference, so.warehouse_id, so.customer_id, so.currency_id, so.user_id, so.issue_date, so.exchange_rate, so.discount, so.subtotal, so.total, so.order_status, so.date_approved, so.migrate_sale_order, so.quote_id,
			c.id AS "customer.id", c.first_name AS "customer.first_name", c.second_name AS "customer.second_name", c.third_name AS "customer.third_name", c.surname AS "customer.surname", c.second_surname AS "customer.second_surname", c.company_name AS "customer.company_name",
			u.id AS "user.id", u.employee_id AS "user.employee_id",
			cu.id AS "currency.id", cu.name AS "currency.name", cu.code AS "currency.code", cu.symbol AS "currency.symbol",
			e.id AS "user.employee.id", e.first_name AS "user.employee.first_name", e.second_name AS "user.employee.second_name", e.third_name AS "user.employee.third_name", e.surname AS "user.employee.surname", e.second_surname AS "user.employee.second_surname",
			w.id AS "warehouse.id", w.branch_office_id AS "warehouse.branch_office_id", w.name AS "warehouse.name",
			bo.id AS "warehouse.branch_office.id", bo.company_id AS "warehouse.branch_office.company_id", bo.name AS "warehouse.branch_office.name",
			co.id AS "warehouse.branch_office.company.id", co.name AS "warehouse.branch_office.company.name", co.ruc AS "warehouse.branch_office.company.ruc",
			COALESCE(q.id, 0) AS "quote.id", COALESCE(q.reference, '') AS "quote.reference"
		FROM sale_orders so
		LEFT JOIN customers c ON so.customer_id=c.id
		LEFT JOIN currencies cu ON so.currency_id=cu.id
		LEFT JOIN warehouses w ON so.warehouse_id=w.id
		LEFT JOIN branch_offices bo ON w.branch_office_id=bo.id
		LEFT JOIN companies co ON bo.company_id=co.id
		LEFT JOIN users u ON so.user_id=u.id
		LEFT JOIN employees e ON u.employee_id=e.id
		LEFT JOIN quotes q ON so.quote_id=q.id
		WHERE so.deleted_at IS NULL
		ORDER BY so.id DESC
	`
	err := r.db.Select(&saleOrders, query)
	if err != nil {
		fmt.Printf("Error obteniendo órdenes de venta: %v\n", err)
		return nil, err
	}

	// Iterar sobre las órdenes de venta y agregar los detalles de cada orden
	for i := range saleOrders {
		var saleOrderDetails []domain.SaleOrderDetails
		// Consulta para obtener los detalles de cada orden de venta
		detailsQuery := `
			SELECT
				sod.id, sod.product_name, sod.sale_order_id, sod.product_id, sod.quantity, sod.price, sod.discount_method, sod.discount, sod.subtotal, sod.total,
				p.id AS "product.id", p.name AS "product.name", p.price AS "product.price", p.cost AS "product.cost"
			FROM sale_order_details sod
			LEFT JOIN products p ON sod.product_id=p.id
			WHERE sod.sale_order_id = $1
		`

		err := r.db.Select(&saleOrderDetails, detailsQuery, saleOrders[i].Id)
		if err != nil {
			fmt.Printf("Error obteniendo detalles de la orden de venta %d: %v\n", saleOrders[i].Id, err)
			return nil, err
		}

		// Asignar los detalles a la orden correspondiente
		saleOrders[i].SaleOrderDetails = saleOrderDetails
	}

	return saleOrders, nil
}

// Obtener una orden de venta por ID junto con sus detalles
func (r *SaleOrderRepository) GetSaleOrderById(orderId int) (*domain.SaleOrders, error) {
	var saleOrders domain.SaleOrders
	query := `
		SELECT
			so.id, so.reference, so.warehouse_id, so.customer_id, so.currency_id, so.user_id, so.issue_date, so.exchange_rate, so.discount, so.subtotal, so.total, so.order_status, so.date_approved, so.migrate_sale_order, so.quote_id,
			c.id AS "customer.id", c.first_name AS "customer.first_name", c.second_name AS "customer.second_name", c.third_name AS "customer.third_name", c.surname AS "customer.surname", c.second_surname AS "customer.second_surname", c.company_name AS "customer.company_name",
			u.id AS "user.id", u.employee_id AS "user.employee_id",
			cu.id AS "currency.id", cu.name AS "currency.name", cu.code AS "currency.code", cu.symbol AS "currency.symbol",
			e.id AS "user.employee.id", e.first_name AS "user.employee.first_name", e.second_name AS "user.employee.second_name", e.third_name AS "user.employee.third_name", e.surname AS "user.employee.surname", e.second_surname AS "user.employee.second_surname",
			w.id AS "warehouse.id", w.branch_office_id AS "warehouse.branch_office_id", w.name AS "warehouse.name",
			bo.id AS "warehouse.branch_office.id", bo.company_id AS "warehouse.branch_office.company_id", bo.name AS "warehouse.branch_office.name",
			co.id AS "warehouse.branch_office.company.id", co.name AS "warehouse.branch_office.company.name", co.ruc AS "warehouse.branch_office.company.ruc",
			COALESCE(q.id, 0) AS "quote.id", COALESCE(q.reference, '') AS "quote.reference"
		FROM sale_orders so
		LEFT JOIN customers c ON so.customer_id=c.id
		LEFT JOIN currencies cu ON so.currency_id=cu.id
		LEFT JOIN warehouses w ON so.warehouse_id=w.id
		LEFT JOIN branch_offices bo ON w.branch_office_id=bo.id
		LEFT JOIN companies co ON bo.company_id=co.id
		LEFT JOIN users u ON so.user_id=u.id
		LEFT JOIN employees e ON u.employee_id=e.id
		LEFT JOIN quotes q ON so.quote_id=q.id
		WHERE so.id = $1 AND so.deleted_at IS NULL
	`
	err := r.db.Get(&saleOrders, query, orderId)
	if err != nil {
		fmt.Printf("Error obteniendo orden de venta con ID %d: %v\n", orderId, err)
		return nil, err
	}

	var saleOrderDetails []domain.SaleOrderDetails
	detailQuery := `
		SELECT
				sod.id, sod.product_name, sod.sale_order_id, sod.product_id, sod.quantity, sod.price, sod.discount_method, sod.discount, sod.subtotal, sod.total,
				p.id AS "product.id", p.name AS "product.name", p.price AS "product.price", p.cost AS "product.cost"
			FROM sale_order_details sod
			LEFT JOIN products p ON sod.product_id=p.id
			WHERE sod.sale_order_id = $1 AND sod.deleted_at IS NULL
	`
	err = r.db.Select(&saleOrderDetails, detailQuery, orderId)
	if err != nil {
		fmt.Printf("Error obteniendo detalles para la orden de venta %d: %v\n", orderId, err)
		return nil, err
	}

	saleOrders.SaleOrderDetails = saleOrderDetails

	return &saleOrders, nil
}

// Actualizar una orden de venta
func (r *SaleOrderRepository) UpdateSaleOrder(order *domain.SaleOrders, details []domain.SaleOrderDetails) error {
	tx, err := r.db.Beginx()
	if err != nil {
		fmt.Printf("Error iniciando transacción: %v\n", err)
		return err
	}

	// Actualizar la orden de venta
	orderQuery := `
		UPDATE sale_orders
		SET
			reference =          :reference,
			warehouse_id =       :warehouse_id,
			customer_id =        :customer_id,
			currency_id =        :currency_id,
			user_id =            :user_id,
			issue_date =         :issue_date,
			exchange_rate =      :exchange_rate,
			discount =           :discount,
			subtotal =           :subtotal,
			total =              :total,
			order_status =       :order_status,
			date_approved =      :date_approved,
			migrate_sale_order = :migrate_sale_order,
			quote_id =           :quote_id,
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
		cadena := fmt.Sprintf("%d", detail.Id)
		longitud := len(cadena)
		if longitud == 13 {
			// Insertar nuevo detalle
			detail.Sale_Order_Id = order.Id
			detailQuery := `
				INSERT INTO sale_order_details (
					product_name, sale_order_id, product_id, quantity, discount_method, discount, price, subtotal, total
				) VALUES (
					:product_name, :sale_order_id, :product_id, :quantity, :discount_method, :discount, :price, :subtotal, :total
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
					product_name = :product_name,
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

// Obtener la última referencia registrada
func (r *SaleOrderRepository) GetLastSaleOrderReference() (string, error) {
	var lastReference string

	query := `SELECT reference FROM sale_orders ORDER BY id DESC LIMIT 1`
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
