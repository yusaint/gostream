package set

type Comparator[T any] interface {
	CompareTo(T) int
}

type Set[T any] interface {
	ToArray() []T
	Add(e T) bool
	Contains(e T) bool
	Size() int
	Remove(e T) bool
	Clear() bool
}
