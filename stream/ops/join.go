package ops

import (
	"github.com/yusaint/gostream/stream/join"
)

type On[L, R any] func(L, R) bool
type Concat[L, R, O any] func(L, R) O
type Miss[I, O any] func(I) O

type _join[L, R, O any] struct {
	AbstractOp
	sinkable   join.Sinkable
	joinStream *_joinStream[R]
	_arr       []O
}

type _joinStream[R any] struct {
	_arr []R
}

func newJoinStream[R any]() *_joinStream[R] {
	return &_joinStream[R]{
		_arr: make([]R, 0),
	}
}

func (j *_joinStream[R]) Accept(a any) error {
	j._arr = append(j._arr, a.(R))
	return nil
}

func (a *_join[L, R, O]) SetJoinStream(sinkable join.Sinkable) {
	a.sinkable = sinkable
}

func (a *_join[L, R, O]) Begin(i int64) {
	a.joinStream = newJoinStream[R]()
	capacity := i
	if capacity <= 0 {
		capacity = 16
	}
	a._arr = make([]O, 0, capacity)
	if err := a.sinkable.Sink(a.joinStream); err != nil {
	}
	a.downstream.Begin(i)
}
func (a *_join[L, R, O]) End() (any, error) {
	for _, o := range a._arr {
		a.downstream.Accept(o)
	}
	return a.downstream.End()
}
func (a *_join[L, R, O]) Accept(a2 any) error {
	panic("you should implement it!")
}

type leftJoin[L, R, O any] struct {
	_join[L, R, O]
	on     On[L, R]
	concat Concat[L, R, O]
	miss   Miss[L, O]
}

func (l *leftJoin[L, R, O]) Accept(a any) error {
	matched := false
	for _, joinElement := range l.joinStream._arr {
		if joined := l.on(a.(L), joinElement); joined {
			matched = true
			l._arr = append(l._arr, l.concat(a.(L), joinElement))
		}
	}
	if !matched {
		l._arr = append(l._arr, l.miss(a.(L)))
	}
	return nil
}

type innerJoin[L, R, O any] struct {
	_join[L, R, O]
	on     On[L, R]
	concat Concat[L, R, O]
}

func (l *innerJoin[L, R, O]) Accept(a any) error {
	for _, joinElement := range l.joinStream._arr {
		if joined := l.on(a.(L), joinElement); joined {
			l._arr = append(l._arr, l.concat(a.(L), joinElement))
		}
	}
	return nil
}

func LeftJoin[L, R, O any](on On[L, R], concat Concat[L, R, O], miss Miss[L, O]) JoinOp {
	return &leftJoin[L, R, O]{
		on:     on,
		concat: concat,
		miss:   miss,
	}
}

func InnerJoin[L, R, O any](on On[L, R], concat Concat[L, R, O]) JoinOp {
	return &innerJoin[L, R, O]{
		on:     on,
		concat: concat,
	}
}
