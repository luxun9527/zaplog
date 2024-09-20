package zaplog

import (
	"github.com/luxun9527/zaplog/report"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"log"
	"testing"
)

func TestConsoleJson(t *testing.T) {
	l, _ := zap.ParseAtomicLevel("debug")
	warnLevel, _ := zap.ParseAtomicLevel("warn")
	c := Config{
		Name:          "test",
		Level:         l,
		Stacktrace:    true,
		AddCaller:     true,
		CallerShip:    0,
		Mode:          "console",
		FileName:      "",
		ErrorFileName: "",
		MaxSize:       0,
		MaxAge:        0,
		MaxBackup:     0,
		Async:         false,
		Json:          true,
		Compress:      false,
		IsReport:      false,
		ReportConfig: report.ReportConfig{
			Type:     "lark",
			Token:    "https://open.feishu.cn/open-apis/bot/v2/hook/71f86ea6-ab9a-4645-b40b-1be00ffe910a",
			ChatID:   0,
			FlushSec: 3,
			MaxCount: 20,
			Level:    warnLevel,
		},

		Color:   true,
		options: nil,
	}
	InitZapLogger(&c)
	DefaultSugarLog.Debugf("test level %v", "debug")
	DefaultSugarLog.Infof("info level %v", "info")
	DefaultSugarLog.Warnf("warn level %v", "warn")
	DefaultSugarLog.Errorf("error level %v", "error")
	DefaultSugarLog.Panicf("panic level %v", "panic")
}

func TestConfig(t *testing.T) {
	l, _ := zap.ParseAtomicLevel("debug")
	warnLevel, _ := zap.ParseAtomicLevel("warn")
	c := Config{
		Name:          "test",
		Level:         l,
		Stacktrace:    true,
		AddCaller:     true,
		CallerShip:    0,
		Mode:          "console",
		FileName:      "",
		ErrorFileName: "",
		MaxSize:       0,
		MaxAge:        0,
		MaxBackup:     0,
		Async:         false,
		Json:          true,
		Compress:      false,
		IsReport:      true,
		ReportConfig: report.ReportConfig{
			Type:     "lark",
			Token:    "https://open.feishu.cn/open-apis/bot/v2/hook/71f86ea6-ab9a-4645-b40b-1be00ffe910a",
			ChatID:   0,
			FlushSec: 3,
			MaxCount: 20,
			Level:    warnLevel,
		},

		Color:   false,
		options: nil,
	}
	logger := c.Build()
	logger.Debug("debug level ", zap.Any("test", "[再见]"))
	logger.Info("info level ", zap.Any("test", "test"))

	for i := 0; i < 4; i++ {
		logger.Error("error level ", zap.Any("test", "test1111"), zap.Any("filed", i))
	}
	defer logger.Sync()
	logger.Panic("error level ", zap.Any("test", "test1111"))

}
func TestViperConfig(t *testing.T) {
	v := viper.New()
	v.SetConfigFile("./config.yaml")
	if err := v.ReadInConfig(); err != nil {
		log.Panicf("read config file failed, err:%v\n", err)
	}
	var c Config
	if err := v.Unmarshal(&c, viper.DecodeHook(StringToLogLevelHookFunc())); err != nil {
		log.Panicf("Unmarshal config file failed, err:%v\n", err)
	}
	InitZapLogger(&c)
	Debug("debug level ", zap.Any("test", "test"))
	Info("info level ", zap.Any("test", "test"))
	Warn("warn level ", zap.Any("test", "test"))
	Error("error level ", zap.Any("test", "test"))
	Panic("panic level ", zap.Any("test", "test"))
	Sync()
}
