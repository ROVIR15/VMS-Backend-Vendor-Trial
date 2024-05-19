package empty

import (
	"github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/pkg/logger"
	"github.com/ILUMINA-Pte-Ltd/VMS-Backend-Vendor-Trial/pkg/logger/models"
)

type emptyLogger struct{}

// emptyLogger implements logger.Logger interface
var _ logger.Logger = (*emptyLogger)(nil)

func (e emptyLogger) Configure(cfg func(internalLog interface{})) {
}

func (e emptyLogger) Debug(args ...interface{}) {
}

func (e emptyLogger) Debugf(template string, args ...interface{}) {
}

func (e emptyLogger) Debugw(msg string, fields logger.Fields) {
}

func (e emptyLogger) LogType() models.LogType {
	return models.Zap
}

func (e emptyLogger) Info(args ...interface{}) {
}

func (e emptyLogger) Infof(template string, args ...interface{}) {
}

func (e emptyLogger) Infow(msg string, fields logger.Fields) {
}

func (e emptyLogger) Warn(args ...interface{}) {
}

func (e emptyLogger) Warnf(template string, args ...interface{}) {
}

func (e emptyLogger) WarnMsg(msg string, err error) {
}

func (e emptyLogger) Error(args ...interface{}) {
}

func (e emptyLogger) Errorw(msg string, fields logger.Fields) {
}

func (e emptyLogger) Errorf(template string, args ...interface{}) {
}

func (e emptyLogger) Err(msg string, err error) {
}

func (e emptyLogger) Fatal(args ...interface{}) {
}

func (e emptyLogger) Fatalf(template string, args ...interface{}) {
}

func (e emptyLogger) Printf(template string, args ...interface{}) {
}

func (e emptyLogger) WithName(name string) {
}
