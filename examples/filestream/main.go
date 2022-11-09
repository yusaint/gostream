package main

import (
	"github.com/yusaint/gostream/gostream-contrib/filestream"
	"github.com/yusaint/gostream/stream"
	"github.com/yusaint/gostream/stream/ops"
	"os"
)

func main() {
	f, err := os.Open("./demo.csv")
	if err != nil {
		panic(err)
	}
	stream.Stream[[]string](filestream.NewCsvFileStream(f)).
		Map(ops.Map(func(row []string) (*Record, error) {
			return &Record{
				R1: row[0],
				R2: row[1],
				R3: row[2],
			}, nil
		})).Filter(ops.Filter(func(r *Record) (bool, error) {
		return r.R1 == "3", nil
	})).Foreach(ops.Print)
}
