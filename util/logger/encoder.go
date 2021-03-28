package logger

import (
	"go.uber.org/zap/zapcore"
	"time"
)

var TimeEncoder = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("02/01/2006 15:04:05"))
}

var ErrorEncoderConfig = zapcore.EncoderConfig{
	TimeKey: "ts",
	LevelKey: "lv",
	NameKey: "n",
	CallerKey: "cl",
	MessageKey: "msg",
	StacktraceKey: "st",
	LineEnding: zapcore.DefaultLineEnding,
	EncodeLevel: zapcore.LowercaseColorLevelEncoder,
	EncodeTime: zapcore.ISO8601TimeEncoder,
	EncodeDuration: zapcore.SecondsDurationEncoder,
	EncodeCaller: zapcore.ShortCallerEncoder,
}

var CommonEncoderConfig = zapcore.EncoderConfig{
	TimeKey: "ts",
	LevelKey: "lv",
	NameKey: "n",
	CallerKey: "cl",
	MessageKey: "msg",
	StacktraceKey: "st",
	LineEnding: zapcore.DefaultLineEnding,
	EncodeLevel: zapcore.CapitalColorLevelEncoder,
	EncodeTime: TimeEncoder,
	EncodeDuration: zapcore.SecondsDurationEncoder,
	EncodeCaller: zapcore.ShortCallerEncoder,
}

var ErrorJsonEncoder = zapcore.NewJSONEncoder(ErrorEncoderConfig)
var CommonConsoleEncoder = zapcore.NewConsoleEncoder(CommonEncoderConfig)
