package logx

import (
	"github.com/ftlynx/tsx"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
	"os"
	"strings"
	"time"
)

type LogConfig struct {
	Level         string
	File          string
	Json          bool
	RotationCount uint
	RotationHour  time.Duration //多少小时滚动文件
}

//转换用户输入的字符串， logrus自身有一个，重写一个如果传入异常，默认为info
func parseLevel(lvl string) logrus.Level {
	switch strings.ToLower(lvl) {
	case "panic":
		return logrus.PanicLevel
	case "fatal":
		return logrus.FatalLevel
	case "error":
		return logrus.ErrorLevel
	case "warn", "warning":
		return logrus.WarnLevel
	case "info":
		return logrus.InfoLevel
	case "debug":
		return logrus.DebugLevel
	default:
		return logrus.InfoLevel
	}
	return logrus.InfoLevel
}

func (l *LogConfig) New() *logrus.Logger {
	log := logrus.New()
	if l.Json {
		log.SetFormatter(&logrus.TextFormatter{})
	} else {
		log.SetFormatter(&logrus.JSONFormatter{})
	}
	log.SetLevel(parseLevel(l.Level))
	if l.File == "" {
		log.SetOutput(os.Stdout)
	} else {
		if err := tsx.CreateParentDir(l.File); err != nil {
			panic(err)
		}
		if l.RotationCount == 0 {
			l.RotationCount = 30
		}
		write, err := rotatelogs.New(
			l.File+".%Y-%m-%d-%H",
			rotatelogs.WithLinkName(l.File),
			rotatelogs.WithRotationTime(l.RotationHour*time.Hour),
			rotatelogs.WithRotationCount(l.RotationCount),
		)
		if err != nil {
			panic(err)
		}
		log.SetOutput(write)
	}
	return log
}
