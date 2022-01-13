package logger

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Fields map[string]interface{}

func (f Fields) ZapFields() []zapcore.Field {
	fields := make([]zapcore.Field, 0)
	for k, v := range f {
		fields = append(fields, zap.Reflect(k, v))
	}
	return fields
}

// Logger this wrap instead of "logrus"
type Logger struct {
	FieldMap Fields
}

func (logger *Logger) WithField(key string, value interface{}) *Logger {
	return &Logger{
		FieldMap: map[string]interface{}{key: value},
	}
}

func (logger *Logger) WithFields(fields Fields) *Logger {
	return &Logger{
		FieldMap: fields,
	}
}

func (logger *Logger) Logf(level Level, format string, args ...interface{}) {
	printCtx(context.TODO(), level, fmt.Sprintf(format, args...), logger.FieldMap.ZapFields()...)
}

func (logger *Logger) Tracef(format string, args ...interface{}) {
	printCtx(context.TODO(), DebugLevel, fmt.Sprintf(format, args...), logger.FieldMap.ZapFields()...)
}

func (logger *Logger) Debugf(format string, args ...interface{}) {
	printCtx(context.TODO(), DebugLevel, fmt.Sprintf(format, args...), logger.FieldMap.ZapFields()...)
}

func (logger *Logger) Infof(format string, args ...interface{}) {
	printCtx(context.TODO(), InfoLevel, fmt.Sprintf(format, args...), logger.FieldMap.ZapFields()...)
}

func (logger *Logger) Printf(format string, args ...interface{}) {
	printCtx(context.TODO(), InfoLevel, fmt.Sprintf(format, args...), logger.FieldMap.ZapFields()...)
}

func (logger *Logger) Warnf(format string, args ...interface{}) {
	printCtx(context.TODO(), WarnLevel, fmt.Sprintf(format, args...), logger.FieldMap.ZapFields()...)
}

func (logger *Logger) Warningf(format string, args ...interface{}) {
	printCtx(context.TODO(), WarnLevel, fmt.Sprintf(format, args...), logger.FieldMap.ZapFields()...)
}

func (logger *Logger) Errorf(format string, args ...interface{}) {
	printCtx(context.TODO(), ErrorLevel, fmt.Sprintf(format, args...), logger.FieldMap.ZapFields()...)
}

func (logger *Logger) Fatalf(format string, args ...interface{}) {
	printCtx(context.TODO(), FatalLevel, fmt.Sprintf(format, args...), logger.FieldMap.ZapFields()...)
}

func (logger *Logger) Panicf(format string, args ...interface{}) {
	printCtx(context.TODO(), FatalLevel, fmt.Sprintf(format, args...), logger.FieldMap.ZapFields()...)
}

func (logger *Logger) Log(level Level, args ...interface{}) {
	printCtx(context.TODO(), level, fmt.Sprint(args...), logger.FieldMap.ZapFields()...)
}

func (logger *Logger) Trace(args ...interface{}) {
	printCtx(context.TODO(), DebugLevel, fmt.Sprint(args...), logger.FieldMap.ZapFields()...)
}

func (logger *Logger) Debug(args ...interface{}) {
	printCtx(context.TODO(), DebugLevel, fmt.Sprint(args...), logger.FieldMap.ZapFields()...)
}

func (logger *Logger) Info(args ...interface{}) {
	printCtx(context.TODO(), InfoLevel, fmt.Sprint(args...), logger.FieldMap.ZapFields()...)
}

func (logger *Logger) Print(args ...interface{}) {
	printCtx(context.TODO(), InfoLevel, fmt.Sprint(args...), logger.FieldMap.ZapFields()...)
}

func (logger *Logger) Warn(args ...interface{}) {
	printCtx(context.TODO(), WarnLevel, fmt.Sprint(args...), logger.FieldMap.ZapFields()...)
}

func (logger *Logger) Warning(args ...interface{}) {
	printCtx(context.TODO(), WarnLevel, fmt.Sprint(args...), logger.FieldMap.ZapFields()...)
}

func (logger *Logger) Error(args ...interface{}) {
	printCtx(context.TODO(), ErrorLevel, fmt.Sprint(args...), logger.FieldMap.ZapFields()...)
}

func (logger *Logger) Fatal(args ...interface{}) {
	printCtx(context.TODO(), FatalLevel, fmt.Sprint(args...), logger.FieldMap.ZapFields()...)
}

