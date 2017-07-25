package database

import (
	"fmt"

	"github.com/jinzhu/gorm"
	// Uses the mysql gorm dialect
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	// DBCon is the connection handle for the database
	DBCon *gorm.DB
)

// ConnectDB enables the connection to the database and returns
// the error if failed to establish a connection
func ConnectDB(username, password, host, port, dbname string) {
	var err error
	DBCon, err = gorm.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, dbname))

	if err != nil {
		panic("Failed to connect to database")
	}
}
