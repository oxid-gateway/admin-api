-- name: GetUpstreamById :one
SELECT * FROM upstreams
WHERE id = $1
LIMIT 1;
