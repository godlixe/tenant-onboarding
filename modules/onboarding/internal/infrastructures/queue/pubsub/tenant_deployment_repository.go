package pubsub

import (
	"context"
	"encoding/json"
	"os"
	"tenant-onboarding/modules/onboarding/internal/domain/entities"

	"cloud.google.com/go/pubsub"
)

type Queuer interface {
	Publish(ctx context.Context, msg []byte) error
}

type TenantDeploymentRepository struct {
	client *pubsub.Client
}

func NewTenantDeploymentRepository(
	client *pubsub.Client,
) *TenantDeploymentRepository {
	return &TenantDeploymentRepository{
		client: client,
	}
}

func (r *TenantDeploymentRepository) PublishTenantDeploymentJob(
	ctx context.Context,
	job *entities.TenantDeploymentJob,
) error {
	topic := r.client.Topic(os.Getenv("DEPLOYMENT_TOPIC"))

	b, err := json.Marshal(job)
	if err != nil {
		return err
	}

	res := topic.Publish(ctx, &pubsub.Message{
		Data: b,
	})

	_, err = res.Get(ctx)
	if err != nil {
		return err
	}

	return nil
}
