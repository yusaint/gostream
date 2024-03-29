package stream

import (
	"github.com/yusaint/gostream/stream/join"
	"github.com/yusaint/gostream/stream/ops"
)

type Streams interface {
	join.Sinkable
	Parallel(...ops.ParallelOption) Streams
	Filter(op ops.Op) Streams
	Reduce(op ops.Op) (any, error)
	Map(op ops.Op) Streams
	Peek(op ops.Op) Streams
	Foreach(op ops.Op) error
	Skip(int64) Streams
	Limit(int64) Streams
	Distinct(...ops.Op) Streams
	Sort(op ops.Op) Streams
	Window(op ops.Op) Streams
	Collect(op ops.Op)
	Group(op ops.Op) Streams
	Join(s Streams, op ops.JoinOp) Streams
}
