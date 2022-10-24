package functions

import "fmt"

func Println(a any) {
	fmt.Printf("%v，%T\n", a, a)
}
