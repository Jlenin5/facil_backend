-- Insertar datos en la tabla plans
INSERT INTO plans (
    title, subtitle, price, max_branch_offices, max_warehouses, max_purchases, max_users, max_products, max_services, max_documents
) VALUES
('Básico', 'Ideal para pequeñas empresas con una sola sucursal', 99.00, 1, 1, 100, 2, 50, 50, 1000),
('Empresarial', 'Para empresas en crecimiento con múltiples sucursales', 299.00, 5, 10, 1000, 50, 500, 500, 10000),
('Corporativo', 'Para grandes empresas con operaciones extensas', 699.00, 20, 50, 5000, 200, 2000, 2000, 50000);

-- Insertar datos en la tabla currencies
INSERT INTO currencies (name, description, code, symbol) VALUES
('US Dollar', 'United States Dollar', 'USD', '$'),
('Euro', 'European Union Currency', 'EUR', '€'),
('British Pound', 'United Kingdom Currency', 'GBP', '£'),
('Japanese Yen', 'Japanese Currency', 'JPY', '¥'),
('Mexican Peso', 'Mexican Currency', 'MXN', '$'),
('Peruvian Sol', 'Currency of Peru', 'PEN', 'S/');

-- Insertar datos en la tabla taxes
INSERT INTO taxes (name, description, rate, tax_type) VALUES
('IGV', 'Impuesto General a las Ventas (18%)', 18.00, B'0'),
('ISC', 'Impuesto Selectivo al Consumo', 10.00, B'0'),
('Percepción', 'Percepción de IGV', 2.00, B'0'),
('Retención', 'Retención del IGV', 3.00, B'0'),
('Impuesto Fijo', 'Ejemplo de impuesto con monto fijo', 5.00, B'1');

-- Insertar datos en la tabla roles
INSERT INTO roles (name, description) VALUES
('Admin', 'Administrador del sistema con todos los permisos'),
('Gerente', 'Encargado de la supervisión general de la empresa'),
('Jefe de Ventas', 'Responsable del área de ventas'),
('Jefe de Almacén', 'Encargado de la gestión de almacenes'),
('Contador', 'Encargado de la contabilidad y finanzas'),
('Recursos Humanos', 'Encargado de la gestión del personal'),
('Soporte Técnico', 'Responsable de la asistencia técnica'),
('Analista de Datos', 'Analiza y genera reportes para la empresa');

-- Insertar datos en la tabla permissions
INSERT INTO permissions (name, description) VALUES
('manage_users', 'Crear, editar y eliminar usuarios'),
('view_reports', 'Ver informes y reportes'),
('edit_financials', 'Editar información financiera'),
('manage_inventory', 'Administrar el inventario'),
('process_sales', 'Procesar ventas'),
('approve_purchases', 'Aprobar compras'),
('manage_roles', 'Administrar roles y permisos'),
('access_sensitive_data', 'Acceder a datos sensibles');

-- Insertar asignaciones de permisos al rol Admin (tendrá todos los permisos)
INSERT INTO role_permissions (role_id, permission_id)
SELECT r.id, p.id
FROM roles r, permissions p
WHERE r.name = 'Admin';

-- Asignar permisos específicos a otros roles
INSERT INTO role_permissions (role_id, permission_id)
SELECT r.id, p.id FROM roles r
JOIN permissions p ON 
(r.name = 'Gerente' AND p.name IN ('view_reports', 'approve_purchases', 'process_sales')) OR
(r.name = 'Jefe de Ventas' AND p.name IN ('process_sales', 'view_reports')) OR
(r.name = 'Jefe de Almacén' AND p.name IN ('manage_inventory', 'approve_purchases')) OR
(r.name = 'Contador' AND p.name IN ('edit_financials', 'view_reports')) OR
(r.name = 'Recursos Humanos' AND p.name IN ('manage_users')) OR
(r.name = 'Soporte Técnico' AND p.name IN ('access_sensitive_data')) OR
(r.name = 'Analista de Datos' AND p.name IN ('view_reports', 'access_sensitive_data'));

-- Insertar datos en la tabla companies
INSERT INTO companies (name, ruc, email, phone, address, status) VALUES
('Ferretería San Juan S.A.', '20123456789', 'contacto@ferresanjuan.com', '015678901', 'Av. Universitaria 1234, Lima, Perú', B'1'),
('Materiales del Norte SAC', '20456789012', 'ventas@matnorte.com', '016543210', 'Calle Comercio 456, Trujillo, Perú', B'1');

-- Insertar datos en la tabla sucursales
INSERT INTO branch_offices (company_id, name, address, phone, status) VALUES
(1, 'Sucursal Lima', 'Av. Aviación 2345, Lima, Perú', '014567890', B'1'),
(1, 'Sucursal Arequipa', 'Calle Mercaderes 789, Arequipa, Perú', '054678901', B'1'),
(1, 'Sucursal Cusco', 'Av. El Sol 555, Cusco, Perú', '084567123', B'1'),
(1, 'Sucursal Chiclayo', 'Calle Balta 789, Chiclayo, Perú', '074678345', B'1'),
(2, 'Sucursal Trujillo', 'Av. América Norte 123, Trujillo, Perú', '044345678', B'1'),
(2, 'Sucursal Piura', 'Av. Grau 432, Piura, Perú', '073567890', B'1'),
(2, 'Sucursal Tacna', 'Jr. Bolognesi 210, Tacna, Perú', '052678234', B'1'),
(2, 'Sucursal Huancayo', 'Calle Real 678, Huancayo, Perú', '064345789', B'1');

-- Insertar datos en la tabla almacenes
INSERT INTO warehouses (branch_office_id, name, description, address, status) VALUES
(1, 'Almacén Central Lima', 'Almacén principal en Lima', 'Av. Aviación 2345, Lima, Perú', B'1'),
(1, 'Depósito Lima', 'Depósito de materiales', 'Calle Industrial 567, Lima, Perú', B'1'),
(2, 'Almacén Arequipa', 'Centro de distribución en el sur', 'Calle Mercaderes 789, Arequipa, Perú', B'1'),
(3, 'Almacén Cusco', 'Depósito de herramientas en Cusco', 'Av. El Sol 555, Cusco, Perú', B'1'),
(4, 'Almacén Chiclayo', 'Depósito de ferretería en Chiclayo', 'Calle Balta 789, Chiclayo, Perú', B'1'),
(5, 'Almacén Trujillo', 'Almacén central en Trujillo', 'Av. América Norte 123, Trujillo, Perú', B'1'),
(6, 'Almacén Piura', 'Depósito de materiales en Piura', 'Av. Grau 432, Piura, Perú', B'1'),
(7, 'Almacén Tacna', 'Centro de distribución en Tacna', 'Jr. Bolognesi 210, Tacna, Perú', B'1'),
(8, 'Almacén Huancayo', 'Depósito de herramientas en Huancayo', 'Calle Real 678, Huancayo, Perú', B'1');

-- Insertar datos en la tabla suscripciones
INSERT INTO subscriptions (company_id, plan_id, end_date, status) VALUES
(1, 1, '2026-03-01 15:04:41.474091', 'active');

-- Insertar datos en la tabla work_areas
INSERT INTO work_areas (name, description) VALUES
('Desarrollo de Software', 'Área encargada del desarrollo de aplicaciones'),
('Recursos Humanos', 'Gestión de personal y administración'),
('Finanzas', 'Gestión financiera y contabilidad'),
('Logística', 'Manejo de inventarios y distribución'),
('Marketing', 'Publicidad y estrategias de mercado');

