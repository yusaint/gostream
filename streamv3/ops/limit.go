package ops

type limit struct {
	AbstractOp
	state int64
	count int64
}

func (f *limit) Begin(i int64) {
	f.state = 0
	f.downstream.Begin(i)
}
func (f *limit) Accept(a any) error {
	if f.state < f.count {
		f.state++
		return f.downstream.Accept(a)
	}
	return nil
}

func NewLimit(count int64) Op {
	return &limit{
		count: count,
	}
}
