-- name: SetCache :execlastid
INSERT INTO caches
  (name, value, updated_at, expires)
VALUES
(?, ?, ?, ?);

-- name: GetCacheByID :one
SELECT * FROM caches
WHERE id = ? LIMIT 1;

-- name: GetCacheByName :one
SELECT * FROM caches
WHERE name = ? LIMIT 1;

-- name: GetCacheAll :many
SELECT * FROM caches;

-- name: GetCacheStale :many
SELECT * FROM caches
WHERE expires < ?;

-- name: DropCacheStale :execrows
DELETE FROM caches
WHERE expires < ?;

-- name: WipeCache :execrows
DELETE FROM caches;