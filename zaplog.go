package zaplog

import "go.uber.org/zap"

const (
	EsModuleKey    = "es"
	KafkaModuleKey = "kafka"
	RedisModuleKey = "redis"
)

var (
	logger   *zap.Logger
	sugarLog *zap.SugaredLogger
)

func InitZapLogger(loggerConfig *Config) {
	logger = loggerConfig.Build()
	sugarLog = logger.Sugar()
}
func Debugf(msg string, fields ...interface{}) {
	sugarLog.Debugf(msg, fields...)
}
func Panicf(msg string, fields ...interface{}) {
	sugarLog.Panicf(msg, fields...)
}
func Infof(msg string, fields ...interface{}) {
	sugarLog.Infof(msg, fields...)
}
func Errorf(msg string, fields ...interface{}) {
	sugarLog.Errorf(msg, fields...)
}
func Warnf(msg string, fields ...interface{}) {
	sugarLog.Warnf(msg, fields...)
}
func GetZapLogger() *zap.Logger {
	return logger
}

func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	logger.Panic(msg, fields...)
}

func DPanic(msg string, fields ...zap.Field) {
	logger.DPanic(msg, fields...)
}
func Sync() error {
	return logger.Sync()
}
