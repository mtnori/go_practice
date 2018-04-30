package controllers

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/mtnori/go_practice/mvc/models"
	"github.com/mtnori/go_practice/mvc/repositories"
)

var tmpl = template.Must(template.ParseGlob("views/*.html.tpl"))

func Index(w http.ResponseWriter, r *http.Request) {
	users := repositories.FindAll()
	err := tmpl.ExecuteTemplate(w, "index", users)
	if err != nil {
		panic(err.Error())
	}
}

func Show(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	user := repositories.FindById(id)
	err := tmpl.ExecuteTemplate(w, "show", user)
	if err != nil {
		panic(err.Error())
	}
}

func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "new", nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	user := repositories.FindById(id)
	err := tmpl.ExecuteTemplate(w, "edit", user)
	if err != nil {
		panic(err.Error())
	}
}

func Insert(w http.ResponseWriter, r *http.Request) {
	newUser := models.User{
		Name: r.FormValue("name"),
		City: r.FormValue("city"),
	}
	// 登録処理
	id := repositories.Insert(newUser)
	http.Redirect(w, r, "/show?id="+strconv.Itoa(id), 301)
}

func Update(w http.ResponseWriter, r *http.Request) {
	id, _ := strconv.Atoi(r.FormValue("uid"))
	updateUser := models.User{
		Id:   id,
		Name: r.FormValue("name"),
		City: r.FormValue("city"),
	}
	// 更新処理
	id = repositories.Update(updateUser)
	http.Redirect(w, r, "/show?id="+strconv.Itoa(id), 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	// 削除処理
	repositories.Delete(id)
	http.Redirect(w, r, "/", 301)
}
