package processor

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"tenant-onboarding/modules/onboarding/internal/domain/entities"
	"tenant-onboarding/modules/onboarding/internal/domain/repositories"
	"tenant-onboarding/modules/onboarding/internal/processor/deployer"
	"tenant-onboarding/modules/onboarding/internal/processor/queue"
	"tenant-onboarding/pkg/events"
	"tenant-onboarding/pkg/events/consumer"
	"tenant-onboarding/providers"

	"github.com/samber/do"
)

// workerFunc executes a job from the queue
func workerFunc(app *providers.App, msg []byte) error {
	var (
		err       error
		tenantJob entities.TenantDeploymentJob
		// eventQueue      *queue.Pubsub   = queue.InitPubsub("tenant_deployment")
		terraformConfig deployer.Config = deployer.Config{
			GoogleServiceAccountAbsolutePath: os.Getenv("GOOGLE_SERVICE_ACCOUNT_ABSOLUTE_PATH"),
			GoogleProjectID:                  os.Getenv("GOOGLE_PROJECT_ID"),
			GoogleDeploymentRegion:           os.Getenv("GOOGLE_DEPLOYMENT_REGION"),
			TerraformAbsolutePath:            os.Getenv("TF_ABSOLUTE_PATH"),
			TerraformExecPath:                os.Getenv("TF_EXEC_PATH"),
			TerraformBackendBucket:           os.Getenv("TF_BACKEND_BUCKET"),
		}
	)

	err = json.Unmarshal(msg, &tenantJob)
	if err != nil {
		return err
	}
	fmt.Println("MSG", string(msg))
	infrastructureRepository, err := do.Invoke[repositories.InfrastructureRepository](app.Injector)
	if err != nil {
		return err
	}

	err = deployer.Deploy(
		context.Background(),
		terraformConfig,
		&tenantJob,
		infrastructureRepository,
	)
	if err != nil {
		return err
	}
	fmt.Println("2 DONE")

	// publish messages to queue
	// eventQueue.Publish(context.Background(), msg)

	return nil
}

func Run(ctx context.Context, app *providers.App) {
	deploymentQueue := queue.InitPubsub("tenant_deployment", app)

	deploymentEventConsumer := consumer.New(
		consumer.WithContext(ctx),
		consumer.WithQueue(deploymentQueue),
		consumer.WithWorkerFunc(workerFunc),
	)

	mockConsumer := consumer.New(
		consumer.WithContext(ctx),
		consumer.WithQueue(&queue.Mock{Data: []byte("hi")}),
		consumer.WithWorkerFunc(func(app *providers.App, msg []byte) error { fmt.Println(string(msg)); return nil }),
	)

	eventProcessor := events.NewProcessor(
		mockConsumer,
		deploymentEventConsumer,
	)

	eventProcessor.Start()
}
