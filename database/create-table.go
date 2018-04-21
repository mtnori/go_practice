package main

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/demo")
	db, err := sql.Open("mysql", "root:root@/demo")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	_, err = db.Exec(`
		CREATE TABLE dept (
			dno VARCHAR(20) PRIMARY KEY,
			dname VARCHAR(20),
			budget NUMERIC(10,2),
			lastupdate DATETIME
		)
	`)
	if err != nil {
		panic(err.Error())
	}
}
