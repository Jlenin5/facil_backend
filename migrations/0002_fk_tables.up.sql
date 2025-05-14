-- Añadir las claves foráneas después de que todas las tablas estén creadas
ALTER TABLE tax_items ADD CONSTRAINT fk_tax_items_tax FOREIGN KEY (tax_id) REFERENCES taxes(id) ON DELETE CASCADE;

ALTER TABLE role_permissions ADD CONSTRAINT fk_role_permissions_role FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE SET NULL;
ALTER TABLE role_permissions ADD CONSTRAINT fk_role_permissions_permission FOREIGN KEY (permission_id) REFERENCES permissions(id) ON DELETE SET NULL;

ALTER TABLE job_positions ADD CONSTRAINT fk_job_positions_work_area FOREIGN KEY (work_area_id) REFERENCES work_areas(id) ON DELETE CASCADE;

ALTER TABLE employees ADD CONSTRAINT fk_employees_warehouse FOREIGN KEY (warehouse_id) REFERENCES warehouses(id) ON DELETE CASCADE;
ALTER TABLE employees ADD CONSTRAINT fk_employees_job_position FOREIGN KEY (job_position_id) REFERENCES job_positions(id) ON DELETE SET NULL;

ALTER TABLE users ADD CONSTRAINT fk_users_company FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE SET NULL;
ALTER TABLE users ADD CONSTRAINT fk_users_role FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE SET NULL;
ALTER TABLE users ADD CONSTRAINT fk_users_employee FOREIGN KEY (employee_id) REFERENCES employees(id) ON DELETE SET NULL;

ALTER TABLE subscriptions ADD CONSTRAINT fk_subscriptions_company FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE SET NULL;
ALTER TABLE subscriptions ADD CONSTRAINT fk_subscriptions_plan FOREIGN KEY (plan_id) REFERENCES plans(id) ON DELETE SET NULL;

ALTER TABLE branch_offices ADD CONSTRAINT fk_branch_offices_company FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE SET NULL;

ALTER TABLE warehouses ADD CONSTRAINT fk_warehouses_company FOREIGN KEY (company_id) REFERENCES companies(id) ON DELETE SET NULL;
ALTER TABLE warehouses ADD CONSTRAINT fk_warehouses_branch_office FOREIGN KEY (branch_office_id) REFERENCES branch_offices(id) ON DELETE SET NULL;

ALTER TABLE exchange_rates ADD CONSTRAINT fk_exchange_rates_base_currency FOREIGN KEY (base_currency_id) REFERENCES currencies(id) ON DELETE CASCADE;
ALTER TABLE exchange_rates ADD CONSTRAINT fk_exchange_rates_target_currency FOREIGN KEY (target_currency_id) REFERENCES currencies(id) ON DELETE CASCADE;
ALTER TABLE exchange_rates ADD CONSTRAINT fk_exchange_rates_created_by FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE SET NULL;
ALTER TABLE exchange_rates ADD CONSTRAINT fk_exchange_rates_updated_by FOREIGN KEY (updated_by) REFERENCES users(id) ON DELETE SET NULL;

ALTER TABLE products ADD CONSTRAINT fk_products_brand FOREIGN KEY (brand_id) REFERENCES brands(id) ON DELETE SET NULL;

ALTER TABLE product_images ADD CONSTRAINT fk_product_images_product FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE SET NULL;

ALTER TABLE product_categories ADD CONSTRAINT fk_product_categories_product FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE SET NULL;
ALTER TABLE product_categories ADD CONSTRAINT fk_product_categories_category FOREIGN KEY (category_id) REFERENCES categories(id) ON DELETE SET NULL;

ALTER TABLE cash_registers ADD CONSTRAINT fk_cash_registers_warehouse FOREIGN KEY (warehouse_id) REFERENCES warehouses(id) ON DELETE CASCADE;
ALTER TABLE cash_registers ADD CONSTRAINT fk_cash_registers_user_open FOREIGN KEY (user_open_id) REFERENCES users(id);
ALTER TABLE cash_registers ADD CONSTRAINT fk_cash_registers_user_close FOREIGN KEY (user_close_id) REFERENCES users(id);

ALTER TABLE cash_movements ADD CONSTRAINT fk_cash_movements_cash_register FOREIGN KEY (cash_register_id) REFERENCES cash_registers(id) ON DELETE CASCADE;
ALTER TABLE cash_movements ADD CONSTRAINT fk_cash_movements_payment_method FOREIGN KEY (payment_method_id) REFERENCES payment_methods(id) ON DELETE CASCADE;
ALTER TABLE cash_movements ADD CONSTRAINT fk_cash_movements_user FOREIGN KEY (user_id) REFERENCES users(id);

