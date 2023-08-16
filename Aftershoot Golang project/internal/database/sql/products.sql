-- products.sql

-- Create products table
-- Include necessary columns such as product ID, name, price, etc.
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    price DECIMAL(10, 2) NOT NULL
);
