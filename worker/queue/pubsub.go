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

func (q *Queue) Init(messages chan<- []byte) {
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
		fmt.Println(err)
		os.Exit(1)
	}

	q.client = pubsubClient
	q.messages = messages
}

func (q *Queue) Pull(ctx context.Context) {
	// topic := q.client.Topic("test-topic")

	sub := q.client.Subscription("test-topic")
	// if err != nil {
	// 	fmt.Println(err)
	// 	os.Exit(1)
	// }
	fmt.Println("fetching data")
	err := sub.Receive(context.Background(), func(ctx context.Context, m *pubsub.Message) {
		fmt.Println("got data, ", string(m.Data))
		q.messages <- m.Data
		m.Ack()
	})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func (q *Queue) Push(ctx context.Context) error {
	topic := q.client.Topic("test-topic")

	res := topic.Publish(ctx, &pubsub.Message{
		Data: []byte("test string"),
	})

	_, err := res.Get(ctx)
	if err != nil {
		return err
	}

	return nil
}
