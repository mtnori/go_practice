// インターフェースは暗黙のうちに満足する
package main

import "fmt"

type I interface {
	M()
}

type T struct {
	S string
}

// このメソッドは構造体TがインターフェースIを実装している
// だが、それを明示的に宣言する必要はない
func (t T) M() {
	fmt.Println(t.S)
}

func main() {
	var i I = T{"hello"}
	i.M()
}
