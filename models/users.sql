-- name: UserCreate :execlastid
INSERT INTO users
  (email, password_hash, data, created_at)
  VALUES
  (?, ?, ?, ?);

-- name: UserCheckPassword :one
SELECT id, email, password_hash, status
  FROM users
  WHERE email = ?
  LIMIT 1;

-- name: UserFindByEmail :one
SELECT id, email, JSON_EXTRACT(data, '$.zips') AS zips, created_at, status
  FROM users
  WHERE email = ?
  LIMIT 1;

-- name: UserExtractZips :one
SELECT JSON_EXTRACT(data, '$.zips')
  FROM users
  WHERE email = ?;

-- name: UserAppendZip :exec
UPDATE users
  SET data = JSON_ARRAY_APPEND(data, '$.zips', ?)
  WHERE email = ?;

-- name: UserSetZips :exec
UPDATE users
  SET data = JSON_SET(data, '$.zips', CAST(? AS JSON))
  WHERE email = ?;

-- name: UserClearZips :exec
UPDATE users
  SET data = JSON_SET(data, '$.zips', JSON_ARRAY())
  WHERE email = ?;

-- name: UsersDelete :exec
DELETE FROM users;