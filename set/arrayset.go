package set

import (
	"sync"
)

type IntArraySet struct {
	ArraySet[int]
}

type ArraySet[T any] struct {
	array []T
	lock  sync.Mutex
}

func (a *ArraySet[T]) New() Set[T] {
	a.array = make([]T, 0, 10)
	return a
}

func (a *ArraySet[T]) ToArray() []T {
	return a.array
}

func (a *ArraySet[T]) indexOf(e any) int {
	for i := 0; i < len(a.array); i++ {

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
	return a.indexOf(e) >= 0
}

func (a *ArraySet[T]) Size() int {
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
	a.array = []T{}
	return true
}
