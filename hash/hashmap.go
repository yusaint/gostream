package hash

import (
	"container/list"
	"github.com/yusaint/gostream/generic"
	"sync"
)

type Comparator interface {
}

type Map[T comparable, R any] struct {
	lock     sync.RWMutex
	hmap     map[T]*list.Element
	l        *list.List
	iterator *list.Element
}

func (m *Map[T, R]) Front() R {
	m.iterator = m.l.Front()
	return m.iterator.Value.(R)
}

func (m *Map[T, R]) Current() R {
	return m.iterator.Value.(R)
}

func (m *Map[T, R]) Next() bool {
	if m.iterator == nil {
		return false
	}
	m.iterator = m.iterator.Next()
	if m.iterator == nil {
		return false
	}
	return true
}

func (m *Map[T, R]) EstimatedSize() int64 {
	return int64(m.Size())
}

func (m *Map[T, R]) ForeachRemaining(sink generic.Consumer) error {
	m.iterator = m.l.Front()
	for {
		isContinue, err := m.TryAdvance(sink)
		if err != nil {
			return err
		} else {
			if !isContinue {
				return nil
			}
		}
	}
}

func (m *Map[T, R]) TryAdvance(sink generic.Consumer) (bool, error) {
	if err := sink.Accept(m.iterator.Value.(R)); err != nil {
		return false, err
	}
	next := m.iterator.Next()
	if next == nil {
		return false, nil
	}
	m.iterator = next
	return true, nil
}

func NewMap[T comparable, R any]() Hash[T, R] {
	return &Map[T, R]{
		hmap: make(map[T]*list.Element),
		l:    list.New(),
	}
}

func (m *Map[T, R]) ToArray() []R {
	values := make([]R, 0, len(m.hmap))
	m.lock.RLock()
	defer m.lock.RUnlock()
	for _, v := range m.hmap {
		values = append(values, v.Value.(R))
	}
	return values
}

func (m *Map[T, R]) Add(k T, v R) bool {
	m.lock.Lock()
	defer m.lock.Unlock()
	_, isOK := m.hmap[k]
	e := m.l.PushBack(v)
	m.hmap[k] = e
	return isOK
}

func (m *Map[T, R]) Contains(k T) bool {
	m.lock.RLock()
	defer m.lock.RUnlock()
	_, isOK := m.hmap[k]
	return isOK
}

func (m *Map[T, R]) Size() int {
	return len(m.hmap)
}

func (m *Map[T, R]) Remove(k T) bool {
	m.lock.Lock()
	defer m.lock.Unlock()
	e, isOK := m.hmap[k]
	m.l.Remove(e)
	delete(m.hmap, k)
	return isOK
}

func (m *Map[T, R]) Clear() bool {
	m.hmap = make(map[T]*list.Element)
	m.l = list.New()
	return true
}
