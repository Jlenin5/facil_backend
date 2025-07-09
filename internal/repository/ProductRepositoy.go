package repository

import (
	"fmt"

	"github.com/Jlenin5/facil_backend/internal/domain"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
)

type ProductRepository struct {
	db *sqlx.DB
}

func NewProductRepository(db *sqlx.DB) *ProductRepository {
	return &ProductRepository{db: db}
}

// Crear un producto
func (r *ProductRepository) CreateProduct(product *domain.Products) error {
	// Inicia una transacción para asegurar atomicidad
	// Crear un mapa para los valores del producto

	// Insertar el producto en la tabla `products`
	queryProduct := `
		INSERT INTO products (
			name, prices_cf, prices_sf, prices_box, featured_pcf, featured_psf, featured_pbox, cost, sku, unit_of_measurement_id, quantity_in_box, created_by
		) VALUES (
			:name, :prices_cf, :prices_sf, :prices_box, :featured_pcf, :featured_psf, :featured_pbox, :cost, :sku, :unit_of_measurement_id, :quantity_in_box, :created_by
		) RETURNING id
	`

	// Finalizar la transacción
	_, err := r.db.NamedExec(queryProduct, product)
	if err != nil {
		fmt.Printf("Error al confirmar la transacción: %v\n", err)
		return err
	}

	return nil
}

// Crear varios productos
func (r *ProductRepository) SaveProducts(products []domain.Products) error {
	// Iniciar una transacción
	tx, err := r.db.Beginx()
	if err != nil {
		return fmt.Errorf("failed to begin transaction: %w", err)
	}

	// Preparar la consulta para insertar productos
	query := `
		INSERT INTO products (sku, name, prices_cf, prices_sf, prices_box, featured_pcf, featured_psf, featured_pbox, cost, created_by)
		VALUES (:sku, :name, :prices_cf, :prices_sf, :prices_box, :featured_pcf, :featured_psf, :featured_pbox, :cost, :created_by)
	`

	// Insertar cada producto
	for _, product := range products {
		_, err := tx.NamedExec(query, map[string]interface{}{
			"sku":            product.SKU,
			"name":           product.Name,
			"prices_cf":      product.Prices_cf,
			"prices_sf":      product.Prices_sf,
			"prices_box":     product.Prices_box,
			"featured_pcf":   product.Featured_Pcf,
			"featured_psf":   product.Featured_Psf,
			"featured_pbox":  product.Featured_Pbox,
			"cost":           product.Cost,
			"created_by":     product.Created_By,
		})
		if err != nil {
			tx.Rollback()
			return fmt.Errorf("failed to insert product %s: %w", product.Name, err)
		}
	}

	// Confirmar la transacción
	if err := tx.Commit(); err != nil {
		return fmt.Errorf("failed to commit transaction: %w", err)
	}

	return nil
}

