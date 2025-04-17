package main

import (
	"time"

	"go.uber.org/zap"
)

// NewLogger 初始化并返回一个zap.Logger实例，同时处理可能的错误
func NewLogger() (*zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{"./myproject.log"}
	return cfg.Build()
}

// main 函数是程序的入口点，它初始化日志记录器，记录访问URL的日志信息，并处理错误
func main() {
	logger, err := NewLogger()
	if err != nil {
		panic(err)
	}
	su := logger.Sugar()
	defer su.Sync()
	url := "https://www.baidu.com"
	su.Info("start visiting url",
		zap.String("url", url),
		zap.Int("attempt", 3),
		zap.Duration("backoff", time.Second),
	)
}
