package workers

import "github.com/samber/do/v2"

var WorkerPackage = do.Package(
	do.Lazy(NewConsumerWorker),
	do.Lazy(NewProducerWorker),
)
