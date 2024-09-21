package zaplog

import (
	"context"
	"go.uber.org/zap"
)

var (
	RedisLogger *redisLogger
)

func init() {
	RedisLogger = &redisLogger{
		logger: DefaultLogger.With(zap.String("module", RedisModuleKey)).Sugar(),
	}
}

type redisLogger struct {
	logger *zap.SugaredLogger
}

func (rl *redisLogger) Printf(ctx context.Context, format string, v ...interface{}) {
	rl.logger.Infof(format, v...)
}
func (rl *redisLogger) Update(logger ...*zap.Logger) {
	if len(logger) == 0 {
		rl.logger = DefaultLogger.With(zap.String("module", RedisModuleKey)).Sugar()
		return
	}
	rl.logger = logger[0].Sugar()
}
