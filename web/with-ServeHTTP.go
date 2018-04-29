// 既存のstructや型に対してServeHTTPメソッドを用意して
// http.Handleに登録できるようにする
package main

import (
	"fmt"
	"net/http"
)

type String string

func (s String) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, s)
}

/*
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World")
}
*/

func main() {
	//http.HandleFunc("/", handler)
	http.Handle("/", String("Hello World."))
	http.ListenAndServe(":8080", nil)
}
