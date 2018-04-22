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

	// Create prepared statement
	stmt, err := db.Prepare("INSERT INTO dept(dno, dname, budget, lastupdate) VALUES (?, ?, ?, now())")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()

	// Insert with parameters
	result, err := stmt.Exec("D1", "Marketing", 10000)
	if err != nil {
		panic(err.Error())
	}

	// Return affected rows count
	n, err := result.RowsAffected()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(n)
}
