package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

/*
error型は組み込みのインターフェース
fmtパッケージは、errorインターフェースを確認する
type error interface {
	Error() string
}
*/

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	// エラーstructのポインタを返す
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}
