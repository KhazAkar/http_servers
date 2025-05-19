-- name: CreateChirp :one
INSERT INTO chirps (body, user_id)
VALUES ($1, $2)
RETURNING id, created_at, updated_at, body, user_id;

-- name: GetOneChirp :one
SELECT * from chirps WHERE id = $1;

-- name: GetAllChirps :many
SELECT
    *
FROM
    chirps
ORDER BY
    created_at ASC;
