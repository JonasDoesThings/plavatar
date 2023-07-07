package utils

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func InitLogger() *zap.SugaredLogger {
	logConfig := zap.NewProductionConfig()
	logConfig.OutputPaths = []string{
		"stdout",
	}
	logConfig.ErrorOutputPaths = []string{
		"stderr",
	}
	logConfig.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	logConfig.Encoding = "console"
	logConfig.EncoderConfig = zapcore.EncoderConfig{
		TimeKey:        "T",
		LevelKey:       "L",
		NameKey:        "N",
		CallerKey:      "C",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "M",
		StacktraceKey:  "S",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.StringDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	var rawLogger, err = logConfig.Build()

	if err != nil {
		fmt.Println(err)
	}

	// sync rawLogger and handle possible errors
	defer func(rawLogger *zap.Logger) {
		err := rawLogger.Sync()
		if err != nil {
			fmt.Println(err)
		}
	}(rawLogger)
	return rawLogger.Sugar()
}
