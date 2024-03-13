package report

import (
	"bufio"
	"github.com/tidwall/pretty"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"sync"
	"time"
)

/*
日志上报，告知异常，具体错误还是通过日志排查
*/

type ImType string

const (
	Tg   ImType = "tg"
	Wx          = "wx"
	Lark        = "lark"
)

// _bufSize 给一个比较大的值，避免write的时候出现flush的情况。
const (
	_bufSize = 1024 * 1024 //1M 内存
)

type ReportConfig struct {
	//上报的类型，目前支持：wx lark tg
	Type string `json:"type"`
	//lark填webhook tg wx填token
	Token string `json:"token"`
	//tg的chatid
	ChatID int64 `json:",optional"`
	//日志刷新的频率 单位秒
	FlushSec int64 `json:",default=3"`
	//最大日志数量
	MaxCount int64 `json:",default=20"`
	//什么级别的日志上报
	Level zap.AtomicLevel `json:"Level"`
}

func NewWriteSyncer(c ReportConfig) zapcore.WriteSyncer {
	var ws zapcore.WriteSyncer
	switch ImType(c.Type) {
	case Wx:
		ws = NewWxWriter(c.Token)
	case Lark:
		ws = NewLarkWriter(c.Token)
	case Tg:
		ws = NewTgWriter(c.Token, c.ChatID)
	default:
		log.Panicf("unsupported report type:%s", c.Type)
	}
	return ws
}

func NewReportWriterBuffer(c ReportConfig) *ReportWriterBuffer {
	ws := NewWriteSyncer(c)
	rwb := &ReportWriterBuffer{
		buf:      bufio.NewWriterSize(ws, _bufSize),
		flushSec: c.FlushSec,
		maxCount: c.MaxCount,
	}
	go rwb.Start()
	return rwb
}

type ReportWriterBuffer struct {
	buf      *bufio.Writer
	count    int64
	flushSec int64
	maxCount int64
	mu       sync.Mutex
}

func (l *ReportWriterBuffer) Start() {
	for {
		time.Sleep(time.Duration(l.flushSec) * time.Second)
		l.Sync()
	}
}

// 这个p会被zap复用，一定要注意。
func (l *ReportWriterBuffer) Write(p []byte) (n int, err error) {
	l.mu.Lock()
	defer l.mu.Unlock()
	data := pretty.Pretty(p)
	l.buf.Write(data)
	l.count++
	if l.count >= l.maxCount {
		l.buf.Flush()
		l.count = 0
	}

	return len(p), nil
}

func (l *ReportWriterBuffer) Sync() error {
	l.mu.Lock()
	defer l.mu.Unlock()
	l.buf.Flush()
	l.count = 0
	return nil
}
