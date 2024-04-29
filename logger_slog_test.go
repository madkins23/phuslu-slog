//go:build go1.21
// +build go1.21

package log

import (
	"log/slog"
	"testing"
)

func TestSlogNormal(t *testing.T) {
	var logger *slog.Logger = (&Logger{
		Level:      InfoLevel,
		TimeField:  "date",
		TimeFormat: "2006-01-02",
		Caller:     1,
	}).Slog()

	logger.Info("hello from slog Info")
	logger.Warn("hello from slog Warn")
	logger.Error("hello from slog Error")
}

func TestSlogAttrs(t *testing.T) {
	var logger *slog.Logger = (&Logger{
		Level:      InfoLevel,
		TimeField:  "date",
		TimeFormat: "2006-01-02",
		Caller:     1,
	}).Slog()

	sublogger := logger.With("logger", "attr_logger").With("everything", 42)
	sublogger.Info("hello from attr slog")
}

func TestSlogGroup(t *testing.T) {
	var logger *slog.Logger = (&Logger{
		Level:      InfoLevel,
		TimeField:  "date",
		TimeFormat: "2006-01-02",
		Caller:     1,
	}).Slog()

	logger1 := logger.WithGroup("g").With("1", "2").With("3", "4")
	logger1.Info("hello from group slog 1")
	logger1.Info("hello from group slog 2")

	logger2 := logger1.WithGroup("g1").With("a", "b").With("c", "d").
		WithGroup("g2").With("foo", "bar").With("bar", "foo").
		WithGroup("g3").With("x", 1).With("y", 2).With("z", 3)
	logger2.Info("hello from group slog 3")
	logger2.Info("hello from group slog 4")

	logger1.Info("hello from group slog 1")
	logger1.Info("hello from group slog 2")
}
