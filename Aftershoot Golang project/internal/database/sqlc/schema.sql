-- schema.sql

-- Create products table
CREATE TABLE products (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    price DECIMAL(10, 2) NOT NULL,
    quantity INT NOT NULL
);


-- Create ratings table
CREATE TABLE ratings (
   id SERIAL PRIMARY KEY,
   product_id INT NOT NULL REFERENCES products(id),
   rating DECIMAL(3, 2) NOT NULL
);

-- Create the users table
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username TEXT NOT NULL,
    password TEXT NOT NULL,
    role TEXT NOT NULL
);

-- customers table
CREATE TABLE customers (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL
);

-- orders table
CREATE TABLE orders (
    id SERIAL PRIMARY KEY,
    product_id INT NOT NULL REFERENCES products(id),
    customer_id INT NOT NULL REFERENCES customers(id),
    quantity INT NOT NULL,
    price INT NOT NULL,
    orderStatus TEXT,
    order_date TIMESTAMP
);
