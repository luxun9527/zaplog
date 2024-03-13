package zaplog

import "go.uber.org/zap"

var (
	l *zap.Logger
)

func InitZapLogger(loggerConfig *Config) {
	l = loggerConfig.Build()
}

func GetZapLogger() *zap.Logger {
	return l
}
