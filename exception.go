package tsx

import (
	"fmt"
	"runtime"
	"time"
)

type exception struct {
	ErrCode  int    `json:"-"`
	ErrMsg   string `json:"-"`
	Time     string `json:"-"`
	Filename string `json:"-"`
	Line     int    `json:"-"`
}

func (e *exception) Error() string {
	return fmt.Sprintf("%s:%d %s", e.Filename, e.Line, e.ErrMsg)
}

//用于获取错误的位置及设置errCode
func Error(err error, errCode ...int) error {
	if err == nil {
		return err
	}
	code := 0
	if len(errCode) > 0 {
		code = errCode[0]
	}

	_, file, line, ok := runtime.Caller(1)
	if !ok {
		return fmt.Errorf("%s %v", "caller fail", err)
	}
	return &exception{
		Time:     time.Now().Local().String(),
		Filename: file,
		Line:     line,
		ErrMsg:   err.Error(),
		ErrCode:  code,
	}
}

