package set

import (
	"sync"
)

type HashSet[T comparable] struct {
	m    map[T]struct{}
	lock sync.Mutex
}

func (a *HashSet[T]) New() Set[T] {
	a.m = make(map[T]struct{})
	return a
}

func (a *HashSet[T]) ToArray() []T {
	arr := make([]T, 0, len(a.m))
	for e, _ := range a.m {
		arr = append(arr, e)
	}
	return arr
}

func (a *HashSet[T]) Add(e T) bool {
	a.lock.Lock()
	defer a.lock.Unlock()

	if _, isOK := a.m[e]; !isOK {
		a.m[e] = struct{}{}
		return true
	}
	return false
}

func (a *HashSet[T]) Contains(e T) bool {
	_, isOK := a.m[e]
	return isOK
}

func (a *HashSet[T]) Size() int {
	return len(a.m)
}

func (a *HashSet[T]) Remove(e T) bool {
	a.lock.Lock()
	defer a.lock.Unlock()
	if _, isOK := a.m[e]; isOK {
		delete(a.m, e)
		return true
	}
	return false
}

func (a *HashSet[T]) Clear() bool {
	a.m = make(map[T]struct{})
	return true
}
