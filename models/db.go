package models

import (
	"database/sql"
	"free-ent-guide-backend/pkg/cred"
)

type RawStore struct {
	*sql.DB
}

func Setup(c cred.Cred) RawStore {
	rawdb, err := sql.Open("mysql", c.DB)
	if err != nil {
		panic(err)
	}
	rawstore := RawStore{rawdb}

	return rawstore
}

type StorageContextType string

const (
	// StorageContextKey     StorageContextType = "store" // deprecated GORM
	SqlcStorageContextKey StorageContextType = "sqlc"
)
