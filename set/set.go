package set

import "github.com/yusaint/gostream/generic"

type Set[T any] interface {
	generic.Splittable
	generic.Collect[T]
	ToArray() []T
	Contains(e T) bool
	Size() int
	Remove(e T) bool
	Clear() bool
}
