package main

import (
	"github.com/yusaint/gostream-contrib/mq"
	"github.com/yusaint/gostream/stream"
	"github.com/yusaint/gostream/stream/ops"
	"log"
	"time"
)

func main() {

	chanStream := mq.NewChannelStream[int](10)
	timer := time.NewTicker(1 * time.Second)
	i := 0
	go func() {
		for {
			select {
			case <-timer.C:
				chanStream.Send(i)
				i++
			}
		}
	}()

	if err := stream.Stream[int](chanStream).Parallel().Foreach(ops.Print); err != nil {
		log.Fatal(err)
	}
}
