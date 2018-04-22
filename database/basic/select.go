package main

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

type Dept struct {
	DNo        string
	DName      string
	Budget     float64
	LastUpdate mysql.NullTime
}

func main() {
	// db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/demo")
	db, err := sql.Open("mysql", "root:root@/demo?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	var dept Dept

	// Select
	err = db.QueryRow(`
		SELECT
			dno
			,dname
			,budget
			,lastupdate
		FROM
			dept
	`).Scan(&(dept.DNo), &(dept.DName), &(dept.Budget), &(dept.LastUpdate))

	if err != nil {
		panic(err.Error())
	}

	fmt.Println(dept)
	fmt.Println(dept.LastUpdate.Value())
}
