package main

import "fmt"

var test = 1

func main() {
	defer fmt.Println("world")
	fmt.Println("hello")
}
