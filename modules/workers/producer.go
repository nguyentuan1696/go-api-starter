package workers

import "github.com/samber/do/v2"

type ProducerWorker struct {
}

func NewProducerWorker(injector do.Injector) (*ProducerWorker, error) {
	return &ProducerWorker{}, nil
}
