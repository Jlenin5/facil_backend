-- SELECT
-- 	er.id, er.base_currency_id, er.target_currency_id, er.exchange_rate, er.created_by, er.updated_by,
-- 	c.id AS "base_currency.id", c.name AS "base_currency.name", c.code AS "base_currency.code", c.symbol AS "base_currency.symbol",
-- 	cu.id AS "target_currency.id", cu.name AS "target_currency.name", cu.code AS "target_currency.code", cu.symbol AS "target_currency.symbol",
-- 	u.id AS "created_by.id", u.employee_id AS "created_by.employee_id",
-- 	e.id AS "created_by.employee.id", e.first_name AS "created_by.employee.first_name", e.second_name AS "created_by.employee.second_name", e.third_name AS "created_by.employee.third_name", e.surname AS "created_by.employee.surname", e.second_surname AS "created_by.employee.second_surname",
-- 	COALESCE(us.id, 0) AS "updated_by.id", us.employee_id AS "updated_by.employee_id"
-- FROM exchange_rates er
-- LEFT JOIN currencies c ON er.base_currency_id=c.id
-- LEFT JOIN currencies cu ON er.target_currency_id=cu.id
-- LEFT JOIN users u ON er.created_by=u.id
-- LEFT JOIN employees e ON u.employee_id=e.id
-- LEFT JOIN users us ON er.updated_by=us.id
-- LEFT JOIN employees em ON us.employee_id=em.id
-- WHERE er.deleted_at IS NULL



-- SELECT
-- 	q.id, q.reference, q.warehouse_id, q.customer_id, q.currency_id, q.user_id, q.issue_date, q.exchange_rate, q.expiration_date, q.approved_by, q.approved_at, q.canceled_by, q.canceled_at, q.discount, q.subtotal, q.total, q.quote_status, q.migrate_quotation,
-- 	c.id AS "customer.id", c.first_name AS "customer.first_name", c.second_name AS "customer.second_name", c.third_name AS "customer.third_name", c.surname AS "customer.surname", c.second_surname AS "customer.second_surname", c.company_name AS "customer.company_name",
-- 	u.id AS "user.id", u.employee_id AS "user.employee_id",
-- 	cu.id AS "currency.id", cu.name AS "currency.name", cu.code AS "currency.code", cu.symbol AS "currency.symbol",
-- 	e.id AS "user.employee.id", e.first_name AS "user.employee.first_name", e.second_name AS "user.employee.second_name", e.third_name AS "user.employee.third_name", e.surname AS "user.employee.surname", e.second_surname AS "user.employee.second_surname",
-- 	w.id AS "warehouse.id", w.branch_office_id AS "warehouse.branch_office_id", w.name AS "warehouse.name",
-- 	bo.id AS "warehouse.branch_office.id", bo.company_id AS "warehouse.branch_office.company_id", bo.name AS "warehouse.branch_office.name",
-- 	co.id AS "warehouse.branch_office.company.id", co.name AS "warehouse.branch_office.company.name", co.ruc AS "warehouse.branch_office.company.ruc"
-- FROM quotations q
-- LEFT JOIN customers c ON q.customer_id=c.id
-- LEFT JOIN currencies cu ON q.currency_id=cu.id
-- LEFT JOIN warehouses w ON q.warehouse_id=w.id
-- LEFT JOIN branch_offices bo ON w.branch_office_id=bo.id
-- LEFT JOIN companies co ON bo.company_id=co.id
-- LEFT JOIN users u ON q.user_id=u.id
-- LEFT JOIN employees e ON u.employee_id=e.id
-- WHERE q.deleted_at IS NULL

-- SELECT
-- 	qd.id, qd.product_name, qd.quotation_id, qd.product_id, qd.quantity, qd.price, qd.discount_method, qd.discount, qd.subtotal, qd.total,
-- 	p.id AS "product.id", p.name AS "product.name", p.price AS "product.price", p.cost AS "product.cost"
-- FROM quotation_details qd
-- LEFT JOIN products p ON qd.product_id=p.id
-- WHERE qd.quotation_id = 1



