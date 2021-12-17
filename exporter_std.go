package logger

import (
	"os"

	"go.uber.org/zap/zapcore"
)

// StdExporter std console exporter
type StdExporter struct{}

// Encoder
/**
 * @Description:
 * @receiver e
 * @return zapcore.Encoder
 */
func (e *StdExporter) Encoder() zapcore.Encoder {
	return fmtEncoder(config.StdFormat, configEncoder())
}

// Writer
/**
 * @Description:
 * @receiver e
 * @return zapcore.WriteSyncer
 */
func (e *StdExporter) Writer() zapcore.WriteSyncer {
	return zapcore.Lock(os.Stdout)
}

// Level
/**
 * @Description:
 * @receiver e
 * @return zapcore.Level
 */
func (e *StdExporter) Level() zapcore.Level {
	return config.GetStdLevel().zapLevel()
}
