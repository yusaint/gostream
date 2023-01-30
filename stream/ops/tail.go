package ops

import "github.com/yusaint/gostream/generic"

type _tail struct {
	AbstractOp
}

func NewTail() Op {
	return &_tail{}
}

func (a *_tail) Begin(i int64)                                 {}
func (a *_tail) End() (any, error)                             { return nil, nil }
func (a *_tail) Accept(a2 any) error                           { return nil }
func (a *_tail) Handle(consumer generic.Consumer) (any, error) { panic("implement me") }
func (a *_tail) Link(next Op)                                  {}
