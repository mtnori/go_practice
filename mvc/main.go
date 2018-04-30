package main

import (
	"html/template"
	"net/http"

	"github.com/gorilla/sessions"
)

var store = sessions.NewCookieStore([]byte("something-very-secret"))
var tmpl = template.Must(template.ParseGlob("views/*.html.tpl"))

func main() {
	// ルーティング設定
	http.HandleFunc("/", Index)
	http.HandleFunc("/show", Show)
	http.HandleFunc("/new", New)
	http.HandleFunc("/edit", Edit)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/delete", Delete)
	http.HandleFunc("/login", Login)
	http.HandleFunc("/logout", Logout)
	// サーバ起動
	http.ListenAndServe(":8080", nil)
}
