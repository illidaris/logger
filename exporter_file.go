package logger

import (
	"log"
	"path"
	"time"

	rotateLogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap/zapcore"
)

// FileExporter std console exporter
type FileExporter struct{}

// Encoder
/**
 * @Description:
 * @receiver e
 * @return zapcore.Encoder
 */
func (e *FileExporter) Encoder() zapcore.Encoder {
	encoderConfig := configEncoder()
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	return fmtEncoder(config.Format, encoderConfig)
}

// Writer
/**
 * @Description:
 * @receiver e
 * @return zapcore.WriteSyncer
 */
func (e *FileExporter) Writer() zapcore.WriteSyncer {
	filename := path.Join(config.FileDirectory, config.FileName)
	hook, err := rotateLogs.New(
		filename+".%Y%m%d%H", // {filename}.{%Y%m%d%H}
		rotateLogs.WithMaxAge(time.Hour*24*time.Duration(config.MaxDays)), // days
		rotateLogs.WithRotationTime(time.Hour*24),                         // split
	)
	if err != nil {
		log.Println("error init writer")
		panic(err)
	}
	return zapcore.AddSync(hook)
}

// Level
/**
 * @Description:
 * @receiver e
 * @return zapcore.Level
 */
func (e *FileExporter) Level() zapcore.Level {
	return config.GetLevel().zapLevel()
}