-- Insertar datos en la tabla job_positions
INSERT INTO job_positions (work_area_id, name, description) VALUES
(1, 'Ingeniero de Software', 'Desarrollo de aplicaciones y sistemas'),
(1, 'Arquitecto de Software', 'Diseño y planificación de sistemas complejos'),
(2, 'Gerente de Recursos Humanos', 'Supervisión del área de RRHH'),
(2, 'Especialista en Reclutamiento', 'Búsqueda y selección de personal'),
(3, 'Analista Financiero', 'Gestión de análisis financiero y reportes'),
(4, 'Coordinador de Logística', 'Gestión de almacenes y transporte'),
(5, 'Especialista en Marketing Digital', 'Estrategias de marketing online');

-- Insertar datos en la tabla employees
INSERT INTO employees (
    first_name, second_name, third_name, surname, second_surname, photo, warehouse_id, 
    document_type, document_number, birth_date, gender, email, phone, address, 
    hire_date, job_position_id, salary, status
) VALUES
('Lenin', 'Josue', NULL, 'Monrroy', 'Vasquez', 'https://example.com/photos/gloria.jpg', 1,
 'dni', '72329210', '2002-01-25', 'M', 'lenin.monrroy@facil.com', '992754901', 
 'Jr. Ayacucho 369, Chorrillos, Lima', '2025-01-01', 1, 4500.00, B'1'),
('Juan', 'Carlos', NULL, 'González', 'Pérez', 'https://example.com/photos/juan.jpg', 1,
 'dni', '12345678', '1985-06-12', 'M', 'juan.gonzalez@example.com', '987654321', 
 'Av. Larco 123, Miraflores, Lima', '2015-08-10', 3, 7500.00, B'1'),
('María', 'Elena', NULL, 'Rodríguez', 'Lopez', 'https://example.com/photos/maria.jpg', 2,
 'dni', '87654321', '1990-03-25', 'F', 'maria.rodriguez@example.com', '987123456', 
 'Calle Los Cedros 456, San Isidro, Lima', '2018-05-20', 5, 4800.00, B'1'),
('Pedro', 'Luis', NULL, 'Fernández', 'Castro', 'https://example.com/photos/pedro.jpg', 3,
 'dni', '56781234', '1982-12-01', 'M', 'pedro.fernandez@example.com', '975312468', 
 'Jr. Amazonas 789, Ate, Lima', '2010-09-15', 6, 5500.00, B'1'),
('Ana', 'Gabriela', NULL, 'Salazar', 'Torres', 'https://example.com/photos/ana.jpg', 4,
 'dni', '34567890', '1995-07-19', 'F', 'ana.salazar@example.com', '965478123', 
 'Urb. Santa Anita, Lima', '2020-02-01', 4, 3200.00, B'1'),
('Luis', 'Eduardo', NULL, 'Cárdenas', 'Reyes', 'https://example.com/photos/luis.jpg', 5,
 'dni', '23456789', '1988-11-30', 'M', 'luis.cardenas@example.com', '951236547', 
 'Av. Grau 456, Barranco, Lima', '2016-07-18', 1, 4600.00, B'1'),
('Carmen', 'Rosa', NULL, 'Mendoza', 'Flores', 'https://example.com/photos/carmen.jpg', 6,
 'dni', '65432178', '1992-05-15', 'F', 'carmen.mendoza@example.com', '941235478', 
 'Calle Primavera 123, Surco, Lima', '2019-11-23', 5, 5200.00, B'1'),
('Jorge', 'Andrés', NULL, 'Vargas', 'Ramos', 'https://example.com/photos/jorge.jpg', 7,
 'dni', '32178945', '1979-09-10', 'M', 'jorge.vargas@example.com', '932145678', 
 'Calle Los Olivos 789, San Juan de Lurigancho, Lima', '2005-03-10', 2, 12000.00, B'1'),
('Patricia', 'Verónica', NULL, 'Ortiz', 'Fernández', 'https://example.com/photos/patricia.jpg', 8,
 'dni', '78965412', '1998-08-20', 'F', 'patricia.ortiz@example.com', '965412378', 
 'Av. Canadá 789, La Victoria, Lima', '2022-06-12', 7, 1800.00, B'1'),
('Ricardo', 'Antonio', NULL, 'Sánchez', 'Gómez', 'https://example.com/photos/ricardo.jpg', 9,
 'dni', '15935746', '1987-04-14', 'M', 'ricardo.sanchez@example.com', '911234567', 
 'Jr. Tupac Amaru 456, Rimac, Lima', '2014-12-05', 1, 6800.00, B'1'),
('Silvia', 'Beatriz', NULL, 'Chávez', 'Rivas', 'https://example.com/photos/silvia.jpg', 1,
 'dni', '25874136', '1993-06-22', 'F', 'silvia.chavez@example.com', '923654789', 
 'Calle San Felipe 321, Magdalena, Lima', '2017-09-30', 6, 5400.00, B'1'),
('José', 'Fernando', NULL, 'Navarro', 'Soto', 'https://example.com/photos/jose.jpg', 2,
 'dni', '75395148', '1980-10-05', 'M', 'jose.navarro@example.com', '987456321', 
 'Av. Universitaria 852, Los Olivos, Lima', '2007-04-28', 3, 8200.00, B'1'),
('Gloria', 'Elisabeth', NULL, 'Rojas', 'Espinoza', 'https://example.com/photos/gloria.jpg', 3,
 'dni', '36985214', '1991-02-08', 'F', 'gloria.rojas@example.com', '975214563', 
 'Jr. Ayacucho 369, Chorrillos, Lima', '2021-01-14', 6, 4500.00, B'1');

-- Insertar datos en la tabla users
INSERT INTO users (password, role_id, username, employee_id, avatar, email, settings, shortcuts)
VALUES (
    '$2a$12$Bs0zCFp19Y7CBjrUuyTlNecseEPNx0pVVjW72jmBMP6P/JbsXyE0e', -- password
    1,
    'Jlenin',
    1,
    '/assets/images/avatars/brian-hughes.jpg',
    'admin@facil.com',
    '{"layout": {}, "theme": {}}'::jsonb,
    '["apps.calendar", "apps.mailbox", "apps.contacts"]'::jsonb
);

-- Insertar datos en la tabla exchange_rates
INSERT INTO exchange_rates (base_currency_id, target_currency_id, exchange_rate, created_by) VALUES
(6, 1, 3.80, 1), -- PEN a USD
(6, 2, 4.10, 1), -- PEN a EUR
(6, 3, 4.80, 1), -- PEN a GBP
(6, 4, 0.028, 1); -- PEN a JPY

