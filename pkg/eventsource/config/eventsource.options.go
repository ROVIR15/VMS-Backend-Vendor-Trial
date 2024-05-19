package config

type EventSourceOptions struct {
	GCPProjectID        string                 `mapstructure:"gcpProjectID"`
	Source              string                 `mapstructure:"source"`
	TopicOptions        map[string]TopicOption `mapstructure:"topics"`
	SubscriptionOptions []SubscriptionOption   `mapstructure:"subscriptions"`
}

type SubscriptionOption string

func (e *SubscriptionOption) String() string {
	return string(*e)
}

type TopicOption string

func (e *TopicOption) String() string {
	return string(*e)
}
