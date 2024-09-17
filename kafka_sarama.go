package zaplog

import "go.uber.org/zap"

// KafkaLogger sarama.Logger=xx
var (
	KafkaLogger = &kafkaLogger{
		logger: logger.With(zap.String("module", KafkaModuleKey)).Sugar(),
	}
)

type kafkaLogger struct {
	logger *zap.SugaredLogger
}

func (k *kafkaLogger) Print(v ...interface{}) {
	k.logger.Info(v)
}
func (k *kafkaLogger) Printf(format string, v ...interface{}) {
	k.logger.Infof(format, v)
}
func (k *kafkaLogger) Println(v ...interface{}) {
	k.logger.Infoln(v)
}
