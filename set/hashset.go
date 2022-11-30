package set

import (
	"github.com/yusaint/gostream/generic"
	"github.com/yusaint/gostream/hash"
)

type HashSet[T comparable] struct {
	hmap hash.Hash[T, T]
}

func NewHashSet[T comparable]() Set[T] {
	return &HashSet[T]{
		hmap: hash.NewMap[T, T](),
	}
}

func (h *HashSet[T]) EstimatedSize() int64 {
	return int64(h.Size())
}

func (h *HashSet[T]) ForeachRemaining(sink generic.Consumer) error {
	h.hmap.Front()
	for {
		isContinue, err := h.TryAdvance(sink)
		if err != nil {
			return err
		} else {
			if !isContinue {
				return nil
			}
		}
	}
}

func (h *HashSet[T]) TryAdvance(sink generic.Consumer) (bool, error) {
	if err := sink.Accept(h.hmap.Current()); err != nil {
		return false, err
	}
	if ele := h.hmap.Next(); ele == false {
		return false, nil
	}
	return true, nil
}

func (h *HashSet[T]) ToArray() []T {
	return h.hmap.ToArray()
}

func (h *HashSet[T]) Add(e T) bool {
	return h.hmap.Add(e, e)
}

func (h *HashSet[T]) Contains(e T) bool {
	return h.hmap.Contains(e)
}

func (h *HashSet[T]) Size() int {
	return h.hmap.Size()
}

func (h *HashSet[T]) Remove(e T) bool {
	return h.hmap.Remove(e)
}

func (h *HashSet[T]) Clear() bool {
	return h.hmap.Clear()
}
