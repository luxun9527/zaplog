对zap日志库的一个的封装，满足开发和生产的需要。

1、开发阶段对idea控制台友好，日志级别彩色，打印行号信息，快速定位到每一行。

2、测试，生产阶段输出json格式对elk等日志收集工具友好。

3、生产阶段能够动态调整日志等级。

4、有报警功能，在指定级别以上的日志，能够发送到企业微信等IM中及时发现问题

5、实现常用中间件sdk的日志输出接口

6、能写到文件中，能够压缩，批量落盘。**(现在k8s基本上都是输出到标准输出中，每个节点一个pod专门收集)**



### 配置文件

```yaml
level: debug #日志等级 debug info warn error
stacktrace: true #默认为true 在error级别及以上显示堆栈
addCaller: true #默认为true增加调用者信息
callerShip: 1 # 默认为3 调用栈深度
mode: console #默认为console 输出到控制台console file
json: false #默认为false是否json格式化
fileName: #可选 file模式参数 输出到指定文件
errorFileName: #可选 file模式参数 错误日志输出到的地方
maxSize: 0 #可选 file模式参数 单文件大小限制 单位MB
maxAge: 0 #可选 file模式参数 文件最大保存时间 单位天
maxBackup: 0 #可选 file模式参数 最大的日志数量
async: false #默认为false file模式参数 是否异步落盘。
compress: false #默认为false file模式参数 是否压缩
console: false #默认为false file模式参数 是否同时输出到控制台
color: true #默认为false输出是否彩色 在开发的时候推荐使用。
port: 34567 #是否开启http热更新日志级别
reportConfig: # 上报配置 warn级别以上报到im工具
  type: lark # 可选 lark(飞书也是这个) wx tg
  token: https://open.feishu.cn/open-apis/bot/v2/hook/71f86ea61212-ab9a23-464512-b40b # lark 飞书填群机器人webhook tg填token wx填key 这个示例地址无效。
  chatID: 0 # tg填chatID 其他不用填
  flushSec: 3 # 刷新间隔单位为秒 开发测试调小一点，生产环境调大一点
  maxCount: 20 #最大缓存数量 达到刷新间隔或最大记录数 触发发送开发测试调小一点，生产环境调大一点
  level: warn # 指定上报级别


```

```go
func TestViperConfig(t *testing.T) {
    v := viper.New()
    v.SetConfigFile("./config.yaml")
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
```

### 1、开发阶段对控制台友好

开发阶段对ide控制台友好，日志级别彩色，打印行号信息快速定位到每一行

```go
func TestDefaultConfig(t *testing.T) {
	Debugf("level %s", "debug")
	Infof("level %s", "info")
	Warnf("level %s", "warn")
	Errorf("level %s", "error")
	Panicf("level %s", "panic")
}

```

```lua
2024-12-14-16:25:13	DEBUG	E:/demoproject/zlog/zap_config_test.go:13	level debug
2024-12-14-16:25:13	INFO	E:/demoproject/zlog/zap_config_test.go:14	level info
2024-12-14-16:25:13	WARN	E:/demoproject/zlog/zap_config_test.go:15	level warn
2024-12-14-16:25:13	ERROR	E:/demoproject/zlog/zap_config_test.go:16	level error
2024-12-14-16:25:13	PANIC	E:/demoproject/zlog/zap_config_test.go:17	level panic
```



### 2、测试，生产阶段输出json格式

2、测试，生产阶段输出json格式，对elk等日志收集工具友好。

```go
InitDefaultLogger(ProdConfig)
ProdConfig.UpdateLevel(zap.DebugLevel)
Debugf("level %s", "debug")
Infof("level %s", "info")
Infof("level %s", "info")
Warnf("level %s", "warn")
Panicf("level %s", "panic")
```

```json
{"level":"debug","time":"2024-12-14-16:08:03","caller":"E:/demoproject/zlog/zap_config_test.go:22","msg":"level debug"}
{"level":"info","time":"2024-12-14-16:08:03","caller":"E:/demoproject/zlog/zap_config_test.go:23","msg":"level info"}
{"level":"info","time":"2024-12-14-16:08:03","caller":"E:/demoproject/zlog/zap_config_test.go:24","msg":"level info"}
{"level":"warn","time":"2024-12-14-16:08:03","caller":"E:/demoproject/zlog/zap_config_test.go:25","msg":"level warn"}
{"level":"panic","time":"2024-12-14-16:08:03","caller":"E:/demoproject/zlog/zap_config_test.go:26","msg":"level panic","stacktrace":"github.com/ikun2021/zlog.TestProdConfig\n\tE:/demoproject/zlog/zap_config_test.go:26\ntesting.tRunner\n\tE:/goroot/src/testing/testing.go:1689"}	
```

### 3、动态调整日志等级

生产阶段能够动态调整日志等级，也是使用zap提供的方法

```go
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
    //动态调整为debug级别
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
```

```json
{"level":"info","time":"2024-12-14-21:52:26","caller":"E:/demoproject/zlog/zaplog.go:78","msg":"log server init success,port:34567"}
{"level":"info","time":"2024-12-14-21:52:26","caller":"E:/demoproject/zlog/zap_config_test.go:36","msg":"level info"}
{"level":"warn","time":"2024-12-14-21:52:26","caller":"E:/demoproject/zlog/zap_config_test.go:37","msg":"level warn"}
{"level":"error","time":"2024-12-14-21:52:26","caller":"E:/demoproject/zlog/zap_config_test.go:38","msg":"level error"}
2024/12/14 21:52:34 update level response map[level:debug]

{"level":"debug","time":"2024-12-14-21:52:34","caller":"E:/demoproject/zlog/zap_config_test.go:56","msg":"level debug"}
{"level":"debug","time":"2024-12-14-21:52:35","caller":"E:/demoproject/zlog/zap_config_test.go:57","msg":"level debug"}
```

4、报警到企业微信等IM中

有报警功能，在指定级别以上的日志，能够发送到企业微信,飞书等IM中及时发现问题

```go
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
```

默认是warn级别的日志往上报。

![img](https://cdn.nlark.com/yuque/0/2024/png/12466223/1734196417421-15f464f8-97cb-40b0-847d-aa29736003fd.png)

### 5、支持中间件sdk日志

实现一些中间件sdk的日志接口，帮助排查问题

```go
//es sdk示例
zlog.DevConfig.UpdateLevel(zapcore.DebugLevel)
esClient, err = elastic.NewClient(
    elastic.SetURL("http://192.168.2.159:9200"),
    elastic.SetBasicAuth("elastic", "123456"),
    elastic.SetSniff(false), // 禁用 Sniffing
    //elastic.SetHealthcheck(false),                                      // 禁用健康检查
    elastic.SetErrorLog(zlog.ErrorEsOlivereLogger), // 启用错误日志
    elastic.SetInfoLog(zlog.InfoEsOlivereLogger),   // 启用信息日志
)
//2024-12-15-00:38:15	DEBUG	E:/gopathdir/pkg/mod/github.com/olivere/elastic/v7@v7.0.32/client.go:847	POST http://192.168.2.159:9200/test-index/_search [status:200, request:0.007s]	{"module": "es"}

```

### 6、输出到文件中

```go
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
}
InitDefaultLogger(DevConfig)
for i := 0; i < 100; i++ {
    for j := 0; j < 10000; j++ {
       Infof("level %s", "info")
       Warnf("level %s", "warn")
       Errorf("level %s", "error")
    }
}
```

![img](https://cdn.nlark.com/yuque/0/2024/png/12466223/1734196509008-4795f6cf-6f9e-44aa-9187-1fe809684463.png)

