package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/demo")
	db, err := sql.Open("mysql", "root:root@/demo")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Insert
	// Resultが戻される
	result, err := db.Exec(`
		INSERT INTO dept(dno, dname, budget) VALUES('D1', 'Marketing', 10000)
  `)
	if err != nil {
		panic(err.Error())
	}

	// AutoIncrementの型で使える
	// 最後に挿入したキーを返す
	id, err := result.LastInsertId()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(id)

	// 影響をあたえた行数を返す
	n, err := result.RowsAffected()
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(n)
}
