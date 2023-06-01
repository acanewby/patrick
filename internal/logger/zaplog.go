package logger

import (
	"go.uber.org/zap"
)

var (
	zapLog   *zap.Logger
	zapSugar *zap.SugaredLogger
)

func init() {
	var err error
	var config zap.Config

	config.Level = zap.NewAtomicLevelAt(zap.DebugLevel)
	config.OutputPaths = []string{"./patrick.log"}
	config.Development = true
	config.DisableCaller = false
	config.DisableStacktrace = false
	config.Encoding = "console"

	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.StacktraceKey = "" // to hide stacktrace info
	config.EncoderConfig = encoderConfig

	zapLog, err = config.Build(zap.AddCallerSkip(1))
	if err != nil {
		panic(err)
	}
	zapSugar = zapLog.Sugar()
}

func Infof(message string, fields ...any) {
	zapSugar.Infof(message, fields...)
}

func Debugf(message string, fields ...any) {
	zapSugar.Debugf(message, fields...)
}

func Errorf(message string, fields ...any) {
	zapSugar.Errorf(message, fields...)
}

func Fatalf(message string, fields ...any) {
	zapSugar.Fatalf(message, fields...)
}
