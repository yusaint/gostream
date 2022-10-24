package generic

type Comparable[T any] interface {
	CompareTo(T) bool
}

type ElemType interface {
	any
}
