package main

import (
	"fmt"
)

// Vertex 構造体です
type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}
