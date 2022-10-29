package ops

import "github.com/yusaint/gostream/stream/functions"

type IteratorFunc[T any] func(T) error

type foreach[T any] struct {
	AbstractOp
	fn    IteratorFunc[T]
	dummy T
}

func (f *foreach[T]) Begin(i int64)      {}
func (f *foreach[T]) End() (any, error)  { return f.dummy, nil }
func (f *foreach[T]) Link(next Op)       {}
func (f *foreach[T]) Accept(a any) error { f.dummy = a.(T); return f.fn(a.(T)) }

func Foreach[T any](fn IteratorFunc[T]) Op {
	return &foreach[T]{
		fn: fn,
	}
}

var Print = Foreach(functions.Print)
