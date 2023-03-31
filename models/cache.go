package models

import (
	"fmt"
	"time"

	"gorm.io/gorm"
)

type Cache struct {
	gorm.Model
	Name    string `gorm:"unique"`
	Value   string `gorm:"type:mediumblob"`
	Expires time.Time
}

func (c Cache) String() string {
	return fmt.Sprintf("Cache %d: %s, expires %s", c.ID, c.Name, c.Expires.Format("06/01/02 15:04"))
}

type Caches []Cache

// ShowCache returns all cached values.
func (cs *Caches) ShowAll(db Store) error {
	tx := db.Find(cs)
	return tx.Error
}

// ShowStale returns all expired cached values.
func (cs *Caches) ShowStale(t time.Time, db Store) error {
	tx := db.Where("expires < ?", t).Find(cs)
	return tx.Error
}

// DropStale removes all cached values with expired dates that have passed.
func (cs *Caches) DropStale(db Store) error {
	tx := db.Unscoped().Delete(cs)
	return tx.Error
}

// DropAll removes all cached values.
func (cs *Caches) DropAll(db Store) error {
	tx := db.Unscoped().Where("ID >= ?", 1).Delete(cs)
	return tx.Error
}
