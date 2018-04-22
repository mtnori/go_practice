package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/demo")
	db, err := sql.Open("mysql", "root:root@/demo?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Select
	rows, err := db.Query(`
		SELECT
			dno
			,dname
			,budget
			,lastupdate
		FROM
			dept
	`)
	if err != nil {
		panic(err.Error())
	}
	defer rows.Close()

	// Get columns
	cols, err := rows.Columns()
	if err != nil {
		panic(err.Error())
	}
	// Show column names
	for _, name := range cols {
		fmt.Println(name)
	}
}
