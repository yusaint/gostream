package set

import (
	"github.com/yusaint/gostream/generic"
	"reflect"
	"sync"
)

type IntArraySet = ArraySet[int]
type StringArraySet = ArraySet[string]

type ArraySet[T any] struct {
	array []T
	lock  sync.RWMutex
	index int
}

func (a *ArraySet[T]) EstimatedSize() int64 {
	return int64(len(a.array))
}

func (a *ArraySet[T]) ForeachRemaining(sink generic.Consumer) error {
	for {
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

func (a *ArraySet[T]) TryAdvance(sink generic.Consumer) (bool, error) {
	if err := sink.Accept(a.array[a.index]); err != nil {
		return false, err
	}
	a.index++
	return a.index < len(a.array), nil
}

func NewArraySet[T any]() Set[T] {
	arr := &ArraySet[T]{
		array: make([]T, 0),
	}
	return arr
}

func (a *ArraySet[T]) ToArray() []T {
	return a.array
}

func (a *ArraySet[T]) indexOf(e any) int {
	for i := 0; i < len(a.array); i++ {
		if reflect.DeepEqual(e, a.array[i]) {
			return i
		}
	}
	return -1
}

func (a *ArraySet[T]) Add(e T) bool {
	a.lock.Lock()
	defer a.lock.Unlock()
	if a.indexOf(e) == -1 {
		a.array = append(a.array, e)
		return true
	}
	return false
}

func (a *ArraySet[T]) Contains(e T) bool {
	a.lock.RLock()
	defer a.lock.RUnlock()
	return a.indexOf(e) >= 0
}

func (a *ArraySet[T]) Size() int {
	a.lock.RLock()
	defer a.lock.RUnlock()
	return len(a.array)
}

func (a *ArraySet[T]) Remove(e T) bool {
	a.lock.Lock()
	defer a.lock.Unlock()

	index := a.indexOf(e)
	if index == -1 {
		return false
	}
	numMoved := len(a.array) - index - 1
	if numMoved == 0 {
		a.array = a.array[0 : len(a.array)-1]
	} else if numMoved == len(a.array)-1 {
		a.array = a.array[1:]
	} else {
		newArray := make([]T, len(a.array)-1)
		copy(newArray[:index], a.array[:index])
		copy(newArray[index:], a.array[index+1:])
		a.array = newArray
	}
	return true
}

func (a *ArraySet[T]) Clear() bool {
	a.lock.Lock()
	defer a.lock.Unlock()
	a.array = []T{}
	return true
}
