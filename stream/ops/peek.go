package ops

type _peek[T any] struct {
	AbstractOp
	fn    IteratorFunc[T]
	dummy T
}

func (f *_peek[T]) Accept(a any) error { f.fn(a.(T)); return f.downstream.Accept(a) }

func Peek[T any](fn IteratorFunc[T]) Op {
	return &_peek[T]{
		fn: fn,
	}
}
