package tsx

import (
	"encoding/json"
	"fmt"
	"runtime"
	"time"
)

type errorInfo struct {
	Time     string `json:"time"`
	Message  string `json:"message"`
	Filename string `json:"filename"`
	Line     int    `json:"line"`
}

func Errx(s string, e error) error {
	return fmt.Errorf("%s %w", s, e)
}

//打印日志
func ErrxError(e error) error {
	if e == nil {
		return e
	}
	_, file, line, ok := runtime.Caller(1)
	if !ok {
		return Errx("caller fail", e)
	}
	msg := errorInfo{
		Time:     time.Now().Local().String(),
		Filename: file,
		Line:     line,
		Message:  e.Error(),
	}
	b, err := json.Marshal(msg)
	if err != nil {
		return Errx(err.Error(), e)
	}
	fmt.Println(string(b))
	return e
}
