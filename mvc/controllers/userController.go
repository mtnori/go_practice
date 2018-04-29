package controllers

import (
	"html/template"
	"net/http"

	"../repositories"
)

var tmpl = template.Must(template.ParseGlob("views/*"))

func Index(w http.ResponseWriter, r *http.Request) {
	users := repositories.FindAll()
	err := tmpl.ExecuteTemplate(w, "index", users)
	if err != nil {
		panic(err)
	}
}
