package ops

type FilterFunction[T any] func(T) (bool, error)

type filter[T any] struct {
	AbstractOp
	fn FilterFunction[T]
}

func (f *filter[T]) Accept(a any) error {
	accept, err := f.fn(a.(T))
	if err != nil {
		return err
	}
	if accept {
		return f.downstream.Accept(a)
	}
	return nil
}

func Filter[T any](fn FilterFunction[T]) Op {
	return &filter[T]{
		fn: fn,
	}
}
