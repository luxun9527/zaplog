package zaplog

import (
	"context"
	"go.uber.org/zap"
)

var (
	RedisLogger = &redisLogger{
		logger: logger.With(zap.String("module", RedisModuleKey)).Sugar(),
	}
)

type redisLogger struct {
	logger *zap.SugaredLogger
}

func (esLog *redisLogger) Printf(ctx context.Context, format string, v ...interface{}) {
	esLog.logger.Infof(format, v...)
}
