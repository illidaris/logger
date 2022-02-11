# logger
该项目是一个基于`go.uber.org/zap`封装的日志框架，便于再开发中快速集成并且使用。
## 快速使用
1. 不生成日志文件，使用console
```go
import "github.com/illidaris/logger"
// 初始化日志管理器
logger.OnlyConsole()
// 打印日志
logger.Info("123")
```
日志样例如下：
```text
2022-02-11T03:45:19.715Z	INFO	only_console/only_console.go:7	123
```
2. 生成日志文件
```go
import "github.com/illidaris/logger"
// 初始化日志管理器
logger.New(nil)
// 打印日志
logger.Info("123")
```
日志样例如下：
```json
{"level":"INFO","datetime":"2022-02-11T03:43:36.048Z","caller":"simple/simple.go:7","message":"123"}
```
## 配置介绍
#### Config结构
```go
type Config struct {
    StdLevel      string `toml:"std_level" json:"std_level"`     // Std Log level.
    StdFormat     string `toml:"std_format" json:"std_format"`   // Std Log format. one of json, text, or console.
    Level         string `toml:"level" json:"level"`             // Log level.
    Format        string `toml:"format" json:"format"`           // Log format. one of json, text, or console.
    FileDirectory string `toml:"file_dir" json:"file_dir"`       // File directory
    FileName      string `toml:"file_name" json:"file_name"`     // Log filename, leave empty to disable file log.
    MaxSize       int    `toml:"max_size" json:"max_size"`       // Max size for a single file, in MB.
    MaxDays       int    `toml:"max_days" json:"max_days"`       // Max log keep days, default is never deleting.
    MaxBackups    int    `toml:"max_backups" json:"max_backups"` // Maximum number of old log files to retain.
    Compress      bool   `toml:"compress" json:"compress"`       // Compress
    EncodeTime    func(time.Time, zapcore.PrimitiveArrayEncoder)
}
```
#### 参数
+ `StdLevel` 终端中输出的日志内容的限制等级（"error","warn","info","debug"）
+ `StdFormat` 终端中输出的日志内容格式（"console","json"）
+ `Level` 日志内容的限制等级（"error","warn","info","debug"）
+ `Format` 日志内容格式（"console","json"）
+ `FileDirectory` 日志输出的文件目标目录
+ `FileName` 文件前缀名
+ `EncodeTime` 时间格式转化函数，默认采用格林威治时间，用于多地区时间统一
```go
func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
    enc.AppendString(t.UTC().Format("2006-01-02T15:04:05.000Z"))
}
```
## 标记字段
目的是为了以后更好的分析日志处理日志，例如在`ELK`中分析日志，在标记加入`session_id`、`trace_id`等等，
方便追踪请求的执行链路，寻求更快的定位问题以及分析代码执行情况。一般我们可以划分横向与纵向两个维度：
+ 横向就是以`session_id`在请求入口植入，将整个请求链路从入口到执行处理到最后返回串联起来，一般采用周期性标记。
+ 纵向就是以`action`或者`step`为分段标记某几个执行函数,一般采用一次性标记。
#### 一次性标记
用于当前日志记录的一次性标记，有效期仅仅再本次标记
```go
package main

import (
	"github.com/illidaris/logger"
	"go.uber.org/zap"
)

func main(){
	logger.OnlyConsole()
	logger.Info("123",zap.String("tag1","tag_value"))
}
```
输出结果：
```text
2022-02-11T06:19:30.114Z	INFO	tag_print_once/tag_print.go:10	123	{"tag1": "tag_value"}
```

#### 周期性标记
借助于`context.context`来实现数据的上下文传递，在ctx的生命周期内，可以多次标记。
```go
package main

import (
	"context"
	"github.com/illidaris/logger"
	"go.uber.org/zap"
)

func main(){
	logger.OnlyConsole()
	ctx:=context.Background()
	ctx=logger.NewContext(ctx,zap.String("tag","tag_value"))
	f1(ctx)
	f2(ctx)
	f3(ctx)
}
func f1(ctx context.Context)  {
	logger.InfoCtx(ctx,"f1 exec....")
}
func f2(ctx context.Context)  {
	logger.InfoCtx(ctx,"f2 exec....")
}
func f3(ctx context.Context)  {
	logger.InfoCtx(ctx,"f3 exec....")
}
```
输出结果：
```text
2022-02-11T06:21:59.242Z	INFO	tag_print_context/tag_print_context.go:19	f1 exec....	{"tag": "tag_value"}
2022-02-11T06:21:59.242Z	INFO	tag_print_context/tag_print_context.go:23	f2 exec....	{"tag": "tag_value"}
2022-02-11T06:21:59.242Z	INFO	tag_print_context/tag_print_context.go:27	f3 exec....	{"tag": "tag_value"}
```