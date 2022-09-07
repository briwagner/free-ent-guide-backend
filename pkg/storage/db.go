package storage

import (
	"errors"
	"free-ent-guide-backend/models"
	"free-ent-guide-backend/pkg/cred"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Store struct {
	*gorm.DB
}

func Setup(c cred.Cred) Store {
	db, err := gorm.Open(mysql.Open(c.DB), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.Cache{})
	store := Store{db}
	return store
}

type StorageContextType string

const StorageContextKey StorageContextType = "store"

// GetCache
func (db Store) GetCache(key string) (string, error) {
	c := models.Cache{}
	// todo: modify query to pass expires time?
	tx := db.First(&c, "name = ?", key)
	if tx.Error != nil {
		return "", tx.Error
	}
	if time.Now().After(c.Expires) {
		// Avoid deleting record, as gorm does soft delete,
		// and we force upsert behavior when setting cache.
		return "", errors.New("cache expired")
	}
	return c.Value, nil
}

// SetCache
func (db Store) SetCache(key string, val string, t time.Duration) error {
	c := models.Cache{
		Name:    key,
		Value:   val,
		Expires: time.Now().Add(t),
	}
	// Force upsert behavior.
	tx := db.Clauses(clause.OnConflict{
		UpdateAll: true,
	}).Create(&c)
	if tx.Error != nil {
		return tx.Error
	}
	return nil
}