-- SELECT
--   p.id, p.name, p.brand_id, p.handle, p.description, p.featured_image_id, p.price, p.cost, p.tax_rate, p.quantity, p.sku, p.width, p.height, p.depth, p.liters, p.weight, p.barcode, p.rating, p.extra_shipping_fee, p.status,
--   COALESCE(b.id, 0) AS "brand.id", b.name AS "brand.name", b.status AS "brand.status",
--   COALESCE(
--     JSON_AGG(
--       JSON_BUILD_OBJECT(
--         'id', c.id,
--         'name', c.name,
--         'status', c.status
--       )
--     ) FILTER (WHERE c.id IS NOT NULL), '[]'
--   ) AS categories
-- FROM products p
-- LEFT JOIN brands b ON p.brand_id=b.id
-- LEFT JOIN product_categories pc ON pc.product_id=p.id
-- LEFT JOIN categories c ON pc.category_id=c.id
-- WHERE p.deleted_at IS NULL
-- GROUP BY p.id, b.id



-- SELECT
--   s.id, s.document_type, s.series, s.number, s.bill, s.warehouse_id, s.customer_id, s.currency_id, s.user_id, s.issue_date, s.exchange_rate, s.discount, s.subtotal, s.total, s.total_paid, s.change, s.sale_status, s.payment_method_id, s.sale_order_id, s.tax_identification, s.retention, s.perception,
--   c.id AS "customer.id", c.first_name AS "customer.first_name", c.second_name AS "customer.second_name", c.third_name AS "customer.third_name", c.surname AS "customer.surname", c.second_surname AS "customer.second_surname", c.company_name AS "customer.company_name",
--   u.id AS "user.id", u.employee_id AS "user.employee_id",
--   cu.id AS "currency.id", cu.name AS "currency.name", cu.code AS "currency.code", cu.symbol AS "currency.symbol",
--   e.id AS "user.employee.id", e.first_name AS "user.employee.first_name", e.second_name AS "user.employee.second_name", e.third_name AS "user.employee.third_name", e.surname AS "user.employee.surname", e.second_surname AS "user.employee.second_surname",
--   w.id AS "warehouse.id", w.branch_office_id AS "warehouse.branch_office_id", w.name AS "warehouse.name",
--   bo.id AS "warehouse.branch_office.id", bo.company_id AS "warehouse.branch_office.company_id", bo.name AS "warehouse.branch_office.name",
--   co.id AS "warehouse.branch_office.company.id", co.name AS "warehouse.branch_office.company.name", co.ruc AS "warehouse.branch_office.company.ruc",
--   COALESCE(so.id, 0) AS "sale_order.id", COALESCE(so.reference, '') AS "sale_order.reference",
--   pm.id AS "payment_method.id", pm.name AS "payment_method.name"
-- FROM sales s
-- LEFT JOIN customers c ON s.customer_id=c.id
-- LEFT JOIN currencies cu ON s.currency_id=cu.id
-- LEFT JOIN warehouses w ON s.warehouse_id=w.id
-- LEFT JOIN branch_offices bo ON w.branch_office_id=bo.id
-- LEFT JOIN companies co ON bo.company_id=co.id
-- LEFT JOIN users u ON s.user_id=u.id
-- LEFT JOIN employees e ON u.employee_id=e.id
-- LEFT JOIN sale_orders so ON s.sale_order_id=so.id
-- LEFT JOIN payment_methods pm ON s.payment_method_id=pm.id
-- WHERE s.deleted_at IS NULL

-- SELECT
--   sd.id, sd.product_name, sd.sale_id, sd.product_id, sd.quantity, sd.price, sd.discount_method, sd.discount, sd.subtotal, sd.total,
--   p.id AS "product.id", p.name AS "product.name", p.price AS "product.price", p.cost AS "product.cost"
-- FROM sale_details sd
-- LEFT JOIN products p ON sd.product_id=p.id
-- WHERE sd.sale_id = $1



