package zaplog

import (
	"github.com/luxun9527/zaplog/report"
	"go.uber.org/zap"
	"testing"
)

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
			Token:    "",
			ChatID:   0,
			FlushSec: 3,
			MaxCount: 20,
			Level:    warnLevel,
		},
		//ReportConfig: report.ReportConfig{
		//	Type:     "tg",
		//	Token:    "6499740288:AAEGZhWULZWto9gjlgxnqwQg1KxVKeJc0Ao",
		//	ChatID:   -1001980672871,
		//	FlushSec: 3,
		//},
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
