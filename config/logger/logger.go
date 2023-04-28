package logger

import (
	"context"
	defaultLog "log"
	"os"
	"strings"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func SetupLogger(ctx context.Context) {
	log, err := newLogger(ctx)
	if err != nil {
		defaultLog.Fatalf("Couldn't initialize logger: %v", err)
	}

	_ = zap.ReplaceGlobals(log)
}

func newLogger(ctx context.Context) (*zap.Logger, error) {
	var config zap.Config

	config = zap.NewProductionConfig()

	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.MessageKey = "message"

	if lvl, exists := os.LookupEnv("LOG_LEVEL"); exists {
		lvl = strings.ToLower(lvl)
		switch lvl {
		case "debug":
			config.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
		case "info":
			config.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
		case "warn":
			config.Level = zap.NewAtomicLevelAt(zapcore.WarnLevel)
		case "error":
			config.Level = zap.NewAtomicLevelAt(zapcore.ErrorLevel)
		case "panic":
			config.Level = zap.NewAtomicLevelAt(zapcore.PanicLevel)
		case "fatal":
			config.Level = zap.NewAtomicLevelAt(zapcore.FatalLevel)
		}
	}

	log, err := config.Build()
	if err != nil {
		return nil, err
	}

	return log, err
}
