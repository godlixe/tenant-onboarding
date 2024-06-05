package consumer

import (
	"context"
	"fmt"
	"os"
	"tenant-onboarding/providers"
)

type Queuer interface {
	Subscribe(context.Context, func(app *providers.App, msg []byte) error) error
}

type EventConsumer struct {
	context  context.Context
	queue    Queuer
	workerFn func(app *providers.App, msg []byte) error
}

type Options func(e *EventConsumer)

func New(opts ...Options) *EventConsumer {
	e := &EventConsumer{}

	for _, opt := range opts {
		opt(e)
	}
	return e
}

func (e *EventConsumer) Start() {
	err := e.queue.Subscribe(e.context, e.workerFn)
	if err != nil {
		fmt.Println("error consuming event: ", err)
		os.Exit(1)
	}
}
