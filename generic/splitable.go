package generic

type Splittable[T any] interface {
	EstimatedSize() int64
	ForeachRemaining(sink Consumer) error
	TryAdvance(sink Consumer) (bool, error)
}
