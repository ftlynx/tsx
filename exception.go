package tsx

import (
	"fmt"
	"net/http"
	"runtime"
	"time"
)

type Exception struct {
	ErrCode  int    `json:"-"`
	ErrMsg   string `json:"-"`
	Time     string `json:"-"`
	Filename string `json:"-"`
	Line     int    `json:"-"`
}

func (e *Exception) Error() string {
	return e.ErrMsg
}

func Err(err error, errCode ...int) error {
	if err == nil {
		return err
	}
	code := http.StatusInternalServerError
	if len(errCode) > 0 {
		code = errCode[0]
	}

	_, file, line, ok := runtime.Caller(1)
	if !ok {
		return fmt.Errorf("%s %v", "caller fail", err)
	}
	return &Exception{
		Time:     time.Now().Local().String(),
		Filename: file,
		Line:     line,
		ErrMsg:   err.Error(),
		ErrCode:  code,
	}
}