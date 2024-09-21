package zaplog

import (
	"go.uber.org/zap"
)

var (
	// String ...
	String = zap.String
	// Any ...
	Any = zap.Any
	// Int64 ...
	Int64 = zap.Int64
	// Int ...
	Int = zap.Int
	// Int32 ...
	Int32 = zap.Int32
	// Uint ...
	Uint = zap.Uint
	// Duration ...
	Duration = zap.Duration
	// Durationp ...
	Durationp = zap.Durationp
	// Object ...
	Object = zap.Object
	// Namespace ...
	Namespace = zap.Namespace
	// Reflect ...
	Reflect = zap.Reflect
	// Skip ...
	Skip = zap.Skip()
	// ByteString ...
	ByteString = zap.ByteString
)

const (
	EsModuleKey    = "es"
	KafkaModuleKey = "kafka"
	RedisModuleKey = "redis"
	GinModuleKey   = "gin"
)

const (
	ConsoleMode = "console"
	FileMode    = "file"
)

var (
	DefaultLogger   *zap.Logger
	DefaultSugarLog *zap.SugaredLogger
)

func init() {
	InitZapLogger(&Config{
		Level:      zap.NewAtomicLevelAt(zap.InfoLevel),
		AddCaller:  true,
		CallerShip: 3,
		Mode:       ConsoleMode,
		Color:      true,
	})
}

func InitZapLogger(loggerConfig *Config) {
	DefaultLogger = loggerConfig.Build()
	DefaultSugarLog = DefaultLogger.Sugar()
}
func Debugf(msg string, fields ...interface{}) {
	DefaultSugarLog.Debugf(msg, fields...)
}
func Panicf(msg string, fields ...interface{}) {
	DefaultSugarLog.Panicf(msg, fields...)
}
func Infof(msg string, fields ...interface{}) {
	DefaultSugarLog.Infof(msg, fields...)
}
func Errorf(msg string, fields ...interface{}) {
	DefaultSugarLog.Errorf(msg, fields...)
}
func Warnf(msg string, fields ...interface{}) {
	DefaultSugarLog.Warnf(msg, fields...)
}
func GetZapLogger() *zap.Logger {
	return DefaultLogger
}

func Info(msg string, fields ...zap.Field) {
	DefaultLogger.Info(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	DefaultLogger.Debug(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	DefaultLogger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	DefaultLogger.Error(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	DefaultLogger.Panic(msg, fields...)
}

func DPanic(msg string, fields ...zap.Field) {
	DefaultLogger.DPanic(msg, fields...)
}
func Sync() error {
	return DefaultLogger.Sync()
}
