package app

import "log/slog"

type LogLevel = slog.Level

type Logger interface {
	Log(ctx Ctx, lvl LogLevel, msg string, args ...any)
}

type logger struct {
	Logger
}

func wrapLogger(l Logger) *logger {
	return &logger{Logger: l}
}

func (l *logger) Debug(ctx Ctx, msg string, args ...any) {
	l.Log(ctx, slog.LevelDebug, msg, args...)
}

func (l *logger) Info(ctx Ctx, msg string, args ...any) {
	l.Log(ctx, slog.LevelInfo, msg, args...)
}

func (l *logger) Warn(ctx Ctx, msg string, args ...any) {
	l.Log(ctx, slog.LevelWarn, msg, args...)
}

func (l *logger) Error(ctx Ctx, msg string, args ...any) {
	l.Log(ctx, slog.LevelError, msg, args...)
}
