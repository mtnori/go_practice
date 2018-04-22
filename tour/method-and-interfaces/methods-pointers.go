package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

// 変数レシーバは元の変数のコピーを操作する
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// ポインタレシーバはレシーバ自身を更新できる
func (v *Vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := Vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
