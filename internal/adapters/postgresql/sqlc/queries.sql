-- name: ListProducts :one
SELECT * FROM products;

-- name: FindProductByID :one
SELECT * FROM products WHERE id = $1;