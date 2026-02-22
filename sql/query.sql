CREATE TABLE products (
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    description VARCHAR NOT NULL,
    stock_quantity INT NOT NULL,
    price DECIMAL(10,2) NOT NULL,
    category_id BIGINT NOT NULL,
    is_available BOOLEAN NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY(category_id) REFERENCES categories(id)
);

CREATE TABLE categories {
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    description VARCHAR NOT NULL
};

CREATE TABLE orders {
    id BIGSERIAL PRIMARY KEY,
    customer_id BIGINT NOT NULL,
    product_id BIGINT NOT NULL,
    quantity INT NOT NULL,
    total_price DECIMAL(10,2) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
};

CREATE TABLE customers {
    id BIGSERIAL PRIMARY KEY,
    name VARCHAR NOT NULL,
    email VARCHAR NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
};

CREATE TABLE order_items {
    id BIGSERIAL PRIMARY KEY,
    order_id BIGINT NOT NULL,
    product_id BIGINT NOT NULL,
    quantity INT NOT NULL,
    unit_price DECIMAL(10,2) NOT NULL,
};

-- First create index on keys
CREATE INDEX IDX_product_category_id ON products(category_id);
CREATE INDEX IDX_order_customer_id ON orders(customer_id);
CREATE INDEX IDX_order_product_id ON orders(product_id);
CREATE INDEX IDX_order_item_order_id ON order_items(order_id);
CREATE INDEX IDX_order_item_product_id ON order_items(product_id);

-- 3.1
SELECT oi.product_id, p.name, c.name SUM(oi.quantity) as quantity_sold FROM order_items AS oi
INNER JOIN SELECT p.id FROM products AS p ON oi.product_id = p.id
INNER JOIN SELECT categories ON p.category_id = c.id
GROUP BY oi.product_id;

-- 3.2
SELECT cu.name, SUM(o.total_price) AS spent FROM orders AS o
INNER JOIN customers AS cu ON cu.id = o.customer_id
GROUP BY o.customer_id
ORDER BY spent DESC
LIMIT 10;

-- 3.3
SELECT oi.order_id, o.quantity, cu.name, p.name
FROM order_id AS oi
INNER JOIN orders AS o ON oi.order_id = o.id
INNER JOIN customers AS cu ON o.customer_id = cu.id
INNER JOIN products AS p ON p.id = oi.product_id;