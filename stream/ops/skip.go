package ops

type skip struct {
	AbstractOp
	state  int64
	offset int64
}

func (f *skip) Accept(a any) error {
	if f.state < f.offset {
		f.state++
		return nil
	}
	return f.downstream.Accept(a)
}

func NewSkip(offset int64) Op {
	return &skip{
		offset: offset,
	}
}
