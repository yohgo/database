package gorm

import (
	"fmt"

	"github.com/jinzhu/gorm"
	// MySQL gorm dialect
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// Database is a container of a configured gorm database.
type Database struct {
	DB *gorm.DB
}

// NewDatabase configures a new gorm database.
func NewDatabase(username, password, host, port, dbName string) (*Database, error) {
	var err error
	var database Database
	// Open database connection
	database.DB, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbName))

	// Check for database open errors
	if err != nil {
		return nil, err
	}

	return &database, nil
}

// Close closes the current connection to a gorm database.
func (factory *Database) Close() error {
	return factory.DB.Close()
}
