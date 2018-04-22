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

	// 1行ずつ取得
	for rows.Next() {
		var dept Dept
		err := rows.Scan(&(dept.DNo), &(dept.DName), &(dept.Budget), &(dept.LastUpdate))
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(dept)
	}

	// 上のイテレーションでエラーがあれば表示
	if err := rows.Err(); err != nil {
		panic(err.Error())
	}
}
