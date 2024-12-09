package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitlizeDb(Donne string) error {
	var err error
	db, err = sql.Open("sqlite3", Donne)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %v", err)
	}

	// Test the connection
	if err = db.Ping(); err != nil {
		return fmt.Errorf("failed to ping database: %v", err)
	}

	return nil
}

func CloseDb() {
	if db != nil {
		db.Close()
	}
}

func GetDb() *sql.DB {
	return db
}
