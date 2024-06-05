package domain

import "context"

// Listener is a struct that contains all listeners for domain events.
// The listeners are put in a map of listeners.
type Listener struct {
	listeners map[string]EventListener
}

func NewListener() *Listener {
	return &Listener{
		listeners: make(map[string]EventListener),
	}
}

// Dispatch dispatches a domain event and handles the event with
// the listener assigned to Listener.
func (l *Listener) Dispatch(ctx context.Context, name string, event Event) {
	_, exists := l.listeners[name]
	if !exists {
		return
	}

	for _, listener := range l.listeners {
		if err := listener.Handle(ctx, event); err != nil {
			return
		}
	}
}

// Register registers a new event listener by a name and the constructor.
func (l *Listener) Register(name string, constructor EventListenerConstructor) error {
	eventListener, err := constructor()
	if err != nil {
		return err
	}

	l.listeners[name] = eventListener

	return nil
}
