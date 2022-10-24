package streamv3

import (
	"github.com/yusaint/gostream/streamv3/ops"
)

type Streams interface {
	Filter(op ops.Op) Streams
	Reduce(op ops.Op) (any, error)
	Map(op ops.Op) Streams
	Foreach(op ops.Op) error
	Skip(int64) Streams
	Limit(int64) Streams
	Distinct(...ops.Op) Streams
	Sort(op ops.Op) Streams
}
