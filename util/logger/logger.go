package logger

import (
	"baseProject/config"
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.SugaredLogger
}

func NewLogger(conf *config.Config) *Logger {
	commonLogPath := fmt.Sprintf("%s/%s-com.log", conf.Project.LogDir, conf.Project.Name)
	errorLogPath := fmt.Sprintf("%s/%s-com.log", conf.Project.LogDir, conf.Project.Name)
	core := zapcore.NewTee(
		zapcore.NewCore(
			CommonConsoleEncoder,
			GetStdoutAndRotatedFileSyncer(commonLogPath),
			CommonLevelEnable,
		),
		zapcore.NewCore(
			ErrorJsonEncoder,
			GetRotatedFileSyncer(errorLogPath),
			ErrorLevelEnable,
		),
	)
	logger := zap.New(
		core,
		zap.AddStacktrace(ErrorLevelEnable),
		zap.AddCaller(),
	)
	sugar := logger.Sugar()
	sugar.Named(conf.Project.Name)
	return &Logger{sugar}
}