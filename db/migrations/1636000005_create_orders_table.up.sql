-- +migrate Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE orders (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    client_id UUID,
    total_price DECIMAL(10, 2),
    created_at TIMESTAMP,
    status_id UUID,
    FOREIGN KEY (client_id) REFERENCES b2b_clients(id),
    FOREIGN KEY (status_id) REFERENCES order_statuses(id)
);

CREATE TABLE order_products (
    order_id UUID,
    product_id UUID,
    quantity INT,
    price INT,
    PRIMARY KEY (order_id, product_id),
    FOREIGN KEY (order_id) REFERENCES orders(id),
    FOREIGN KEY (product_id) REFERENCES products(id)
);