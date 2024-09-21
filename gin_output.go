package zaplog

import (
	"go.uber.org/zap"
)

var (
	GinOutPut *ginOutput
)

func init() {
	GinOutPut = &ginOutput{
		logger: DefaultLogger.With(zap.String("module", GinModuleKey)).Sugar(),
	}
}

type ginOutput struct {
	logger *zap.SugaredLogger
}

func (gl *ginOutput) Write(p []byte) (n int, err error) {
	gl.logger.Info(string(p))
	return len(p), nil
}
