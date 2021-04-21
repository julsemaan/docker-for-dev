package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	db, err := sql.Open("postgres", os.Getenv("PG_CONNECTION_STRING"))
	if err != nil {
		panic(err)
	}
	defer db.Close()

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("DROP TABLE testing")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS testing(ID SERIAL PRIMARY KEY, NAME TEXT)")
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`INSERT INTO TESTING (NAME) VALUES('banana')`)
	if err != nil {
		panic(err)
	}

	rows, err := db.Query("select NAME from testing")
	if err != nil {
		panic(err)
	}

	if !rows.Next() {
		panic("No record to read")
	}

	var s string
	err = rows.Scan(&s)
	if err != nil {
		panic(err)
	}

	fmt.Println("helloooo", s)
}
