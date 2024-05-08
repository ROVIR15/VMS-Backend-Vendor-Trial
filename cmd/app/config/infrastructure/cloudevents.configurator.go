package infrastructure

import (
	"emperror.dev/errors"
	"github.com/ILUMINA-Pte-Ltd/PrimeCRM-Backend-Service/pkg/eventsource"
)

func (ic *infrastructureConfigurator) configCloudevents() (*eventsource.CloudEventsClient, error) {
	cloudEventsClient, err := eventsource.NewCloudEventsClient(ic.logger, ic.config.EventSourceOptions)
	if err != nil {
		return nil, errors.WrapIf(err, "eventsource.NewCloudEventsClient")
	}

	ic.logger.Infof("cloudevents service established")

	return cloudEventsClient, nil
}
