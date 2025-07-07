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
	tx, err := r.db.Beginx()
	if err != nil {
		fmt.Printf("Error iniciando la transacción: %v\n", err)
		return err
	}

	// Crear un mapa para los valores del producto
	productMap := map[string]interface{}{
		"name":               product.Name,
		"brand_id":           product.Brand_Id,
		"handle":             product.Handle,
		"description":        product.Description,
		"tags":               product.Tags,
		"featured_image":     product.FeaturedImage.String,
		"images":             product.Images,
		"prices_cf":          product.Prices_cf,
		"prices_sf":          product.Prices_sf,
		"prices_box":         product.Prices_box,
		"cost":               product.Cost,
		"tax_rate":           product.TaxRate,
		"quantity":           product.Quantity,
		"sku":                product.SKU,
		"width":              product.Width,
		"height":             product.Height,
		"depth":              product.Depth,
		"liters":             product.Liters,
		"weight":             product.Weight,
		"barcode":            product.Barcode,
		"rating":             product.Rating,
		"extra_shipping_fee": product.ExtraShippingFee,
		"status":             product.Status,
	}

	// Insertar el producto en la tabla `products`
	queryProduct := `
		INSERT INTO products (
			name, brand_id, handle, description, tags, featured_image, images, prices_cf, prices_sf, prices_box, cost, tax_rate, quantity, sku, width, height, depth, liters, weight, barcode, rating, extra_shipping_fee, status
		) VALUES (
			:name, :brand_id, :handle, :description, :tags, :featured_image, :images, :prices_cf, :prices_sf, :prices_box, :cost, :tax_rate, :quantity, :sku, :width, :height, :depth, :liters, :weight, :barcode, :rating, :extra_shipping_fee, :status
		) RETURNING id
	`

	// Ejecutar la consulta con los datos nombrados
	var productId int
	stmt, err := tx.PrepareNamed(queryProduct) // Preparar la consulta nombrada
	if err != nil {
		tx.Rollback()
		fmt.Printf("Error preparando la consulta: %v\n", err)
		return err
	}
	defer stmt.Close()

	err = stmt.Get(&productId, productMap) // Ejecutar la consulta y obtener el ID
	if err != nil {
		tx.Rollback()
		fmt.Printf("Error insertando el producto: %v\n", err)
		return err
	}

	// Insertar las categorías en la tabla `product_categories`
	queryCategories := `
		INSERT INTO product_categories (product_id, category_id)
		VALUES (:product_id, :category_id)
	`

	for _, category := range product.Categories {
		_, err := tx.NamedExec(queryCategories, map[string]interface{}{
			"product_id":  productId,
			"category_id": category.Id,
		})
		if err != nil {
			tx.Rollback()
			fmt.Printf("Error insertando categoría: %v\n", err)
			return err
		}
	}

	// Finalizar la transacción
	err = tx.Commit()
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
		INSERT INTO products (name, prices_cf, cost, quantity)
		VALUES (:name, :prices_cf, :cost, :quantity)
	`

	// Insertar cada producto
	for _, product := range products {
		_, err := tx.NamedExec(query, map[string]interface{}{
			"name":     product.Name,
			"prices_cf":    product.Prices_cf,
			"cost":     product.Cost,
			"quantity": product.Quantity,
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
			p.id, p.name, p.brand_id, p.handle, p.description, p.tags,
			p.featured_image, p.images, p.prices_cf, p.prices_sf, p.prices_box,
			p.cost, p.tax_rate, p.quantity, p.sku, p.width, p.height, p.depth,
			p.liters, p.weight, p.barcode, p.rating, p.extra_shipping_fee, p.status,
			p.created_by, p.updated_by,
			COALESCE(b.id, 0) AS "brand.id", COALESCE(b.name, '') AS "brand.name"
		FROM products p
		LEFT JOIN brands b ON p.brand_id=b.id
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
			p.id, p.name, p.brand_id, p.handle, p.description, p.tags, p.featured_image, p.images, p.prices_cf, p.cost, p.tax_rate, p.quantity, p.sku, p.width, p.height, p.depth, p.liters, p.weight, p.barcode, p.rating, p.extra_shipping_fee, p.status, p.created_by, p.updated_by,
			COALESCE(b.id, 0) AS "brand.id", COALESCE(b.name, '') AS "brand.name"
		FROM products p
		LEFT JOIN brands b ON p.brand_id=b.id
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
			p.id, p.name, p.brand_id, p.handle, p.description, p.featured_image, p.prices_cf, p.cost, p.tax_rate, p.quantity, p.sku, p.width, p.height, p.depth, p.liters, p.weight, p.barcode, p.rating, p.extra_shipping_fee, p.status,
			b.id AS "brand.id", b.name AS "brand.name"
		FROM products p
		LEFT JOIN brands b ON p.brand_id=b.id
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
	// Inicia una transacción para asegurar atomicidad
	tx, err := r.db.Beginx()
	if err != nil {
		fmt.Printf("Error iniciando la transacción: %v\n", err)
		return err
	}

	// Crear un mapa para los valores del producto
	productMap := map[string]interface{}{
		"id":                 product.Id,
		"name":               product.Name,
		"brand_id":           product.Brand_Id,
		"handle":             product.Handle,
		"description":        product.Description,
		"featured_image":     product.FeaturedImage.String,
		"prices_cf":              product.Prices_cf,
		"cost":               product.Cost,
		"tax_rate":           product.TaxRate,
		"quantity":           product.Quantity,
		"sku":                product.SKU,
		"width":              product.Width,
		"height":             product.Height,
		"depth":              product.Depth,
		"liters":             product.Liters,
		"weight":             product.Weight,
		"barcode":            product.Barcode,
		"rating":             product.Rating,
		"extra_shipping_fee": product.ExtraShippingFee,
		"status":             product.Status,
	}

	// Actualizar el producto en la tabla `products`
	queryProduct := `
		UPDATE products
		SET 
			name =               :name,
			brand_id =           :brand_id,
			handle =             :handle,
			description =        :description,
			featured_image =  :featured_image.String,
			prices_cf =              :prices_cf,
			cost =               :cost,
			tax_rate =           :tax_rate,
			quantity =           :quantity,
			sku =                :sku,
			width =              :width,
			height =             :height,
			depth =              :depth,
			liters =             :liters,
			weight =             :weight,
			barcode =            :barcode,
			rating =             :rating,
			extra_shipping_fee = :extra_shipping_fee,
			status =             :status,
			updated_at =         NOW()
		WHERE id = :id AND deleted_at IS NULL
	`

	stmt, err := tx.PrepareNamed(queryProduct)
	if err != nil {
		tx.Rollback()
		fmt.Printf("Error preparando la consulta: %v\n", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(productMap)
	if err != nil {
		tx.Rollback()
		fmt.Printf("Error actualizando el producto: %v\n", err)
		return err
	}

	// Actualizar categorías: Eliminar las existentes e insertar las nuevas
	queryDeleteCategories := `
		DELETE FROM product_categories WHERE product_id = :product_id
	`
	_, err = tx.NamedExec(queryDeleteCategories, map[string]interface{}{
		"product_id": product.Id,
	})
	if err != nil {
		tx.Rollback()
		fmt.Printf("Error eliminando categorías previas: %v\n", err)
		return err
	}

	queryInsertCategories := `
		INSERT INTO product_categories (product_id, category_id)
		VALUES (:product_id, :category_id)
	`
	for _, category := range product.Categories {
		_, err := tx.NamedExec(queryInsertCategories, map[string]interface{}{
			"product_id":  product.Id,
			"category_id": category.Id,
		})
		if err != nil {
			tx.Rollback()
			fmt.Printf("Error insertando nueva categoría: %v\n", err)
			return err
		}
	}

	// Confirmar la transacción
	err = tx.Commit()
	if err != nil {
		fmt.Printf("Error al confirmar la transacción: %v\n", err)
		return err
	}

	return nil
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