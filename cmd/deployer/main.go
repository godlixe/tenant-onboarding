package main

import (
	"context"
	"fmt"
	"os"
	"tenant-onboarding/worker/queue"
	"tenant-onboarding/worker/worker"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	var tenantQueue queue.Queue
	var worker worker.Worker
	messages := make(chan []byte, 100)

	tenantQueue.Init(messages)

	ctx, _ := context.WithTimeout(context.Background(), 10)
	for i := 0; i < 1; i++ {
		tenantQueue.Push(ctx)
	}
	worker.Init(messages)

	tenantQueue.Pull(ctx)
}
