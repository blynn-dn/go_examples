package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

var (
	id   int
	name string
)

type testTable struct {
	id   int
	name string
}

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

	multipleRows(db)

	data := multipleRows2(db)
	for _, row := range data {
		fmt.Printf("row: %d, %s\n", row.id, row.name)
	}

	id, err := addItem(db, testTable{3, "item 3"})
	if err != nil {
		log.Fatalf("error adding item: %v", err)
	}
	fmt.Printf("added item id: %d", id)
}

func multipleRows(db *sql.DB) {
	rows, err := db.Query("select id, name from test_table where id > $1", 0)
	if err != nil {
		log.Fatalf("query error %s", err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatalf("scan error %s", err)
		}
		log.Println("result: ", id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatalf("error %s", err)
	}
}

func multipleRows2(db *sql.DB) []testTable {
	var data []testTable

	rows, err := db.Query("select id, name from test_table where id > $1", 0)
	if err != nil {
		log.Fatalf("query error %s", err)
	}
	defer rows.Close()
	for rows.Next() {
		var r testTable
		err := rows.Scan(&r.id, &r.name)
		if err != nil {
			log.Fatalf("scan error %s\n", err)
		}
		data = append(data, r)
	}
	err = rows.Err()
	if err != nil {
		log.Fatalf("error %s\n", err)
	}

	fmt.Printf("data: %v\n", data)

	return data
}

func addItem(db *sql.DB, item testTable) (int64, error) {
	result, err := db.Exec("insert into test_table(id, name) values($1, $2)", item.id, item.name)
	if err != nil {
		return 0, fmt.Errorf("addItem: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addItem: %v", err)
	}

	return id, nil
}
