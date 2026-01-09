package workers

import "github.com/samber/do/v2"

type ConsumerWorker struct {
}

func NewConsumerWorker(injector do.Injector) (*ConsumerWorker, error) {
	return &ConsumerWorker{}, nil
}
