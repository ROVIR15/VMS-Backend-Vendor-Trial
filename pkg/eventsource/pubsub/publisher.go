package pubsub

import (
	"context"
	"fmt"

	"github.com/ROVIR15/VMS-Backend-Vendor-Trial/pkg/eventsource"
	"github.com/ROVIR15/VMS-Backend-Vendor-Trial/pkg/logger"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	cecontext "github.com/cloudevents/sdk-go/v2/context"
	uuid "github.com/satori/go.uuid"
)

type EventPublisher struct {
	logger logger.Logger
	client *eventsource.CloudEventsClient
}

func NewEventPublisher(logger logger.Logger, client *eventsource.CloudEventsClient) *EventPublisher {
	return &EventPublisher{
		logger: logger,
		client: client,
	}
}

func (ep *EventPublisher) PublishEvent(ctx context.Context, topic, eventType string, data interface{}) error {
	event := cloudevents.NewEvent()
	event.SetType(eventType)
	event.SetID(uuid.NewV4().String())
	event.SetSource("deltahq-hotel-service")

	if err := event.SetData(cloudevents.ApplicationJSON, data); err != nil {
		ep.logger.Errorf("failed to set data: %v", err)
		return err
	}

	fmt.Println("debug", event)
	sendCtx := cecontext.WithTopic(ctx, topic)

	result := ep.client.Send(sendCtx, event)
	if cloudevents.IsUndelivered(result) {
		ep.logger.Printf("failed to send: %v", result)
		return result
	} else {
		ep.logger.Printf("sent, accepted: %t", cloudevents.IsACK(result))
	}

	return nil
}
