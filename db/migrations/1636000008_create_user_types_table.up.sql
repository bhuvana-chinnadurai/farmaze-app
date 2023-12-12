CREATE TABLE user_types (
    id SERIAL PRIMARY KEY,
    type_name VARCHAR(20) UNIQUE NOT NULL
);

INSERT INTO user_types (type_name) VALUES
    ('customer'),
    ('b2b'),
    ('vendor');

