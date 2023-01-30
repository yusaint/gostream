package ops

import (
	"github.com/yusaint/gostream/generic"
)

type head struct {
	AbstractOp
	spl generic.Splittable
}

func (f *head) Handle(c generic.Consumer) (any, error) {
	f.Begin(f.spl.EstimatedSize())
	if err := f.spl.ForeachRemaining(c); err != nil {
		return nil, err
	}
	return f.End()
}

func NewHead(spl generic.Splittable) Op {
	return &head{
		spl: spl,
	}
}