ALTER TABLE quotes ADD CONSTRAINT fk_quotes_warehouse FOREIGN KEY (warehouse_id) REFERENCES warehouses(id) ON DELETE SET NULL;
ALTER TABLE quotes ADD CONSTRAINT fk_quotes_customer FOREIGN KEY (customer_id) REFERENCES customers(id) ON DELETE CASCADE;
ALTER TABLE quotes ADD CONSTRAINT fk_quotes_currency FOREIGN KEY (currency_id) REFERENCES currencies(id) ON DELETE CASCADE;
ALTER TABLE quotes ADD CONSTRAINT fk_quotes_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE SET NULL;
ALTER TABLE quotes ADD CONSTRAINT fk_quotes_approved_by FOREIGN KEY (approved_by) REFERENCES users(id);
ALTER TABLE quotes ADD CONSTRAINT fk_quotes_canceled_by FOREIGN KEY (canceled_by) REFERENCES users(id);

ALTER TABLE quote_details ADD CONSTRAINT fk_quote_details_quote FOREIGN KEY (quote_id) REFERENCES quotes(id) ON DELETE CASCADE;
ALTER TABLE quote_details ADD CONSTRAINT fk_quote_details_product FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE;

ALTER TABLE sale_orders ADD CONSTRAINT fk_sale_orders_warehouse FOREIGN KEY (warehouse_id) REFERENCES warehouses(id) ON DELETE CASCADE;
ALTER TABLE sale_orders ADD CONSTRAINT fk_sale_orders_customer FOREIGN KEY (customer_id) REFERENCES customers(id) ON DELETE CASCADE;
ALTER TABLE sale_orders ADD CONSTRAINT fk_sale_orders_currency FOREIGN KEY (currency_id) REFERENCES currencies(id) ON DELETE CASCADE;
ALTER TABLE sale_orders ADD CONSTRAINT fk_sale_orders_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE;
ALTER TABLE sale_orders ADD CONSTRAINT fk_sale_orders_quote FOREIGN KEY (quote_id) REFERENCES quotes(id) ON DELETE CASCADE;

ALTER TABLE sale_order_details ADD CONSTRAINT fk_sale_order_details_sale_order FOREIGN KEY (sale_order_id) REFERENCES sale_orders(id) ON DELETE CASCADE;
ALTER TABLE sale_order_details ADD CONSTRAINT fk_sale_order_details_product FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE;

ALTER TABLE shipping_details ADD CONSTRAINT fk_shipping_details_order FOREIGN KEY (order_id) REFERENCES sale_orders(id) ON DELETE CASCADE;

ALTER TABLE purchase_orders ADD CONSTRAINT fk_purchase_orders_warehouse FOREIGN KEY (warehouse_id) REFERENCES warehouses(id) ON DELETE CASCADE;
ALTER TABLE purchase_orders ADD CONSTRAINT fk_purchase_orders_supplier FOREIGN KEY (supplier_id) REFERENCES suppliers(id) ON DELETE CASCADE;
ALTER TABLE purchase_orders ADD CONSTRAINT fk_purchase_orders_currency FOREIGN KEY (currency_id) REFERENCES currencies(id) ON DELETE CASCADE;
ALTER TABLE purchase_orders ADD CONSTRAINT fk_purchase_orders_created_by FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE SET NULL;
ALTER TABLE purchase_orders ADD CONSTRAINT fk_purchase_orders_approved_by FOREIGN KEY (approved_by) REFERENCES users(id) ON DELETE SET NULL;

ALTER TABLE purchase_order_details ADD CONSTRAINT fk_purchase_order_details_purchase_order FOREIGN KEY (purchase_order_id) REFERENCES purchase_orders(id) ON DELETE CASCADE;
ALTER TABLE purchase_order_details ADD CONSTRAINT fk_purchase_order_details_product FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE;

ALTER TABLE purchases ADD CONSTRAINT fk_purchases_warehouse FOREIGN KEY (warehouse_id) REFERENCES warehouses(id) ON DELETE CASCADE;
ALTER TABLE purchases ADD CONSTRAINT fk_purchases_supplier FOREIGN KEY (supplier_id) REFERENCES suppliers(id) ON DELETE CASCADE;
ALTER TABLE purchases ADD CONSTRAINT fk_purchases_currency FOREIGN KEY (currency_id) REFERENCES currencies(id) ON DELETE CASCADE;
ALTER TABLE purchases ADD CONSTRAINT fk_purchases_purchase_order FOREIGN KEY (purchase_order_id) REFERENCES purchase_orders(id) ON DELETE CASCADE;
ALTER TABLE purchases ADD CONSTRAINT fk_purchases_payment_method FOREIGN KEY (payment_method_id) REFERENCES payment_methods(id) ON DELETE CASCADE;
ALTER TABLE purchases ADD CONSTRAINT fk_purchases_created_by FOREIGN KEY (created_by) REFERENCES users(id) ON DELETE CASCADE;

