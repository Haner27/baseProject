package logger

import (
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

func GetRotatedFileSyncer(logPath string) zapcore.WriteSyncer {
	hook := &lumberjack.Logger{
		Filename: logPath,
		MaxSize: 128,
		MaxBackups: 30,
		MaxAge: 7,
		Compress: true,
	}
	return zapcore.AddSync(hook)
}

func GetStdoutAndRotatedFileSyncer(logPath string) zapcore.WriteSyncer {
	return zapcore.NewMultiWriteSyncer(
		zapcore.AddSync(os.Stdout),
		zapcore.AddSync(GetRotatedFileSyncer(logPath)),
	)
}
