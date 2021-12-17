package logger

import "go.uber.org/zap"

var (
	config    *Config     // config
	ctxLogger *zap.Logger // log core key-value from context
)

// Key context logger core
type Key string

const (
	CtxLogger Key = "_CtxLogger"
)
