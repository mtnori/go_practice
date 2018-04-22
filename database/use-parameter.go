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

	// Insert with parameters
	result, err := db.Exec(`
		INSERT INTO dept(dno, dname, budget, lastupdate) VALUES (?, ?, ?, now())
  `, "D1", "Marketing", 10000)
	if err != nil {
		panic(err.Error())
	}

	// 影響をあたえた行数を返す
	n, err := result.RowsAffected()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(n)
}
