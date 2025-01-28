package models

import (
	"database/sql"
	"free-ent-guide-backend/pkg/cred"
	"time"

	_ "github.com/go-sql-driver/mysql"
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

	// Log all queries.
	// loggerAdapter := zerologadapter.New(zerolog.New(os.Stdout))
	// rawdb = sqldblogger.OpenDriver(c.DB, rawdb.Driver(), loggerAdapter)

	rawstore := RawStore{rawdb}

	return rawstore
}

type StorageContextType string

const (
	SqlcStorageContextKey StorageContextType = "sqlc"
)
