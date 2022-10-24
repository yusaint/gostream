package ops

type MapFunction[T, R any] func(T) (R, error)

type _map[T, R any] struct {
	AbstractOp
	fn MapFunction[T, R]
}

func (f *_map[T, R]) Accept(a any) error {
	if r, err := f.fn(a.(T)); err != nil {
		return err
	} else {
		return f.downstream.Accept(r)
	}
}

func Map[T, R any](fn MapFunction[T, R]) Op {
	return &_map[T, R]{
		fn: fn,
	}
}