-- Insertar datos en la tabla categories
INSERT INTO categories (name, description, status) VALUES
('Herramientas Manuales', 'Categoría para herramientas como martillos, destornilladores, llaves, etc.', B'1'),
('Herramientas Eléctricas', 'Categoría para taladros, sierras eléctricas, lijadoras y otras herramientas eléctricas.', B'1'),
('Materiales de Construcción', 'Materiales como cemento, ladrillos, arena y piedra.', B'1'),
('Pinturas y Acabados', 'Pinturas, barnices, brochas, rodillos y accesorios para acabados.', B'1'),
('Plomería', 'Tubos, conexiones, válvulas, llaves de paso y artículos de plomería.', B'1'),
('Electricidad', 'Cables, enchufes, interruptores, luminarias y artículos eléctricos.', B'1'),
('Jardinería', 'Herramientas de jardinería, fertilizantes, semillas y macetas.', B'1'),
('Seguridad Industrial', 'Cascos, guantes, gafas de seguridad y ropa reflectante.', B'1'),
('Adhesivos y Selladores', 'Siliconas, pegamentos, masillas y selladores de todo tipo.', B'1'),
('Tornillería y Fijaciones', 'Tornillos, clavos, pernos, tuercas y arandelas.', B'1'),
('Cerrajería', 'Cerraduras, candados, llaves y accesorios relacionados.', B'1'),
('Ferretería General', 'Bisagras, rieles, soportes y artículos generales.', B'1'),
('Automatización y Domótica', 'Motores, sistemas de automatización y dispositivos de domótica.', B'1'),
('Maquinaria Pesada', 'Pequeña maquinaria como mezcladoras, compactadoras y otros equipos.', B'1'),
('Accesorios para Baño y Cocina', 'Grifería, accesorios para baño y artículos para cocina.', B'1'),
('Vidrio y Acrílicos', 'Cristales, acrílicos y materiales relacionados.', B'1'),
('Equipos de Medición', 'Cintas métricas, niveles, calibradores y otros instrumentos de medición.', B'1'),
('Maderas y Derivados', 'Tablas, triplay, MDF y otros derivados de madera.', B'1'),
('Sistemas de Riego', 'Tuberías, aspersores y equipos para riego.', B'1'),
('Calentadores y Climatización', 'Calentadores, ventiladores y aire acondicionado.', B'1');

-- Insertar datos en la tabla brands
INSERT INTO brands (name, description, logo_url, website_url, status) VALUES
('DeWalt', 'Fabricante de herramientas eléctricas y accesorios para profesionales y aficionados.', 'https://example.com/logos/dewalt.png', 'https://www.dewalt.com', B'1'),
('Bosch', 'Marca reconocida por herramientas eléctricas, electrodomésticos y tecnología industrial.', 'https://example.com/logos/bosch.png', 'https://www.bosch.com', B'1'),
('Stanley', 'Líder en herramientas manuales, medidores y productos de almacenamiento.', 'https://example.com/logos/stanley.png', 'https://www.stanleytools.com', B'1'),
('Makita', 'Proveedor global de herramientas eléctricas y equipos industriales.', 'https://example.com/logos/makita.png', 'https://www.makitatools.com', B'1'),
('Black+Decker', 'Fabricante de herramientas eléctricas, jardinería y electrodomésticos.', 'https://example.com/logos/blackdecker.png', 'https://www.blackanddecker.com', B'1'),
('Truper', 'Empresa líder en herramientas manuales, eléctricas y accesorios de ferretería.', 'https://example.com/logos/truper.png', 'https://www.truper.com', B'1'),
('Klein Tools', 'Fabricante de herramientas manuales y equipos eléctricos de alta calidad.', 'https://example.com/logos/klein.png', 'https://www.kleintools.com', B'1'),
('Hilti', 'Proveedor especializado en herramientas y sistemas de fijación para construcción.', 'https://example.com/logos/hilti.png', 'https://www.hilti.com', B'1'),
('Fischer', 'Expertos en soluciones de fijación y anclajes para construcción.', 'https://example.com/logos/fischer.png', 'https://www.fischer.com', B'1'),
('RIDGID', 'Marca especializada en herramientas para plomería y construcción.', 'https://example.com/logos/ridgid.png', 'https://www.ridgid.com', B'1'),
('Irwin', 'Fabricante de herramientas de corte, sujeción y perforación.', 'https://example.com/logos/irwin.png', 'https://www.irwin.com', B'1'),
('3M', 'Líder en productos industriales como adhesivos, cintas, abrasivos y más.', 'https://example.com/logos/3m.png', 'https://www.3m.com', B'1'),
('Troy-Bilt', 'Especialista en herramientas y equipos de jardinería.', 'https://example.com/logos/troybilt.png', 'https://www.troybilt.com', B'1'),
('Husqvarna', 'Fabricante de herramientas para jardinería, agricultura y construcción.', 'https://example.com/logos/husqvarna.png', 'https://www.husqvarna.com', B'1'),
('Craftsman', 'Marca confiable en herramientas manuales, eléctricas y almacenamiento.', 'https://example.com/logos/craftsman.png', 'https://www.craftsman.com', B'1'),
('Milwaukee', 'Proveedor de herramientas eléctricas y accesorios de alta calidad.', 'https://example.com/logos/milwaukee.png', 'https://www.milwaukeetool.com', B'1'),
('Metabo', 'Especialista en herramientas eléctricas y accesorios industriales.', 'https://example.com/logos/metabo.png', 'https://www.metabo.com', B'1'),
('Bostitch', 'Fabricante de herramientas de fijación como engrampadoras y clavos.', 'https://example.com/logos/bostitch.png', 'https://www.bostitch.com', B'1'),
('Simpson Strong-Tie', 'Líder en soluciones de anclaje y conectores estructurales.', 'https://example.com/logos/simpson.png', 'https://www.strongtie.com', B'1'),
('Paslode', 'Fabricante de herramientas de fijación neumáticas y a gas.', 'https://example.com/logos/paslode.png', 'https://www.paslode.com', B'1');

-- Insertar datos en la tabla products
INSERT INTO products (
    name, brand_id, handle, description, featured_image_id, price, cost, tax_rate, 
    quantity, sku, width, height, depth, weight,  extra_shipping_fee
) VALUES ('Mountain View - Canvas Print', 1, 'mountain-view-canvas-print', 'Beautiful mountain scenery captured at sunrise.', 'img-01', 20.00, 22.00, 10, 5, 'MV001', 50, 70, 2, 1.5, 2),
('Golden Forest - Canvas Print', 1, 'golden-forest-canvas-print', 'A breathtaking autumn forest with golden leaves.', 'img-02', 25.00, 27.50, 10, 3, 'GF002', 60, 80, 2, 2, 2.5),
('Misty Mountains - Canvas Print', 2, 'misty-mountains-canvas-print', 'A serene view of misty mountains fading into the sky.', 'img-03', 30.00, 33.00, 10, 4, 'MM003', 70, 90, 2, 2.5, 3),
('Snowy Peaks - Canvas Print', 2, 'snowy-peaks-canvas-print', 'A dramatic capture of snow-covered mountain peaks.', 'img-04', 28.00, 30.80, 10, 2, 'SP004', 60, 80, 2, 2.2, 2.8),
('City Sunset - Canvas Print', 3, 'city-sunset-canvas-print', 'A vibrant sunset reflecting on city buildings.', 'img-05', 22.00, 24.20, 10, 5, 'CS005', 55, 75, 2, 1.8, 2),
('Lake Serenity - Canvas Print', 3, 'lake-serenity-canvas-print', 'A peaceful lake surrounded by mountains.', 'img-06', 26.00, 28.60, 10, 3, 'LS006', 60, 80, 2, 2, 2.5),
('Rustic Road - Canvas Print', 4, 'rustic-road-canvas-print', 'A lonely road cutting through a red canyon.', 'img-07', 23.00, 25.30, 10, 4, 'RR007', 65, 85, 2, 2.3, 2.2),
('Glacier Escape - Canvas Print', 4, 'glacier-escape-canvas-print', 'An icy blue glacier with a frozen landscape.', 'img-08', 32.00, 35.20, 10, 2, 'GE008', 70, 100, 2, 3, 3.5),
('A Walk Amongst Friends - Canvas Print', 1, 'a-walk-amongst-friends-canvas-print', 'Officia amet eiusmod eu sunt tempor voluptate laboris velit nisi amet enim proident et. Consequat laborum non eiusmod cillum eu exercitation. Qui adipisicing est fugiat eiusmod esse. Sint aliqua cupidatat pariatur mollit ad est proident reprehenderit. Eiusmod adipisicing laborum incididunt sit aliqua ullamco.', '77a10fce', 9.309, 10.24, 10, 3, 'A445BV', 22, 24, 15, 3, 3);

