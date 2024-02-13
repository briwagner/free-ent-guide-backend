package models

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"free-ent-guide-backend/models/modelstore"
	"time"
)

type Cache struct {
	ID        int64
	Name      string
	Value     string
	UpdatedAt time.Time
	Expires   time.Time
}

const expiresDur = time.Hour * 1

func (c *Cache) Insert(db *modelstore.Queries) error {
	t := time.Now()
	vals := modelstore.SetCacheParams{
		Name:      sql.NullString{String: c.Name, Valid: true},
		Value:     sql.NullString{String: c.Value, Valid: true},
		UpdatedAt: sql.NullTime{Time: t, Valid: true},
	}
	// Set default expires, unless set on demand.
	if c.Expires.IsZero() {
		vals.Expires = sql.NullTime{Time: t.Add(expiresDur), Valid: true}
	} else {
		vals.Expires = sql.NullTime{Time: c.Expires, Valid: true}
	}

	id, err := db.SetCache(context.Background(), vals)
	if err != nil {
		return err
	}
	c.ID = id
	return nil
}

func (c *Cache) GetByID(id int64, db *modelstore.Queries) error {
	ret, err := db.GetCacheByID(context.Background(), id)
	if err != nil {
		return err
	}
	c.FromDB(ret)
	return nil
}

func (c *Cache) GetByName(name string, db *modelstore.Queries) error {
	if name == "" {
		return errors.New("no cache name passed")
	}
	ret, err := db.GetCacheByName(context.Background(), sql.NullString{String: name, Valid: true})
	if err != nil {
		return err
	}
	c.FromDB(ret)
	return nil
}

// FromDB converts a sqlc model to our Cache model.
func (c *Cache) FromDB(ms modelstore.Cach) {
	c.ID = ms.ID
	c.Name = ms.Name.String
	c.Value = ms.Value.String
	c.UpdatedAt = ms.UpdatedAt.Time
	c.Expires = ms.Expires.Time
}

func (c Cache) String() string {
	return fmt.Sprintf("Cache %d: %s, expires %s", c.ID, c.Name, c.Expires.Format("06/01/02 15:04"))
}

type Caches []Cache

// ShowCache returns all cached values.
func (cs *Caches) ShowAll(db *modelstore.Queries) error {
	ret, err := db.GetCacheAll(context.Background())
	if err != nil {
		return err
	}

	for _, val := range ret {
		var c Cache
		c.FromDB(val)
		*cs = append(*cs, c)
	}
	return nil
}

// ShowStale returns all expired cached values.
func (cs *Caches) ShowStale(t time.Time, db *modelstore.Queries) error {
	if t.IsZero() {
		return errors.New("invalid time: nil")
	}
	ret, err := db.GetCacheStale(context.Background(), sql.NullTime{Time: t, Valid: true})
	if err != nil {
		return err
	}
	for _, val := range ret {
		var c Cache
		c.FromDB(val)
		*cs = append(*cs, c)
	}
	return nil
}

// DropStale removes all cached values with expired dates.
func (cs *Caches) DropStale(t time.Time, db *modelstore.Queries) (int64, error) {
	if t.IsZero() {
		return 0, errors.New("invalid time: nil")
	}
	return db.DropCacheStale(context.Background(), sql.NullTime{Time: t, Valid: true})
}

// DropAll removes all cached values.
func (cs *Caches) DropAll(db *modelstore.Queries) (int64, error) {
	return db.WipeCache(context.Background())
}
