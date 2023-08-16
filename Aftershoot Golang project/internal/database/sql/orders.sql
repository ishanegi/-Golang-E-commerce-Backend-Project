-- orders.sql

-- Create orders table
-- Include necessary columns such as order ID, customer ID, product ID, etc.
CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    customer_id INT NOT NULL,
    product_id INT NOT NULL,
    quantity INT NOT NULL,
    -- Add other columns as needed
);
