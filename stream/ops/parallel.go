package ops

import "golang.org/x/exp/rand"

// WorkerStrategy define the concurrent model
type WorkerStrategy int

const (
	// BufferPoolStrategy will trigger one goroutine for all the incoming message from upstream
	BufferPoolStrategy WorkerStrategy = 1
	// FixedPoolStrategy will set an upper size for the worker pool
	// will use a producer-consumer model to deal with message stream, need a hasher to make the stream data more consistent
	// or we can use a work steel method to speed up the consumer
	// todo propagate the parallel flag to downstream and enable all parallel process in all following nodes transparently
	FixedPoolStrategy WorkerStrategy = 2
)

type parallelConfig struct {
	strategy          WorkerStrategy
	poolSize          int
	enableWorkerSteel bool
}

type ParallelOption func(config *parallelConfig)

func WithFixedPool(size int) ParallelOption {
	return func(config *parallelConfig) {
		config.strategy = FixedPoolStrategy
		config.poolSize = size
	}
}

type _parallel struct {
	AbstractOp
	impl worker
}

func (p *_parallel) Link(next Op) { p.downstream = next; p.impl.Link(next) }
func (p *_parallel) Accept(a any) error {
	p.impl.AcceptMessage(a)
	return nil
}

type worker interface {
	Link(downstream Op)
	AcceptMessage(m any)
}

type bufferPoolWorker struct {
	downstream Op
}

func (b *bufferPoolWorker) Link(downstream Op) {
	b.downstream = downstream
}

func (b *bufferPoolWorker) AcceptMessage(m any) {
	go b.downstream.Accept(m)
}

type Hasher func(m any, size int) int

var defaultHasher = func(m any, size int) int {
	return rand.Intn(size)
}

type fixedPoolWorker struct {
	downstream Op
	poolSize   int
	hasher     Hasher
	runq       []chan any
}

func (f *fixedPoolWorker) Link(downstream Op) {
	f.downstream = downstream
}

func (f *fixedPoolWorker) loop() {
	for i := 0; i < f.poolSize; i++ {
		f.runq[i] = make(chan any, 1)
		go func(index int) {
			for {
				select {
				case m := <-f.runq[index]:
					f.downstream.Accept(m)
				}
			}
		}(i)
	}
}

func newFixedPoolWorker(size int) *fixedPoolWorker {
	w := &fixedPoolWorker{
		poolSize: size,
		hasher:   defaultHasher,
		runq:     make([]chan any, size),
	}
	w.loop()
	return w
}

func (f *fixedPoolWorker) AcceptMessage(m any) {
	index := f.hasher(m, f.poolSize)
	f.runq[index] <- m
}

func Parallel(options ...ParallelOption) Op {
	cfg := parallelConfig{
		strategy: BufferPoolStrategy,
	}
	for _, opt := range options {
		opt(&cfg)
	}
	switch cfg.strategy {
	case FixedPoolStrategy:
		return &_parallel{
			impl: newFixedPoolWorker(cfg.poolSize),
		}
	case BufferPoolStrategy:
		return &_parallel{
			impl: &bufferPoolWorker{},
		}
	default:
		return &_parallel{
			impl: &bufferPoolWorker{},
		}
	}
}
