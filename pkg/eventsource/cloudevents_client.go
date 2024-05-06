package eventsource

import (
	"context"

	"github.com/ROVIR15/VMS-Backend-Vendor-Trial/pkg/eventsource/config"
	"github.com/ROVIR15/VMS-Backend-Vendor-Trial/pkg/logger"
	cepubsub "github.com/cloudevents/sdk-go/protocol/pubsub/v2"
	cloudevents "github.com/cloudevents/sdk-go/v2"
)

type CloudEventsClient struct {
	cloudevents.Client
	logger logger.Logger
}

func NewCloudEventsClient(logger logger.Logger, options *config.EventSourceOptions) (*CloudEventsClient, error) {

	ceClient := &CloudEventsClient{logger: logger}

	var opts []cepubsub.Option
	opts = append(opts, cepubsub.WithProjectID(options.GCPProjectID))
	for _, subscription := range options.SubscriptionOptions {
		opts = append(opts, cepubsub.WithSubscriptionID(subscription.String()))
	}

	t, err := cepubsub.New(context.Background(), opts...)
	if err != nil {
		ceClient.logger.Fatalf("failed to create cloudevents pubsub transport, %s", err.Error())
		return nil, err
	}

	c, err := cloudevents.NewClient(t)
	if err != nil {
		ceClient.logger.Fatalf("failed to create cloudevents client, %s", err.Error())
		return nil, err
	}

	ceClient.Client = c

	return ceClient, nil
}
