package zlog

import (
	"io"

	"go.uber.org/atomic"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"

	"github.com/frank-yf/go-tool/internal/pkg/cjson"
)

var (
	debug = atomic.NewBool(false)
)

func UseDebug() {
	debug.Store(true)
}

// Initial 创建日志实例并替换 zap.L 全局调用
func Initial(options ...zap.Option) {
	l := zap.Must(New(options...))
	zap.ReplaceGlobals(l)
}

// New 创建日志实例
func New(options ...zap.Option) (*zap.Logger, error) {
	cfg := defaultZapConfig()
	return cfg.Build(options...)
}

func defaultZapConfig() (cfg zap.Config) {
	if debug.Load() {
		cfg = zap.NewDevelopmentConfig()
		cfg.EncoderConfig.FunctionKey = "func"     // 打印调用函数
		cfg.EncoderConfig.EncodeTime = timeEncoder // 日志时间戳使用指定格式
	} else {
		cfg = zap.NewProductionConfig()
		cfg.EncoderConfig.TimeKey = ""
		cfg.EncoderConfig.NameKey = ""
		cfg.EncoderConfig.CallerKey = ""
	}

	cfg.Sampling = nil                                               // 禁用采样
	cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder // 日志级别设为大写，使用不同的输出颜色
	cfg.EncoderConfig.EncodeDuration = zapcore.StringDurationEncoder // 使用 time.Duration.String() 输出
	cfg.EncoderConfig.NewReflectedEncoder = jsonEncoder              // 配置json序列化方式

	return
}

var timeEncoder = zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000")

func jsonEncoder(w io.Writer) zapcore.ReflectedEncoder {
	return cjson.NewEncoder(w, cjson.SetEscapeHTML(false))
}
