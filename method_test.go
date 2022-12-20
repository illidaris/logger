package logger

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/spf13/cast"
	"go.uber.org/zap/zapcore"
)

func TestDebug_Func(t *testing.T) {
	methodProxy(func() {
		Debug("debug")
	})
}

func BenchmarkDebug(b *testing.B) {
	num := 10
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		out := fmt.Sprintf("%d", num)
		methodProxy(func() {
			Debug(out)
		})
	}
}

func ExampleDebug() {
	Debug("example")
}

func TestDebugPanicCtx(t *testing.T) {
	//OnlyConsole()
	methodProxyWithCtx(func(ctx context.Context) {
		DebugCtx(ctx, "debug with ctx")
		// sub, cancel := context.WithTimeout(ctx, time.Second*2)
		// time.Sleep(time.Second * 3)
		// defer cancel()
		go func(c context.Context) {
			defer func() {
				if p := recover(); p != nil {
					DebugCtx(c, "xxx")
				}
			}()
			a := []int{1, 2}
			time.Sleep(time.Second * 2)
			DebugCtx(c, "g run~~~")
			time.Sleep(time.Second * 1)
			DebugCtx(c, cast.ToString(a[5]))
		}(context.TODO())
		<-time.After(time.Second * 5)
		DebugCtx(ctx, "over")
	})
}

func TestDebugCtx(t *testing.T) {
	methodProxyWithCtx(func(ctx context.Context) {
		DebugCtx(ctx, "debug with ctx")
	})
}

func BenchmarkDebugCtx(b *testing.B) {
	num := 10
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		out := fmt.Sprintf("%d", num)
		methodProxyWithCtx(func(ctx context.Context) {
			DebugCtx(ctx, out)
		})
	}
}

func ExampleDebugCtx() {
	DebugCtx(context.TODO(), "example")
}

func methodProxyWithCtx(f func(ctx context.Context)) {
	//OnlyConsole()
	ctx := context.Background()
	ctx = NewContext(ctx, []zapcore.Field{
		{Key: "test1", Type: zapcore.StringType, String: "test123456"},
	}...)
	ctx = NewContext(ctx, []zapcore.Field{
		{Key: "test2", Type: zapcore.StringType, String: "test1234567"},
	}...)
	f(ctx)
}

func methodProxy(f func()) {
	OnlyConsole()
	f()
}
