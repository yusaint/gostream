package stream

import (
	"github.com/yusaint/gostream/generic"
	"github.com/yusaint/gostream/stream/ops"
)

type Pipeline struct {
	op         ops.Op
	upstream   *Pipeline
	downstream *Pipeline
}

func (p *Pipeline) Peek(op ops.Op) Streams {
	return p.build(op)
}

func (p *Pipeline) Join(sinkable Streams, op ops.JoinOp) Streams {
	op.SetJoinStream(sinkable)
	return p.build(op)
}

func (p *Pipeline) Sink(consumer generic.Consumer) error {
	var pe = p.build(ops.NewTail())
	for ; pe.upstream != nil; pe = pe.upstream {
	}
	_, err := pe.op.Handle(consumer)
	return err
}

func (p *Pipeline) Group(op ops.Op) Streams {
	return p.build(op)
}

func (p *Pipeline) Collect(op ops.Op) {
	p.build(op).evaluate()
}

func (p *Pipeline) Window(op ops.Op) Streams {
	return p.build(op)
}

func (p *Pipeline) Parallel(options ...ops.ParallelOption) Streams {
	return p.build(ops.Parallel(options...))
}

func (p *Pipeline) Skip(i int64) Streams {
	return p.build(ops.NewSkip(i))
}

func (p *Pipeline) Limit(i int64) Streams {
	return p.build(ops.NewLimit(i))
}

func (p *Pipeline) Distinct(oplist ...ops.Op) Streams {
	if len(oplist) == 0 {
		return p.build(ops.NewComparableDistinct())
	} else {
		return p.build(oplist[0])
	}
}

func (p *Pipeline) Sort(op ops.Op) Streams {
	return p.build(op)
}

func (p *Pipeline) build(op ops.Op) *Pipeline {
	cur := &Pipeline{
		op: op,
	}
	cur.upstream = p
	p.op.Link(cur.op)
	p.downstream = cur
	return cur
}

func (p *Pipeline) evaluate() (any, error) {
	var pe = p.build(ops.NewTail())
	for ; pe.upstream != nil; pe = pe.upstream {
	}
	return pe.op.Handle(pe.op)
}

func (p *Pipeline) Filter(op ops.Op) Streams {
	return p.build(op)
}

func (p *Pipeline) Map(op ops.Op) Streams {
	return p.build(op)
}

func (p *Pipeline) Reduce(op ops.Op) (any, error) {
	return p.build(op).evaluate()
}

func (p *Pipeline) Foreach(op ops.Op) error {
	_, err := p.build(op).evaluate()
	return err
}

// Stream ...
func Stream[T any](spl generic.Splittable) Streams {
	return &Pipeline{
		op: ops.NewHead(spl),
	}
}
