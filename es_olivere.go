package zaplog

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

/*
elastic.SetErrorLog(EsErrorLog), // 启用错误日志
elastic.SetInfoLog(EsInfoLog),  // 启用信息日志
*/

var (
	ErrorEsLogger = &esLogger{
		logger: DefaultLogger.With(zap.String("module", EsModuleKey)).Sugar(),
		level:  zapcore.ErrorLevel,
	}
	InfoEsLogger = &esLogger{
		logger: DefaultLogger.With(zap.String("module", EsModuleKey)).Sugar(),
		level:  zapcore.ErrorLevel,
	}
)

type esLogger struct {
	logger *zap.SugaredLogger
	level  zapcore.Level
}

func (esLog *esLogger) Printf(format string, v ...interface{}) {
	if esLog.level == zapcore.InfoLevel {
		esLog.logger.Infof(format, v...)
	} else {
		esLog.logger.Errorf(format, v...)
	}
}
