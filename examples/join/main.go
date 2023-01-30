package main

import (
	"github.com/yusaint/gostream/arrays"
	"github.com/yusaint/gostream/stream"
	"github.com/yusaint/gostream/stream/functions"
	"github.com/yusaint/gostream/stream/ops"
	"log"
)

type LeftStream struct {
	Id   int
	Name string
}

type RightStream struct {
	Id  int
	Age int
}

type JoinStream struct {
	Id   int
	Name string
	Age  int
}

func main() {
	log.Println("start inner join")
	innerJoin()
	log.Println("start left join")
	leftJoin()
}

var FilterLeftStream = func() ops.Op {
	return ops.Filter(func(t *LeftStream) (bool, error) {
		return t.Id >= 2, nil
	})
}()

func innerJoin() {
	var left = []*LeftStream{{Id: 1, Name: "name1"}, {Id: 2, Name: "name2"}, {Id: 3, Name: "name3"}, {Id: 4, Name: "name4"}}
	var right = []*RightStream{{Id: 1, Age: 18}, {Id: 2, Age: 19}, {Id: 3, Age: 20}}
	var ls = stream.Stream[*LeftStream](arrays.Of(left...))
	var rs = stream.Stream[*RightStream](arrays.Of(right...))

	ls.Peek(ops.Peek(functions.Noop)).Parallel().Filter(FilterLeftStream).Join(rs, ops.InnerJoin(func(l *LeftStream, r *RightStream) bool {
		return l.Id == r.Id
	}, func(l *LeftStream, r *RightStream) *JoinStream {
		return &JoinStream{
			Id:   l.Id,
			Name: l.Name,
			Age:  r.Age,
		}
	})).Foreach(ops.PrintJson)
}

func leftJoin() {
	var left = []*LeftStream{{Id: 1, Name: "name1"}, {Id: 2, Name: "name2"}, {Id: 3, Name: "name3"}, {Id: 4, Name: "name4"}}
	var right = []*RightStream{{Id: 1, Age: 18}, {Id: 2, Age: 19}, {Id: 3, Age: 20}, {Id: 5, Age: 20}}
	var ls = stream.Stream[*LeftStream](arrays.Of(left...))
	var rs = stream.Stream[*RightStream](arrays.Of(right...))

	ls.Peek(ops.Peek(functions.Noop)).Join(rs, ops.LeftJoin(func(l *LeftStream, r *RightStream) bool {
		return l.Id == r.Id
	}, func(l *LeftStream, r *RightStream) *JoinStream {
		return &JoinStream{
			Id:   l.Id,
			Name: l.Name,
			Age:  r.Age,
		}
	}, func(l *LeftStream) *JoinStream {
		return &JoinStream{
			Id:   l.Id,
			Name: l.Name,
			Age:  0,
		}
	})).Foreach(ops.PrintJson)
}
