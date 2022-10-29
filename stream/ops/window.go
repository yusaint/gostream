package ops

type windowType int

const (
	useFlipWindow windowType = 1
)

type windowCfg struct {
	windowType windowType
	windowSize int
}

type WindowOption func(cfg *windowCfg)

func WithFlipWindow(size int) WindowOption {
	return func(cfg *windowCfg) {
		cfg.windowSize = size
		cfg.windowType = useFlipWindow
	}
}

type _window struct {
	AbstractOp
	impl window
}

type window interface {
	IsFull(any) bool
	WindowContents() []any
	MoveOn(any)
	Size() int
}

type flipWindow struct {
	size      int
	container []any
}

func (f *flipWindow) Size() int {
	return f.size
}
func (f *flipWindow) MoveOn(t any) {
	f.container = make([]any, 0, f.size)
	f.container = append(f.container, t)
}

func (f *flipWindow) IsFull(t any) bool {
	if len(f.container) >= f.size {
		return true
	} else {
		f.container = append(f.container, t)
		return false
	}
}

func (f *flipWindow) WindowContents() []any {
	return f.container
}
func (f *_window) Begin(i int64)     {}
func (f *_window) End() (any, error) { return f.downstream.End() }
func (f *_window) Accept(a any) error {
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

func Window(opts ...WindowOption) Op {
	cfg := windowCfg{
		windowType: useFlipWindow,
	}
	for _, opt := range opts {
		opt(&cfg)
	}
	switch cfg.windowType {
	case useFlipWindow:
		return &_window{
			impl: &flipWindow{
				size:      cfg.windowSize,
				container: make([]any, 0, cfg.windowSize),
			},
		}
	default:
		panic("unknown window type")
	}
}
