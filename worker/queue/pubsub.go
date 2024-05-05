package queue

import (
	"context"
	"fmt"
	"os"

	"cloud.google.com/go/pubsub"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/option"
)

type Queue struct {
	client   *pubsub.Client
	messages chan<- []byte
}

func Init(messages chan<- []byte) *Queue {
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

	q := &Queue{}

	q.client = pubsubClient
	q.messages = messages

	return q
}

func (q *Queue) Pull(ctx context.Context) error {
	sub := q.client.Subscription("test-topic")
	err := sub.Receive(ctx, func(ctx context.Context, m *pubsub.Message) {
		fmt.Println("got data, ", string(m.Data))
		q.messages <- m.Data
		m.Ack()
	})
	if err != nil {
		return err
	}

	return nil
}

func (q *Queue) Push(ctx context.Context, msg []byte) error {
	topic := q.client.Topic("test-topic")

	res := topic.Publish(ctx, &pubsub.Message{
		Data: msg,
	})

	_, err := res.Get(ctx)
	if err != nil {
		return err
	}

	return nil
}
