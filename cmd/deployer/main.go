package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"tenant-onboarding/internal/domain/users/entity"
	"tenant-onboarding/worker/worker"

	"github.com/joho/godotenv"
)

func deployerFunc(msg []byte) {
	var tenantData entity.Tenant

	err := json.Unmarshal(msg, &tenantData)
	if err != nil {
		fmt.Println(err)
		return
	}

	worker.Deploy(context.Background(), &tenantData)
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	base := worker.Init(
		make(chan []byte),
		worker.WithWorkerNum(1),
		worker.WithWorkerTimeout(1000),
		worker.WithWorkerFunc(deployerFunc),
	)

	base.StartTest()
}
