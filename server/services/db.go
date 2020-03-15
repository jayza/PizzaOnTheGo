package services

import (
	"database/sql"
	"log"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jayza/pizzaonthego/models"

	// Mysql Driver for database connection pool
	_ "github.com/go-sql-driver/mysql"
)

// Database struct
type Database struct {
	DB   *sql.DB
	Mock sqlmock.Sqlmock
}

// Db connection pool
var Db Database

// InitDB initializes the Database connection pool and instantiates the Database struct.
func InitDB(db *sql.DB) Database {
	Db = Database{DB: db}
	return Db
}

// InitMockDB initializes the Database connection pool and instantiates the Database struct.
func InitMockDB(db *sql.DB, mock sqlmock.Sqlmock) Database {
	Db = Database{DB: db, Mock: mock}
	return Db
}

// NewDB creates the Database connection pool.
func NewDB(env models.Env) Database {
	if env.Mock == false {
		db, err := sql.Open("mysql", "root:password@tcp(mariadb)/pizzaonthego")

		if err != nil {
			log.Panic(err)
		}

		return InitDB(db)
	}

	db, mock, err := sqlmock.New()

	if err != nil {
		env.T.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return InitMockDB(db, mock)
}

// Fields is a helper function that spreads out the fields that map to the model.
func (d *Database) Fields(fields ...interface{}) []interface{} {
	return fields
}

// Params is a helper function that spreads out the parameters for the MySQL query.
func (d *Database) Params(args ...interface{}) []interface{} {
	return args
}

// Find queries a row from the database.
// The fields passed are pointing to a model
// Theres no need for preparing statements since go does it under the covers @http://go-database-sql.org/prepared.html
func (d *Database) Find(query string, params []interface{}, fields []interface{}) error {
	return d.DB.QueryRow(query, params...).Scan(fields...)
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
