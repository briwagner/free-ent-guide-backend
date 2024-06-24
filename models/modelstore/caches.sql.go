// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.25.0
// source: caches.sql

package modelstore

import (
	"context"
	"database/sql"
)

const dropCacheStale = `-- name: DropCacheStale :execrows
DELETE FROM caches
WHERE expires < ?
`

func (q *Queries) DropCacheStale(ctx context.Context, expires sql.NullTime) (int64, error) {
	result, err := q.db.ExecContext(ctx, dropCacheStale, expires)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}

const getCacheAll = `-- name: GetCacheAll :many
SELECT id, name, value, updated_at, expires FROM caches
`

func (q *Queries) GetCacheAll(ctx context.Context) ([]Caches, error) {
	rows, err := q.db.QueryContext(ctx, getCacheAll)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Caches
	for rows.Next() {
		var i Caches
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Value,
			&i.UpdatedAt,
			&i.Expires,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const getCacheByID = `-- name: GetCacheByID :one
SELECT id, name, value, updated_at, expires FROM caches
WHERE id = ? LIMIT 1
`

func (q *Queries) GetCacheByID(ctx context.Context, id int64) (Caches, error) {
	row := q.db.QueryRowContext(ctx, getCacheByID, id)
	var i Caches
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Value,
		&i.UpdatedAt,
		&i.Expires,
	)
	return i, err
}

const getCacheByName = `-- name: GetCacheByName :one
SELECT id, name, value, updated_at, expires FROM caches
WHERE name = ? LIMIT 1
`

func (q *Queries) GetCacheByName(ctx context.Context, name sql.NullString) (Caches, error) {
	row := q.db.QueryRowContext(ctx, getCacheByName, name)
	var i Caches
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Value,
		&i.UpdatedAt,
		&i.Expires,
	)
	return i, err
}

const getCacheStale = `-- name: GetCacheStale :many
SELECT id, name, value, updated_at, expires FROM caches
WHERE expires < ?
`

func (q *Queries) GetCacheStale(ctx context.Context, expires sql.NullTime) ([]Caches, error) {
	rows, err := q.db.QueryContext(ctx, getCacheStale, expires)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Caches
	for rows.Next() {
		var i Caches
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Value,
			&i.UpdatedAt,
			&i.Expires,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const setCache = `-- name: SetCache :execlastid
INSERT INTO caches
  (name, value, updated_at, expires)
VALUES
(?, ?, ?, ?)
`

type SetCacheParams struct {
	Name      sql.NullString
	Value     sql.NullString
	UpdatedAt sql.NullTime
	Expires   sql.NullTime
}

func (q *Queries) SetCache(ctx context.Context, arg SetCacheParams) (int64, error) {
	result, err := q.db.ExecContext(ctx, setCache,
		arg.Name,
		arg.Value,
		arg.UpdatedAt,
		arg.Expires,
	)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

const wipeCache = `-- name: WipeCache :execrows
DELETE FROM caches
`

func (q *Queries) WipeCache(ctx context.Context) (int64, error) {
	result, err := q.db.ExecContext(ctx, wipeCache)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected()
}
