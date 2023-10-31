-- +migrate Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE b2b_clients (
                             id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
                             company_name VARCHAR(255),
                             contact_name VARCHAR(255),
                             email VARCHAR(255),
                             phone_number VARCHAR(20)
);