package main

import (
	"fmt"
)

func main() {
	s := []int{2, 3, 5, 7, 11, 13}

	// 下限は0、上限はスライスの長さを指定する
	s = s[1:4]
	fmt.Println(s)

	s = s[:2]
	fmt.Println(s)

	s = s[1:]
	fmt.Println(s)
}
