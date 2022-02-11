package main

import (
	"github.com/illidaris/logger"
	"go.uber.org/zap"
)

func main() {
	logger.OnlyConsole()
	logger.Info("123", zap.String("tag1", "tag_value"))
}
