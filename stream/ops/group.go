package ops

import "sync"

type KeyFunc[T any, R comparable] func(T) (R, error)

type _group[T any, R comparable] struct {
	AbstractOp
	fn     KeyFunc[T, R]
	locker sync.RWMutex
	values map[R][]T
}

func (f *_group[T, R]) Begin(i int64) {
	f.values = make(map[R][]T)
}
func (f *_group[T, R]) End() (any, error) {
	f.downstream.Begin(int64(len(f.values)))
	if err := f.downstream.Accept(f.values); err != nil {
		return nil, err
	}
	return f.downstream.End()
}
func (f *_group[T, R]) Accept(a any) (err error) {
	key, err := f.fn(a.(T))
	if err != nil {
		return err
	}
	f.locker.Lock()
	defer f.locker.Unlock()
	if _, isOK := f.values[key]; !isOK {
		f.values[key] = make([]T, 0, 10)
	}
	f.values[key] = append(f.values[key], a.(T))
	return nil
}

func NewGroup[T any, R comparable](keyFunc KeyFunc[T, R]) Op {
	return &_group[T, R]{
		fn: keyFunc,
	}
}
