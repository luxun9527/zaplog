package zlog

import (
	"encoding/json"
	"github.com/ikun2021/zlog/report"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"io"
	"log"
	"net/http"
	"testing"
	"time"
)

func TestDefaultConfig(t *testing.T) {
	DevConfig.UpdateLevel(zap.DebugLevel)
	Debugf("level %s", "debug")
	Infof("level %s", "info")
	Warnf("level %s", "warn")
	Errorf("level %s", "error")
	Panicf("level %s", "panic")
}
func TestProdConfig(t *testing.T) {
	InitDefaultLogger(ProdConfig)
	ProdConfig.UpdateLevel(zap.DebugLevel)
	Debugf("level %s", "debug")
	Infof("level %s", "info")
	Warnf("level %s", "warn")
	Errorf("level %s", "error")
	Panicf("level %s", "panic")
}
func TestUpdateLogger(t *testing.T) {
	ProdConfig.Port = 34567
	InitDefaultLogger(ProdConfig)
	//默认是info
	Debugf("level %s", "debug")
	Infof("level %s", "info")
	Warnf("level %s", "warn")
	Errorf("level %s", "error")
	//	curl -X PUT localhost:8080/log/level -d level=debug
	// 支持 form和json格式
	req, _ := http.NewRequest(http.MethodPut, "http://localhost:34567/log/level?level=debug", nil)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Panicf("update level error %s", err.Error())
	}
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Panicf("read response error %s", err.Error())
	}
	defer resp.Body.Close()
	var m = map[string]interface{}{}
	if err := json.Unmarshal(data, &m); err != nil {
		log.Panicf("unmarshal response error %s", err.Error())
	}
	log.Printf("update level response %v", m)
	Debugf("level %s", "debug")
	Debugf("level %s", "debug")
}
func TestReportLogger(t *testing.T) {
	ProdConfig.ReportConfig = &report.ReportConfig{
		Type:  "lark",
		Token: "https://open.feishu.cn/open-apis/bot/v2/hook/71fxxxxxxxxxxxxxxxxxxxxxxxxxxx",
	}
	InitDefaultLogger(ProdConfig)
	//默认是info
	Debugf("level %s", "debug")
	Infof("level %s", "info")
	Warnf("level %s", "warn")
	Errorf("level %s", "error")
	Panicf("level %s", "panic")

}

func TestCtxLog(t *testing.T) {
	ctx := InitTrace()
	DebugfCtx(ctx, "level %s", "debug")                 //level info traceId 355dca29445483687de01f8b31a375f9
	InfofCtx(ctx, "level %s", "info")                   //level warn traceId 355dca29445483687de01f8b31a375f9
	WarnfCtx(ctx, "level %s", "warn")                   //level error traceId 355dca29445483687de01f8b31a375f9
	ErrorfCtx(ctx, "level %s", "error")                 //level error traceId 355dca29445483687de01f8b31a375f9
	DebugCtx(ctx, "level ", zap.String("key", "value")) //level debug traceId 355dca29445483687de01f8b31a375f9
	InfoCtx(ctx, "level ", zap.String("key", "value"))  //level info traceId 355dca29445483687de01f8b31a375f9
	WarnCtx(ctx, "level ", zap.String("key", "value"))  //level 	{"key": "value", "traceId": "355dca29445483687de01f8b31a375f9"}
	ErrorCtx(ctx, "level ", zap.String("key", "value")) //level 	{"key": "value", "traceId": "355dca29445483687de01f8b31a375f9"}
}

func TestFile(t *testing.T) {
	DevConfig = &Config{
		Level:      zap.NewAtomicLevelAt(zap.InfoLevel),
		Stacktrace: false,
		AddCaller:  true,
		CallerShip: 1,
		Mode:       FileMode,
		FileName:   "./log/test.log",
		//ErrorFileName: "./log/err.log",
		MaxSize:   1,
		MaxAge:    0,
		MaxBackup: 5,
		Json:      true,
	}
	InitDefaultLogger(DevConfig)
	for i := 0; i < 100; i++ {
		Infof("level %s", "info")
		Warnf("level %s", "warn")
		Errorf("level %s", "error")

	}
	//for i := 0; i < 100; i++ {
	//	for j := 0; j < 10000; j++ {
	//		Infof("level %s", "info")
	//		Warnf("level %s", "warn")
	//		Errorf("level %s", "error")
	//	}
	//}
}

func TestViperConfig(t *testing.T) {
	v := viper.New()
	v.SetConfigFile("./config.toml")
	if err := v.ReadInConfig(); err != nil {
		log.Panicf("read config file failed, err:%v", err)
	}
	var c Config
	if err := v.Unmarshal(&c, viper.DecodeHook(StringToLogLevelHookFunc())); err != nil {
		log.Panicf("Unmarshal config file failed, err:%v", err)
	}

	InitDefaultLogger(&c)
	Debug("debug level ", String("test", "test"))
	Info("info level ", Duration("Duration", time.Second))
	Warn("warn level ", Any("test", "test"))
	Error("error level ", Any("test", "test"))
	Panic("panic level ", Any("test", "test"))
	Sync()
}
