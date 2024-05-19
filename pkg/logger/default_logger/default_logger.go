package defaultLogger

import (
	"os"

	"github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/pkg/constants"
	"github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/pkg/logger"
	"github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/pkg/logger/config"
	"github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/pkg/logger/logrous"
	"github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/pkg/logger/models"
	"github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/pkg/logger/zap"
)

var Logger logger.Logger

func SetupDefaultLogger() {
	logType := os.Getenv("LOGGER_TYPE")

	switch logType {
	case "logrus":
		Logger = logrous.NewLogrusLogger(
			&config.LogOptions{LogType: models.Logrus, CallerEnabled: false},
			constants.Dev,
		)
	default:
		Logger = zap.NewZapLogger(
			&config.LogOptions{LogType: models.Zap, CallerEnabled: false},
			constants.Dev,
		)
	}
}