ALTER TABLE purchase_details ADD CONSTRAINT fk_purchase_details_purchase_order FOREIGN KEY (purchase_id) REFERENCES purchase_orders(id) ON DELETE CASCADE;
ALTER TABLE purchase_details ADD CONSTRAINT fk_purchase_details_product FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE;

ALTER TABLE purchase_order_payments ADD CONSTRAINT fk_purchase_order_payments_payment_method FOREIGN KEY (payment_method_id) REFERENCES payment_methods(id) ON DELETE CASCADE;
ALTER TABLE purchase_order_payments ADD CONSTRAINT fk_purchase_order_payments_purchase_order FOREIGN KEY (purchase_order_id) REFERENCES purchase_orders(id) ON DELETE CASCADE;
ALTER TABLE purchase_order_payments ADD CONSTRAINT fk_purchase_order_payments_currency FOREIGN KEY (currency_id) REFERENCES currencies(id) ON DELETE CASCADE;

ALTER TABLE sales ADD CONSTRAINT fk_sales_warehouse FOREIGN KEY (warehouse_id) REFERENCES warehouses(id) ON DELETE CASCADE;
ALTER TABLE sales ADD CONSTRAINT fk_sales_customer FOREIGN KEY (customer_id) REFERENCES customers(id);
ALTER TABLE sales ADD CONSTRAINT fk_sales_currency FOREIGN KEY (currency_id) REFERENCES currencies(id) ON DELETE CASCADE;
ALTER TABLE sales ADD CONSTRAINT fk_sales_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE SET NULL;
ALTER TABLE sales ADD CONSTRAINT fk_sales_payment_method FOREIGN KEY (payment_method_id) REFERENCES payment_methods(id) ON DELETE CASCADE;
ALTER TABLE sales ADD CONSTRAINT fk_sales_sale_order FOREIGN KEY (sale_order_id) REFERENCES sale_orders(id) ON DELETE CASCADE;

ALTER TABLE sale_details ADD CONSTRAINT fk_sale_details_sale FOREIGN KEY (sale_id) REFERENCES sales(id);
ALTER TABLE sale_details ADD CONSTRAINT fk_sale_details_product FOREIGN KEY (product_id) REFERENCES products(id);

ALTER TABLE opportunity_tracking ADD CONSTRAINT fk_opportunity_tracking_customer FOREIGN KEY (customer_id) REFERENCES customers(id) ON DELETE CASCADE;
ALTER TABLE opportunity_tracking ADD CONSTRAINT fk_opportunity_tracking_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE SET NULL;

ALTER TABLE purchase_requests ADD CONSTRAINT fk_purchase_requests_supplier FOREIGN KEY (supplier_id) REFERENCES suppliers(id) ON DELETE CASCADE;
ALTER TABLE purchase_requests ADD CONSTRAINT fk_purchase_requests_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE SET NULL;
ALTER TABLE purchase_requests ADD CONSTRAINT fk_purchase_requests_approved_by FOREIGN KEY (approved_by) REFERENCES users(id) ON DELETE SET NULL;

ALTER TABLE purchase_request_details ADD CONSTRAINT fk_purchase_request_details_purchase_request FOREIGN KEY (purchase_request_id) REFERENCES purchase_requests(id) ON DELETE CASCADE;
ALTER TABLE purchase_request_details ADD CONSTRAINT fk_purchase_request_details_product FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE;

ALTER TABLE stock_control ADD CONSTRAINT fk_stock_control_warehouse FOREIGN KEY (warehouse_id) REFERENCES warehouses(id) ON DELETE CASCADE;
ALTER TABLE stock_control ADD CONSTRAINT fk_stock_control_product FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE;

ALTER TABLE inventory_movements ADD CONSTRAINT fk_inventory_movements_warehouse FOREIGN KEY (warehouse_id) REFERENCES warehouses(id) ON DELETE CASCADE;
ALTER TABLE inventory_movements ADD CONSTRAINT fk_inventory_movements_product FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE;
ALTER TABLE inventory_movements ADD CONSTRAINT fk_inventory_movements_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE SET NULL;

ALTER TABLE audit_log ADD CONSTRAINT fk_audit_log_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE SET NULL;

ALTER TABLE keys ADD CONSTRAINT fk_keys_system FOREIGN KEY (system_id) REFERENCES systems(id);
ALTER TABLE keys ADD CONSTRAINT fk_keys_created_by FOREIGN KEY (created_by) REFERENCES users(id);
ALTER TABLE keys ADD CONSTRAINT fk_keys_updated_by FOREIGN KEY (updated_by) REFERENCES users(id);

ALTER TABLE notifications ADD CONSTRAINT fk_notifications_user FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE;