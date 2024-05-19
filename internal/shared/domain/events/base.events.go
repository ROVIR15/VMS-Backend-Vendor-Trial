package events

import "github.com/google/uuid"

type BaseEvent struct {
	topic     string
	eventType string `default:"com.deltahq.event.v1.message.published"`
	id        string
}

var _ Event = (*BaseEvent)(nil)

func NewBaseEvent(topic, eventType string) *BaseEvent {
	return &BaseEvent{
		topic:     topic,
		eventType: eventType,
		id:        uuid.New().String(),
	}
}

func (e *BaseEvent) Topic() string {
	return e.topic
}

func (e *BaseEvent) EventType() string {
	return e.eventType
}

func (e *BaseEvent) ID() string {
	return e.id
}
