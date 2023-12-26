package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// New create config, split with level
func New(cfg *Config) {
	if cfg == nil {
		cfg = defaultConfig()
	}
	config = cfg
	exps := NewLevelExporters(cfg)
	NewLogger(exps...)
}

// NewOne create config, split with only date
func NewOne(cfg *Config) {
	if cfg == nil {
		cfg = defaultConfig()
	}
	config = cfg
	exps := NewExporters(cfg)
	NewLogger(exps...)
}

// OnlyConsole for test or develop
func OnlyConsole() {
	if config == nil {
		config = &Config{
			StdLevel:  "debug",
			StdFormat: "console",
		}
	}
	NewLogger(&StdExporter{})
}

// NewLogger build new logger
func NewLogger(exporters ...IExporter) {
	coreTree := make([]zapcore.Core, 0)
	for _, exp := range exporters {
		coreTree = append(coreTree, zapcore.NewCore(exp.Encoder(), exp.Writer(), exp.Level()))
	}
	cores := zapcore.NewTee(coreTree...)
	// build log
	lg := zap.New(cores, zap.AddCaller(), zap.AddCallerSkip(0))
	ctxLogger = lg.WithOptions(zap.AddCallerSkip(2))
	zap.ReplaceGlobals(lg)
}
