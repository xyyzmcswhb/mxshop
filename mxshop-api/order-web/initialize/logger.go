package initialize

import "go.uber.org/zap"

// 日志的初始化
func InitLogger() {
	logger, _ := zap.NewDevelopment()
	zap.ReplaceGlobals(logger)
}
