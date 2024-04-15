package queue

import (
	"context"
	"encoding/json"
	"tenant-onboarding/internal/domain/users/entity"

	"cloud.google.com/go/pubsub"
)

type TenantPublisher struct {
	client *pubsub.Client
}

func NewTenantPublisher(
	client *pubsub.Client,
) *TenantPublisher {
	return &TenantPublisher{
		client: client,
	}
}

func (p *TenantPublisher) Publish(
	ctx context.Context,
	tenant *entity.Tenant,
) error {
	b, err := json.Marshal(tenant)
	if err != nil {
		return err
	}

	topic := p.client.Topic("test-topic")
	res := topic.Publish(
		ctx,
		&pubsub.Message{
			Data: b,
		},
	)

	_, err = res.Get(ctx)
	if err != nil {
		return err
	}
	return nil
}
