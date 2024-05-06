package worker

import (
	"context"
	"fmt"
	"os"
)

type QueueItf interface {
	// Pull pulls a message indefinitely
	// and returns an error if encountered.
	Pull(context.Context, func([]byte) error) error

	// Push pushes a message to the queue.
	Push(context.Context, []byte) error
}

type WorkerItf interface {
	Work(context.Context, func([]byte))
}

type WorkerBase struct {
	// Queue containing the jobs.
	queue QueueItf

	// Function to execute a job from the queue.
	workerFunc func([]byte) error

	// Defines timeout.
	workerTimeout uint
}

type Options func(wb *WorkerBase)

func Init(
	opts ...Options,
) *WorkerBase {
	w := &WorkerBase{}

	// use options
	for _, opt := range opts {
		opt(w)
	}

	return w
}

// Pull messages indefinitely
// and exits if an error occurs.
func (w *WorkerBase) Start(ctx context.Context) {
	err := w.queue.Pull(ctx, w.workerFunc)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Test queue by pushing messages and
// calling Start() to receive and process messages.
func (w *WorkerBase) StartTest() {
	ctx, _ := context.WithCancel(context.Background())
	w.queue.Push(ctx, []byte("Sample message"))
	w.Start(context.Background())
}
