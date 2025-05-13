-- Plans Table
CREATE TABLE plans (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    title VARCHAR(100) NOT NULL,
    subtitle VARCHAR(200) DEFAULT NULL,
    price DECIMAL(10,2) NOT NULL,
    max_branch_offices SMALLINT NOT NULL CHECK (max_branch_offices >= 0),
    max_warehouses SMALLINT NOT NULL CHECK (max_warehouses >= 0),
    max_purchases SMALLINT NOT NULL CHECK (max_purchases >= 0),
    max_users SMALLINT NOT NULL CHECK (max_users >= 0),
    max_products SMALLINT NOT NULL CHECK (max_products >= 0),
    max_services SMALLINT NOT NULL CHECK (max_services >= 0),
    max_documents INT NOT NULL CHECK (max_documents >= 0),
    status BIT(1) NOT NULL DEFAULT B'1', -- Usar BIT(1) con valor predeterminado de 1
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

-- Subscriptions Table
CREATE TABLE subscriptions (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    company_id INT,
    plan_id INT,
    start_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    end_date TIMESTAMP NOT NULL,
    status VARCHAR(10) NOT NULL CHECK (status IN ('active', 'expired', 'canceled')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

-- Currencies Table
CREATE TABLE currencies (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT DEFAULT NULL,
    code VARCHAR(10) NOT NULL,
    symbol VARCHAR(10) NOT NULL,
    status BIT(1) NOT NULL DEFAULT B'1', -- Usar BIT(1) con valor predeterminado de 1
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

-- Taxes Table
CREATE TABLE taxes (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(50) NOT NULL, -- IGV, ISC, Percepción, Retención, etc.
    description TEXT DEFAULT NULL, -- Explicación del impuesto
    rate DECIMAL(10,2) NOT NULL, -- % de impuesto (Ejemplo: 18.00 para IGV)
    tax_type  BIT(1) NOT NULL DEFAULT B'1', -- percentage(0) o fixed(1)
    status BIT(1) NOT NULL DEFAULT B'1', -- Usar BIT(1) con valor predeterminado de 1
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

CREATE TABLE tax_items (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    tax_id INT NOT NULL,
    reference_table VARCHAR(20) NOT NULL CHECK (reference_table IN ('sale_orders', 'sale_order_details', 'purchase_orders', 'purchase_order_details', 'sales', 'sale_details')),
    reference_id INT NOT NULL, -- ID de la orden o venta
    amount DECIMAL(10,2) NOT NULL, -- Valor del impuesto calculado
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

-- Roles Table
CREATE TABLE roles (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    description TEXT
);

-- Permissions Table
CREATE TABLE permissions (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    description TEXT
);

-- Role Permissions Table
CREATE TABLE role_permissions (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    role_id INT NULL,
    permission_id INT NULL
);

-- Companies Table
CREATE TABLE companies (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(200) NOT NULL,
    logo VARCHAR(150) DEFAULT NULL,
    ruc CHAR(11) NOT NULL,
    email VARCHAR(255),
    phone VARCHAR(15),
    web_site VARCHAR(100) DEFAULT NULL,
    address TEXT NOT NULL,
    status BIT(1) NOT NULL DEFAULT B'1',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

-- Branch Office Table
CREATE TABLE branch_offices (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    company_id INT,
    name VARCHAR(200) NOT NULL,
    description TEXT DEFAULT NULL,
    address TEXT NOT NULL,
    phone VARCHAR(15),
    status BIT(1) NOT NULL DEFAULT B'1',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

-- Warehouses Table
CREATE TABLE warehouses (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    company_id INT,
    branch_office_id INT,
    name VARCHAR(200) NOT NULL,
    description TEXT DEFAULT NULL,
    address TEXT,
    phone VARCHAR(15) DEFAULT NULL,
    status BIT(1) NOT NULL DEFAULT B'1',
    CHECK (
        (branch_office_id IS NOT NULL AND company_id IS NULL) OR
        (branch_office_id IS NULL AND company_id IS NOT NULL)
    ),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

-- Work Areas Table
CREATE TABLE work_areas (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT DEFAULT NULL,
    status BIT(1) NOT NULL DEFAULT B'1',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

-- Job Positions Table
CREATE TABLE job_positions (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    work_area_id INT NOT NULL,
    name VARCHAR(100) NOT NULL,
    description TEXT DEFAULT NULL,
    status BIT(1) NOT NULL DEFAULT B'1',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

-- Employees Table
CREATE TABLE employees (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    first_name VARCHAR(50) NOT NULL,
    second_name VARCHAR(50),
    third_name VARCHAR(50),
    surname VARCHAR(50),
    second_surname VARCHAR(50),
    photo VARCHAR(150),
    warehouse_id INTEGER NOT NULL,
    document_type VARCHAR(3) NOT NULL CHECK (document_type IN ('dni', 'ce')),
    document_number VARCHAR(9) NOT NULL,
    birth_date DATE NOT NULL,
    gender CHAR(1) CHECK (gender IN ('M', 'F')) NOT NULL,
    email VARCHAR(100) DEFAULT NULL,
    phone CHAR(9) DEFAULT NULL,
    address VARCHAR(200),
    hire_date DATE NOT NULL,
    job_position_id INT,
    salary NUMERIC(10,2) CHECK (salary >= 0) NOT NULL,
    status BIT(1) NOT NULL DEFAULT B'1', -- Usar BIT(1) con valor predeterminado de 1
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

-- Users Table
CREATE TABLE users (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    company_id INT,
    password VARCHAR(255) NOT NULL,
    role_id INTEGER NULL,
    username VARCHAR(100) NOT NULL,
    employee_id INTEGER DEFAULT NULL,
    avatar VARCHAR(150),
    email VARCHAR(100) NOT NULL,
    settings JSONB DEFAULT NULL, -- DEFAULT '{}'::jsonb
    shortcuts JSONB DEFAULT NULL, -- DEFAULT '[]'::jsonb
    status BIT(1) NOT NULL DEFAULT B'1', -- Usar BIT(1) con valor predeterminado de 1
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

-- Exchange Rates Table
CREATE TABLE exchange_rates (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    base_currency_id INT NOT NULL,
    target_currency_id INT NOT NULL,
    exchange_rate DECIMAL(18,6) NOT NULL, -- Tipo de cambio con 6 decimales de precisión
    created_by INT NOT NULL,
    updated_by INT DEFAULT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

-- Categories Table
CREATE TABLE categories (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT DEFAULT NULL,
    status BIT(1) NOT NULL DEFAULT B'1', -- Usar BIT(1) con valor predeterminado de 1
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

-- Brands Table
CREATE TABLE brands (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT DEFAULT NULL,
    logo_url VARCHAR(255) DEFAULT NULL,
    website_url VARCHAR(255) DEFAULT NULL,
    status BIT(1) NOT NULL DEFAULT B'1', -- Usar BIT(1) con valor predeterminado de 1
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

-- Products Table
CREATE TABLE products (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(150) DEFAULT NULL,
    brand_id INTEGER NULL,
    handle VARCHAR(255) DEFAULT NULL,
    description TEXT DEFAULT NULL,
    featured_image_id VARCHAR(10) DEFAULT NULL,
    price DECIMAL(12,2) NOT NULL CHECK (price >= 0),
    cost DECIMAL(12,2) NOT NULL CHECK (cost >= 0),
    tax_rate DECIMAL(5, 2) DEFAULT NULL,
    quantity  DECIMAL(12,2) NOT NULL DEFAULT 0,
    sku VARCHAR(50) DEFAULT NULL,
    width DECIMAL(12,2) DEFAULT 0,
    height DECIMAL(12,2) DEFAULT 0,
    depth DECIMAL(12,2) DEFAULT 0,
    liters DECIMAL(12,2) DEFAULT 0,
    weight DECIMAL(12,2) DEFAULT 0,
    barcode VARCHAR(100) DEFAULT NULL,
    rating DECIMAL(4, 2) DEFAULT 0,
    extra_shipping_fee  DECIMAL(12,2) DEFAULT 0,
    status BIT(1) NOT NULL DEFAULT B'1', -- Usar BIT(1) con valor predeterminado de 1
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

-- Services Table
CREATE TABLE services (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(200) NOT NULL,
    description TEXT,
    price DECIMAL(10, 2) NOT NULL,
    cost DECIMAL(10, 2),
    duration INTERVAL,
    requirements TEXT DEFAULT NULL,
    rating DECIMAL(4, 2) DEFAULT 0,
    status BIT(1) NOT NULL DEFAULT B'1', -- Usar BIT(1) con valor predeterminado de 1
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

-- Customers Table
CREATE TABLE customers (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    first_name VARCHAR(50),
    second_name VARCHAR(50),
    third_name VARCHAR(50),
    surname VARCHAR(50),
    second_surname VARCHAR(50),
    company_name VARCHAR(100),
    document_type VARCHAR(3) NOT NULL CHECK (document_type IN ('dni', 'ruc', 'ce')),
    document_number VARCHAR(12) NOT NULL,
    email VARCHAR(50) NOT NULL,
    address VARCHAR(200),
    phone CHAR(9),
    status BIT(1) NOT NULL DEFAULT B'1', -- Usar BIT(1) con valor predeterminado de 1
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

-- Suppliers Table
CREATE TABLE suppliers (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    ruc CHAR(11) NOT NULL,
    email VARCHAR(100),
    phone VARCHAR(20),
    web_site VARCHAR(100) DEFAULT NULL,
    address TEXT,
    status BIT(1) NOT NULL DEFAULT B'1', -- Usar BIT(1) con valor predeterminado de 1
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

-- Product Images Table
CREATE TABLE product_images (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    product_id INT NULL,
    url VARCHAR(150) NOT NULL,
    featured VARCHAR(10)
);

-- Product Categories Table
CREATE TABLE product_categories (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    product_id INT,
    category_id INT
);

-- Payment Methods Table
CREATE TABLE payment_methods (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(50) NOT NULL,
    description TEXT DEFAULT NULL,
    status BIT(1) NOT NULL DEFAULT B'1', -- Usar BIT(1) con valor predeterminado de 1
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

-- Cash Registers Table
CREATE TABLE cash_registers (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    warehouse_id INT NOT NULL,
    user_open_id INT NOT NULL,
    user_close_id INT,
    opening_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    closing_date TIMESTAMP DEFAULT NULL,
    initial_amount DECIMAL(12,2) NOT NULL,
    closing_amount DECIMAL(12,2) DEFAULT NULL,
    difference DECIMAL(12,2) DEFAULT NULL, -- Diferencia entre la caja y lo calculado
    status VARCHAR(10) NOT NULL CHECK (status IN ('open', 'closed')) DEFAULT 'open',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

-- Cash Movements Table
CREATE TABLE cash_movements (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    cash_register_id INT NOT NULL,
    movement_type VARCHAR(10) NOT NULL CHECK (movement_type IN ('income', 'expense')),
    payment_method_id INT,
    amount DECIMAL(12,2) NOT NULL,
    description VARCHAR(255),
    user_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Quotations Table
CREATE TABLE quotes (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,                
    reference CHAR(8) NOT NULL,
    warehouse_id INT,
    customer_id INT NOT NULL,
    currency_id INT NOT NULL,
    user_id INT NOT NULL,
    issue_date TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    exchange_rate DECIMAL(10, 4) DEFAULT 1,
    expiration_date TIMESTAMP DEFAULT NULL,
    approved_by INT,
    approved_at TIMESTAMP,
    canceled_by INT,
    canceled_at TIMESTAMP,
    discount DECIMAL(12,2) DEFAULT 0 CHECK (discount >= 0),
    subtotal DECIMAL(12,2) NOT NULL,
    total DECIMAL(12,2) NOT NULL,
    quote_status VARCHAR(20) NOT NULL CHECK (quote_status IN ('issued', 'pending', 'approved', 'rejected', 'canceled')) DEFAULT 'issued', -- Estado
    migrate_quote BIT(1) NOT NULL DEFAULT B'0',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

-- Quotation Details Table
CREATE TABLE quote_details (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    product_name VARCHAR(150) DEFAULT NULL,
    quote_id INT NOT NULL,
    product_id INT NOT NULL,
    quantity  DECIMAL(12,2) NOT NULL CHECK (quantity > 0),
    price DECIMAL(12,2) NOT NULL CHECK (price >= 0),
    discount_method BIT(1) NOT NULL DEFAULT B'1',
    discount DECIMAL(12,2) DEFAULT 0 CHECK (discount >= 0),
    subtotal DECIMAL(12,2) NOT NULL,
    total DECIMAL(12,2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

-- Sale Orders Table
CREATE TABLE sale_orders (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    reference CHAR(8) NOT NULL,
    warehouse_id INTEGER NOT NULL,
    customer_id INTEGER NOT NULL,
    currency_id INT NOT NULL,
    user_id INTEGER NOT NULL,
    issue_date TIMESTAMP NOT NULL,
    exchange_rate  DECIMAL(12,2) DEFAULT 1.0,
    discount  DECIMAL(12,2),
    subtotal  DECIMAL(12,2) NOT NULL,
    total  DECIMAL(12,2) NOT NULL,
    order_status VARCHAR(10) NOT NULL CHECK (order_status IN ('issued', 'approved', 'canceled')),
    date_approved DATE DEFAULT NULL,
    migrate_sale_order BIT(1) NOT NULL DEFAULT B'0',
    quote_id INT DEFAULT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

-- Sale Order Details Table
CREATE TABLE sale_order_details (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    product_name VARCHAR(150) DEFAULT NULL,
    sale_order_id INT,
    product_id INT,
    quantity  DECIMAL(12,2) NOT NULL,
    discount_method BIT(1) NOT NULL DEFAULT B'1',
    discount  DECIMAL(12,2) DEFAULT NULL,
    price DECIMAL(12,2) NOT NULL CHECK (price >= 0),
    subtotal  DECIMAL(12,2) NOT NULL,
    total  DECIMAL(12,2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

-- Shipping Details Table
CREATE TABLE shipping_details (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    order_id INTEGER NOT NULL,
    tracking VARCHAR(50),
    carrier VARCHAR(50),
    weight  DECIMAL(12,2),
    fee  DECIMAL(12,2),
    date TIMESTAMP NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Purchase Orders Table
CREATE TABLE purchase_orders (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    reference CHAR(8) NOT NULL,
    description TEXT DEFAULT NULL,
    warehouse_id INTEGER NOT NULL,
    supplier_id INTEGER NOT NULL,
    supplier_document VARCHAR(20) NULL,
    exchange_rate  DECIMAL(12,2),
    discount  DECIMAL(12,2),
    user_id INTEGER NULL,
    issue_date TIMESTAMP NOT NULL,
    tax  DECIMAL(12,2) NOT NULL,
    subtotal  DECIMAL(12,2) NOT NULL,
    total  DECIMAL(12,2) NOT NULL,
    order_status VARCHAR(15) NOT NULL CHECK (order_status IN ('issued', 'paid', 'unpaid', 'pending', 'canceled', 'approved', 'partial_paid', 'shipped')) NOT NULL DEFAULT 'issued', -- Estado de la orden
    date_approved DATE DEFAULT NULL,
    migrate_purchase BIT(1) NOT NULL DEFAULT B'0', -- Usar BIT(1) con valor predeterminado de 0
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

-- Purchase Order Details Table
CREATE TABLE purchase_order_details (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    purchase_order_id INT,
    product_id INT,
    quantity  DECIMAL(12,2) NOT NULL,
    price DECIMAL(12,2) NOT NULL CHECK (price >= 0),
    discount  DECIMAL(12,2) DEFAULT 0,
    tax_rate DECIMAL(5, 2),
    total  DECIMAL(12,2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

-- Purchase Order Payments Table
CREATE TABLE purchase_order_payments (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    payment_method_id INT NOT NULL,
    purchase_order_id INT,
    payment_date TIMESTAMP NOT NULL,
    amount  DECIMAL(12,2) NOT NULL,
    currency_id INT,
    exchange_rate  DECIMAL(12,2) DEFAULT 1.0, -- Tipo de cambio (útil para pagos en moneda extranjera)
    reference_number VARCHAR(50), -- Número de referencia (por ejemplo, número de operación bancaria)
    status VARCHAR(20) NOT NULL CHECK (status IN ('pending', 'completed', 'failed', 'refunded')), -- Estado del pago
    notes TEXT, -- Notas adicionales
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

-- Sales Table
CREATE TABLE sales (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    document_type VARCHAR(10) NOT NULL CHECK (document_type IN ('ticket', 'invoice')),
    series VARCHAR(4) DEFAULT NULL,
    number INT DEFAULT NULL,
    bill VARCHAR(12) DEFAULT NULL,
    issue_date TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    warehouse_id INTEGER NOT NULL,
    customer_id INT NOT NULL,
    currency_id INT,
    user_id INT NOT NULL,
    exchange_rate  DECIMAL(12,2) DEFAULT 1.0,
    discount  DECIMAL(12,2) DEFAULT NULL,
    subtotal DECIMAL(10,2) NOT NULL,
    total DECIMAL(10,2) NOT NULL,
    total_paid DECIMAL(10,2) NOT NULL,
    change DECIMAL(10,2) NOT NULL,
    sale_status VARCHAR(10) NOT NULL CHECK (sale_status IN ('issued', 'paid', 'unpaid', 'pending', 'canceled')) NOT NULL DEFAULT 'issued',
    payment_method_id INT NOT NULL,
    sale_order_id INT DEFAULT NULL,
    tax_identification VARCHAR(20) DEFAULT NULL,
    retention DECIMAL(10,2) DEFAULT NULL,
    perception DECIMAL(10,2) DEFAULT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

-- Sale Details Table
CREATE TABLE sale_details (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    product_name VARCHAR(150) DEFAULT NULL,
    sale_id INT NOT NULL,
    product_id INT NOT NULL,
    quantity DECIMAL(10,2) NOT NULL,
    discount_method BIT(1) NOT NULL DEFAULT B'1',
    discount  DECIMAL(12,2) DEFAULT NULL,
    price DECIMAL(10,2) NOT NULL,
    subtotal DECIMAL(10,2) NOT NULL,
    total DECIMAL(10,2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

-- Opportunity Tracking Table
CREATE TABLE opportunity_tracking (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    customer_id INT NOT NULL,
    user_id INT NOT NULL, -- Comercial asignado
    title VARCHAR(200) NOT NULL,
    description TEXT,
    status VARCHAR(50) CHECK (status IN ('open', 'in_progress', 'won', 'lost')) NOT NULL DEFAULT 'open',
    expected_revenue DECIMAL(10,2), -- Ingreso estimado
    probability INT CHECK (probability BETWEEN 0 AND 100), -- Probabilidad de cierre
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

-- Purchase Requests Table
CREATE TABLE purchase_requests (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    reference CHAR(8) NOT NULL,
    supplier_id INT NOT NULL,
    user_id INT NOT NULL, -- Usuario que solicita
    request_date DATE DEFAULT CURRENT_DATE,
    expected_delivery_date DATE,
    status VARCHAR(50) CHECK (status IN ('pending', 'approved', 'rejected')) NOT NULL DEFAULT 'pending',
    priority VARCHAR(20) CHECK (priority IN ('low', 'medium', 'high')) DEFAULT 'medium',
    total_cost DECIMAL(12,2) DEFAULT 0.00,
    approved_by INT,
    approval_date TIMESTAMP,
    notes TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

-- Purchase Requests Details Table
CREATE TABLE purchase_request_details (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    purchase_request_id INT NOT NULL,
    product_id INT NOT NULL,
    quantity DECIMAL(10,2) NOT NULL CHECK (quantity > 0),
    unit_price DECIMAL(10,2),
    estimated_cost DECIMAL(10,2),
    subtotal DECIMAL(12,2),
    received_quantity INT DEFAULT 0,
    status VARCHAR(50) CHECK (status IN ('pending', 'received', 'canceled')) DEFAULT 'pending',
    comments TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

-- Stock Control Table
CREATE TABLE stock_control (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    warehouse_id INT NOT NULL,
    product_id INT NOT NULL,
    current_stock INT NOT NULL CHECK (current_stock >= 0),
    current_booking INT NOT NULL CHECK (current_booking >= 0),
    min_stock INT CHECK (min_stock >= 0),
    max_stock INT CHECK (max_stock > min_stock),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

-- Inventory Movements Table
CREATE TABLE inventory_movements (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    warehouse_id INT NOT NULL,
    product_id INT NOT NULL,
    movement_type VARCHAR(50) CHECK (movement_type IN ('in', 'out', 'adjustment')) NOT NULL,
    quantity INT NOT NULL CHECK (quantity > 0),
    reference CHAR(8), -- Puede ser una compra, venta u otro documento
    user_id INT NOT NULL, -- Quién hizo el movimiento
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    deleted_at TIMESTAMP DEFAULT NULL
);

-- Audit Log Table
CREATE TABLE audit_log (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    table_name VARCHAR(100) NOT NULL,
    record_id INT NOT NULL,
    action VARCHAR(10) NOT NULL CHECK (action IN ('INSERT', 'UPDATE', 'DELETE')),
    old_data JSONB,
    new_data JSONB,
    user_id INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Systems Table
CREATE TABLE systems (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    description TEXT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Keys Table
CREATE TABLE keys (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    system_id INT, -- Odoo, Nubefact, SUNAT, etc.
    name VARCHAR(100) NOT NULL,
    description TEXT,
    config_key VARCHAR(100) NOT NULL,
    config_value TEXT NOT NULL,
    created_by INT,
    updated_by INT,
    status BIT(1) NOT NULL DEFAULT B'1',
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- Notifications Table
CREATE TABLE notifications (
    id INT GENERATED ALWAYS AS IDENTITY PRIMARY KEY,
    user_id INT,
    message TEXT NOT NULL,
    is_read BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);