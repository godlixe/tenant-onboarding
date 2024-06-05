package queue

import (
	"context"
	"tenant-onboarding/providers"
)

type Mock struct {
	app  *providers.App
	Data []byte
}

func (m *Mock) Subscribe(ctx context.Context, workerFunc func(app *providers.App, data []byte) error) error {
	workerFunc(m.app, m.Data)
	return nil
}

func (m *Mock) Publish(ctx context.Context, msg []byte) error {
	m.Data = []byte("test-message")
	return nil
}
