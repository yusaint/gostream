package ops

import (
	"github.com/yusaint/gostream/generic"
)

type head[T any] struct {
	AbstractOp
	spl generic.Splittable[T]
}

func (f *head[T]) Handle() (any, error) {
	f.Begin(f.spl.EstimatedSize())
	if err := f.spl.ForeachRemaining(f); err != nil {
		return nil, err
	}
	return f.End()
}

func NewHead[T any](spl generic.Splittable[T]) Op {
	return &head[T]{
		spl: spl,
	}
}
