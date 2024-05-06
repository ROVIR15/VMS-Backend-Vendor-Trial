package infrastructure

import (
	"emperror.dev/errors"
	"github.com/ROVIR15/VMS-Backend-Vendor-Trial/pkg/eventsource"
)

func (ic *infrastructureConfigurator) configCloudevents() (*eventsource.CloudEventsClient, error) {
	cloudEventsClient, err := eventsource.NewCloudEventsClient(ic.logger, ic.config.EventSourceOptions)
	if err != nil {
		return nil, errors.WrapIf(err, "eventsource.NewCloudEventsClient")
	}

	ic.logger.Infof("cloudevents service established")

	return cloudEventsClient, nil
}
