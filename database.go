package database

// Database is an interface of a configured application database.
type Database interface {
	NewDatabase(username, password, host, port, dbName string) (*Database, error)
	Close() error
}
