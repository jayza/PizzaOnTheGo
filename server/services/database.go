package services

import "database/sql"

var db *sql.DB
var err error

// Db ...
func Db() *sql.DB {
	return db
}

// Connect ...
func Connect() *sql.DB {
	db, err = sql.Open("mysql", "root:password@tcp(mariadb)/pizzaonthego")

	if err != nil {
		panic(err.Error())
	}

	return db
}

// Close ...
func Close() {
	defer db.Close()
}
