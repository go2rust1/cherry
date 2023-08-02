package logger

import (
	"github.com/go2rust1/cherry/src/conf"
	"github.com/go2rust1/cherry/src/mod/logger"
)

var option = []logger.TeeRotateOption{
	{
		Filename: conf.InfoFile,
		Rop: logger.RotateOption{
			MaxSize:    conf.InfoMaxSize,
			MaxAge:     conf.InfoMaxAge,
			MaxBackups: conf.InfoMaxBackups,
			Compress:   conf.InfoCompress,
		},
		Lef: func(lvl logger.Level) bool {
			return lvl <= logger.InfoLevel
		},
	},
	{
		Filename: conf.ErrorFile,
		Rop: logger.RotateOption{
			MaxSize:    conf.ErrorMaxSize,
			MaxAge:     conf.ErrorMaxAge,
			MaxBackups: conf.ErrorMaxBackups,
			Compress:   conf.ErrorCompress,
		},
		Lef: func(lvl logger.Level) bool {
			return lvl > logger.InfoLevel
		},
	},
}

var Logger = logger.NewTeeWithRotate(option)
