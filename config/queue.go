package config

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/pubsub"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

func InitPubsub() *pubsub.Client {
	creds, err := google.FindDefaultCredentials(context.Background())
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	pubsubClient, err := pubsub.NewClient(
		context.Background(),
		os.Getenv("GOOGLE_PROJECT_ID"),
		option.WithCredentialsJSON(creds.JSON),
	)
	if err != nil {
		os.Exit(1)
	}

	return pubsubClient
}

// func (q *Pubsub) Subscribe(topic string, ctx context.Context, workerFunc func(data []byte) error) error {
// 	sub := q.client.Subscription(topic)
// 	err := sub.Receive(ctx, func(ctx context.Context, m *pubsub.Message) {
// 		err := workerFunc(m.Data)
// 		if err != nil {
// 			m.Nack()
// 		}
// 		m.Ack()
// 	})
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (q *Pubsub) Publish(topic string, ctx context.Context, msg []byte) error {
// 	topic := q.client.Topic(q.topic)

// 	res := topic.Publish(ctx, &pubsub.Message{
// 		Data: msg,
// 	})

// 	_, err := res.Get(ctx)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }
