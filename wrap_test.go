package logger

import "testing"

func TestLogger_Info(t *testing.T) {
	OnlyConsole()
	l := Logger{}
	l.Info(1, 2, 3, 4, false)
}
