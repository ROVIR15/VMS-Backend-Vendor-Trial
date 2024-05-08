package pubsub

import (
	"context"

	"github.com/ILUMINA-Pte-Ltd/PrimeCRM-Backend-Service/pkg/logger"
	cloudevents "github.com/cloudevents/sdk-go/v2"
)

type EventSubscriber struct {
	logger        logger.Logger
	client        cloudevents.Client
	EventHandlers map[string]func(cloudevents.Event) error
}

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

func (es *EventSubscriber) handleEvents(event cloudevents.Event) error {
	handler, ok := es.EventHandlers[event.Type()]
	if !ok {
		return nil
	}
	return handler(event)
}
