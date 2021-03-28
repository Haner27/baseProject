package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var CommonLevelEnable = zap.LevelEnablerFunc(func(level zapcore.Level) bool {
	return level < zapcore.ErrorLevel
})

var ErrorLevelEnable = zap.LevelEnablerFunc(func(level zapcore.Level) bool {
	return level >= zapcore.ErrorLevel
})