-- Insertar datos en la tabla product_categories
INSERT INTO product_categories (product_id, category_id) VALUES
(1, 1),
(1, 2);

-- Insertar datos en la tabla product_images
INSERT INTO product_images (product_id, url, featured) VALUES
(1, 'uploads/images/products/01-320x200.jpg', 'img-01'),
(2, 'uploads/images/products/02-320x200.jpg', 'img-02'),
(3, 'uploads/images/products/03-320x200.jpg', 'img-03'),
(4, 'uploads/images/products/04-320x200.jpg', 'img-04'),
(5, 'uploads/images/products/10-320x200.jpg', 'img-05'),
(6, 'uploads/images/products/11-512x512.jpg', 'img-06'),
(7, 'uploads/images/products/12-512x512.jpg', 'img-07'),
(8, 'uploads/images/products/13-160x160.jpg', 'img-08'),
(9, 'uploads/images/products/+++¡.png', '77a10fce');

-- Insertar datos en la tabla customers
INSERT INTO customers (
    first_name, second_name, third_name, surname, second_surname, 
    company_name, document_type, document_number, email, address, phone
) VALUES
('Juan', 'Carlos', NULL, 'Pérez', 'Gómez', NULL, 'dni', '12345678', 'juan.perez@gmail.com', 'Av. Los Olivos 123, Lima', '987654321'),
('María', 'Luisa', NULL, 'Torres', 'Ramírez', NULL, 'dni', '87654321', 'maria.torres@gmail.com', 'Calle La Marina 456, Arequipa', '912345678'),
(NULL, NULL, NULL, NULL, NULL, 'Inversiones El Buen Precio SAC', 'ruc', '20481234567', 'contacto@buenprecio.com', 'Jr. Las Flores 789, Trujillo', '923456789'),
('Luis', 'Alberto', 'Manuel', 'García', 'López', NULL, 'dni', '23456789', 'luis.garcia@gmail.com', 'Av. Brasil 321, Lima', '934567890'),
('Ana', 'Sofía', NULL, 'Martínez', 'Díaz', NULL, 'dni', '34567890', 'ana.martinez@gmail.com', 'Calle Los Pinos 654, Cusco', '945678901'),
(NULL, NULL, NULL, NULL, NULL, 'Tech Solutions SAC', 'ruc', '20551234567', 'info@techsolutions.com', 'Av. La Paz 987, Arequipa', '956789012'),
('Carlos', 'Andrés', NULL, 'Fernández', 'Vargas', NULL, 'dni', '45678901', 'carlos.fernandez@gmail.com', 'Jr. San Martín 123, Trujillo', '967890123'),
('Lucía', 'Isabel', NULL, 'Hernández', 'Morales', NULL, 'dni', '56789012', 'lucia.hernandez@gmail.com', 'Av. Los Alamos 456, Lima', '978901234'),
(NULL, NULL, NULL, NULL, NULL, 'Construcciones Modernas SAC', 'ruc', '20661234567', 'contacto@construccionesmodernas.com', 'Calle Los Olivos 789, Cusco', '989012345'),
('Pedro', 'José', NULL, 'Sánchez', 'Ruiz', NULL, 'dni', '67890123', 'pedro.sanchez@gmail.com', 'Av. Los Jardines 321, Arequipa', '990123456'),
('Elena', 'María', NULL, 'Díaz', 'Gómez', NULL, 'dni', '78901234', 'elena.diaz@gmail.com', 'Jr. Las Rosas 654, Trujillo', '901234567'),
(NULL, NULL, NULL, NULL, NULL, 'Importaciones Globales SAC', 'ruc', '20771234567', 'info@importacionesglobales.com', 'Av. Los Pinos 987, Lima', '912345678'),
('Miguel', 'Ángel', NULL, 'López', 'García', NULL, 'dni', '89012345', 'miguel.lopez@gmail.com', 'Calle Los Laureles 123, Cusco', '923456789'),
('Carmen', 'Rosa', NULL, 'Gómez', 'Fernández', NULL, 'dni', '90123456', 'carmen.gomez@gmail.com', 'Av. Los Claveles 456, Arequipa', '934567890'),
(NULL, NULL, NULL, NULL, NULL, 'Logística Rápida SAC', 'ruc', '20881234567', 'contacto@logisticarapida.com', 'Jr. Las Palmeras 789, Trujillo', '945678901');

-- Insertar datos en la tabla suppliers
INSERT INTO suppliers (ruc, name, email, phone, address) VALUES 
('12345678901', 'Ferro Perú S.A.', 'contacto@ferroperu.com', '+51 1 234 5678', 'Av. Argentina 1543, Callao, Lima, Perú'),
('12345678902', 'Sodimac Perú S.A.', 'ventas@sodimac.com.pe', '+51 1 411 6000', 'Av. Javier Prado Este 4200, Surco, Lima, Perú'),
('12345678903', 'Maestro Perú S.A.', 'info@maestro.com.pe', '+51 1 614 6000', 'Av. La Marina 2350, San Miguel, Lima, Perú'),
('12345678904', 'Disensa Perú', 'atencion@disensa.com.pe', '+51 1 719 2000', 'Av. Alfredo Mendiola 5545, Los Olivos, Lima, Perú'),
('12345678905', 'Promart Homecenter', 'soporte@promart.com.pe', '+51 1 619 1616', 'Av. Tomás Marsano 3675, Surquillo, Lima, Perú'),
('12345678906', 'Ferretería EPA', 'contacto@epa.com.pe', '+51 1 705 0000', 'Av. Nicolás Ayllón 5777, Ate, Lima, Perú'),
('12345678907', 'Cementos Pacasmayo', 'ventas@pacasmayo.com.pe', '+51 44 608 400', 'Av. Pacasmayo 230, Trujillo, La Libertad, Perú'),
('12345678908', 'Fierros & Aceros S.A.', 'ventas@fierrosyacerossa.com', '+51 1 336 7890', 'Av. Industrial 4230, San Martín de Porres, Lima, Perú'),
('12345678909', 'Siderperu', 'clientes@siderperu.com.pe', '+51 44 481 110', 'Av. Néstor Gambetta 1234, Chimbote, Áncash, Perú'),
('12345678910', 'Construrama Perú', 'info@construrama.com.pe', '+51 1 345 6789', 'Jr. Puno 1350, Cercado de Lima, Lima, Perú'),
('12345678911', 'Indeco Perú', 'ventas@indeco.com.pe', '+51 1 213 7000', 'Av. Canadá 3350, San Luis, Lima, Perú'),
('12345678912', 'Ferrotodo S.A.C.', 'ventas@ferrotodo.com.pe', '+51 1 555 7777', 'Av. Colonial 1498, Cercado de Lima, Lima, Perú'),
('12345678913', 'Metales Peruanos S.A.', 'info@metalesperuanos.com', '+51 1 619 8000', 'Jr. Parinacochas 565, La Victoria, Lima, Perú'),
('12345678914', 'Grupo Fierro', 'contacto@grupofierro.com.pe', '+51 1 567 8900', 'Av. Faucett 3000, Callao, Lima, Perú'),
('12345678915', 'Grupo Ferretero S.A.C.', 'ventas@grupoferretero.com.pe', '+51 1 700 1234', 'Av. Los Héroes 350, San Juan de Miraflores, Lima, Perú'),
('12345678916', 'Ferretería del Norte', 'norte@ferreterianorte.com.pe', '+51 44 345 6789', 'Av. España 1245, Trujillo, La Libertad, Perú'),
('12345678917', 'Aceros Arequipa', 'info@acerosarequipa.com', '+51 54 284 400', 'Av. Aviación 1000, Cerro Colorado, Arequipa, Perú'),
('12345678918', 'Tornillos y Herramientas S.A.C.', 'ventas@tyh.com.pe', '+51 1 765 4321', 'Av. Nicolás de Piérola 250, Cercado de Lima, Lima, Perú');

