package mysql

import (
	"database/sql"
	"fmt"

	// MySQL go sql driver
	_ "github.com/go-sql-driver/mysql"
)

// Database is a container of a configured mysql database.
type Database struct {
	DB *sql.DB
}

// NewDatabase configures a new mysql database.
func NewDatabase(username, password, host, port, dbName string) (*Database, error) {
	var err error
	var database Database
	// Open database connection
	database.DB, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbName))

	// Check for database open errors
	if err != nil {
		return nil, err
	}

	// Check to verify database connection
	if err = database.DB.Ping(); err != nil {
		return nil, err
	}

	return &database, nil
}

// Close closes the current connection to a mysql database.
func (factory *Database) Close() error {
	return factory.DB.Close()
}
