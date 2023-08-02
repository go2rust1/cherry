package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

type Option = zap.Option

type LevelEnablerFunc func(lvl Level) bool

type RotateOption struct {
	MaxSize    int
	MaxAge     int
	MaxBackups int
	Compress   bool
}

type TeeRotateOption struct {
	Filename string
	Rop      RotateOption
	Lef      LevelEnablerFunc
}

func NewTeeWithRotate(tops []TeeRotateOption, opts ...Option) *Logger {
	var cores []zapcore.Core
	cfg := zap.NewProductionConfig()
	cfg.EncoderConfig.TimeKey = "time"
	cfg.EncoderConfig.MessageKey = "msg"
	cfg.EncoderConfig.LevelKey = "level"
	cfg.EncoderConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format("2006-01-02T15:04:05.000Z0700"))
	}
	// comment it if you want short caller path
	cfg.EncoderConfig.EncodeCaller = zapcore.FullCallerEncoder
	for _, top := range tops {
		top := top

		lv := zap.LevelEnablerFunc(func(lvl Level) bool {
			return top.Lef(lvl)
		})

		w := zapcore.AddSync(&lumberjack.Logger{
			Filename:   top.Filename,       // 日志文件的位置
			MaxSize:    top.Rop.MaxSize,    // 在进行切割之前, 日志文件的最大大小(以MB为单位)
			MaxBackups: top.Rop.MaxBackups, // 保留旧文件的最大个数
			MaxAge:     top.Rop.MaxAge,     // 保留旧文件的最大天数
			Compress:   top.Rop.Compress,   // 是否压缩/归档旧文件
		})

		core := zapcore.NewCore(
			zapcore.NewJSONEncoder(cfg.EncoderConfig),
			zapcore.AddSync(w),
			lv,
		)
		cores = append(cores, core)
	}

	// default print all level message to console
	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(cfg.EncoderConfig),
		zapcore.AddSync(os.Stdout),
		DebugLevel,
	)
	cores = append(cores, core)

	logger := &Logger{
		l: zap.New(zapcore.NewTee(cores...), opts...).
			WithOptions(
				zap.AddCaller(),
				zap.AddCallerSkip(1),
			),
	}
	return logger
}
