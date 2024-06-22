package queue

import (
	"context"
	"fmt"
	"os"
	"tenant-onboarding/providers"

	"cloud.google.com/go/pubsub"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

type Pubsub struct {
	client *pubsub.Client
	topic  string
	app    *providers.App
}

func InitPubsub(topic string, app *providers.App) *Pubsub {
	creds, err := google.FindDefaultCredentials(context.Background())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println(string(creds.JSON))

	pubsubClient, err := pubsub.NewClient(
		context.Background(),
		os.Getenv("GOOGLE_PROJECT_ID"),
		option.WithCredentialsJSON(creds.JSON),
	)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	q := &Pubsub{
		client: pubsubClient,
		topic:  topic,
		app:    app,
	}

	return q
}

func (q *Pubsub) Subscribe(ctx context.Context, workerFunc func(app *providers.App, data []byte) error) error {
	sub := q.client.Subscription(q.topic)
	fmt.Println("ABC", sub)
	err := sub.Receive(ctx, func(ctx context.Context, m *pubsub.Message) {
		err := workerFunc(q.app, m.Data)
		if err != nil {
			fmt.Println("NACK", err)
			m.Ack()
		}
		fmt.Println("ACK")
		m.Ack()
	})
	if err != nil {
		return err
	}

	return nil
}

func (q *Pubsub) Publish(ctx context.Context, msg []byte) error {
	topic := q.client.Topic(q.topic)

	res := topic.Publish(ctx, &pubsub.Message{
		Data: msg,
	})

	_, err := res.Get(ctx)
	if err != nil {
		return err
	}

	return nil
}
