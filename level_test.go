package logger

import (
	"context"
	"errors"
	"go.uber.org/zap/zapcore"
	"testing"
)

func TestLevel_CtxPrintf(t *testing.T) {
	OnlyConsole()
	ctx := context.Background()
	ctx = NewContext(ctx, []zapcore.Field{
		{Key: "test1", Type: zapcore.StringType, String: "test123456"},
	}...)
	ctx = NewContext(ctx, []zapcore.Field{
		{Key: "test2", Type: zapcore.StringType, String: "test1234567"},
	}...)
	DebugLevel.CtxPrintf(ctx, "a is  %s", errors.New("zui le"))
	InfoLevel.CtxPrintf(ctx, "a is  %s", errors.New("zui le"))
	WarnLevel.CtxPrintf(ctx, "a is  %s", errors.New("zui le"))
	ErrorLevel.CtxPrintf(ctx, "a is  %s", errors.New("zui le"))
	FatalLevel.CtxPrintf(ctx, "a is  %s", errors.New("zui le"))
}
