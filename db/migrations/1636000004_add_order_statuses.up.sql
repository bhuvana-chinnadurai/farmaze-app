-- +migrate Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE order_statuses (
    id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
    name VARCHAR(255) NOT NULL
);

INSERT INTO order_statuses (name) VALUES 
    ('ordered'),
    ('dispatched'),
    ('delivered');