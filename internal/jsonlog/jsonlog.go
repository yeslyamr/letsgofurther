package jsonlog

import (
	"context"
	"log/slog"
	"os"
	"runtime/debug"
)

const (
	LevelFatal = slog.Level(12)
)

var LevelNames = map[slog.Leveler]string{
	LevelFatal: "FATAL",
}

func Init() *slog.JSONHandler {
	opts := &slog.HandlerOptions{
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.LevelKey {
				level := a.Value.Any().(slog.Level)
				levelLabel, exists := LevelNames[level]
				if !exists {
					levelLabel = level.String()
				}
				a.Value = slog.StringValue(levelLabel)
			}
			return a
		},
	}

	handler := slog.NewJSONHandler(os.Stderr, opts)
	logger := slog.New(handler)
	slog.SetDefault(logger)
	return handler
}

func Info(message string, properties map[string]string) {
	slog.Info(message, slog.Any("properties", properties))
}
func Error(err error, properties map[string]string) {
	slog.Error(err.Error(), slog.Any("properties", properties))
}
func Fatal(err error, properties map[string]string) {
	trace := string(debug.Stack())
	slog.Log(context.Background(), LevelFatal, err.Error(), slog.Any("properties", properties), slog.Any("trace", trace))
	os.Exit(1) // For entries at the FATAL level, we also terminate the application.
}