-- Insertar datos en la tabla payment_methods
INSERT INTO payment_methods (name, description, status)
VALUES 
('Efectivo', 'Pago en efectivo', B'1'),
('Tarjeta de Crédito', 'Pago con tarjeta de crédito', B'1'),
('Transferencia Bancaria', 'Pago por transferencia bancaria', B'1'),
('Yape', 'Pago mediante Yape', B'1'),
('Plin', 'Pago mediante Plin', B'1');

-- Insertar datos en la tabla registros de caja
INSERT INTO cash_registers (warehouse_id, user_open_id, initial_amount)
VALUES (1, 1, 100.00);

-- Insertar datos en la tabla movimientos de caja
INSERT INTO cash_movements (cash_register_id, movement_type, payment_method_id, amount, description, user_id)
VALUES (1, 'income', 2, 150.00, 'Venta de productos', 1);

-- Insertar datos en la tabla quotes
INSERT INTO quotes (
    reference, customer_id, user_id, warehouse_id, currency_id, exchange_rate, issue_date, expiration_date, quote_status,
    approved_by, approved_at, canceled_by, canceled_at, subtotal, discount, total
) VALUES
('CT-00001', 1, 1, 1, 6, 1.0, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'pending', NULL, NULL, NULL, NULL, 200.00, 10.00, 209.00),
('CT-00002', 3, 1, 2, 6, 1.0, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'approved', 1, CURRENT_TIMESTAMP, NULL, NULL, 450.00, 20.00, 472.00),
('CT-00003', 5, 1, 3, 2, 4.1, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'pending', NULL, NULL, NULL, NULL, 600.00, 30.00, 627.00),
('CT-00004', 7, 1, 4, 1, 1.0, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'rejected', NULL, NULL, 1, CURRENT_TIMESTAMP, 300.00, 15.00, 313.50),
('CT-00005', 9, 1, 5, 5, 3.7, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'pending', NULL, NULL, NULL, NULL, 520.00, 25.00, 544.50),
('CT-00006', 11, 1, 6, 3, 0.85, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'approved', 1, CURRENT_TIMESTAMP, NULL, NULL, 750.00, 35.00, 786.25),
('CT-00007', 13, 1, 7, 4, 155.0, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'canceled', NULL, NULL, 1, CURRENT_TIMESTAMP, 180.00, 9.00, 189.10),
('CT-00008', 15, 1, 8, 6, 1.0, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, 'pending', NULL, NULL, NULL, NULL, 310.00, 16.00, 323.40);

-- Insertar datos en la tabla quote_details
INSERT INTO quote_details (
    quote_id, product_id, quantity, price, discount, subtotal, total
) VALUES
(1, 1, 2, 20.00, 2.00, 40.00, 38.00),
(1, 2, 3, 25.00, 3.00, 75.00, 72.00),
(2, 3, 1, 30.00, 0.00, 30.00, 30.00),
(2, 4, 2, 28.00, 2.00, 56.00, 54.00),
(3, 5, 4, 22.00, 4.00, 88.00, 84.00),
(3, 6, 2, 26.00, 2.00, 52.00, 50.00),
(4, 7, 3, 23.00, 3.00, 69.00, 66.00),
(5, 8, 1, 32.00, 0.00, 32.00, 32.00),
(6, 9, 5, 9.31, 1.00, 46.55, 45.55),
(7, 1, 2, 20.00, 2.00, 40.00, 38.00),
(8, 2, 3, 25.00, 3.00, 75.00, 72.00);

-- Insertar datos en la tabla sale_orders
INSERT INTO sale_orders (
    reference, warehouse_id, customer_id, user_id, quote_id, issue_date, currency_id, exchange_rate, discount, subtotal, total, order_status, date_approved
) VALUES
('OV-00001', 1, 1, 1, NULL, CURRENT_TIMESTAMP, 6, 3.74, 2.00, 50.00, 66.00, 'approved', CURRENT_DATE),
('OV-00002', 1, 2, 1, 1, CURRENT_TIMESTAMP, 6, 3.74, 0.00, 100.00, 118.00, 'canceled', NULL),
('OV-00003', 1, 3, 1, NULL, CURRENT_TIMESTAMP, 6, 3.74, 5.00, 200.00, 241.00, 'issued', NULL),
('OV-00004', 1, 1, 1, 2, CURRENT_TIMESTAMP, 6, 3.74, 10.00, 300.00, 342.00, 'approved', CURRENT_DATE);

-- Insertar datos en la tabla sale_order_details
INSERT INTO sale_order_details (
    product_name, sale_order_id, product_id, quantity, discount_method, discount, price, subtotal, total
) VALUES
('A Walk Amongst Friends - Canvas Print', 1, 1, 1, B'0', NULL, 10.24, 10.24, 10.24),
('A Walk Amongst Friends - Canvas Print', 2, 1, 2, B'0', NULL, 10.24, 20.48, 20.48),
('A Walk Amongst Friends - Canvas Print', 3, 1, 5, B'0', NULL, 10.24, 51.20, 51.20);

-- Insertar datos en la tabla purchase_orders
INSERT INTO purchase_orders (reference, warehouse_id, supplier_id, currency_id, exchange_rate, discount, issue_date, tax, subtotal, total, order_status, created_by, migrate_purchase)
VALUES
('OC-00001', 1, 1, 6, 3.75, 5.00, '2024-07-01 10:00:00', 10.00, 180.00, 198.00, 'canceled', 1, B'0'),
('OC-00002', 2, 2, 6, 3.70, 0.00, '2024-07-03 15:30:00', 8.00, 240.00, 259.20, 'approved', 1, B'0'),
('OC-00003', 3, 3, 6, 3.80, 10.00, '2024-07-05 12:45:00', 9.50, 500.00, 547.50, 'partial', 1, B'1'),
('OC-00004', 4, 4, 6, 3.85, 3.00, '2024-07-07 09:20:00', 7.00, 300.00, 321.00, 'rejected', 1, B'0'),
('OC-00005', 5, 5, 6, 3.78, 5.00, '2024-07-09 14:10:00', 6.50, 420.00, 447.30, 'approved', 1, B'0'),
('OC-00006', 6, 6, 6, 3.79, 8.00, '2024-07-11 17:50:00', 8.20, 350.00, 378.70, 'partial', 1, B'1'),
('OC-00007', 7, 7, 6, 3.76, 2.00, '2024-07-13 11:05:00', 7.80, 275.00, 296.45, 'received', 1, B'0'),
('OC-00008', 8, 8, 6, 3.77, 0.00, '2024-07-15 13:25:00', 8.50, 600.00, 648.00, 'rejected', 1, B'0'),
('OC-00009', 9, 9, 6, 3.74, 4.00, '2024-07-17 16:40:00', 9.00, 480.00, 523.20, 'approved', 1, B'1'),
('OC-00010', 2, 10, 6, 3.80, 6.00, '2024-07-19 10:10:00', 7.50, 550.00, 591.25, 'received', 1, B'0');

