# go_postgres
go postgress example

## Setup postgres DB
<details>
  <summary>Expand to see details of installing postgres docker container</summary>

Install `postgres` as a docker container:
```
docker pull postgres
```

Configure the db on the postgres container:
```
docker run -itd -e POSTGRES_USER=test -e POSTGRES_PASSWORD=test -p 5432:5432 -v /data:/var/lib/postgresql/data --name postgresql postgres
```

Install `libpq` is not already installed. The following is an example (your results may vary):
```
brew doctor
brew update
brew install libpq

brew link --force libpq
```

Verfify access to DB with `psql`
```
PGPASSWORD=baeldung psql -U test -c '\l' -h localhost
Password for user test:
                                                   List of databases
   Name    | Owner | Encoding | Locale Provider |  Collate   |   Ctype    | ICU Locale | ICU Rules | Access privileges
-----------+-------+----------+-----------------+------------+------------+------------+-----------+-------------------
 postgres  | test  | UTF8     | libc            | en_US.utf8 | en_US.utf8 |            |           |
 template0 | test  | UTF8     | libc            | en_US.utf8 | en_US.utf8 |            |           | =c/test          +
           |       |          |                 |            |            |            |           | test=CTc/test
 template1 | test  | UTF8     | libc            | en_US.utf8 | en_US.utf8 |            |           | =c/test          +
           |       |          |                 |            |            |            |           | test=CTc/test
 test      | test  | UTF8     | libc            | en_US.utf8 | en_US.utf8 |            |           |
(4 rows)
```

```
createdb -h localhost -U test testdb
```

```
PGPASSWORD=test psql -U test  -h localhost testdb
psql (16.0)
Type "help" for help.

testdb=# create table test_table(id int primary key not null, name text not null);
CREATE TABLE
```

</details>

## Initial go steps
* create main.go -- this will be a simple go file with a main 
  ```
  package main
    
  import "fmt"
    
  func main() {
     fmt.Printf("test\n")
  }
  ```
* run mod init 
  ```shell
  gp mod init
  ```
* Install package `pq`:
  ```shell
  go get github.com/lib/pq
  go: downloading github.com/lib/pq v1.10.9
  go: added github.com/lib/pq v1.10.9
  ```

## Simple query example
```go
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
```