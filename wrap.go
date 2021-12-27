package logger

import (
	"context"
	"fmt"
)

// Logger this wrap instead of "logrus"
type Logger struct{}

func (logger *Logger) Logf(level Level, format string, args ...interface{}) {
	printCtx(context.TODO(), level, fmt.Sprintf(format, args...))
}

func (logger *Logger) Tracef(format string, args ...interface{}) {
	printCtx(context.TODO(), DebugLevel, fmt.Sprintf(format, args...))
}

func (logger *Logger) Debugf(format string, args ...interface{}) {
	printCtx(context.TODO(), DebugLevel, fmt.Sprintf(format, args...))
}

func (logger *Logger) Infof(format string, args ...interface{}) {
	printCtx(context.TODO(), InfoLevel, fmt.Sprintf(format, args...))
}

func (logger *Logger) Printf(format string, args ...interface{}) {
	printCtx(context.TODO(), InfoLevel, fmt.Sprintf(format, args...))
}

func (logger *Logger) Warnf(format string, args ...interface{}) {
	printCtx(context.TODO(), WarnLevel, fmt.Sprintf(format, args...))
}

func (logger *Logger) Warningf(format string, args ...interface{}) {
	printCtx(context.TODO(), WarnLevel, fmt.Sprintf(format, args...))
}

func (logger *Logger) Errorf(format string, args ...interface{}) {
	printCtx(context.TODO(), ErrorLevel, fmt.Sprintf(format, args...))
}

func (logger *Logger) Fatalf(format string, args ...interface{}) {
	printCtx(context.TODO(), FatalLevel, fmt.Sprintf(format, args...))
}

func (logger *Logger) Panicf(format string, args ...interface{}) {
	printCtx(context.TODO(), FatalLevel, fmt.Sprintf(format, args...))
}

func (logger *Logger) Log(level Level, args ...interface{}) {
	printCtx(context.TODO(), level, fmt.Sprint(args...))
}

func (logger *Logger) Trace(args ...interface{}) {
	printCtx(context.TODO(), DebugLevel, fmt.Sprint(args...))
}

func (logger *Logger) Debug(args ...interface{}) {
	printCtx(context.TODO(), DebugLevel, fmt.Sprint(args...))
}

func (logger *Logger) Info(args ...interface{}) {
	printCtx(context.TODO(), InfoLevel, fmt.Sprint(args...))
}

func (logger *Logger) Print(args ...interface{}) {
	printCtx(context.TODO(), InfoLevel, fmt.Sprint(args...))
}

func (logger *Logger) Warn(args ...interface{}) {
	printCtx(context.TODO(), WarnLevel, fmt.Sprint(args...))
}

func (logger *Logger) Warning(args ...interface{}) {
	printCtx(context.TODO(), WarnLevel, fmt.Sprint(args...))
}

func (logger *Logger) Error(args ...interface{}) {
	printCtx(context.TODO(), ErrorLevel, fmt.Sprint(args...))
}

func (logger *Logger) Fatal(args ...interface{}) {
	printCtx(context.TODO(), FatalLevel, fmt.Sprint(args...))
}

func (logger *Logger) Panic(args ...interface{}) {
	printCtx(context.TODO(), FatalLevel, fmt.Sprint(args...))
}

func (logger *Logger) Logln(level Level, args ...interface{}) {
	printCtx(context.TODO(), level, fmt.Sprint(args...))
}

func (logger *Logger) Traceln(args ...interface{}) {
	printCtx(context.TODO(), DebugLevel, fmt.Sprint(args...))
}

func (logger *Logger) Debugln(args ...interface{}) {
	printCtx(context.TODO(), DebugLevel, fmt.Sprint(args...))
}

func (logger *Logger) Infoln(args ...interface{}) {
	printCtx(context.TODO(), InfoLevel, fmt.Sprint(args...))
}

func (logger *Logger) Println(args ...interface{}) {
	printCtx(context.TODO(), InfoLevel, fmt.Sprint(args...))
}

func (logger *Logger) Warnln(args ...interface{}) {
	printCtx(context.TODO(), WarnLevel, fmt.Sprint(args...))
}

func (logger *Logger) Warningln(args ...interface{}) {
	printCtx(context.TODO(), WarnLevel, fmt.Sprint(args...))
}

func (logger *Logger) Errorln(args ...interface{}) {
	printCtx(context.TODO(), ErrorLevel, fmt.Sprint(args...))
}

func (logger *Logger) Fatalln(args ...interface{}) {
	printCtx(context.TODO(), FatalLevel, fmt.Sprint(args...))
}

func (logger *Logger) Panicln(args ...interface{}) {
	printCtx(context.TODO(), FatalLevel, fmt.Sprint(args...))
}
