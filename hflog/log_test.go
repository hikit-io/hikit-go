package hflog

import (
	"testing"
)

func TestLogger(t *testing.T) {
	t.Run("nil debug", func(t *testing.T) {
		Debug(nil, "")
	})
	t.Run("debug", func(t *testing.T) {
		Debug(nil, "")
	})
	t.Run("nil info", func(t *testing.T) {
		Info(nil, "")
	})
	t.Run("info", func(t *testing.T) {
		Info(nil, "")
	})
	t.Run("nil Error", func(t *testing.T) {
		Error(nil, "")
	})
	t.Run("Error", func(t *testing.T) {
		Error(nil, "")
	})
	ReplaceLoggerKey("hflog test")
	t.Run("nil Warn", func(t *testing.T) {
		Warn(nil, "")
	})
	t.Run("Warn", func(t *testing.T) {
		Warn(nil, "")
	})
}
