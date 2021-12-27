package logger

import (
	"fmt"
	"testing"
)

func TestLogger_Info(t *testing.T) {
	OnlyConsole()
	l := &Logger{}
	l.WithField("x", []interface{}{22.31, "yyy"}).Info(1, 2, 3, 4, false, []interface{}{11.11, "x"})
	l.WithFields(Fields{"b": 88, "a": []interface{}{99.99, "www"}}).Info(1, 2, 3, 4, false, []interface{}{11.11, "x"})
}

func TestLogger_Warn(t *testing.T) {
	OnlyConsole()
	l := &Logger{}
	l.WithField("x", []interface{}{22.31, "yry"}).Warn(1, 2, 3, 4, false, []interface{}{11.11, "x"})
}

func TestFields_ZapFields(t *testing.T) {
	l := &Logger{}
	println(fmt.Sprint(l.FieldMap.ZapFields()))
}
