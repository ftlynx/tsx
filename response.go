package tsx

import "net/http"

type CodeValue struct {
	httpCode int
	msg      string
}

var CodeMap = map[int]CodeValue{
	http.StatusBadRequest:          {http.StatusBadRequest, "参数错误"},
	http.StatusUnauthorized:        {http.StatusUnauthorized, "认证失败"},
	http.StatusForbidden:           {http.StatusForbidden, "权限拒绝"},
	http.StatusNotFound:            {http.StatusNotFound, http.StatusText(http.StatusNotFound)},
	http.StatusMethodNotAllowed:    {http.StatusMethodNotAllowed, http.StatusText(http.StatusMethodNotAllowed)},
	http.StatusInternalServerError: {http.StatusInternalServerError, "服务器内部错误"},
}

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
	if exception, ok := err.(*Exception); ok {
		message := "未定义的code"
		httpCode := http.StatusInternalServerError
		if exception.ErrCode == 0 {
			exception.ErrCode = http.StatusInternalServerError
		}
		if _, ok := CodeMap[exception.ErrCode]; ok {
			message = CodeMap[exception.ErrCode].msg
			httpCode = CodeMap[exception.ErrCode].httpCode
		}
		return &Response{
			HttpCode:  httpCode,
			ErrCode:   exception.ErrCode,
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
