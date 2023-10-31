-- +migrate Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE products (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255),
    price DECIMAL(10, 2),
    description TEXT,
    available_quantity INT DEFAULT 0,
    category VARCHAR(255)
);

