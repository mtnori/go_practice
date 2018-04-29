package main

import (
	"net/http"

	"./controllers"
)

func main() {
	http.HandleFunc("/", controllers.Index)
	http.ListenAndServe(":8080", nil)
}
