package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	connStr := "postgres://test:test@localhost:5432/testdb?sslmode=disable"

	// connect to the db
	db, err := sql.Open("postgres", connStr)

	// defer the close so that the db connection is closed if/when the main() exits
	defer db.Close()

	if err != nil {
		log.Fatal(err)
	}

	// try to ping
	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	// simple query
	var s sql.NullString
	err = db.QueryRow("select name from test_table").Scan(&s)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("results: %s\n", s.String)
}
