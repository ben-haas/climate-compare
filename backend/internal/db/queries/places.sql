-- name: CreatePlace :one
INSERT INTO places (name, country, latitude, longitude, altitude)
VALUES ($1, $2, $3, $4, $5)
RETURNING *;

-- name: GetPlace :one
SELECT *
FROM places
WHERE id = $1
LIMIT 1;

-- name: ListPlaces :many
SELECT *
FROM places
ORDER BY id
LIMIT $1 OFFSET $2;

-- name: DeletePlace :exec
DELETE FROM places
WHERE id = $1;