package ops

import "github.com/yusaint/gostream/stream/functions"

type ReduceFunction[T, R any] func(R, T) (R, error)

type _reduce[T, R any] struct {
	AbstractOp
	fn    ReduceFunction[T, R]
	state R
}

func (f *_reduce[T, R]) Begin(i int64)     {}
func (f *_reduce[T, R]) End() (any, error) { return f.state, nil }
func (f *_reduce[T, R]) Accept(a any) (err error) {
	f.state, err = f.fn(f.state, a.(T))
	return
}

func NewReduce[T, R any](init R, fn ReduceFunction[T, R]) Op {
	return &_reduce[T, R]{
		fn:    fn,
		state: init,
	}
}

var (
	IntSum = func() Op {
		return NewReduce[int, int](0, func(e1, e2 int) (int, error) {
			return functions.Sum(e1, e2), nil
		})
	}()
)
