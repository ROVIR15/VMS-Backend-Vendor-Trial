package config

type EventSourceOptions struct {
	GCPProjectID        string               `mapstructure:"gcpProjectID"`
	TopicOptions        []TopicOption        `mapstructure:"topics"`
	SubscriptionOptions []SubscriptionOption `mapstructure:"subscriptions"`
}

type SubscriptionOption string

func (e *SubscriptionOption) String() string {
	return string(*e)
}

type TopicOption string

func (e *TopicOption) String() string {
	return string(*e)
}
