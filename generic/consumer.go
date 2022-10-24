package generic

type Consumer interface {
	Accept(any) error
}
