package models

import (
	"database/sql"
	"free-ent-guide-backend/pkg/cred"
	"time"
)

type RawStore struct {
	*sql.DB
}

func Setup(c cred.Cred) RawStore {
	rawdb, err := sql.Open("mysql", c.DB)
	if err != nil {
		panic(err)
	}
	rawdb.SetMaxOpenConns(2)
	rawdb.SetMaxIdleConns(2)
	rawdb.SetConnMaxIdleTime(time.Minute * 60)

	rawstore := RawStore{rawdb}

	return rawstore
}

type StorageContextType string

const (
	// StorageContextKey     StorageContextType = "store" // deprecated GORM
	SqlcStorageContextKey StorageContextType = "sqlc"
)
