package main

import "github.com/illidaris/logger"

func main() {
	logger.New(nil)
	logger.Info("123")
	logger.Error("XXX")
	logger.Warn("AAAA")
	logger.Debug("HHH")
}
