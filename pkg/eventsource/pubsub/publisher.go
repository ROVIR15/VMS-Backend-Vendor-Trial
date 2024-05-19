package pubsub

import (
	"context"

	"github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/pkg/eventsource"
	esConfig "github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/pkg/eventsource/config"
	esConstants "github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/pkg/eventsource/constants"
	"github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/pkg/logger"
	cloudevents "github.com/cloudevents/sdk-go/v2"
	cecontext "github.com/cloudevents/sdk-go/v2/context"
)

type EventPublisher struct {
	logger       logger.Logger
	client       *eventsource.CloudEventsClient
	source       string
	topicOptions map[string]esConfig.TopicOption
}

var _ eventsource.EventPublisher = (*EventPublisher)(nil)

func NewEventPublisher(logger logger.Logger, client *eventsource.CloudEventsClient, source string, topicOptions map[string]esConfig.TopicOption) *EventPublisher {
	return &EventPublisher{
		logger:       logger,
		client:       client,
		source:       source,
		topicOptions: topicOptions,
	}
}

func (ep *EventPublisher) PublishEvent(ctx context.Context, topic, eventType, eventID string, data interface{}) error {
	event := cloudevents.NewEvent()
	// Set the event context
	event.SetSpecVersion(cloudevents.VersionV1)
	event.SetSource(ep.source)
	event.SetID(eventID)
	event.SetType(esConstants.DeltaHQEventType + eventType)

	// Set the event data
	if err := event.SetData(cloudevents.ApplicationJSON, data); err != nil {
		ep.logger.Errorf("failed to set data: %v", err)
		return err
	}

	eventTopic := ep.getTopic(topic)

	sendCtx := cecontext.WithTopic(ctx, eventTopic)

	result := ep.client.Send(sendCtx, event)
	if cloudevents.IsUndelivered(result) {
		ep.logger.Printf("failed to send event: %v", result)
		return result
	} else {
		ep.logger.Printf("sent, accepted: %t", cloudevents.IsACK(result))
	}

	return nil
}

func (ep *EventPublisher) getTopic(topic string) string {
	if topicOption, ok := ep.topicOptions[topic]; ok {
		return topicOption.String()
	}

	return topic
}
