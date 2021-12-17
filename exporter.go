package logger

import (
	"time"

	"github.com/illidaris/core"
	"go.uber.org/zap/zapcore"
)

type IExporter interface {
	Encoder() zapcore.Encoder
	Writer() zapcore.WriteSyncer
	Level() zapcore.Level
}

func NewExporters() []IExporter {
	return []IExporter{
		&FileExporter{},
		&StdExporter{},
	}
}

// fmtEncoder choose format eg: json/console
func fmtEncoder(format string, cfg zapcore.EncoderConfig) zapcore.Encoder {
	switch format {
	case "console":
		return zapcore.NewConsoleEncoder(cfg)
	case "json":
		return zapcore.NewJSONEncoder(cfg)
	default:
		return zapcore.NewConsoleEncoder(cfg)
	}
}

// configEncoder config Encoder
func configEncoder() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:       core.Datetime.String(),
		LevelKey:      core.LevelKey.String(),
		NameKey:       core.NameKey.String(),
		CallerKey:     core.Caller.String(),
		FunctionKey:   core.FunctionKey.String(),
		MessageKey:    core.Message.String(),
		StacktraceKey: core.StacktraceKey.String(),
		LineEnding:    core.LineEnding.String(),
		EncodeLevel:   zapcore.LowercaseLevelEncoder,
		EncodeTime: func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.UTC().Format("2006-01-02T15:04:05.000Z"))
		},
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}
}
