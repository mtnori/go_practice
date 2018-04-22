package main

import (
	"fmt"
)

func main() {
	i, j := 42, 2701

	p := &i
	fmt.Println(*p) // ポインタを通してiを取得
	*p = 21
	fmt.Println(i)

	p = &j
	*p = *p / 37 // ポインタを通してjを計算
	fmt.Println(j)
}
