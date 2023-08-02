package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	DebugLevel  = zap.DebugLevel  // -1
	InfoLevel   = zap.InfoLevel   // 0, default level
	WarnLevel   = zap.WarnLevel   // 1
	ErrorLevel  = zap.ErrorLevel  // 2
	DPanicLevel = zap.DPanicLevel // 3, used in development log
	PanicLevel  = zap.PanicLevel  // 4  logs a message, then panics
	FatalLevel  = zap.FatalLevel  // 5  logs a message, then calls os.Exit(1)
)

type (
	Field = zap.Field
	Level = zapcore.Level
)

type Logger struct {
	l     *zap.Logger
	level Level
}

func (l *Logger) Debug(msg string, fields ...Field) {
	l.l.Debug(msg, fields...)
}

func (l *Logger) Info(msg string, fields ...Field) {
	l.l.Info(msg, fields...)
}

func (l *Logger) Warn(msg string, fields ...Field) {
	l.l.Warn(msg, fields...)
}

func (l *Logger) Error(msg string, fields ...Field) {
	l.l.Error(msg, fields...)
}

func (l *Logger) DPanic(msg string, fields ...Field) {
	l.l.DPanic(msg, fields...)
}

func (l *Logger) Panic(msg string, fields ...Field) {
	l.l.Panic(msg, fields...)
}

func (l *Logger) Fatal(msg string, fields ...Field) {
	l.l.Fatal(msg, fields...)
}
