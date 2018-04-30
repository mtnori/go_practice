package main

import (
	"net/http"

	"github.com/mtnori/go_practice/mvc/controllers"
)

func main() {
	// ルーティング設定
	http.HandleFunc("/", controllers.Index)
	http.HandleFunc("/show", controllers.Show)
	http.HandleFunc("/new", controllers.New)
	http.HandleFunc("/edit", controllers.Edit)
	http.HandleFunc("/insert", controllers.Insert)
	http.HandleFunc("/update", controllers.Update)
	http.HandleFunc("/delete", controllers.Delete)
	// サーバ起動
	http.ListenAndServe(":8080", nil)
}
