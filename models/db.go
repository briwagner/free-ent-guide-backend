package models

import (
	"database/sql"
	"free-ent-guide-backend/pkg/cred"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
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
