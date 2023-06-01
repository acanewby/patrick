package common

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
)

var (
	zapLog   *zap.Logger
	zapSugar *zap.SugaredLogger
	logLevel = zap.NewAtomicLevelAt(zap.InfoLevel)
)

func init() {
	var err error
	var config zap.Config

	encoderConfig := zap.NewDevelopmentEncoderConfig()
	encoderConfig.StacktraceKey = "" // to hide stacktrace info

	config.Level = logLevel
	config.OutputPaths = []string{"./patrick.log"}
	config.Development = true
	config.DisableCaller = false
	config.DisableStacktrace = false
	config.Encoding = "console"
	config.EncoderConfig = encoderConfig

	if zapLog, err = config.Build(zap.AddCallerSkip(1)); err != nil {
		panic(err)
	}
	zapSugar = zapLog.Sugar()

}

func LogLevel() zapcore.Level {
	return zapSugar.Level()
}

func SetLogLevel(level string) {

	var l zap.AtomicLevel
	var err error

	if l, err = zap.ParseAtomicLevel(strings.ToLower(level)); err != nil {
		LogErrorf(ErrorTemplateParseError, err)
		os.Exit(EXIT_CODE_CONFIGURATION_ERROR)
	} else {
		LogInfof(LogTemplateSettingLogLevel, l)
		logLevel.SetLevel(l.Level())
		LogInfof(LogTemplateSetLogLevel, l)
	}

}

func LogInfof(message string, fields ...any) {
	zapSugar.Infof(message, fields...)
}

func LogDebugf(message string, fields ...any) {
	zapSugar.Debugf(message, fields...)
}

func LogWarnf(message string, fields ...any) {
	zapSugar.Warnf(message, fields...)
}

func LogErrorf(message string, fields ...any) {
	zapSugar.Errorf(message, fields...)
}

func LogFatalf(message string, fields ...any) {
	zapSugar.Fatalf(message, fields...)
}
