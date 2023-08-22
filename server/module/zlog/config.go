package zlog

import (
	"gopkg.in/natefinch/lumberjack.v2"
)

type Config struct {
	// 因为zaplog没有分割日志文件的功能，所以采用lumberjack
	logRoll *lumberjack.Logger `yaml:"lumberjack"`
}

var logConfig = Config{
	logRoll: &lumberjack.Logger{
		Filename:   "./log/test.log",
		MaxSize:    100,
		MaxAge:     1,
		MaxBackups: 60,
		Compress:   false,
	},
}