// Obtener todos los productos
func (r *ProductRepository) GetAllProducts() ([]domain.Products, error) {
	var products []domain.Products
	query := `
		SELECT
			p.id, p.name, p.brand_id, p.unit_of_measurement_id, p.handle, p.description, p.tags,
			p.featured_image, p.images, p.prices_cf, p.prices_sf, p.prices_box, p.featured_pcf, p.featured_psf, p.featured_pbox, p.quantity_in_box, p.created_at, p.updated_at,
			p.cost, p.tax_rate, p.quantity, p.sku, p.width, p.height, p.depth,
			p.liters, p.weight, p.barcode, p.rating, p.extra_shipping_fee, p.status,
			p.created_by, p.updated_by,
			COALESCE(b.id, 0) AS "brand.id", COALESCE(b.name, '') AS "brand.name",
			COALESCE(uom.id, 0) AS "unit_of_measurement.id", COALESCE(uom.name, '') AS "unit_of_measurement.name", COALESCE(uom.shortcut, '') AS "unit_of_measurement.shortcut"
		FROM products p
		LEFT JOIN brands b ON p.brand_id=b.id
		LEFT JOIN units_of_measurement uom ON p.unit_of_measurement_id=uom.id
		WHERE p.deleted_at IS NULL
		ORDER BY p.id DESC
	`
	err := r.db.Select(&products, query)
	if err != nil {
		fmt.Printf("Error obteniendo productos: %v\n", err)
		return nil, err
	}

	// Iterar sobre los productos
	for i := range products {
		// Obtener categorías del producto
		var categories []domain.CategoryReducedData
		err := r.db.Select(&categories, `
			SELECT c.id, c.name
			FROM categories c
			INNER JOIN product_categories pc ON c.id = pc.category_id
			WHERE pc.product_id = $1
		`, products[i].Id)
		if err != nil {
			fmt.Printf("Error obteniendo categorías para el producto %d: %v\n", products[i].Id, err)
			return nil, err
		}
		products[i].Categories = categories

		// Obtener stock del producto
		var stock int
		err = r.db.Get(&stock, `
			SELECT
				COALESCE(SUM(current_stock), 0)
			FROM stock_control
			WHERE product_id = $1
		`, products[i].Id)
		if err != nil {
			fmt.Printf("Error obteniendo total de stock para producto %d: %v\n", products[i].Id, err)
			return nil, err
		}
		products[i].Stock = stock

		// Obtener booking del producto
		var booking []domain.Booking
		err = r.db.Select(&booking, `
			SELECT 
				warehouse_id,
				SUM(total) AS total
			FROM (
				-- Compras pendientes
				SELECT 
					po.warehouse_id AS warehouse_id,
					SUM(pod.quantity) AS total
				FROM 
					purchase_orders po
				JOIN 
					purchase_order_details pod ON po.id = pod.purchase_order_id
				WHERE 
					po.order_status = 'pending' AND pod.product_id = $1
				GROUP BY 
					po.warehouse_id

				UNION ALL

				-- Ventas pendientes
				SELECT 
					so.warehouse_id AS warehouse_id,
					SUM(sod.quantity) AS total
				FROM 
					sale_orders so
				JOIN 
					sale_order_details sod ON so.id = sod.sale_order_id
				WHERE 
					so.order_status = 'pending' AND sod.product_id = $1
				GROUP BY 
					so.warehouse_id
			) AS combined
			GROUP BY 
				warehouse_id
		`, products[i].Id)
		if err != nil {
			fmt.Printf("Error obteniendo booking para el producto %d: %v\n", products[i].Id, err)
			return nil, err
		}
		products[i].Booking = booking
	}

	return products, nil
}

// Obtener un producto por Id
func (r *ProductRepository) GetProductById(productId int) (*domain.Products, error) {
	var product domain.Products
	query := `
		SELECT
			p.id, p.name, p.brand_id, p.unit_of_measurement_id, p.handle, p.description, p.tags,
			p.featured_image, p.images, p.prices_cf, p.prices_sf, p.prices_box, p.featured_pcf, p.featured_psf, p.featured_pbox, p.quantity_in_box,
			p.cost, p.tax_rate, p.quantity, p.sku, p.width, p.height, p.depth,
			p.liters, p.weight, p.barcode, p.rating, p.extra_shipping_fee, p.status,
			p.created_by, p.updated_by,
			COALESCE(b.id, 0) AS "brand.id", COALESCE(b.name, '') AS "brand.name",
			COALESCE(uom.id, 0) AS "unit_of_measurement.id", COALESCE(uom.name, '') AS "unit_of_measurement.name", COALESCE(uom.shortcut, '') AS "unit_of_measurement.shortcut"
		FROM products p
		LEFT JOIN brands b ON p.brand_id=b.id
		LEFT JOIN units_of_measurement uom ON p.unit_of_measurement_id=uom.id
		WHERE p.id = $1 AND p.deleted_at IS NULL
		ORDER BY p.id DESC
	`
	err := r.db.Get(&product, query, productId)
	if err != nil {
		fmt.Printf("Error obteniendo producto con ID %d: %v\n", productId, err)
		return nil, err
	}

	// Query para obtener las categorías relacionadas con el producto
	var categories []domain.CategoryReducedData
	err = r.db.Select(&categories, `
		SELECT c.id, c.name
		FROM categories c
		INNER JOIN product_categories pc ON c.id = pc.category_id
		WHERE pc.product_id = $1
	`, productId)
	if err != nil {
		fmt.Printf("Error obteniendo categorías para el producto %d: %v\n", productId, err)
		return nil, err
	}
	product.Categories = categories

	// Obtener stock del producto
	var stock int
	err = r.db.Get(&stock, `
			SELECT
				COALESCE(SUM(current_stock), 0)
			FROM stock_control
			WHERE product_id = $1
		`, productId)
	if err != nil {
		fmt.Printf("Error obteniendo total de stock para producto %d: %v\n", productId, err)
		return nil, err
	}
	product.Stock = stock

	// Obtener booking del producto
	var booking []domain.Booking
	err = r.db.Select(&booking, `
		SELECT 
			warehouse_id,
			SUM(total) AS total
		FROM (
			-- Compras pendientes
			SELECT 
				po.warehouse_id AS warehouse_id,
				SUM(pod.quantity) AS total
			FROM 
				purchase_orders po
			JOIN 
				purchase_order_details pod ON po.id = pod.purchase_order_id
			WHERE 
				po.order_status = 'pending' AND pod.product_id = $1
			GROUP BY 
				po.warehouse_id

			UNION ALL

			-- Ventas pendientes
			SELECT 
				so.warehouse_id AS warehouse_id,
				SUM(sod.quantity) AS total
			FROM 
				sale_orders so
			JOIN 
				sale_order_details sod ON so.id = sod.sale_order_id
			WHERE 
				so.order_status = 'pending' AND sod.product_id = $1
			GROUP BY 
				so.warehouse_id
		) AS combined
		GROUP BY 
			warehouse_id
	`, productId)
	if err != nil {
		fmt.Printf("Error obteniendo booking para el producto %d: %v\n", productId, err)
		return nil, err
	}
	product.Booking = booking

	return &product, nil
}

