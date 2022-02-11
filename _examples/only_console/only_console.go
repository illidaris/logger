package main

import "github.com/illidaris/logger"

func main() {
	logger.OnlyConsole()
	logger.Info("123")
}
