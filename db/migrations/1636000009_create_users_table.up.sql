CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(50) UNIQUE NOT NULL,
    password VARCHAR(60) NOT NULL,
    role VARCHAR(20) NOT NULL,
    user_type_id INT NOT NULL,
    CONSTRAINT fk_user_type
        FOREIGN KEY (user_type_id) 
        REFERENCES user_types (id)
);

