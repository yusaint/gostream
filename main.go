package main

import (
	"errors"
	"fmt"
	"github.com/yusaint/gostream/arrays"
	"github.com/yusaint/gostream/gostream-contrib/filestream"
	"github.com/yusaint/gostream/streamv/functions"
	"github.com/yusaint/gostream/streamv/ops"
	"os"
)

func main() {
	cc()
}

func bb() {
	stream := streamv.Stream[int](arrays.Of(1, 2, 3))
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

func cc() {
	f, err := os.Open("./demo.csv")
	if err != nil {
		panic(err)
	}
	streamv.Stream[[]string](filestream.NewCsvFileStream(f)).Foreach(ops.Print)
}
