package repositories

import (
	"../models"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// DBの接続設定
// TODO 設定ファイルにしたい
func dbConn() (db *sql.DB) {
	dbDriver := "mysql"
	dbUser := "root"
	dbPass := "root"
	dbName := "goblog"
	db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func FindAll() []models.User {
	// Select
	db := dbConn()
	rows, err := db.Query("SELECT * FROM users ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	users := []models.User{}
	// 一行ずつ取得
	for rows.Next() {
		var user models.User
		err := rows.Scan(&(user.Id), &(user.Name), &(user.City))
		if err != nil {
			panic(err.Error())
		}
		users = append(users, user)
	}
	// 上のイテレーションでエラーがあれば表示
	if err := rows.Err(); err != nil {
		panic(err.Error())
	}
	// 結果を返す
	return users
}
