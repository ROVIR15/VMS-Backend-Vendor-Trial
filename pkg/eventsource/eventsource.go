package eventsource

import "context"

type EventPublisher interface {
	PublishEvent(ctx context.Context, topic, eventType string, data interface{}) error
}

type EventSubscriber interface {
	SubscribeEvent(ctx context.Context) error
}
