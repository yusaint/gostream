package ops

import (
	"golang.org/x/exp/slices"
)

type Comparator[T any] func(T, T) bool

type _sort[T any] struct {
	AbstractOp
	fn  Comparator[T]
	arr []T
}

func (f *_sort[T]) Begin(i int64) {
	capacity := 10
	if i > 0 {
		capacity = int(i)
	}
	f.arr = make([]T, 0, capacity)
	f.downstream.Begin(i)
}
func (f *_sort[T]) End() (any, error) {
	slices.SortFunc(f.arr, f.fn)
	f.downstream.Begin(int64(len(f.arr)))
	for _, a := range f.arr {
		if err := f.downstream.Accept(a); err != nil {
			return nil, err
		}
	}
	return f.downstream.End()
}
func (f *_sort[T]) Accept(a any) error { f.arr = append(f.arr, a.(T)); return nil }

func NewSort[T any](fn func(T, T) bool) Op {
	return &_sort[T]{
		fn: fn,
	}
}