-- Insertar datos en la tabla purchase_order_details
INSERT INTO purchase_order_details (purchase_order_id, product_id, quantity, price, subtotal, total)
VALUES
(1, 1, 5, 20.00, 100.00, 100.00),
(1, 2, 3, 25.00, 75.00, 75.00),
(2, 3, 4, 30.00, 120.00, 120.00),
(2, 4, 2, 28.00, 56.00, 56.00),
(3, 5, 5, 22.00, 110.00, 110.00),
(3, 6, 3, 26.00, 78.00, 78.00),
(4, 7, 4, 23.00, 92.00, 92.00),
(4, 8, 2, 32.00, 64.00, 64.00),
(5, 9, 3, 9.31, 27.93, 27.93),
(6, 1, 6, 20.00, 120.00, 120.00),
(7, 2, 5, 25.00, 125.00, 125.00),
(8, 3, 7, 30.00, 210.00, 210.00),
(9, 4, 4, 28.00, 112.00, 112.00),
(10, 5, 6, 22.00, 132.00, 132.00);

-- Insertar datos en la tabla purchases
INSERT INTO purchases (
    reference, invoice_number, supplier_id, warehouse_id, currency_id, exchange_rate, 
    purchase_status, purchase_order_id, issue_date, received_date, payment_date, 
    discount, subtotal, tax, total, total_paid, change, payment_method_id, 
    created_by, document_attachment, notes
)
VALUES
('CP-00001', 'INV-001', 3, 3, 6, 3.80, 'partial', 3, '2024-07-05', '2024-07-06', NULL, 10.00, 500.00, 9.50, 547.50, 0.00, 0.00, 1, 1, NULL, NULL),
('CP-00002', 'INV-002', 6, 6, 6, 3.79, 'partial', 6, '2024-07-11', '2024-07-12', NULL, 8.00, 350.00, 8.20, 378.70, 0.00, 0.00, 1, 1, NULL, NULL),
('CP-00003', 'INV-003', 7, 7, 6, 3.76, 'received', 7, '2024-07-13', '2024-07-14', '2024-07-15', 2.00, 275.00, 7.80, 296.45, 296.45, 0.00, 2, 1, NULL, NULL),
('CP-00004', 'INV-004', 9, 9, 6, 3.74, 'received', 9, '2024-07-17', '2024-07-18', '2024-07-19', 4.00, 480.00, 9.00, 523.20, 523.20, 0.00, 2, 1, NULL, NULL),
('CP-00005', 'INV-005', 10, 2, 6, 3.80, 'received', 10, '2024-07-19', '2024-07-20', '2024-07-21', 6.00, 550.00, 7.50, 591.25, 591.25, 0.00, 1, 1, NULL, NULL),
('CP-00006', 'INV-006', 2, 2, 6, 3.70, 'received', 2, '2024-07-03', '2024-07-04', '2024-07-05', 0.00, 240.00, 8.00, 259.20, 259.20, 0.00, 2, 1, NULL, NULL),
('CP-00007', 'INV-007', 1, 1, 6, 3.75, 'received', 1, '2024-07-01', '2024-07-02', '2024-07-03', 5.00, 180.00, 10.00, 198.00, 198.00, 0.00, 2, 1, NULL, NULL),
('CP-00008', 'INV-008', 4, 4, 6, 3.85, 'canceled', 4, '2024-07-07', NULL, NULL, 3.00, 300.00, 7.00, 321.00, 0.00, 0.00, NULL, 1, NULL, NULL),
('CP-00009', 'INV-009', 8, 8, 6, 3.77, 'canceled', 8, '2024-07-15', NULL, NULL, 0.00, 600.00, 8.50, 648.00, 0.00, 0.00, NULL, 1, NULL, NULL),
('CP-00010', 'INV-010', 5, 5, 6, 3.78, 'draft', 5, '2024-07-09', NULL, NULL, 5.00, 420.00, 6.50, 447.30, 0.00, 0.00, NULL, 1, NULL, NULL),
('CP-00011', 'INV-011', 2, 2, 6, 3.70, 'paid', 2, '2024-07-03', '2024-07-04', '2024-07-05', 0.00, 240.00, 8.00, 259.20, 259.20, 0.00, 2, 1, NULL, 'Pago adelantado'),
('CP-00012', 'INV-012', 6, 6, 6, 3.79, 'unpaid', 6, '2024-07-11', NULL, NULL, 8.00, 350.00, 8.20, 378.70, 0.00, 0.00, 1, 1, NULL, 'Pendiente de abono');

-- Insertar datos en la tabla purchase_details
INSERT INTO purchase_details (purchase_id, product_id, quantity, price, discount, subtotal, total)
VALUES
(1, 5, 5, 22.00, 0.00, 110.00, 110.00),
(1, 6, 3, 26.00, 0.00, 78.00, 78.00),
(2, 1, 6, 20.00, 0.00, 120.00, 120.00),
(3, 2, 5, 25.00, 0.00, 125.00, 125.00),
(4, 4, 4, 28.00, 0.00, 112.00, 112.00),
(5, 5, 6, 22.00, 0.00, 132.00, 132.00),
(6, 3, 4, 30.00, 0.00, 120.00, 120.00),
(6, 4, 2, 28.00, 0.00, 56.00, 56.00),
(7, 1, 5, 20.00, 0.00, 100.00, 100.00),
(7, 2, 3, 25.00, 0.00, 75.00, 75.00),
(8, 7, 4, 23.00, 0.00, 92.00, 92.00),
(8, 8, 2, 32.00, 0.00, 64.00, 64.00),
(9, 3, 7, 30.00, 0.00, 210.00, 210.00),
(10, 9, 3, 9.31, 0.00, 27.93, 27.93),
(11, 3, 4, 30.00, 0.00, 120.00, 120.00),
(11, 4, 2, 28.00, 0.00, 56.00, 56.00),
(12, 1, 6, 20.00, 0.00, 120.00, 120.00),
(12, 6, 2, 24.00, 0.00, 48.00, 48.00); 

-- Insertar datos en la tabla purchase_order_payments
INSERT INTO purchase_order_payments (payment_method_id, purchase_order_id, payment_date, amount, currency_id, exchange_rate, reference_number, status, notes) VALUES
(1, 1, '2023-10-01 10:00:00', 1500.00, 6, 1.0, NULL, 'completed', 'Pago en efectivo en tienda'),
(2, 2, '2023-10-02 15:30:00', 2500.00, 6, 1.0, '123456789', 'completed', 'Pago con tarjeta Visa terminada en 1234'),
(3, 3, '2023-10-03 09:45:00', 3000.00, 6, 1.0, '987654321', 'completed', 'Transferencia desde BCP'),
(4, 4, '2023-10-04 12:15:00', 500.00, 6, 1.0, 'YP123456', 'completed', 'Pago con Yape desde el número 999888777'),
(5, 5, '2023-10-05 14:20:00', 750.00, 6, 1.0, 'PL987654', 'completed', 'Pago con Plin desde el número 999111222');

