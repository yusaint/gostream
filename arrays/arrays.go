package arrays

import (
	"github.com/yusaint/gostream/generic"
)

type Array[T generic.ElemType] struct {
	arr   []T
	index int
}

func (a *Array[T]) EstimatedSize() int64 {
	return int64(len(a.arr))
}

func (a *Array[T]) ForeachRemaining(sink generic.Consumer) error {
	for {
		if a.EstimatedSize() == 0 {
			return nil
		}
		isContinue, err := a.TryAdvance(sink)
		if err != nil {
			return err
		} else {
			if !isContinue {
				return nil
			}
		}
	}
}

func (a *Array[T]) TryAdvance(sink generic.Consumer) (bool, error) {
	if err := sink.Accept(a.arr[a.index]); err != nil {
		return false, err
	}
	a.index++
	return a.index < len(a.arr), nil
}

func Of[T generic.ElemType](e ...T) *Array[T] {
	array := &Array[T]{
		arr: make([]T, 0, 10),
	}
	for _, v := range e {
		array.arr = append(array.arr, v)
	}
	return array
}
