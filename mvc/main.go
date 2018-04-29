package main

import (
	"net/http"

	"github.com/mtnori/go_practice/mvc/controllers"
)

func main() {
	http.HandleFunc("/", controllers.Index)
	http.ListenAndServe(":8080", nil)
}
