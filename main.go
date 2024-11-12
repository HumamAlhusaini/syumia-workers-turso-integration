// main.go
package main

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/tursodatabase/libsql-client-go/libsql"
)

func main() {
	url := "libsql://my-db-humamalhusaini.turso.io?authToken=eyJhbGciOiJFZERTQSIsInR5cCI6IkpXVCJ9.eyJpYXQiOjE3MzEyNzcwODUsImlkIjoiNzMwZDZiZDAtYzlmNS00MjFkLWE5MTQtNGNmYjZkOTlhNjU0In0.zkSvvLcXOAgvRNqJrY-Bi8bZjynZiMF3EHhKSpwIWeNWznhlR17mmSWsKNZLaCw_3wvuYOYqXSBYOJrPd248DQ"

	db, err := sql.Open("libsql", url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed to open db %s: %s", url, err)
		os.Exit(1)
	} else {
		fmt.Println("success to open db")
	}

	fmt.Println("success to open db")

	defer db.Close()
}
