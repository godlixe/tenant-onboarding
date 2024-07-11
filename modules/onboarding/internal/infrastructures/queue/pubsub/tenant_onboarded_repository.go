package pubsub

import (
	"context"
	"encoding/json"
	"os"
	"tenant-onboarding/pkg/deployer/types"

	"cloud.google.com/go/pubsub"
)

type TenantOnboardedRepository struct {
	client *pubsub.Client
}

func NewTenantOnboardedRepository(
	client *pubsub.Client,
) *TenantOnboardedRepository {
	return &TenantOnboardedRepository{
		client: client,
	}
}

func (r *TenantOnboardedRepository) PublishTenantOnboarded(
	ctx context.Context,
	job *types.TenantOnboardedEvent,
) error {
	topic := r.client.Topic(os.Getenv("TENANT_ONBOARDED_TOPIC"))

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
