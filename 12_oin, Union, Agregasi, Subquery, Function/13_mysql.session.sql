CREATE TABLE operators (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE product_types (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
CREATE TABLE product_descriptions (
    id INT AUTO_INCREMENT PRIMARY KEY,
    description TEXT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
CREATE TABLE products (
    id INT AUTO_INCREMENT PRIMARY KEY,
    product_type_id INT,
    operator_id INT,
    code VARCHAR(50) NOT NULL,
    name VARCHAR(100) NOT NULL,
    status SMALLINT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    FOREIGN KEY (product_type_id) REFERENCES product_types(id),
    FOREIGN KEY (operator_id) REFERENCES operators(id)
);
CREATE TABLE payment_methods (
    id INT AUTO_INCREMENT PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    status SMALLINT NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
CREATE TABLE users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    status SMALLINT NOT NULL,
    dob DATE NOT NULL,
    gender CHAR(1) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);
CREATE TABLE transactions (
    id INT AUTO_INCREMENT PRIMARY KEY,
    user_id INT,
    payment_method_id INT,
    status VARCHAR(10) NOT NULL,
    total_only INT NOT NULL,
    total_price NUMERIC(25,2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (payment_method_id) REFERENCES payment_methods(id)
);
CREATE TABLE transaction_details (
    transaction_id INT,
    product_id INT,
    qty INT NOT NULL,
    price NUMERIC(25,2) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    PRIMARY KEY (transaction_id, product_id),
    FOREIGN KEY (transaction_id) REFERENCES transactions(id),
    FOREIGN KEY (product_id) REFERENCES products(id)
);
INSERT INTO operators (name, created_at, updated_at) VALUES
('Operator A', NOW(), NOW()),
('Operator B', NOW(), NOW()),
('Operator C', NOW(), NOW()),
('Operator D', NOW(), NOW()),
('Operator E', NOW(), NOW());
INSERT INTO product_types (name, created_at, updated_at) VALUES
('Type 1', NOW(), NOW()),
('Type 2', NOW(), NOW()),
('Type 3', NOW(), NOW());
-- Insert 2 products with product type id = 1, and operator id = 3
INSERT INTO products (product_type_id, operator_id, code, name, status, created_at, updated_at) VALUES
(1, 3, 'P001', 'Product 1', 1, NOW(), NOW()),
(1, 3, 'P002', 'Product 2', 1, NOW(), NOW());

-- Insert 3 products with product type id = 2, and operator id = 1
INSERT INTO products (product_type_id, operator_id, code, name, status, created_at, updated_at) VALUES
(2, 1, 'P003', 'Product 3', 1, NOW(), NOW()),
(2, 1, 'P004', 'Product 4', 1, NOW(), NOW()),
(2, 1, 'P005', 'Product 5', 1, NOW(), NOW());

-- Insert 3 products with product type id = 3, and operator id = 4
INSERT INTO products (product_type_id, operator_id, code, name, status, created_at, updated_at) VALUES
(3, 4, 'P006', 'Product 6', 1, NOW(), NOW()),
(3, 4, 'P007', 'Product 7', 1, NOW(), NOW()),
(3, 4, 'P008', 'Product 8', 1, NOW(), NOW());
INSERT INTO product_descriptions (id, description, created_at, updated_at) VALUES
(1, 'Description for Product 1', NOW(), NOW()),
(2, 'Description for Product 2', NOW(), NOW()),
(3, 'Description for Product 3', NOW(), NOW()),
(4, 'Description for Product 4', NOW(), NOW()),
(5, 'Description for Product 5', NOW(), NOW()),
(6, 'Description for Product 6', NOW(), NOW()),
(7, 'Description for Product 7', NOW(), NOW()),
(8, 'Description for Product 8', NOW(), NOW());

INSERT INTO payment_methods (name, status, created_at, updated_at) VALUES
('Credit Card', 1, NOW(), NOW()),
('PayPal', 1, NOW(), NOW()),
('Bank Transfer', 1, NOW(), NOW());

INSERT INTO users (status, dob, gender, created_at, updated_at) VALUES
(1, '1990-01-01', 'M', NOW(), NOW()),
(1, '1992-02-02', 'F', NOW(), NOW()),
(1, '1994-03-03', 'M', NOW(), NOW()),
(1, '1996-04-04', 'F', NOW(), NOW()),
(1, '1998-05-05', 'M', NOW(), NOW());

INSERT INTO transactions (user_id, payment_method_id, status, total_only, total_price, created_at) VALUES
(1, 1, 'Completed', 1, 100.00, NOW()),
(1, 2, 'Pending', 1, 50.00, NOW()),
(1, 3, 'Completed', 1, 150.00, NOW()),
(2, 1, 'Completed', 1, 200.00, NOW()),
(2, 2, 'Pending', 1, 75.00, NOW()),
(2, 3, 'Completed', 1, 125.00, NOW()),
(3, 1, 'Completed', 1, 300.00, NOW()),
(3, 2, 'Pending', 1, 50.00, NOW()),
(3, 3, 'Completed', 1, 175.00, NOW()),
(4, 1, 'Completed', 1, 400.00, NOW()),
(4, 2, 'Pending', 1, 60.00, NOW()),
(4, 3, 'Completed', 1, 80.00, NOW()),
(5, 1, 'Completed', 1, 500.00, NOW()),
(5, 2, 'Pending', 1, 90.00, NOW()),
(5, 3, 'Completed', 1, 110.00, NOW());

-- This is a generalized example; it should be adjusted to actual transaction IDs and product IDs.
INSERT INTO transaction_details (transaction_id, product_id, qty, price, created_at, updated_at) VALUES
(1, 1, 1, 20.00, NOW(), NOW()),
(1, 2, 2, 30.00, NOW(), NOW()),
(1, 3, 3, 50.00, NOW(), NOW());
-- Repeat similarly for other transactions and products.

SELECT name FROM users WHERE gender = 'M';

SELECT * FROM products WHERE id = 3;

SELECT * FROM users WHERE created_at >= DATE_SUB(NOW(), INTERVAL 7 DAY) AND name LIKE '%a%';

SELECT COUNT(*) FROM users WHERE gender = 'F';

SELECT * FROM users ORDER BY name ASC;

SELECT * FROM products LIMIT 5;

UPDATE products SET name = 'product dummy' WHERE id = 1;

UPDATE transaction_details SET qty = 3 WHERE product_id = 1;

DELETE FROM products WHERE id = 1;

DELETE FROM products WHERE product_type_id = 1;

SELECT * FROM transactions WHERE user_id IN (1, 2);

SELECT SUM(total_price) AS total_price FROM transactions WHERE user_id = 1;

SELECT COUNT(*) AS total_transactions FROM transactions t
JOIN products p ON t.id = p.id
WHERE p.product_type_id = 2;

SELECT p.*, pt.name AS product_type_name FROM products p
JOIN product_types pt ON p.product_type_id = pt.id;

SELECT t.*, p.name AS product_name, u.name AS user_name FROM transactions t
JOIN products p ON t.product_id = p.id
JOIN users u ON t.user_id = u.id;

CREATE TRIGGER after_transaction_delete
AFTER DELETE ON transactions
FOR EACH ROW
BEGIN
    DELETE FROM transaction_details WHERE transaction_id = OLD.id;
END;

CREATE TRIGGER after_transaction_detail_delete
AFTER DELETE ON transaction_details
FOR EACH ROW
BEGIN
    UPDATE transactions SET total_qty = total_qty - OLD.qty WHERE id = OLD.transaction_id;
END;

SELECT * FROM products WHERE id NOT IN (SELECT product_id FROM transaction_details);

