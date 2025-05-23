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

-- name: ListUpstreams :many
SELECT * FROM upstreams
WHERE name like $3
LIMIT $1 
OFFSET $2;

-- name: GetUserUpstreams :many
SELECT up.* FROM upstreams up
LEFT JOIN users_upstreams uu
ON uu.upstream_id = up.id
WHERE up.name LIKE $3 AND uu.user_id = $3
LIMIT $1 
OFFSET $2;

-- name: CountUpstreams :one
SELECT COUNT(*) FROM upstreams
WHERE name like $1;

-- name: GetUpstreamUsers :many
SELECT us.* FROM users us
LEFT JOIN users_upstreams uu
ON uu.user_id = us.id
WHERE uu.upstream_id = $3
LIMIT $1 
OFFSET $2;

-- name: CountUpstreamUsers :one
SELECT COUNT(us.*) FROM users us
LEFT JOIN users_upstreams uu
ON uu.user_id = us.id
WHERE uu.upstream_id = $1;

-- name: GetUserByUsername :one
SELECT * FROM users
WHERE username = $1
LIMIT 1;

-- name: GetUserById :one
SELECT * FROM users
WHERE id = $1
LIMIT 1;

-- name: CreateUser :one
INSERT INTO users (name, username, email) VALUES ($1, $2, $3) RETURNING *;

-- name: LinkUserToUpstream :exec
INSERT INTO users_upstreams (user_id, upstream_id) VALUES ($1, $2);

-- name: GetUserLinkUpstream :one
SELECT us.* FROM users us
LEFT JOIN users_upstreams uu
ON uu.user_id = $2
WHERE uu.upstream_id = $1
LIMIT 1;

-- name: ListUsers :many
-- SELECT * FROM users
-- WHERE name like $3 or username like $3 or email like $3
-- LIMIT $1 
-- OFFSET $2;

-- name: CountUsers :one
-- SELECT COUNT(*) FROM users
-- WHERE name like $1 or username like $1 or email like $1;
