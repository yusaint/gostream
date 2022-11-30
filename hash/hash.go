package hash

type Hash[T comparable, R any] interface {
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
