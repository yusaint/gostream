package main

import (
	"errors"
	"github.com/yusaint/gostream/arrays"
	"github.com/yusaint/gostream/streamv3"
	"github.com/yusaint/gostream/streamv3/functions"
	"github.com/yusaint/gostream/streamv3/ops"
	"testing"
)

func Benchmark_Stream(b *testing.B) {
	for n := 0; n < b.N; n++ {
		stream := streamv3.Stream[int](arrays.Of(1, 2, 3))
		_, err := stream.
			Filter(ops.Filter(func(t int) (bool, error) { return t > -1, errors.New("!!!!") })).
			Distinct().
			Skip(0).Limit(30).
			Sort(ops.NewSort(functions.IntLte)).
			Reduce(ops.IntSum)
		if err != nil {
			b.Error(err.Error())
		}
	}

}
