package models

import (
	"database/sql"
	"errors"
	"free-ent-guide-backend/pkg/cred"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type Store struct {
	*gorm.DB
}

type RawStore struct {
	*sql.DB
}

func Setup(c cred.Cred) (Store, RawStore) {
	rawdb, err := sql.Open("mysql", c.DB)
	if err != nil {
		panic(err)
	}
	rawstore := RawStore{rawdb}

	db, err := gorm.Open(mysql.Open(c.DB), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	// db.AutoMigrate(
	// 	&Cache{},
	// 	&User{},
	// 	&UserZip{},
	// 	&MLBGame{},
	// 	&MLBTeam{},
	// 	&NHLGame{},
	// 	&NHLTeam{},
	// )
	store := Store{db}
	return store, rawstore
}

type StorageContextType string

const (
	StorageContextKey     StorageContextType = "store"
	SqlcStorageContextKey StorageContextType = "sqlc"
)

// GetCache
func (db Store) GetCache(key string) (string, error) {
	c := Cache{}
	tx := db.First(&c, "name = ? AND expires > ?", key, time.Now())
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
	c := Cache{
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
