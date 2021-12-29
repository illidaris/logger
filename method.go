package logger

import (
	"context"
	"go.uber.org/zap/zapcore"
)

// Debug
/**
 * @Description:
 * @param msg
 */
func Debug(msg string, fields ...zapcore.Field) {
	printCtx(context.TODO(), DebugLevel, msg, fields...)
}

// DebugCtx
/**
 * @Description:
 * @param ctx
 * @param msg
 */
func DebugCtx(ctx context.Context, msg string, fields ...zapcore.Field) {
	printCtx(ctx, DebugLevel, msg, fields...)
}

// Info
/**
 * @Description:
 * @param msg
 */
func Info(msg string, fields ...zapcore.Field) {
	printCtx(context.TODO(), InfoLevel, msg, fields...)
}

// InfoCtx
/**
 * @Description:
 * @param ctx
 * @param msg
 */
func InfoCtx(ctx context.Context, msg string, fields ...zapcore.Field) {
	printCtx(ctx, InfoLevel, msg, fields...)
}

// Warn
/**
 * @Description:
 * @param msg
 */
func Warn(msg string, fields ...zapcore.Field) {
	printCtx(context.TODO(), WarnLevel, msg, fields...)
}

// WarnCtx
/**
 * @Description:
 * @param ctx
 * @param msg
 */
func WarnCtx(ctx context.Context, msg string, fields ...zapcore.Field) {
	printCtx(ctx, WarnLevel, msg, fields...)
}

// Error
/**
 * @Description:
 * @param msg
 */
func Error(msg string, fields ...zapcore.Field) {
	printCtx(context.TODO(), ErrorLevel, msg, fields...)
}

// ErrorCtx
/**
 * @Description:
 * @param ctx
 * @param msg
 */
func ErrorCtx(ctx context.Context, msg string, fields ...zapcore.Field) {
	printCtx(ctx, ErrorLevel, msg, fields...)
}

// Fatal
/**
 * @Description:
 * @param msg
 */
func Fatal(msg string, fields ...zapcore.Field) {
	printCtx(context.TODO(), FatalLevel, msg, fields...)
}

// FatalCtx
/**
 * @Description:
 * @param ctx
 * @param msg
 */
func FatalCtx(ctx context.Context, msg string, fields ...zapcore.Field) {
	printCtx(ctx, FatalLevel, msg, fields...)
}
