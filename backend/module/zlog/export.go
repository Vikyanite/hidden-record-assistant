package zlog

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Info(msg string) {
	do(zap.InfoLevel, "", msg)
}

func Debug(msg string) {
	do(zap.DebugLevel, "", msg)
}

func Warn(msg string) {
	do(zap.WarnLevel, "", msg)
}

func Error(msg string) {
	do(zap.ErrorLevel, "", msg)
}

func Panic(msg string) {
	do(zap.PanicLevel, "", msg)
}

func Fatal(msg string) {
	do(zap.FatalLevel, "", msg)
}

func Infof(tmplt string, args ...interface{}) {
	do(zap.InfoLevel, tmplt, args...)
}

func Debugf(tmplt string, args ...interface{}) {
	do(zap.DebugLevel, tmplt, args...)
}

func Warnf(tmplt string, args ...interface{}) {
	do(zap.WarnLevel, tmplt, args...)
}

func Errorf(tmplt string, args ...interface{}) {
	do(zap.ErrorLevel, tmplt, args...)
}

func Panicf(tmplt string, args ...interface{}) {
	do(zap.PanicLevel, tmplt, args...)
}

func Fatalf(tmplt string, args ...interface{}) {
	do(zap.FatalLevel, tmplt, args...)
}

func do(l zapcore.Level, tmplt string, args ...interface{}) {
	// 留个field字段便于以后要插东西
	var field []zapcore.Field
	msg := formatMsg(tmplt, args)
	switch l {
	case zapcore.DebugLevel:
		logger.Debug(msg, field...)
	case zapcore.InfoLevel:
		logger.Info(msg, field...)
	case zapcore.WarnLevel:
		logger.Warn(msg, field...)
	case zapcore.ErrorLevel:
		logger.Error(msg, field...)
	case zapcore.PanicLevel:
		logger.Panic(msg, field...)
	case zapcore.FatalLevel:
		logger.Fatal(msg, field...)
	}

}

func formatMsg(template string, args []interface{}) string {
	if template == "" {
		return args[0].(string)
	}
	return fmt.Sprintf(template, args...)
}
