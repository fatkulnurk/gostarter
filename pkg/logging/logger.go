package logging

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"sync"
	"time"
)

var (
	log  *zap.Logger
	once sync.Once
)

func InitLogger() {
	once.Do(func() {
		logFileName := fmt.Sprintf("logs/log-%s.log", time.Now().Format("2006-01-02"))

		if err := os.MkdirAll("logs", 0755); err != nil {
			panic(fmt.Sprintf("failed to create logs directory: %v", err))
		}

		logFile, err := os.OpenFile(logFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			panic(fmt.Sprintf("failed to open log file: %v", err))
		}

		encoderCfg := zap.NewProductionEncoderConfig()
		encoderCfg.TimeKey = "timestamp"
		encoderCfg.EncodeTime = zapcore.ISO8601TimeEncoder
		encoder := zapcore.NewConsoleEncoder(encoderCfg)

		fileSyncer := zapcore.AddSync(logFile)
		stdoutSyncer := zapcore.AddSync(os.Stdout)

		core := zapcore.NewTee(
			zapcore.NewCore(encoder, stdoutSyncer, zap.InfoLevel),
			zapcore.NewCore(encoder, fileSyncer, zap.InfoLevel),
		)

		log = zap.New(core, zap.AddCaller())
	})
}

// GetLogger Shortcut access ke zap logger
func GetLogger() *zap.Logger {
	if log == nil {
		InitLogger()
	}
	return log
}

func Debug(msg string, fields ...zap.Field) {
	GetLogger().Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	GetLogger().Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	GetLogger().Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	GetLogger().Error(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	GetLogger().Fatal(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	GetLogger().Panic(msg, fields...)
}

func Infof(format string, args ...interface{}) {
	Info(fmt.Sprintf(format, args...))
}
