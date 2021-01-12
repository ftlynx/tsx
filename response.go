package tsx

import "net/http"

type CodeValue struct {
	HttpCode int
	Msg      string
}

var CodeMap map[int]CodeValue

type Response struct {
	HttpCode  int         `json:"-"`
	ErrCode   int         `json:"code"` //自定义code
	ErrMsg    string      `json:"-"`
	Message   string      `json:"message,omitempty"` //错误提示
	Data      interface{} `json:"data,omitempty"`
	RequestId string      `json:"request_id"` //请求id
}

func (r *Response) Error() string {
	return r.ErrMsg
}

func NewOk(data ...interface{}) *Response {
	var d interface{}
	if len(data) > 0 {
		d = data[0]
	}
	return &Response{
		HttpCode:  http.StatusOK,
		ErrCode:   http.StatusOK,
		Message:   "ok",
		ErrMsg:    "",
		Data:      d,
		RequestId: "",
	}
}

func NewResponse(err error) *Response {
	if exception, ok := err.(*exception); ok {
		message := "未定义的code"
		httpCode := http.StatusInternalServerError
		if exception.ErrCode == 0 {
			exception.ErrCode = http.StatusInternalServerError
		}
		if _, ok := CodeMap[exception.ErrCode]; ok {
			message = CodeMap[exception.ErrCode].Msg
			httpCode = CodeMap[exception.ErrCode].HttpCode
		}
		return &Response{
			HttpCode:  httpCode,
			ErrCode:   exception.ErrCode,
			ErrMsg:    exception.ErrMsg,
			Message:   message,
			Data:      nil,
			RequestId: "",
		}
	}

	return &Response{
		HttpCode: http.StatusInternalServerError,
		ErrCode:  http.StatusInternalServerError,
		Message:  "未知错误",
		ErrMsg:   err.Error(),
	}
}
