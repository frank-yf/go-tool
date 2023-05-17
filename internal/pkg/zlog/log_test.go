package zlog_test

import (
	"go.uber.org/zap"

	"github.com/frank-yf/go-tool/internal/pkg/zlog"
)

func ExampleInitial() {
	zlog.Initial()

	zap.L().Debug("this an debug message") // cannot output
	zap.L().Info("this an info message")

	zap.S().Debug("this an debug message for sugar") // cannot output
	zap.S().Info("this an info message for sugar")

	_ = zap.L().Sync()
	// Output:
	// [34mINFO[0m	this an info message
	// [34mINFO[0m	this an info message for sugar
}

func ExampleUseDebug() {
	zlog.UseDebug()
	zlog.Initial()

	zap.L().Debug("debug message")
	zap.L().Info("info message")

	zap.S().Debug("debug message for sugar")
	zap.S().Info("info message for sugar")

	_ = zap.L().Sync()
	// Output:
	// [35mDEBUG[0m	zlog/log_test.go:28	github.com/frank-yf/go-tool/internal/pkg/zlog_test.ExampleUseDebug	debug message
	// [34mINFO[0m	zlog/log_test.go:29	github.com/frank-yf/go-tool/internal/pkg/zlog_test.ExampleUseDebug	info message
	// [35mDEBUG[0m	zlog/log_test.go:31	github.com/frank-yf/go-tool/internal/pkg/zlog_test.ExampleUseDebug	debug message for sugar
	// [34mINFO[0m	zlog/log_test.go:32	github.com/frank-yf/go-tool/internal/pkg/zlog_test.ExampleUseDebug	info message for sugar
}
