package ops

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"runtime"
	"runtime/debug"
)

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
	strategy WorkerStrategy
	poolSize int
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

func (p *_parallel) End() (any, error) {
	if err := p.impl.Wait(); err != nil {
		return nil, err
	}
	return p.downstream.End()
}
func (p *_parallel) Link(next Op) { p.downstream = next; p.impl.Link(next) }
func (p *_parallel) Accept(a any) error {
	p.impl.AcceptMessage(a)
	return nil
}

type worker interface {
	Link(downstream Op)
	AcceptMessage(m any)
	Wait() error
}

type bufferPoolWorker struct {
	downstream Op
	eg         errgroup.Group
}

func (b *bufferPoolWorker) Wait() error {
	return b.eg.Wait()
}

func (b *bufferPoolWorker) Link(downstream Op) {
	b.downstream = downstream
}

func (b *bufferPoolWorker) AcceptMessage(m any) {
	b.eg.Go(func() (err error) {
		defer func() {
			if e := recover(); e != nil {
				err = fmt.Errorf("%v", e)
				debug.PrintStack()
			}
		}()
		err = b.downstream.Accept(m)
		return
	})
}

func newFixedPoolWorker(size int) *bufferPoolWorker {
	w := &bufferPoolWorker{}
	w.eg.SetLimit(size)
	return w
}

func Parallel(options ...ParallelOption) Op {
	cfg := parallelConfig{
		strategy: FixedPoolStrategy,
		poolSize: runtime.NumCPU(),
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
