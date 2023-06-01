package logger

import (
	"github.com/acanewby/patrick/internal/common"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"strings"
)

var (
	zapLog   *zap.Logger
	zapSugar *zap.SugaredLogger
	lvl      = zap.NewAtomicLevelAt(zap.InfoLevel)
)

func init() {
	var err error
	var config zap.Config

	config.Level = lvl
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

func LogLevel() zapcore.Level {
	return zapSugar.Level()
}

func SetLogLevel(level string) {

	var l zap.AtomicLevel
	var err error

	if l, err = zap.ParseAtomicLevel(strings.ToLower(level)); err != nil {
		Errorf(common.ErrorTemplateParseError, err)
		os.Exit(common.EXIT_CODE_CONFIGURATION_ERROR)
	} else {
		Infof(common.LogTemplateSettingLogLevel, l)
		lvl.SetLevel(l.Level())
		Infof(common.LogTemplateSetLogLevel, l)
	}

}

func Infof(message string, fields ...any) {
	zapSugar.Infof(message, fields...)
}

func Debugf(message string, fields ...any) {
	zapSugar.Debugf(message, fields...)
}

func Warnf(message string, fields ...any) {
	zapSugar.Warnf(message, fields...)
}

func Errorf(message string, fields ...any) {
	zapSugar.Errorf(message, fields...)
}

func Fatalf(message string, fields ...any) {
	zapSugar.Fatalf(message, fields...)
}
