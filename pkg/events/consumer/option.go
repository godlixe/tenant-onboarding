package consumer

import (
	"context"
	"tenant-onboarding/providers"
)

func WithContext(context context.Context) func(e *EventConsumer) {
	return func(e *EventConsumer) {
		e.context = context
	}
}

func WithQueue(queue Queuer) func(e *EventConsumer) {
	return func(e *EventConsumer) {
		e.queue = queue
	}
}

func WithWorkerFunc(fn func(app *providers.App, msg []byte) error) func(e *EventConsumer) {
	return func(e *EventConsumer) {
		e.workerFn = fn
	}
}
