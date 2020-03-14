package services

import (
	"database/sql"
	"log"

	// Mysql Driver for database connection pool
	_ "github.com/go-sql-driver/mysql"
	"github.com/jayza/pizzaonthego/helpers"
)

// Database struct
type Database struct {
	DB *sql.DB
}

// Db connection pool
var Db Database

// InitDB initializes the Database connection pool and instantiates the Database struct.
func InitDB(dataSourceName string) {
	var err error

	Db.DB, err = sql.Open("mysql", dataSourceName)

	if err != nil {
		log.Panic(err)
	}

	// Add retries
	// if err = Db.Ping(); err != nil {
	// 	log.Panic(err)
	// }
}

// Fields is a helper function that spreads out the fields that map to the model.
func (d *Database) Fields(fields ...interface{}) []interface{} {
	return fields
}

// Params is a helper function that spreads out the parameters for the MySQL query.
func (d *Database) Params(args ...interface{}) []interface{} {
	return args
}

// Row is used to determine errors for every row in the query result.
func (d *Database) Row(e error) (err error) {
	switch err {
	case nil:
		return nil
	case sql.ErrNoRows:
		return helpers.NewHTTPError(err, 404, "Not Found")
	default:
		return helpers.NewHTTPError(err, 500, "Internal Server Error")
	}
}

// Find queries a row from the database.
func (d *Database) Find(query string, params []interface{}, fields []interface{}) error {
	stmt, err := d.DB.Prepare(query)

	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	return d.Row(stmt.QueryRow(params...).Scan(fields...))
}

// All ...
// func (d *Database) All(query string, params []interface{}, fields []interface{}) error {
// 	res, err := d.DB.Query(query, params...)

// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	defer res.Close()

// 	for res.Next() {
// 		d.Row(res.Scan(fields...))
// 	}
// }
