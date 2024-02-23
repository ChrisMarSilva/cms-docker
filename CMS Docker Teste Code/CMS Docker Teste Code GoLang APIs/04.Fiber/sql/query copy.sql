
/*PostgreSQL*/
/*
-- name: GetAuthor :one
SELECT * FROM authors WHERE id = $1 LIMIT 1;

-- name: ListAuthors :many
SELECT * FROM authors ORDER BY name;

-- name: CreateAuthor :one
INSERT INTO authors (name, bio) VALUES ($1, $2) RETURNING *;

-- name: UpdateAuthor :one
UPDATE authors set name = $2, bio = $3 WHERE id = $1 RETURNING *;

-- name: DeleteAuthor :exec
DELETE FROM authors WHERE id = $1;
*/


/*MySQL*/
/*
-- name: GetAuthor :one
SELECT * FROM authors WHERE id = ? LIMIT 1;

-- name: ListAuthors :many
SELECT * FROM authors ORDER BY name;

-- name: CreateAuthor :execresult
INSERT INTO authors (name, bio) VALUES (?, ?);

-- name: UpdateAuthor :exec
UPDATE authors set name = ?, bio = ? WHERE id = ?;

-- name: DeleteAuthor :exec
DELETE FROM authors WHERE id = ?;
*/


/*SQLite*/
/*
-- name: GetAuthor :one
SELECT * FROM authors WHERE id = ? LIMIT 1;

-- name: ListAuthors :many
SELECT * FROM authors ORDER BY name;

-- name: CreateAuthor :one
INSERT INTO authors (name, bio) VALUES (?, ?) RETURNING *;

-- name: UpdateAuthor :exec
UPDATE authors set name = ?, bio = ? WHERE id = ?;

-- name: DeleteAuthor :exec
DELETE FROM authors WHERE id = ?;
*/