package set

import "github.com/yusaint/gostream/generic"

type Comparator[T any] interface {
	CompareTo(T) int
}

type Set[T any] interface {
	generic.Splittable
	ToArray() []T
	Add(e T) bool
	Contains(e T) bool
	Size() int
	Remove(e T) bool
	Clear() bool
}
