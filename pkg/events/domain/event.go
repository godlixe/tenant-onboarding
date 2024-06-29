package domain

import "context"

type EventAttributes struct {
	EventName string `json:"event"`
}

type Event struct {
	Data       any             `json:"data"`
	Attributes EventAttributes `json:"attributes"`
}

// type Event interface {
// 	OccuredOn() time.Time
// 	JSON() ([]byte, error)
// }

type EventListener interface {
	Name() string
	Handle(ctx context.Context, event Event) error
}

type EventListenerConstructor = func() (EventListener, error)

type EventService interface {
	Dispatch(ctx context.Context, name string, payload Event)
	Register(name string, listenersConstructor []EventListenerConstructor)
}
