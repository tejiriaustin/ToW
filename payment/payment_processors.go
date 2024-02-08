package payment

type Processor struct{}

func NewPaymentProcessor() *Processor {
	return &Processor{}
}

type ProcessorProvider interface{}
