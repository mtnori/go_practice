package main

import (
	"database/sql"
	"fmt"
	"math/rand"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/demo")
	db, err := sql.Open("mysql", "root:root@/demo?parseTime=true")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	// Create randum seed
	rand.Seed(time.Now().UnixNano())

	// Start transaction
	tx, err := db.Begin()
	if err != nil {
		panic(err.Error())
	}

	// Rollback setting
	defer func() {
		if err := recover(); err != nil {
			if err := tx.Rollback(); err != nil {
				panic(err.Error())
			}
			fmt.Println("Rollbacked")
		}
	}()

	// Insert
	_, err = db.Exec(`
		INSERT INTO dept(dno, dname, budget, lastupdate) VALUES ('D1', 'Marketing', 10000, now())
	`)
	if err != nil {
		panic(err.Error())
	}

	// Panic
	if n := rand.Intn(10); n < 5 {
		panic("opps!!")
	}

	// Commit
	if err := tx.Commit(); err != nil {
		panic(err.Error())
	}
	fmt.Println("Commited")
}
