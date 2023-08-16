-- customers.sql

-- Create customers table
-- Include necessary columns such as customer ID, name, email, etc.
CREATE TABLE customers (
    id SERIAL PRIMARY KEY,
    name TEXT NOT NULL,
    email TEXT NOT NULL,
    -- Add other columns as needed
);
