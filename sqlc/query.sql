-- name: GetUpstreamById :one
SELECT * FROM upstreams
WHERE id = $1
LIMIT 1;

-- name: GetUpstreamConflic :one
SELECT * FROM upstreams
WHERE name = $1
LIMIT 1;

-- name: CreateUpstream :one
INSERT INTO upstreams (name) VALUES ($1) RETURNING *;
