package logger

import (
	"context"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/illidaris/core"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func TestNewLogger(t *testing.T) {
	cfg := zapcore.EncoderConfig{
		TimeKey:       core.Datetime.String(),
		LevelKey:      core.LevelKey.String(),
		NameKey:       core.NameKey.String(),
		CallerKey:     core.Caller.String(),
		FunctionKey:   core.FunctionKey.String(),
		MessageKey:    core.Message.String(),
		StacktraceKey: core.StacktraceKey.String(),
		LineEnding:    core.LineEnding.String(),
		EncodeLevel:   nil,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			var cstZone = time.FixedZone("CST", 8*3600)
			enc.AppendString(strconv.FormatInt(t.In(cstZone).Unix(), 10))
		},
		EncodeDuration:   zapcore.MillisDurationEncoder,
		EncodeCaller:     nil,
		ConsoleSeparator: "|",
	}
	NewLogger(&CustomerExporter{
		EncoderConfig: cfg,
	})
	InfoCtx(context.TODO(), "xxx#xxx#xxx", zap.String("", "进入房间"))
}

type CustomerExporter struct {
	zapcore.EncoderConfig
}

func (e *CustomerExporter) Encoder() zapcore.Encoder {
	return zapcore.NewConsoleEncoder(e.EncoderConfig)
}

func (e *CustomerExporter) Writer() zapcore.WriteSyncer {
	return zapcore.Lock(os.Stdout)
}

func (e *CustomerExporter) Level() zapcore.LevelEnabler {
	return DebugLevel.zapLevel()
}
