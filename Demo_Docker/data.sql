CREATE TABLE product (
    productId SERIAL PRIMARY KEY,
    productName VARCHAR(255) NOT NULL,
    price NUMERIC NOT NULL,
    descriptions TEXT
);

INSERT INTO product (productName, price, descriptions) VALUES
('Product 1', 100.0, 'Description of product 1'),
('Product 2', 200.0, 'Description of product 2'),
('Product 3', 300.0, 'Description of product 3');