func (logger *Logger) Panic(args ...interface{}) {
	printCtx(context.TODO(), FatalLevel, fmt.Sprint(args...), logger.FieldMap.ZapFields()...)
}

func (logger *Logger) Logln(level Level, args ...interface{}) {
	printCtx(context.TODO(), level, fmt.Sprint(args...), logger.FieldMap.ZapFields()...)
}

func (logger *Logger) Traceln(args ...interface{}) {
	printCtx(context.TODO(), DebugLevel, fmt.Sprint(args...), logger.FieldMap.ZapFields()...)
}

func (logger *Logger) Debugln(args ...interface{}) {
	printCtx(context.TODO(), DebugLevel, fmt.Sprint(args...), logger.FieldMap.ZapFields()...)
}

func (logger *Logger) Infoln(args ...interface{}) {
	printCtx(context.TODO(), InfoLevel, fmt.Sprint(args...), logger.FieldMap.ZapFields()...)
}

func (logger *Logger) Println(args ...interface{}) {
	printCtx(context.TODO(), InfoLevel, fmt.Sprint(args...), logger.FieldMap.ZapFields()...)
}

func (logger *Logger) Warnln(args ...interface{}) {
	printCtx(context.TODO(), WarnLevel, fmt.Sprint(args...), logger.FieldMap.ZapFields()...)
}

func (logger *Logger) Warningln(args ...interface{}) {
	printCtx(context.TODO(), WarnLevel, fmt.Sprint(args...), logger.FieldMap.ZapFields()...)
}

func (logger *Logger) Errorln(args ...interface{}) {
	printCtx(context.TODO(), ErrorLevel, fmt.Sprint(args...), logger.FieldMap.ZapFields()...)
}

func (logger *Logger) Fatalln(args ...interface{}) {
	printCtx(context.TODO(), FatalLevel, fmt.Sprint(args...), logger.FieldMap.ZapFields()...)
}

func (logger *Logger) Panicln(args ...interface{}) {
	printCtx(context.TODO(), FatalLevel, fmt.Sprint(args...), logger.FieldMap.ZapFields()...)
}

func (logger *Logger) LogCtxf(ctx context.Context, level Level, format string, args ...interface{}) {
	printCtx(ctx, level, fmt.Sprintf(format, args...), logger.FieldMap.ZapFields()...)
}

func (logger *Logger) TraceCtxf(ctx context.Context, format string, args ...interface{}) {
	printCtx(ctx, DebugLevel, fmt.Sprintf(format, args...), logger.FieldMap.ZapFields()...)
}

func (logger *Logger) DebugCtxf(ctx context.Context, format string, args ...interface{}) {
	printCtx(ctx, DebugLevel, fmt.Sprintf(format, args...), logger.FieldMap.ZapFields()...)
}

func (logger *Logger) InfoCtxf(ctx context.Context, format string, args ...interface{}) {
	printCtx(ctx, InfoLevel, fmt.Sprintf(format, args...), logger.FieldMap.ZapFields()...)
}

func (logger *Logger) PrintCtxf(ctx context.Context, format string, args ...interface{}) {
	printCtx(ctx, InfoLevel, fmt.Sprintf(format, args...), logger.FieldMap.ZapFields()...)
}

func (logger *Logger) WarnCtxf(ctx context.Context, format string, args ...interface{}) {
	printCtx(ctx, WarnLevel, fmt.Sprintf(format, args...), logger.FieldMap.ZapFields()...)
}

func (logger *Logger) WarningCtxf(ctx context.Context, format string, args ...interface{}) {
	printCtx(ctx, WarnLevel, fmt.Sprintf(format, args...), logger.FieldMap.ZapFields()...)
}

func (logger *Logger) ErrorCtxf(ctx context.Context, format string, args ...interface{}) {
	printCtx(ctx, ErrorLevel, fmt.Sprintf(format, args...), logger.FieldMap.ZapFields()...)
}

func (logger *Logger) FatalCtxf(ctx context.Context, format string, args ...interface{}) {
	printCtx(ctx, FatalLevel, fmt.Sprintf(format, args...), logger.FieldMap.ZapFields()...)
}

func (logger *Logger) PanicCtxf(ctx context.Context, format string, args ...interface{}) {
	printCtx(ctx, FatalLevel, fmt.Sprintf(format, args...), logger.FieldMap.ZapFields()...)
}