// Obtener una venta por IDs
func (r *ProductRepository) GetProductsByIds(ids []int) ([]domain.Products, error) {
	var products []domain.Products

	query := `
		SELECT
			p.id, p.name, p.brand_id, p.unit_of_measurement_id, p.handle, p.description, p.tags,
			p.featured_image, p.images, p.prices_cf, p.prices_sf, p.prices_box, p.featured_pcf, p.featured_psf, p.featured_pbox, p.quantity_in_box, p.created_at, p.updated_at,
			p.cost, p.tax_rate, p.quantity, p.sku, p.width, p.height, p.depth,
			p.liters, p.weight, p.barcode, p.rating, p.extra_shipping_fee, p.status,
			p.created_by, p.updated_by
		FROM products p
		WHERE p.id = ANY($1)
	`

	err := r.db.Select(&products, query, pq.Array(ids))
	if err != nil {
		return nil, err
	}

	return products, nil
}

// Actualizar un producto
func (r *ProductRepository) UpdateProduct(product *domain.Products) error {

	// Actualizar el producto en la tabla `products`
	query := `
		UPDATE products
		SET 
			name =                      :name,
			prices_cf =                 :prices_cf,
			prices_sf =                 :prices_sf,
			prices_box =                :prices_box,
			featured_pcf =              :featured_pcf,
			featured_psf =              :featured_psf,
			featured_pbox =             :featured_pbox,
			cost =                      :cost,
			sku =                       :sku,
			unit_of_measurement_id =    :unit_of_measurement_id,
			quantity_in_box =           :quantity_in_box,
			updated_by =                :updated_by,
			updated_at =                NOW()
		WHERE id = :id AND deleted_at IS NULL
	`

	_, err := r.db.NamedExec(query, product)
	return err
}

// Eliminar un producto por Id (eliminación lógica)
func (r *ProductRepository) DeleteProductById(id int) error {
	query := `
		UPDATE products
		SET deleted_at = NOW()
		WHERE id = $1 AND deleted_at IS NULL
	`
	_, err := r.db.Exec(query, id)
	if err != nil {
		fmt.Printf("Error eliminando el producto con Id %d: %v\n", id, err)
		return err
	}
	return nil
}

// Eliminar múltiples prductos por Ids (eliminación lógica)
func (r *ProductRepository) DeleteProductsByIds(ids []int) error {
	query := `UPDATE products SET deleted_at = NOW() WHERE id = ANY($1)`
	_, err := r.db.Exec(query, pq.Array(ids))
	return err
}

func (r *ProductRepository) DeleteProductImageById(imageId int) error {
	query := `
		DELETE FROM product_images
		WHERE id = $1
	`

	_, err := r.db.Exec(query, imageId)
	if err != nil {
		fmt.Printf("Error al eliminar la imagen del producto: %v\n", err)
		return err
	}

	return nil
}