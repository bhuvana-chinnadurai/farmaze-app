-- +migrate Up
ALTER TABLE products
    ADD COLUMN unit VARCHAR(255);