package worker

import (
	"context"
	"fmt"
	"os"
	"tenant-onboarding/worker/queue"
)

var workerBase WorkerBase

type QueueItf interface {
	// Pull pulls a message indefinitely
	// and returns an error if encountered.
	Pull(context.Context) error

	// Push pushes a message to the queue.
	Push(context.Context, []byte) error
}

type WorkerItf interface {
	Work(context.Context, func([]byte))
}

type WorkerBase struct {
	// WorkerNum defines the number of goroutines
	// the program will create.
	WorkerNum uint

	// Defines timeout.
	WorkerTimeout uint

	// Function that will run for every message.
	workerFunc func([]byte)

	// Message channel.
	messageChan chan []byte

	// interfaces
	queue QueueItf
}

type Dependencies interface {
	Init(workerBase *WorkerBase)
}

type Options func(wb *WorkerBase)

func Init(
	messages chan []byte,
	opts ...Options,
) *WorkerBase {
	defaultQueue := queue.Init(messages)

	w := &WorkerBase{
		queue:     defaultQueue,
		WorkerNum: 1,
		workerFunc: func(msg []byte) {
			fmt.Printf("this message is from the default workerFunc: %s\n", string(msg))
		},
		messageChan: messages,
	}
	// use options

	for _, opt := range opts {
		opt(w)
	}

	return w
}

// Pull messages indefinitely
// and exits if an error occurs.
func (w *WorkerBase) Start(ctx context.Context) {
	err := w.queue.Pull(ctx)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// Test queue by pushing messages and
// calling Start() to receive and process messages.
func (w *WorkerBase) StartTest() {
	ctx, _ := context.WithCancel(context.Background())

	w.queue.Push(ctx, []byte("hi"))

	var i uint

	for i = 0; i < w.WorkerNum; i++ {
		go func() {
			for msg := range w.messageChan {
				w.workerFunc(msg)
			}
		}()
	}

	w.Start(context.Background())
}
