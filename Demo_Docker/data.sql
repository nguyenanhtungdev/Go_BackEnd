CREATE TABLE product (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    price NUMERIC NOT NULL,
    description TEXT
);

INSERT INTO product (name, price, description) VALUES
('Product 1', 100.0, 'Description of product 1'),
('Product 2', 200.0, 'Description of product 2');
