package main

import (
	"context"
	"github.com/illidaris/logger"
	"go.uber.org/zap"
)

func main() {
	logger.OnlyConsole()
	ctx := context.Background()
	ctx = logger.NewContext(ctx, zap.String("tag", "tag_value"))
	f1(ctx)
	f2(ctx)
	f3(ctx)
}

func f1(ctx context.Context) {
	logger.InfoCtx(ctx, "f1 exec....")
}

func f2(ctx context.Context) {
	logger.InfoCtx(ctx, "f2 exec....")
}

func f3(ctx context.Context) {
	logger.InfoCtx(ctx, "f3 exec....")
}
