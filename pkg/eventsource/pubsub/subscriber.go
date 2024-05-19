package pubsub

import (
	"context"

	"github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/pkg/eventsource"
	esConstants "github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/pkg/eventsource/constants"
	"github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/pkg/logger"
	cloudevents "github.com/cloudevents/sdk-go/v2"
)

type EventSubscriber struct {
	logger        logger.Logger
	client        cloudevents.Client
	EventHandlers map[string]func(context.Context, cloudevents.Event) error
}

var _ eventsource.EventSubscriber = (*EventSubscriber)(nil)

func NewEventSubscriber(client cloudevents.Client) *EventSubscriber {
	return &EventSubscriber{
		client: client,
	}
}

func (es *EventSubscriber) SubscribeEvent(ctx context.Context) error {
	err := es.client.StartReceiver(ctx, es.handleEvents)
	if err != nil {
		es.logger.Fatalf("failed to start pubsub receiver, %v", err)
		return err
	}
	return nil
}

func (es *EventSubscriber) handleEvents(ctx context.Context, event cloudevents.Event) error {
	handler, ok := es.EventHandlers[event.Type()]
	if !ok {
		return nil
	}
	return handler(ctx, event)
}

func (es *EventSubscriber) RegisterEventHandler(eventType string, handler func(context.Context, cloudevents.Event) error) {
	if es.EventHandlers == nil {
		es.EventHandlers = make(map[string]func(context.Context, cloudevents.Event) error)
	}
	es.EventHandlers[esConstants.DeltaHQEventType+eventType] = handler
}
