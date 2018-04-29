package controllers

import (
	"html/template"
	"net/http"

	"github.com/mtnori/go_practice/mvc/repositories"
)

var tmpl = template.Must(template.ParseGlob("views/*.html.tpl"))

func Index(w http.ResponseWriter, r *http.Request) {
	users := repositories.FindAll()
	err := tmpl.ExecuteTemplate(w, "index", users)
	if err != nil {
		panic(err)
	}
}
