package functions

import "fmt"

func Println[T any](a T) {
	fmt.Printf("%v，%T\n", a, a)
}

var Print = func(e any) error { Println(e); return nil }
