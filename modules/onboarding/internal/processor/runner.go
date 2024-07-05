package processor

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"tenant-onboarding/modules/onboarding/internal/processor/pipeline"
	"tenant-onboarding/modules/onboarding/internal/processor/queue"
	"tenant-onboarding/pkg/deployer"
	"tenant-onboarding/pkg/deployer/types"
	"tenant-onboarding/pkg/events"
	"tenant-onboarding/pkg/events/consumer"
	"tenant-onboarding/pkg/framework"
	"tenant-onboarding/providers"
)

// workerFunc executes a job from the queue
func workerFunc(app *providers.App, msg []byte) error {
	var (
		err             error
		tenantJob       types.TenantDeploymentJob
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

	terraformDeploymentPipeline := pipeline.NewTerraformDeployer(
		app,
		terraformConfig,
	)

	deploymentPipeline := deployer.NewDeploymentPipeline(
		terraformDeploymentPipeline,
	)

	err = deploymentPipeline.Start(context.TODO(), tenantJob)
	if err != nil {
		return err
	}

	return nil
}

func Run(ctx context.Context, app *providers.App) {

	var deploymentTopic *queue.Pubsub
	if framework.CheckIntegratedMode() {
		deploymentTopic = queue.InitPubsub(os.Getenv("BILLING_PAID_TOPIC_SUBSCRIPTION"), app)
	} else {
		deploymentTopic = queue.InitPubsub(os.Getenv("DEPLOYMENT_TOPIC_SUBSCRIPTION"), app)

	}

	deploymentEventConsumer := consumer.New(
		consumer.WithContext(ctx),
		consumer.WithQueue(deploymentTopic),
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
