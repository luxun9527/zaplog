package zaplog

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

var (
	GinOutPut *ginOutput
)

func init() {
	GinOutPut = &ginOutput{
		logger: DefaultLogger.With(zap.String("module", GinModuleKey)).WithOptions(zap.AddCallerSkip(1)),
	}
}
func GetGinLogger(conf ...gin.LoggerConfig) gin.HandlerFunc {
	if len(conf) == 0 {
		return gin.LoggerWithConfig(gin.LoggerConfig{
			Formatter: LogFormatter,
			Output:    GinOutPut,
		})
	}
	return gin.LoggerWithConfig(conf[0])
}

var LogFormatter = func(param gin.LogFormatterParams) string {
	if param.Latency > time.Minute {
		param.Latency = param.Latency.Truncate(time.Second)
	}
	h := gin.H{
		"statusCode": param.StatusCode,
		"path":       param.Path,
		"method":     param.Method,
		"latency":    param.Latency,
		"clientIp":   param.ClientIP,
	}
	data, _ := json.Marshal(h)
	return string(data)
}

type ginOutput struct {
	logger *zap.Logger
}

func (gl *ginOutput) Write(p []byte) (n int, err error) {
	param := gin.H{}
	if err := json.Unmarshal(p, &param); err != nil {
		return 0, err
	}
	gl.logger.Info("",
		Any("statusCode", param["statusCode"]),
		Any("path", param["path"]),
		Any("method", param["method"]),
		Any("latency", param["latency"]),
		Any("clientIp", param["clientIp"]),
	)
	return len(p), nil
}
func (g *ginOutput) Update(logger ...*zap.Logger) {
	if len(logger) == 0 {
		g.logger = DefaultLogger.With(zap.String("module", GinModuleKey)).WithOptions(zap.AddCallerSkip(1))
		return
	}
	g.logger = logger[0]
}
