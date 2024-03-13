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

func Info(msg string, fields ...zap.Field) {
	l.Info(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	l.Debug(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	l.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	l.Error(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	l.Panic(msg, fields...)
}

func DPanic(msg string, fields ...zap.Field) {
	l.DPanic(msg, fields...)
}
func Sync() error {
	return l.Sync()
}
