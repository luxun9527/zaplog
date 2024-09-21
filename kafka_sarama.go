package zaplog

import "go.uber.org/zap"

// KafkaSaramaLogger sarama.Logger=xx
var (
	KafkaSaramaLogger *kafkaSaramaLogger
)

func init() {
	KafkaSaramaLogger = &kafkaSaramaLogger{
		logger: DefaultLogger.With(zap.String("module", KafkaModuleKey)).Sugar(),
	}
}

type kafkaSaramaLogger struct {
	logger *zap.SugaredLogger
}

func (k *kafkaSaramaLogger) Print(v ...interface{}) {
	k.logger.Info(v)
}
func (k *kafkaSaramaLogger) Printf(format string, v ...interface{}) {
	k.logger.Infof(format, v)
}
func (k *kafkaSaramaLogger) Println(v ...interface{}) {
	k.logger.Infoln(v)
}
func (k *kafkaSaramaLogger) Update(logger ...*zap.Logger) {
	if len(logger) == 0 {
		k.logger = DefaultLogger.With(zap.String("module", KafkaModuleKey)).Sugar()
		return
	}
	k.logger = logger[0].Sugar()
}
