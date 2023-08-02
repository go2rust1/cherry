package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
	"time"
)

type TeeOption struct {
	W   io.Writer
	Lef LevelEnablerFunc
}

func NewTee(tops []TeeOption, opts ...Option) *Logger {
	var cores []zapcore.Core
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02T15:04:05.000Z0700"))
	}
	for _, top := range tops {
		top := top

		if top.W == nil {
			panic("the writer is nil")
		}

		lv := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
			return top.Lef(lvl)
		})

		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(cfg.EncoderConfig),
			zapcore.AddSync(top.W),
			lv,
		)
		cores = append(cores, core)
	}

	logger := &Logger{
		l: zap.New(zapcore.NewTee(cores...), opts...),
	}
	return logger
}
