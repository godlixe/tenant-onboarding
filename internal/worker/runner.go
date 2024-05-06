package worker

import (
	"context"
	"encoding/json"
	"tenant-onboarding/internal/domain/users/entity"
	"tenant-onboarding/internal/worker/deployer"
	"tenant-onboarding/internal/worker/queue"
	"tenant-onboarding/pkg/worker"
)

// workerFunc executes a job from the queue
func workerFunc(msg []byte) error {
	var (
		err        error
		tenantData entity.Tenant
	)

	err = json.Unmarshal(msg, &tenantData)
	if err != nil {
		return err
	}

	err = deployer.Deploy(context.Background(), &tenantData)
	if err != nil {
		return err
	}

	return nil
}

func Run() {
	queue := queue.Init()

	base := worker.Init(
		worker.WithQueue(queue),
		worker.WithWorkerFunc(workerFunc),
	)
	base.StartTest()
}