-- Insertar datos en la tabla sales
INSERT INTO sales (document_type, series, number, bill, issue_date, warehouse_id, customer_id, currency_id, user_id, exchange_rate, discount, subtotal, total, total_paid, change, sale_status, payment_method_id, sale_order_id)
VALUES
('ticket', 'B001', 1000001, 'B001-1000001', NOW(), 1, 1, 6, 1, 3.80, 5.00, 100.00, 113.00, 113.00, 0, 'issued', 1, NULL),
('invoice', 'F001', 1000001, 'F001-1000001', NOW(), 1, 2, 1, 1, 1.00, 0.00, 200.00, 236.00, 240.00, 4, 'paid', 2, NULL),
('ticket', 'B001', 1000002, 'B001-1000002', NOW(), 2, 3, 2, 1, 1.10, 10.00, 150.00, 166.00, 170.00, 4, 'unpaid', 3, 1),
('invoice', 'F001', 1000002, 'F001-1000002', NOW(), 1, 4, 5, 1, 17.50, 0.00, 180.00, 212.40, 212.40, 0, 'issued', 1, 2),
('ticket', 'B001', 1000003, 'B001-1000003', NOW(), 3, 5, 6, 1, 3.80, 5.00, 120.00, 135.60, 135.60, 0, 'pending', 2, NULL),
('invoice', 'F001', 1000003, 'F001-1000003', NOW(), 2, 6, 3, 1, 1.20, 0.00, 90.00, 106.20, 106.20, 0, 'canceled', 3, NULL),
('invoice', 'F001', 1000004, 'F001-1000004', NOW(), 3, 7, 4, 1, 0.85, 2.00, 300.00, 354.00, 360.00, 6, 'issued', 1, 3),
('invoice', 'F001', 1000005, 'F001-1000005', NOW(), 1, 8, 6, 1, 3.80, 0.00, 110.00, 124.30, 124.30, 0, 'paid', 2, NULL);

-- Insertar datos en la tabla sale_details
INSERT INTO sale_details (product_name, sale_id, product_id, quantity, discount_method, discount, price, subtotal, total)
VALUES
('Laptop HP', 1, 1, 2, B'1', 5.00, 50.00, 100.00, 113.00),
('Mouse Logitech', 1, 2, 1, B'1', 0.00, 20.00, 20.00, 22.60),
('Monitor LG', 2, 3, 1, B'1', 0.00, 200.00, 200.00, 236.00),
('Teclado Mecánico', 3, 4, 3, B'0', 10.00, 50.00, 150.00, 166.00),
('Silla Gamer', 4, 5, 1, B'1', 0.00, 180.00, 180.00, 212.40),
('Impresora Epson', 5, 6, 2, B'1', 5.00, 60.00, 120.00, 135.60),
('Tablet Samsung', 6, 7, 1, B'1', 0.00, 90.00, 90.00, 106.20),
('Disco Duro SSD', 7, 8, 3, B'0', 2.00, 100.00, 300.00, 354.00),
('Auriculares Sony', 8, 9, 2, B'1', 0.00, 55.00, 110.00, 124.30);

-- Insertar datos en la tabla opportunity_tracking
INSERT INTO opportunity_tracking (customer_id, user_id, title, status) VALUES
(1, 1, 'Compra de herramientas eléctricas', 'open'),
(2, 1, 'Adquisición de materiales de construcción', 'lost'),
(3, 1, 'Pedido de maquinaria pesada', 'won'),
(4, 1, 'Solicitud de cotización para ferretería', 'lost'),
(5, 1, 'Proyecto de remodelación de oficinas', 'in_progress'),
(6, 1, 'Compra de equipos de protección personal', 'won'),
(7, 1, 'Solicitud de compra de cemento y fierros', 'open'),
(8, 1, 'Pedido de herramientas manuales', 'lost'),
(9, 1, 'Cotización de pintura y acabados', 'won'),
(10, 1, 'Requerimiento de tuberías y conexiones', 'in_progress');

-- Insertar datos en la tabla purchase_requests
INSERT INTO purchase_requests (reference, supplier_id, user_id, request_date, expected_delivery_date, status, priority, total_cost, approved_by, approval_date, notes)
VALUES
('PR000001', 1, 1, '2024-02-01', '2024-02-10', 'pending', 'high', 1500.00, NULL, NULL, 'Solicitud urgente para proyecto A'),
('PR000002', 2, 1, '2024-02-02', '2024-02-15', 'approved', 'medium', 2450.50, 1, '2024-02-03 10:30:00', 'Pedido aprobado para almacén'),
('PR000003', 3, 1, '2024-02-03', '2024-02-18', 'pending', 'low', 780.00, NULL, NULL, 'Compra de materiales básicos'),
('PR000004', 1, 1, '2024-02-05', '2024-02-12', 'approved', 'high', 3120.75, 1, '2024-02-06 15:00:00', 'Solicitud para nueva obra'),
('PR000005', 4, 1, '2024-02-06', '2024-02-20', 'rejected', 'medium', 930.50, 1, '2024-02-07 12:45:00', 'Rechazado por presupuesto insuficiente'),
('PR000006', 5, 1, '2024-02-07', '2024-02-25', 'pending', 'low', 1200.00, NULL, NULL, 'Pedido de herramientas'),
('PR000007', 2, 1, '2024-02-08', '2024-02-14', 'approved', 'medium', 5670.90, 1, '2024-02-09 09:20:00', 'Material aprobado para producción'),
('PR000008', 3, 1, '2024-02-09', '2024-02-22', 'pending', 'high', 850.75, NULL, NULL, 'Solicitud para mantenimiento'),
('PR000009', 4, 1, '2024-02-10', '2024-02-28', 'rejected', 'medium', 2100.30, 1, '2024-02-11 14:00:00', 'Rechazado por duplicidad'),
('PR000010', 1, 1, '2024-02-11', '2024-03-01', 'approved', 'high', 4730.00, 1, '2024-02-12 11:10:00', 'Pedido importante para obra nueva');

INSERT INTO purchase_request_details (purchase_request_id, product_id, quantity, unit_price, estimated_cost, received_quantity, status, comments)
VALUES
(1, 2, 10, 150.00, 1500.00, 0, 'pending', 'Pendiente de aprobación'),
(2, 3, 5, 490.10, 2450.50, 5, 'received', 'Recibido sin problemas'),
(3, 1, 15, 52.00, 780.00, 0, 'pending', 'Esperando confirmación de proveedor'),
(4, 4, 8, 390.00, 3120.75, 8, 'received', 'Entrega completada correctamente'),
(5, 5, 7, 133.50, 930.50, 0, 'canceled', 'Pedido cancelado por error en solicitud'),
(6, 2, 12, 100.00, 1200.00, 0, 'pending', 'Pendiente de stock'),
(7, 3, 6, 945.15, 5670.90, 6, 'received', 'Material recibido con una unidad defectuosa'),
(8, 1, 5, 170.15, 850.75, 0, 'pending', 'Esperando autorización'),
(9, 4, 10, 210.03, 2100.30, 0, 'canceled', 'Orden duplicada, cancelada'),
(10, 5, 20, 236.50, 4730.00, 20, 'received', 'Pedido recibido correctamente y en buen estado');

