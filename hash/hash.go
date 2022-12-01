package hash

import "github.com/yusaint/gostream/generic"

type Hash[T comparable, R any] interface {
	generic.Splittable
	ToArray() []R
	Add(T, R) bool
	Contains(e T) bool
	Size() int
	Remove(e T) bool
	Clear() bool
	Front() R
	Next() bool
	Current() R
}
