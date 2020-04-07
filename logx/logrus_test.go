package logx

import (
	"testing"
	"time"
	"github.com/ftlynx/tsx"
)

func TestNew(t *testing.T) {
	log := LogConfig{
		Level:         "debug",
		File:          "logs/test.log",
		Json:          false,
		RotationCount: 30,
		RotationHour:  24,
	}
	l := log.New()
	for {
		l.Infof("%v", tsx.Millisecond())
		time.Sleep(2 * time.Second)
	}
}
