-- +migrate Up
CREATE TABLE procurement (
                             id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
                             product_id UUID NOT NULL,
                             quantity INT NOT NULL,
                             created_at TIMESTAMP DEFAULT NOW(),
                             FOREIGN KEY (product_id) REFERENCES products(id)
);