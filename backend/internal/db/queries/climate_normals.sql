-- name: CreateClimate :one
INSERT INTO climate_normals (place_id, month, tavg, tmin, tmax, prcp, wspd, pres, tsun, last_updated)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
RETURNING *;

-- name: GetClimate :one
SELECT *
FROM climate_normals
WHERE place_id = $1
LIMIT 1;