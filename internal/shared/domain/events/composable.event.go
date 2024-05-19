package events

import (
	"context"

	"github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/pkg/eventsource"
)

type Event interface {
	Topic() string
	EventType() string
	ID() string
}

type DomainHasEvents struct {
	Events []Event
}

func NewDomainHasEvents() *DomainHasEvents {
	return &DomainHasEvents{}
}

func (d *DomainHasEvents) AddEvent(e Event) {
	d.Events = append(d.Events, e)
}

func (d *DomainHasEvents) PublishEvents(ctx context.Context, publisher eventsource.EventPublisher) error {
	for _, e := range d.Events {
		if err := publisher.PublishEvent(ctx, e.Topic(), e.EventType(), e.ID(), e); err != nil {
			return err
		}
	}
	return nil
}
