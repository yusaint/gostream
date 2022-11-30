package ops

import (
	"github.com/yusaint/gostream/generic"
)

type _collect[T any] struct {
	AbstractOp
	c generic.Collect[T]
}

func (f *_collect[T]) Begin(i int64)     {}
func (f *_collect[T]) End() (any, error) { return f.c, nil }
func (f *_collect[T]) Accept(a any) (err error) {
	f.c.Add(a.(T))
	return nil
}

func NewCollect[T any](c generic.Collect[T]) Op {
	return &_collect[T]{
		c: c,
	}
}
