package main

import (
	"database/sql"
	"log"

	// because stackoverflow told me to
	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

// InitDB handles the initial DB connection
func InitDB(dataSourceName string) {
	var err error
	db, err = sql.Open("mysql", dataSourceName)

	if err != nil {
		log.Panic(err)
	}

	if err = db.Ping(); err != nil {
		log.Panic(err)
	}
}
