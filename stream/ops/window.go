package ops

func UseFlipWindow[T any](size int) Op {
	w := &_window[T]{}
	w.impl = &flipWindow[T]{
		size:      size,
		container: make([]T, 0, size),
	}
	return w
}

type _window[T any] struct {
	AbstractOp
	impl window[T]
}

type window[T any] interface {
	IsFull(T) bool
	WindowContents() []T
	MoveOn(T)
	Size() int
}

type flipWindow[T any] struct {
	size      int
	container []T
}

func (f *flipWindow[T]) Size() int {
	return f.size
}
func (f *flipWindow[T]) MoveOn(t T) {
	f.container = make([]T, 0, f.size)
	f.container = append(f.container, t)
}

func (f *flipWindow[T]) IsFull(t T) bool {
	if len(f.container) >= f.size {
		return true
	} else {
		f.container = append(f.container, t)
		return false
	}
}

func (f *flipWindow[T]) WindowContents() []T {
	return f.container
}
func (f *_window[T]) Begin(i int64)     {}
func (f *_window[T]) End() (any, error) { return f.downstream.End() }
func (f *_window[T]) Accept(a any) error {
	isFull := f.impl.IsFull(a)
	if !isFull {
		return nil
	}
	f.downstream.Begin(int64(f.impl.Size()))
	if err := f.downstream.Accept(f.impl.WindowContents()); err != nil {
		return err
	}
	f.impl.MoveOn(a)
	return nil
}
