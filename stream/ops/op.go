package ops

import (
	"github.com/yusaint/gostream/generic"
	"github.com/yusaint/gostream/stream/join"
)

type Op interface {
	Begin(int64)
	End() (any, error)
	Accept(any) error
	Handle(consumer generic.Consumer) (any, error)
	Link(next Op)
}

type JoinOp interface {
	Op
	SetJoinStream(sinkable join.Sinkable)
}

type AbstractOp struct {
	downstream Op
	e          error
}

func (a *AbstractOp) Begin(i int64)                                 { a.downstream.Begin(i) }
func (a *AbstractOp) End() (any, error)                             { return a.downstream.End() }
func (a *AbstractOp) Accept(a2 any) error                           { return a.downstream.Accept(a2) }
func (a *AbstractOp) Handle(consumer generic.Consumer) (any, error) { panic("implement me") }
func (a *AbstractOp) Link(next Op)                                  { a.downstream = next }
