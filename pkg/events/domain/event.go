package domain

import (
	"context"
	"time"
)

type Event interface {
	OccuredOn() time.Time
	JSON() ([]byte, error)
}

type EventListener interface {
	Name() string
	Handle(ctx context.Context, event Event) error
}

type EventListenerConstructor = func() (EventListener, error)

type EventService interface {
	Dispatch(ctx context.Context, name string, payload Event)
	Register(name string, listenersConstructor []EventListenerConstructor)
}
