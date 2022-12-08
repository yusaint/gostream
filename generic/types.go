package generic

type ElemType interface {
	any
}

type Collect[T any] interface {
	Add(e T) bool
}
