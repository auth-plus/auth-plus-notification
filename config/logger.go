package config

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// GetLogger exports a logger
func GetLogger() *zap.Logger {
	core := getCore()
	tree := zapcore.NewTee(core)
	logger := zap.New(tree)
	defer logger.Sync()
	return logger
}

func getProductionCore() zapcore.Core {
	encoderCfg := zap.NewProductionEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encoder := zapcore.NewJSONEncoder(encoderCfg)
	priority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	return zapcore.NewCore(encoder, zapcore.Lock(os.Stderr), priority)
}

func getTestCore() zapcore.Core {
	encoderCfg := zap.NewDevelopmentEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encoder := zapcore.NewConsoleEncoder(encoderCfg)
	priority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl == zapcore.DebugLevel
	})
	return zapcore.NewCore(encoder, zapcore.Lock(os.Stdout), priority)
}

func getDevCore() zapcore.Core {
	encoderCfg := zap.NewDevelopmentEncoderConfig()
	encoderCfg.TimeKey = "timestamp"
	encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
	encoder := zapcore.NewConsoleEncoder(encoderCfg)
	priority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.InfoLevel
	})
	return zapcore.NewCore(encoder, zapcore.Lock(os.Stdout), priority)
}

func getCore() zapcore.Core {
	env := GetEnv()
	if env.App.Env == "production" {
		return getProductionCore()
	}
	if env.App.Env == "test" {
		return getTestCore()
	}
	return getDevCore()
}
