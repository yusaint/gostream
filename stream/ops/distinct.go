package ops

import "github.com/yusaint/gostream/set"

type comparableDistinct struct {
	AbstractOp
	set map[any]bool
}

func (f *comparableDistinct) Accept(a any) error {
	if _, isOK := f.set[a]; !isOK {
		f.set[a] = true
		return f.downstream.Accept(a)
	}
	return nil
}

func NewComparableDistinct() Op {
	return &comparableDistinct{
		set: make(map[any]bool),
	}
}

type comparatorDistinct[T any, R comparable] struct {
	AbstractOp
	set      set.Set[R]
	hashcode HashCodeFunc[T, R]
}

func (f *comparatorDistinct[T, R]) Accept(a any) error {
	code, err := f.hashcode(a.(T))
	if err != nil {
		return err
	}
	if !f.set.Contains(code) {
		f.set.Add(code)
		return f.downstream.Accept(a)
	}
	return nil
}

type HashCodeFunc[T any, R comparable] func(T) (R, error)

func NewDistinct[T any, R comparable](hashcode HashCodeFunc[T, R]) Op {
	return &comparatorDistinct[T, R]{
		set:      set.NewHashSet[R](),
		hashcode: hashcode,
	}
}
