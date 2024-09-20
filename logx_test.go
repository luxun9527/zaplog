package zaplog

import (
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"testing"
)

func TestLogx(t *testing.T) {
	var c Config
	conf.MustLoad("./config.yaml", &c)
	//c.ReportConfig = report.ReportConfig{
	//	Type:     "wx",
	//	Token:    "cd29951d-d5ba-4784-b844-67ffe9fca84e",
	//	ChatID:   0,
	//	FlushSec: 3,
	//}
	InitZapLogger(&c)
	logx.SetWriter(NewZapWriter(DefaultLogger))
	logx.Debugw("this is debug level ", logx.Field("key", "value"))
	logx.Infow("this is info level ", logx.Field("key", "value"))
	//sloww为warn
	logx.Sloww("this is warn level ", logx.Field("key", "value"))
	defer logx.Close()
	//server为panic
	logx.Severe("this is panic level ")

}
