// +build ignore
package main

import (
	"fmt"
)

// Pi は数値です
const Pi float = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("Happy", Pi, "Day")

	const Truth = true
	fmt.Println("Go rules?", Truth)
}
