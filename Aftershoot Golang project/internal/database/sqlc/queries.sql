-- queries.sql

-- Fetch a list of products along with their average customer ratings
-- Use a JOIN to combine products and ratings
-- Calculate the average rating for each product
SELECT
    p.id,
    p.name,
    p.price,
    p.quantity,
    AVG(r.rating) AS avg_rating
FROM products p
LEFT JOIN ratings r ON p.id = r.product_id
GROUP BY p.id, p.name, p.price, p.quantity
ORDER BY avg_rating DESC;



-- Check if a user has admin role
-- Returns true if the user has admin role, false otherwise
-- @param userID: ID of the user
SELECT EXISTS(
    SELECT * FROM users WHERE id = $1 AND role = 'admin'
) AS is_admin;


SELECT
    c.id,
    c.name,
    c.email,
    COUNT(o.id) AS order_count
FROM
    customers c
JOIN
    orders o ON c.id = o.customer_id
GROUP BY
    c.id, c.name, c.email
ORDER BY
    order_count DESC
LIMIT
    3;


-- Fetch orders along with details about the products and customers involved
SELECT
    o.id AS order_id,
    o.order_date,
    c.id AS customer_id,
    c.name AS customer_name,
    c.email AS customer_email,
    p.id AS product_id,
    p.name AS product_name,
    p.price AS product_price,
    p.quantity AS product_quantity
FROM
    orders o
JOIN
    customers c ON o.customer_id = c.id
JOIN
    products p ON o.product_id = p.id;

