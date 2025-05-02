-- name: GetUpstreamById :one
SELECT * FROM upstreams
WHERE id = $1
LIMIT 1;

-- name: GetUpstreamConflic :one
SELECT * FROM upstreams
WHERE name = $2 and id <> $1
LIMIT 1;

-- name: CreateUpstream :one
INSERT INTO upstreams (name) VALUES ($1) RETURNING *;

-- name: UpdateUpstream :exec
UPDATE upstreams SET name = $2 WHERE id = $1;

-- name: DeleteUpstream :one
DELETE FROM upstreams WHERE id = $1 RETURNING *;