-- Insertar datos en la tabla stock_control
INSERT INTO stock_control (warehouse_id, product_id, current_stock, current_booking) VALUES
(1, 3, 50, 20),
(2, 5, 30, 10),
(3, 7, 20, 19),
(4, 1, 10, 42),
(5, 9, 40, 8),
(6, 2, 20, 45),
(7, 4, 30, 18),
(8, 6, 20, 39),
(9, 8, 30, 31),
(1, 3, 20, 32);

-- Insertar datos en la tabla inventory_movements
INSERT INTO inventory_movements (warehouse_id, product_id, movement_type, quantity, user_id) VALUES
(1, 2, 'in', 100, 1),
(2, 4, 'out', 50, 1),
(3, 6, 'in', 200, 1),
(4, 8, 'out', 150, 1),
(5, 1, 'in', 300, 1),
(6, 3, 'out', 120, 1),
(7, 5, 'in', 80, 1),
(8, 7, 'out', 90, 1),
(9, 9, 'in', 250, 1),
(1, 2, 'out', 70, 1);

-- Insertar datos en la tabla systems
INSERT INTO systems (name, description) VALUES
('Nubefact', 'Integración con el sistema de facturación electrónica Nubefact.'),
('Odoo', 'Integración con el ERP Odoo para sincronización de datos.'),
('SUNAT', 'Conexión directa con SUNAT para validación de comprobantes.'),
('SMTP', 'Configuración del servidor de correo para envíos automáticos.'),
('Pasarela de Pagos', 'Pasarelas de pago como Culqi, MercadoPago y PayU.'),
('Sistema Propio', 'Configuraciones internas del sistema ERP.'),
('Certificados Digitales', 'Llaves y certificados digitales del sistema.');

-- Insertar datos en la tabla keys
INSERT INTO keys (system_id, name, description, config_key, config_value, status) VALUES
(1, 'Nubefact API URL', 'URL base para conexión a la API de Nubefact', 'nubefact_api_url', 'https://api.nubefact.com/api/v1/e421b983-cc37-45ca-8bbf-b4b5e4af3136', B'1'),
(1, 'Nubefact API Token', 'Token de autenticación para emisión de comprobantes en Nubefact', 'nubefact_api_token', '9a048f724546488eb4e3da7a46e27b110732865e97de49c5a07f930c57a45ffc', B'1'),
(2, 'Odoo Base URL', 'URL base para conexión al servidor Odoo XML-RPC', 'odoo_base_url', 'https://erp.miempresa.com/xmlrpc/2/object', B'1'),
(2, 'Odoo Database Name', 'Nombre de la base de datos de Odoo', 'odoo_db_name', 'erp_miempresa', B'1'),
(2, 'Odoo Username', 'Usuario administrador de Odoo', 'odoo_username', 'admin', B'1'),
(2, 'Odoo Password', 'Contraseña del usuario de Odoo', 'odoo_password', 'odooPass123', B'1'),
(3, 'SUNAT SOL Username', 'Usuario SOL para conexión a SUNAT', 'sunat_sol_username', '20555555555MODDATOS', B'1'),
(3, 'SUNAT SOL Password', 'Contraseña SOL del usuario para SUNAT', 'sunat_sol_password', 'miClaveSegura123', B'1'),
(3, 'SUNAT URL Beta', 'URL del servicio web de SUNAT en ambiente de prueba', 'sunat_beta_url', 'https://e-beta.sunat.gob.pe/ol-ti-itcpfegem-beta/billService', B'1'),
(3, 'SUNAT URL Producción', 'URL del servicio web de SUNAT en producción', 'sunat_prod_url', 'https://e-factura.sunat.gob.pe/ol-ti-itcpfegem/billService', B'1'),
(3, 'SUNAT Certificate Path', 'Ruta del certificado digital para SUNAT', 'sunat_cert_path', '/certs/sunat_cert.pfx', B'1'),
(3, 'SUNAT Certificate Password', 'Contraseña del certificado digital SUNAT', 'sunat_cert_password', 'certPassword123', B'1'),
(4, 'SMTP Host', 'Servidor de correo SMTP', 'smtp_host', 'smtp.miempresa.com', B'1'),
(4, 'SMTP Port', 'Puerto del servidor SMTP', 'smtp_port', '587', B'1'),
(4, 'SMTP User', 'Correo electrónico utilizado para enviar los correos', 'smtp_user', 'noreply@miempresa.com', B'1'),
(4, 'SMTP Password', 'Contraseña del correo SMTP', 'smtp_password', 'correoSuperSeguro123', B'1'),
(4, 'SMTP Encryption', 'Método de cifrado utilizado (TLS o SSL)', 'smtp_encryption', 'TLS', B'1'),
(5, 'Culqi Public Key', 'Llave pública para el cliente en Culqi', 'culqi_public_key', 'pk_test_CULQIPUBLICKEY', B'1'),
(5, 'Culqi Private Key', 'Llave secreta para operaciones en Culqi', 'culqi_private_key', 'sk_test_CULQIPRIVATEKEY', B'1'),
(5, 'Culqi API URL', 'URL base para el servicio de Culqi', 'culqi_api_url', 'https://api.culqi.com/v2', B'1'),
(5, 'MercadoPago Public Key', 'Llave pública para el cliente en MercadoPago', 'mercadopago_public_key', 'PUBLIC_MP_KEY', B'1'),
(5, 'MercadoPago Access Token', 'Token de acceso privado de MercadoPago', 'mercadopago_access_token', 'ACCESS_TOKEN_MP', B'1'),
(6, 'Sistema Nombre', 'Nombre de la empresa o sistema', 'system_name', 'ERP Mi Empresa', B'1'),
(6, 'Sistema Modo', 'Modo del sistema (produccion/desarrollo)', 'system_mode', 'produccion', B'1'),
(6, 'JWT Secret', 'Llave secreta utilizada para firmar los JWT tokens', 'jwt_secret', 'jwtSuperSecret123456', B'1'),
(6, 'Token Expiración', 'Tiempo de expiración del token JWT en minutos', 'jwt_expiration_minutes', '120', B'1'),
(6, 'URL Frontend', 'URL del frontend del sistema', 'frontend_url', 'https://miempresa.app', B'1'),
(6, 'URL Backend', 'URL del backend del sistema', 'backend_url', 'https://api.miempresa.com', B'1'),
(7, 'Certificado SSL Path', 'Ruta donde se almacena el certificado SSL', 'ssl_cert_path', '/certs/ssl_cert.pem', B'1'),
(7, 'Llave Privada SSL Path', 'Ruta donde se almacena la llave privada del SSL', 'ssl_private_key_path', '/certs/ssl_key.pem', B'1'),
(7, 'Certificado SSL Password', 'Contraseña para el acceso al certificado SSL si aplica', 'ssl_cert_password', 'sslPassword123', B'1'),
(7, 'Certificado SUNAT PFX', 'Ruta del archivo PFX del certificado digital SUNAT', 'sunat_cert_pfx_path', '/certs/sunat_cert.pfx', B'1'),
(7, 'Password SUNAT PFX', 'Password del archivo PFX de SUNAT', 'sunat_cert_pfx_password', 'pfxPassword321', B'1');