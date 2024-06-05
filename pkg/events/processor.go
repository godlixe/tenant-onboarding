package events

type Consumer interface {
	Start()
}

type EventProcessor struct {
	Consumers []Consumer
}

func NewProcessor(consumers ...Consumer) *EventProcessor {
	p := &EventProcessor{}
	p.Consumers = append(p.Consumers, consumers...)
	return p
}

func (p *EventProcessor) Start() {

	// Start starts consuming events from event consumers
	// in a new goroutine asynchronously.
	for _, consumer := range p.Consumers {
		go consumer.Start()
	}
}