-- SELECT
--   so.id, so.reference, so.warehouse_id, so.customer_id, so.currency_id, so.user_id, so.issue_date, so.exchange_rate, so.discount, so.subtotal, so.total, so.order_status, so.date_approved, so.migrate_sale_order, so.quote_id,
--   c.id AS "customer.id", c.first_name AS "customer.first_name", c.second_name AS "customer.second_name", c.third_name AS "customer.third_name", c.surname AS "customer.surname", c.second_surname AS "customer.second_surname", c.company_name AS "customer.company_name",
--   u.id AS "user.id", u.employee_id AS "user.employee_id",
--   cu.id AS "currency.id", cu.name AS "currency.name", cu.code AS "currency.code", cu.symbol AS "currency.symbol",
--   e.id AS "user.employee.id", e.first_name AS "user.employee.first_name", e.second_name AS "user.employee.second_name", e.third_name AS "user.employee.third_name", e.surname AS "user.employee.surname", e.second_surname AS "user.employee.second_surname",
--   w.id AS "warehouse.id", w.branch_office_id AS "warehouse.branch_office_id", w.name AS "warehouse.name",
--   bo.id AS "warehouse.branch_office.id", bo.company_id AS "warehouse.branch_office.company_id", bo.name AS "warehouse.branch_office.name",
--   co.id AS "warehouse.branch_office.company.id", co.name AS "warehouse.branch_office.company.name", co.ruc AS "warehouse.branch_office.company.ruc",
--   COALESCE(q.id, 0) AS "quote.id", COALESCE(q.reference, '') AS "quote.reference"
-- FROM sale_orders so
-- LEFT JOIN customers c ON so.customer_id=c.id
-- LEFT JOIN currencies cu ON so.currency_id=cu.id
-- LEFT JOIN warehouses w ON so.warehouse_id=w.id
-- LEFT JOIN branch_offices bo ON w.branch_office_id=bo.id
-- LEFT JOIN companies co ON bo.company_id=co.id
-- LEFT JOIN users u ON so.user_id=u.id
-- LEFT JOIN employees e ON u.employee_id=e.id
-- LEFT JOIN quotes q ON so.quote_id=q.id
-- WHERE so.id = $1 AND so.deleted_at IS NULL

-- SELECT
--   sod.id, sod.product_name, sod.sale_order_id, sod.product_id, sod.quantity, sod.price, sod.discount_method, sod.discount, sod.subtotal, sod.total,
--   p.id AS "product.id", p.name AS "product.name", p.price AS "product.price", p.cost AS "product.cost"
-- FROM sale_order_details sod
-- LEFT JOIN products p ON sod.product_id=p.id
-- WHERE sod.sale_order_id = $1



-- SELECT
--   s.id, s.company_id, s.plan_id, s.start_date, s.end_date, s.status,
--   c.id AS "company.id", c.name AS "company.name",
--   p.id AS "plan.id", p.title AS "plan.title"
-- FROM subscriptions s
-- LEFT JOIN companies c ON s.company_id=c.id
-- LEFT JOIN plans p ON s.plan_id = p.id
-- WHERE s.deleted_at IS NULL



-- SELECT
--   cr.id, cr.warehouse_id, cr.user_open_id, cr.user_close_id, cr.opening_date, cr.closing_date, cr.initial_amount, cr.closing_amount, cr.difference, cr.status,
--   w.id AS "warehouse.id", w.name AS "warehouse.name",
--   uo.id AS "user_open.id", uo.employee_id AS "user_open.employee_id",
--   eo.id AS "user_open.employee.id", eo.first_name AS "user_open.employee.first_name", eo.second_name AS "user_open.employee.second_name", eo.third_name AS "user_open.employee.third_name", eo.surname AS "user_open.employee.surname", eo.second_surname AS "user_open.employee.second_surname",
--   COALESCE(uc.id, 0) AS "user_close.id", COALESCE(uc.employee_id, 0) AS "user_close.employee_id",
--   COALESCE(ec.id, 0) AS "user_close.employee.id", COALESCE(ec.first_name, '') AS "user_close.employee.first_name", COALESCE(ec.second_name, '') AS "user_close.employee.second_name", COALESCE(ec.third_name, '') AS "user_close.employee.third_name", COALESCE(ec.surname, '') AS "user_close.employee.surname", COALESCE(ec.second_surname, '') AS "user_close.employee.second_surname"
-- FROM cash_registers cr
-- LEFT JOIN warehouses w ON cr.warehouse_id=w.id
-- LEFT JOIN users uo ON cr.user_open_id=uo.id
-- LEFT JOIN employees eo ON uo.employee_id=eo.id
-- LEFT JOIN users uc ON cr.user_close_id=uc.id
-- LEFT JOIN employees ec ON uc.employee_id=ec.id
-- WHERE cr.id = $1 AND cr.deleted_at IS NULL



