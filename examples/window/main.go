package main

import (
	"fmt"
	"github.com/yusaint/gostream-contrib/filestream"
	"github.com/yusaint/gostream/stream"
	"github.com/yusaint/gostream/stream/ops"
	"log"
	"strings"
)

type Item struct {
	ItemType string
	ItemID   string
}

func main() {
	fileUrl := "xx"
	csvStreamFile, err := filestream.NewCsvFileStreamByUrl(fileUrl)
	if err != nil {
		log.Printf(err.Error())
		return
	}
	err = stream.Stream[[]string](csvStreamFile).
		Skip(1).
		Filter(ops.Filter(func(records []string) (bool, error) {
			if len(records) == 0 || len(strings.TrimSpace(records[0])) == 0 {
				return false, nil
			}
			return true, nil
		})).
		Map(ops.Map(func(records []string) (*Item, error) {
			return &Item{
				ItemType: "1",
				ItemID:   strings.TrimSpace(records[0]),
			}, nil
		})).
		Window(ops.UseFlipWindow[*Item](10)).
		Parallel().
		Map(ops.Map(func(items []*Item) ([]*Item, error) {
			return items, nil
		})).
		Filter(ops.Filter(func(items []*Item) (bool, error) {
			return len(items) > 0, nil
		})).
		Foreach(ops.Foreach(func(items []*Item) error {
			panic("error occurs")
			fmt.Printf("successfully consume message %v", items)
			return nil
		}))
	fmt.Println("result:", err.Error())
}
