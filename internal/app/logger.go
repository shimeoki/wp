package app

import "log/slog"

type LogLevel = slog.Level

type Logger interface {
	Log(ctx Ctx, lvl LogLevel, msg string, args ...any)
}

func Debug(l Logger, ctx Ctx, msg string, args ...any) {
	l.Log(ctx, slog.LevelDebug, msg, args...)
}

func Info(l Logger, ctx Ctx, msg string, args ...any) {
	l.Log(ctx, slog.LevelInfo, msg, args...)
}

func Warn(l Logger, ctx Ctx, msg string, args ...any) {
	l.Log(ctx, slog.LevelWarn, msg, args...)
}

func Error(l Logger, ctx Ctx, msg string, args ...any) {
	l.Log(ctx, slog.LevelError, msg, args...)
}