-- SELECT
--   cm.id, cm.cash_register_id, cm.movement_type, cm.payment_method_id, cm.amount, cm.description, cm.user_id,
--   cr.id AS "cash_register.id",
--   pm.id AS "payment_method.id", pm.name AS "payment_method.name",
--   u.id AS "user.id", u.employee_id AS "user.employee_id",
--   eo.id AS "user.employee.id", eo.first_name AS "user.employee.first_name", eo.second_name AS "user.employee.second_name", eo.third_name AS "user.employee.third_name", eo.surname AS "user.employee.surname", eo.second_surname AS "user.employee.second_surname"
-- FROM cash_movements cm
-- LEFT JOIN cash_registers cr ON cm.cash_register_id=cr.id
-- LEFT JOIN payment_methods pm ON cm.payment_method_id=pm.id
-- LEFT JOIN users u ON cm.user_id=u.id
-- LEFT JOIN employees eo ON u.employee_id=eo.id
-- WHERE cm.id = $1



-- SELECT
--   jp.id, jp.work_area_id, jp.name, jp.description, jp.status,
--   wa.id AS "work_area.id", wa.name AS "work_area.name"
-- FROM job_positions jp
-- LEFT JOIN work_areas wa ON jp.work_area_id=wa.id
-- WHERE jp.id = $1



-- SELECT 
--   e.id, e.first_name, e.second_name, e.third_name, e.surname, e.second_surname, e.photo, 
--   e.warehouse_id, e.document_type, e.document_number, e.birth_date, e.gender, 
--   e.email, e.phone, e.address, e.hire_date, e.job_position_id, e.salary, e.status,
--   w.id AS "warehouse.id", w.name AS "warehouse.name", w.address AS "warehouse.address", w.status AS "warehouse.status", w.branch_office_id AS "warehouse.branch_office_id",
--   b.id AS "warehouse.branch_office.id", b.name AS "warehouse.branch_office.name", b.address AS "warehouse.branch_office.address", b.phone AS "warehouse.branch_office.phone", b.company_id AS "warehouse.branch_office.company_id",
--   c.id AS "warehouse.branch_office.company.id", c.name AS "warehouse.branch_office.company.name", c.ruc AS "warehouse.branch_office.company.ruc", c.phone AS "warehouse.branch_office.company.phone", c.address AS "warehouse.branch_office.company.address",
--   jp.id AS "job_position.id", jp.work_area_id AS "job_position.work_area_id", jp.name AS "job_position.name",
--   wa.id AS "job_position.work_area.id", wa.name AS "job_position.work_area.name"
-- FROM employees e
-- LEFT JOIN warehouses w ON e.warehouse_id = w.id AND w.deleted_at IS NULL
-- LEFT JOIN branch_offices b ON w.branch_office_id = b.id AND b.deleted_at IS NULL
-- LEFT JOIN companies c ON b.company_id = c.id AND c.deleted_at IS NULL
-- LEFT JOIN job_positions jp ON e.job_position_id = jp.id AND jp.deleted_at IS NULL
-- LEFT JOIN work_areas wa ON jp.work_area_id = wa.id AND wa.deleted_at IS NULL
-- WHERE e.deleted_at IS NULL



-- SELECT 
--   k.id, k.system_id, k.name, k.description, k.config_key, k.config_value, k.created_by, k.updated_by, k.status,
--   COALESCE(s.id, 0) AS "system.id", COALESCE(s.name, '') AS "system.name"
-- FROM keys k
-- LEFT JOIN systems s ON k.system_id = s.id
-- ORDER BY s.id DESC



SELECT
  c.id, c.name, c.logo, c.ruc, c.email, c.phone, c.web_site, c.address, c.status,
  COALESCE(p.id, 0) AS "subscription.plan.id", COALESCE(p.title, '') AS "subscription.plan.title"
FROM companies c
INNER JOIN subscriptions s ON c.id=s.company_id
INNER JOIN plans p ON s.plan_id=p.id
WHERE c.deleted_at IS NULL