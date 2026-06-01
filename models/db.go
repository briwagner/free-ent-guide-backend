package models

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type StorageContextType string

const (
	SqlcStorageContextKey StorageContextType = "sqlc"
)

type RawStore struct {
	*sql.DB
}

func Setup(sqlString string) RawStore {
	rawdb, err := sql.Open("mysql", sqlString)
	if err != nil {
		panic(err)
	}
	rawdb.SetMaxOpenConns(2)
	rawdb.SetMaxIdleConns(2)
	rawdb.SetConnMaxIdleTime(time.Minute * 60)

	// Log all queries.
	// loggerAdapter := zerologadapter.New(zerolog.New(os.Stdout))
	// rawdb = sqldblogger.OpenDriver(c.DB, rawdb.Driver(), loggerAdapter)
	return RawStore{rawdb}
}
