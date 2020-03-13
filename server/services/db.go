package services

import (
	"database/sql"
	"log"

	// Mysql Driver for database connection pool
	_ "github.com/go-sql-driver/mysql"
)

// Db connection pool
var Db *sql.DB
var err error

// InitDB ...
func InitDB(dataSourceName string) {
	Db, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		log.Panic(err)
	}

	// Add retries
	// if err = Db.Ping(); err != nil {
	// 	log.Panic(err)
	// }
}
