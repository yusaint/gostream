package join

import "github.com/yusaint/gostream/generic"

type Sinkable interface {
	Sink(consumer generic.Consumer) error
}
