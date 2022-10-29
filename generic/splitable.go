package generic

type Splittable interface {
	EstimatedSize() int64
	ForeachRemaining(sink Consumer) error
	TryAdvance(sink Consumer) (bool, error)
}
