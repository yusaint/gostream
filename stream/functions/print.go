package functions

import (
	"encoding/json"
	"fmt"
)

func Println[T any](a T) {
	fmt.Printf("%v，%T\n", a, a)
}

var Print = func(e any) error { Println(e); return nil }
var Noop = func(e any) error { return nil }
var PrintJson = func(e any) error { v, _ := json.Marshal(e); fmt.Println(string(v)); return nil }
