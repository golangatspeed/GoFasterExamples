package main

import "fmt"

type Database struct {
	ConnString string
}

// NewDatabase is a naive factory to create a database connection
// it returns a pointer to share the address of the connection info
func NewDatabase(server, username, password string) *Database {
	if db != nil {
		return db
	}
	connString := fmt.Sprintf("%s@%s:%s", username, server, password)
	return &Database{connString}
}

var db *Database

func main() {
	db = NewDatabase("localhost", "joeblogs", "password")
	fmt.Printf("Connection String: %+v\n", db)
}