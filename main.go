package main

import (
	"errors"
	"fmt"
	"github.com/yusaint/gostream/arrays"
	"github.com/yusaint/gostream/streamv3"
	"github.com/yusaint/gostream/streamv3/functions"
	"github.com/yusaint/gostream/streamv3/ops"
)

func main() {
	bb()
}

func bb() {
	stream := streamv3.Stream[int](arrays.Of(1, 2, 3))
	sum, err := stream.
		Filter(ops.Filter(func(t int) (bool, error) { return t > -1, errors.New("aaaaa") })).
		Distinct().
		Skip(0).Limit(30).
		Sort(ops.NewSort(functions.IntLte)).
		Reduce(ops.IntSum)
	if err != nil {
		fmt.Println("!!!", err.Error(), sum)
	} else {
		fmt.Println(sum)
	}
}
