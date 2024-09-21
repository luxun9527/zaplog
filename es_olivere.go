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
	ErrorEsOlivereLogger *esOlivereLogger
	InfoEsOlivereLogger  *esOlivereLogger
)

func init() {
	ErrorEsOlivereLogger = &esOlivereLogger{
		logger: DefaultLogger.With(zap.String("module", EsModuleKey)).Sugar(),
		level:  zapcore.ErrorLevel,
	}
	InfoEsOlivereLogger = &esOlivereLogger{
		logger: DefaultLogger.With(zap.String("module", EsModuleKey)).Sugar(),
		level:  zapcore.ErrorLevel,
	}
}

type esOlivereLogger struct {
	logger *zap.SugaredLogger
	level  zapcore.Level
}

func (esLog *esOlivereLogger) Printf(format string, v ...interface{}) {
	if esLog.level == zapcore.InfoLevel {
		esLog.logger.Infof(format, v...)
	} else {
		esLog.logger.Errorf(format, v...)
	}
}
func (esLog *esOlivereLogger) Update(logger ...*zap.Logger) {
	if len(logger) == 0 {
		esLog.logger = DefaultLogger.With(zap.String("module", EsModuleKey)).Sugar()
		return
	}
	esLog.logger = logger[0].Sugar()
}
