package logger

import (
	"log"
	"path"
	"strings"
	"time"

	rotateLogs "github.com/lestrrat-go/file-rotatelogs"
	"go.uber.org/zap/zapcore"
)

func NewLevelFileExporter(level zapcore.Level, cfg *Config) *LevelFileExporter {
	exp := &LevelFileExporter{
		level: level,
		EncoderConfig: zapcore.EncoderConfig{
			EncodeTime: cfg.EncodeTime,
		}}
	return exp
}

// LevelFileExporter file console exporter
type LevelFileExporter struct {
	level zapcore.Level
	zapcore.EncoderConfig
}

// Encoder
/**
 * @Description:
 * @receiver e
 * @return zapcore.Encoder
 */
func (e *LevelFileExporter) Encoder() zapcore.Encoder {
	encoderConfig := configEncoder()
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	if e.EncodeTime != nil {
		encoderConfig.EncodeTime = e.EncodeTime
	}
	return fmtEncoder(config.Format, encoderConfig)
}

// Writer
/**
 * @Description:
 * @receiver e
 * @return zapcore.WriteSyncer
 */
func (e *LevelFileExporter) Writer() zapcore.WriteSyncer {
	filename := path.Join(config.FileDirectory, config.FileName)
	nameFmt := strings.Join([]string{filename, e.level.String(), "%Y%m%d%H"}, ".")
	hook, err := rotateLogs.New(
		nameFmt, // {filename}.{%Y%m%d%H}
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
func (e *LevelFileExporter) Level() zapcore.LevelEnabler {
	return e
}

// Enabled level enable
func (e *LevelFileExporter) Enabled(lvl zapcore.Level) bool {
	return e.level == lvl
}
