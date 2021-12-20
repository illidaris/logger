# logger
```go
package main

import (
	"context"
	"github.com/illidaris/logger"
)

func main() {
	logger.OnlyConsole()
	
	ctx:=context.TODO()
	logger.InfoCtx(ctx,"info log")
	logger.Info("info log")
}
```