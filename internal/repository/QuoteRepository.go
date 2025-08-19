package repository

import (
	"database/sql"
	"fmt"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/jmoiron/sqlx"
)

type QuoteRepository struct {
	db *sqlx.DB
}

func NewQuoteRepository(db *sqlx.DB) *QuoteRepository {
	return &QuoteRepository{db: db}
}

// Crear una cotización junto con sus detalles
func (r *QuoteRepository) CreateQuote(quote *domain.Quotes, details []domain.QuoteDetails) error {
	tx, err := r.db.Beginx()
	if err != nil {
		fmt.Printf("Error iniciando transacción: %v\n", err)
		return err
	}

	// Descomponer la estructura quote en sus campos individuales
	quoteMap := map[string]interface{}{
		"reference":       quote.Reference,
		"warehouse_id":    quote.Warehouse_Id,
		"customer_id":     quote.Customer_Id,
		"currency_id":     quote.Currency_Id,
		"user_id":         quote.User_Id,
		"issue_date":      quote.Issue_Date,
		"exchange_rate":   quote.Exchange_Rate,
		"expiration_date": quote.Expiration_Date,
		"approved_by":     quote.Approved_By,
		"approved_at":     quote.Approved_At,
		"canceled_by":     quote.Canceled_By,
		"canceled_at":     quote.Canceled_At,
		"discount":        quote.Discount,
		"subtotal":        quote.Subtotal,
		"total":           quote.Total,
		"quote_status":    quote.Quote_Status,
		"migrate_quote":   quote.Migrate_Quote,
	}

	// Insertar la cotización
	quoteQuery := `
		INSERT INTO quotes (
			reference, warehouse_id, customer_id, currency_id, user_id, issue_date, exchange_rate, expiration_date, approved_by, approved_at, canceled_by, canceled_at, discount, subtotal, total, quote_status, migrate_quote
		) VALUES (
			:reference, :warehouse_id, :customer_id, :currency_id, :user_id, :issue_date, :exchange_rate, :expiration_date, :approved_by, :approved_at, :canceled_by, :canceled_at, :discount, :subtotal, :total, :quote_status, :migrate_quote
		) RETURNING id
	`

	var quoteId int
	stmt, err := tx.PrepareNamed(quoteQuery) // Preparar la consulta nombrada
	if err != nil {
		tx.Rollback()
		fmt.Printf("Error preparando la consulta: %v\n", err)
		return err
	}
	defer stmt.Close()

	err = stmt.Get(&quoteId, quoteMap) // Ejecutar la consulta y obtener el ID
	if err != nil {
		tx.Rollback()
		fmt.Printf("Error insertando la cotización: %v\n", err)
		return err
	}

	// Insertar los detalles de la cotización
	for _, detail := range details {
		detail.Quote_Id = quoteId
		detailQuery := `
			INSERT INTO quote_details (
				product_name, quote_id, product_id, quantity, price, discount_method, discount, subtotal, total
			) VALUES (
				:product_name, :quote_id, :product_id, :quantity, :price, :discount_method, :discount, :subtotal, :total
			)
		`
		_, err = tx.NamedExec(detailQuery, detail)
		if err != nil {
			fmt.Printf("Error creando el detalle de la cotización: %v\n", err)
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

// Obtener todas las cotizaciones
func (r *QuoteRepository) GetAllQuotes() ([]domain.Quotes, error) {
	var quotes []domain.Quotes
	query := `
		SELECT
			q.id, q.reference, q.warehouse_id, q.customer_id, q.currency_id, q.user_id, q.issue_date, q.exchange_rate, q.expiration_date, q.approved_by, q.approved_at, q.canceled_by, q.canceled_at, q.discount, q.subtotal, q.total, q.quote_status, q.migrate_quote,
			c.id AS "customer.id", c.names AS "customer.names", c.surname AS "customer.surname", c.second_surname AS "customer.second_surname", c.company_name AS "customer.company_name",
			u.id AS "user.id", u.employee_id AS "user.employee_id",
			cu.id AS "currency.id", cu.name AS "currency.name", cu.code AS "currency.code", cu.symbol AS "currency.symbol",
			e.id AS "user.employee.id", e.names AS "user.employee.names", e.surname AS "user.employee.surname", e.second_surname AS "user.employee.second_surname",
			w.id AS "warehouse.id", w.branch_office_id AS "warehouse.branch_office_id", w.name AS "warehouse.name",
			bo.id AS "warehouse.branch_office.id", bo.company_id AS "warehouse.branch_office.company_id", bo.name AS "warehouse.branch_office.name",
			co.id AS "warehouse.branch_office.company.id", co.name AS "warehouse.branch_office.company.name", co.ruc AS "warehouse.branch_office.company.ruc"
		FROM quotes q
		LEFT JOIN customers c ON q.customer_id=c.id
		LEFT JOIN currencies cu ON q.currency_id=cu.id
		LEFT JOIN warehouses w ON q.warehouse_id=w.id
		LEFT JOIN branch_offices bo ON w.branch_office_id=bo.id
		LEFT JOIN companies co ON bo.company_id=co.id
		LEFT JOIN users u ON q.user_id=u.id
		LEFT JOIN employees e ON u.employee_id=e.id
		WHERE q.deleted_at IS NULL
		ORDER BY q.id DESC
	`
	err := r.db.Select(&quotes, query)
	if err != nil {
		fmt.Printf("Error obteniendo cotizaciones: %v\n", err)
		return nil, err
	}

	// Iterar sobre las cotizaciones y agregar los detalles de cada cotización
	for i := range quotes {
		var quoteDetails []domain.QuoteDetails
		// Consulta para obtener los detalles de cada cotización
		detailsQuery := `
			SELECT
				qd.id, qd.product_name, qd.quote_id, qd.product_id, qd.quantity, qd.price, qd.discount_method, qd.discount, qd.subtotal, qd.total,
				p.id AS "product.id", p.name AS "product.name", p.featured_pcf AS "product.featured_pcf", p.cost AS "product.cost"
			FROM quote_details qd
			LEFT JOIN products p ON qd.product_id=p.id
			WHERE qd.quote_id = $1
		`

		err := r.db.Select(&quoteDetails, detailsQuery, quotes[i].Id)
		if err != nil {
			fmt.Printf("Error obteniendo detalles de la cotización %d: %v\n", quotes[i].Id, err)
			return nil, err
		}

		// Asignar los detalles a la cotización correspondiente
		quotes[i].QuoteDetails = quoteDetails
	}

	return quotes, nil
}

// Obtener una cotización por ID junto con sus detalles
func (r *QuoteRepository) GetQuoteById(quoteId int) (*domain.Quotes, error) {
	var quotes domain.Quotes
	query := `
		SELECT
			q.id, q.reference, q.warehouse_id, q.customer_id, q.currency_id, q.user_id, q.issue_date, q.exchange_rate, q.expiration_date, q.approved_by, q.approved_at, q.canceled_by, q.canceled_at, q.discount, q.subtotal, q.total, q.quote_status, q.migrate_quote,
			c.id AS "customer.id", c.names AS "customer.names", c.surname AS "customer.surname", c.second_surname AS "customer.second_surname", c.company_name AS "customer.company_name",
			u.id AS "user.id", u.employee_id AS "user.employee_id",
			cu.id AS "currency.id", cu.name AS "currency.name", cu.code AS "currency.code", cu.symbol AS "currency.symbol",
			e.id AS "user.employee.id", e.names AS "user.employee.names", e.surname AS "user.employee.surname", e.second_surname AS "user.employee.second_surname",
			w.id AS "warehouse.id", w.branch_office_id AS "warehouse.branch_office_id", w.name AS "warehouse.name",
			bo.id AS "warehouse.branch_office.id", bo.company_id AS "warehouse.branch_office.company_id", bo.name AS "warehouse.branch_office.name",
			co.id AS "warehouse.branch_office.company.id", co.name AS "warehouse.branch_office.company.name", co.ruc AS "warehouse.branch_office.company.ruc"
		FROM quotes q
		LEFT JOIN customers c ON q.customer_id=c.id
		LEFT JOIN currencies cu ON q.currency_id=cu.id
		LEFT JOIN warehouses w ON q.warehouse_id=w.id
		LEFT JOIN branch_offices bo ON w.branch_office_id=bo.id
		LEFT JOIN companies co ON bo.company_id=co.id
		LEFT JOIN users u ON q.user_id=u.id
		LEFT JOIN employees e ON u.employee_id=e.id
		WHERE q.id = $1 AND q.deleted_at IS NULL
	`
	err := r.db.Get(&quotes, query, quoteId)
	if err != nil {
		fmt.Printf("Error obteniendo cotización con ID %d: %v\n", quoteId, err)
		return nil, err
	}
	var quoteDetails []domain.QuoteDetails
	detailQuery := `
		SELECT
			qd.id, qd.product_name, qd.quote_id, qd.product_id, qd.quantity, qd.price, qd.discount_method, qd.discount, qd.subtotal, qd.total,
			p.id AS "product.id", p.name AS "product.name", p.featured_pcf AS "product.featured_pcf", p.cost AS "product.cost"
		FROM quote_details qd
		LEFT JOIN products p ON qd.product_id=p.id
		WHERE qd.quote_id = $1 AND qd.deleted_at IS NULL
	`
	err = r.db.Select(&quoteDetails, detailQuery, quoteId)
	if err != nil {
		fmt.Printf("Error obteniendo detalles para la cotización %d: %v\n", quoteId, err)
		return nil, err
	}

	quotes.QuoteDetails = quoteDetails

	return &quotes, nil
}

// Actualizar una cotización
func (r *QuoteRepository) UpdateQuote(quote *domain.Quotes, details []domain.QuoteDetails) error {
	tx, err := r.db.Beginx()
	if err != nil {
		fmt.Printf("Error iniciando transacción: %v\n", err)
		return err
	}

	// Actualizar la cotización
	quoteQuery := `
		UPDATE quotes
		SET
			reference = :reference,
			warehouse_id = :warehouse_id,
			customer_id = :customer_id,
			currency_id = :currency_id,
			user_id = :user_id,
			issue_date = :issue_date,
			exchange_rate = :exchange_rate,
			expiration_date = :expiration_date,
			approved_by = :approved_by,
			approved_at = :approved_at,
			canceled_by = :canceled_by,
			canceled_at = :canceled_at,
			discount = :discount,
			subtotal = :subtotal,
			total = :total,
			quote_status = :quote_status,
			migrate_quote = :migrate_quote,
			updated_at = NOW()
		WHERE id = :id
	`
	_, err = tx.NamedExec(quoteQuery, quote)
	if err != nil {
		fmt.Printf("Error actualizando la cotización: %v\n", err)
		tx.Rollback()
		return err
	}

	// Actualizar o insertar los detalles de la cotización
	for _, detail := range details {
		cadena := fmt.Sprintf("%d", detail.Id)
		longitud := len(cadena)
		if longitud == 13 {
			// Insertar nuevo detalle
			detail.Quote_Id = quote.Id
			detailQuery := `
				INSERT INTO quote_details (
					product_name, quote_id, product_id, quantity, price, discount_method, discount, subtotal, total
				) VALUES (
					:product_name, :quote_id, :product_id, :quantity, :price, :discount_method, :discount, :subtotal, :total
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
				UPDATE quote_details
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

// Obtener la última referencia registrada
func (r *QuoteRepository) GetLastQuoteReference() (string, error) {
	var lastReference string

	query := `SELECT reference FROM quotes ORDER BY id DESC LIMIT 1`
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
