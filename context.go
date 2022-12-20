package logger

import (
	"context"
	"fmt"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// NewContext
/**
 * @Description:
 * @param ctx
 * @param fields
 * @return context.Context
 */
func NewContext(ctx context.Context, fields ...zapcore.Field) context.Context {
	if logInstance := WithContext(ctx); logInstance == nil {
		fmt.Println("NewContext() failed")
		return ctx
	}
	if logInstance := WithContext(ctx).With(fields...); logInstance == nil {
		fmt.Println("NewContext() failed")
		return ctx
	}
	return context.WithValue(ctx, CtxLogger, WithContext(ctx).With(fields...))
}

// WithContext
/**
 * @Description:
 * @param ctx
 * @return *zap.Logger
 */
func WithContext(ctx context.Context) *zap.Logger {
	if ctx == nil {
		return ctxLogger
	}
	if l, ok := ctx.Value(CtxLogger).(*zap.Logger); ok {
		return l
	}
	return ctxLogger
}

// debugCtxFunc
/**
 * @Description:
 * @param ctx
 * @return func(msg string, fields ...zapcore.Field)
 */
func debugCtxFunc(ctx context.Context) func(msg string, fields ...zapcore.Field) {
	if log := WithContext(ctx); log != nil {
		return log.Debug
	}
	return DefaultPrint
}

// infoCtxFunc
/**
 * @Description:
 * @param ctx
 * @return func(msg string, fields ...zapcore.Field)
 */
func infoCtxFunc(ctx context.Context) func(msg string, fields ...zapcore.Field) {
	if log := WithContext(ctx); log != nil {
		return log.Info
	}
	return DefaultPrint
}

// warnCtxFunc
/**
 * @Description:
 * @param ctx
 * @return func(msg string, fields ...zapcore.Field)
 */
func warnCtxFunc(ctx context.Context) func(msg string, fields ...zapcore.Field) {
	if log := WithContext(ctx); log != nil {
		return log.Warn
	}
	return DefaultPrint
}

// errorCtxFunc
/**
 * @Description:
 * @param ctx
 * @return func(msg string, fields ...zapcore.Field)
 */
func errorCtxFunc(ctx context.Context) func(msg string, fields ...zapcore.Field) {
	if log := WithContext(ctx); log != nil {
		return log.Error
	}
	return DefaultPrint
}

// criticalCtxFunc
/**
 * @Description:
 * @param ctx
 * @return func(msg string, fields ...zapcore.Field)
 */
func criticalCtxFunc(ctx context.Context) func(msg string, fields ...zapcore.Field) {
	if log := WithContext(ctx); log != nil {
		return log.Fatal
	}
	return DefaultPrint
}

// printCtx
/**
 * @Description:
 * @param ctx
 * @param level
 * @param msg
 * @param fields
 */
func printCtx(ctx context.Context, level Level, msg string, fields ...zapcore.Field) {
	if f, ok := zapPrints[level]; ok {
		f(ctx)(msg, fields...)
	} else {
		println("map can not found log func")
	}
}

func DefaultPrint(msg string, fields ...zapcore.Field) {
	fmt.Printf("%s,%v \n", msg, fields)
}

func LogFrmCtx(ctx context.Context, level Level, msg string, fields ...zapcore.Field) {
	if f, ok := zapPrints[level]; ok {
		f(ctx)(msg, fields...)
	} else {
		println("map can not found log func")
	}
}
