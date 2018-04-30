package repositories

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/mtnori/go_practice/mvc/models"
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
	rows, err := db.Query(`SELECT * FROM users ORDER BY id DESC`)
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

func FindById(id string) models.User {
	db := dbConn()

	var user models.User
	err := db.QueryRow(`SELECT * FROM users WHERE id = ?`, id).Scan(&(user.Id), &(user.Name), &(user.City))
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()

	return user
}

func Insert(user models.User) int {
	db := dbConn()

	stmt, err := db.Prepare("INSERT INTO users(name, city) VALUES (?, ?)")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	defer db.Close()

	result, err := stmt.Exec(user.Name, user.City)
	if err != nil {
		panic(err.Error())
	}
	log.Println("INSERT: Name: " + user.Name + " | City: " + user.City)

	// AutoIncrementの型で使える
	// 最後に挿入したキーを返す
	id, err := result.LastInsertId()
	if err != nil {
		panic(err.Error())
	}

	return int(id)
}

func Update(user models.User) int {
	db := dbConn()

	stmt, err := db.Prepare("UPDATE users SET name = ?, city = ? WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	defer db.Close()

	_, err = stmt.Exec(user.Name, user.City, user.Id)
	if err != nil {
		panic(err.Error())
	}
	log.Println("UPDATE: Name: " + user.Name + " | City: " + user.City)

	return user.Id
}

func Delete(id string) {
	db := dbConn()

	stmt, err := db.Prepare("DELETE FROM users WHERE id = ?")
	if err != nil {
		panic(err.Error())
	}
	defer stmt.Close()
	defer db.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		panic(err.Error())
	}
	log.Println("DELETE")
}
