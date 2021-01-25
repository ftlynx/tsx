package tsx

import "net/http"

const RequestIdKey = "request_id" //用来记录每次请求ID的key

type CodeValue struct {
	HttpCode int
	Msg      string
}

var CodeMap map[int]CodeValue //定义code

type Response struct {
	HttpCode  int         `json:"-"`
	ErrCode   int         `json:"code"` //自定义code
	Success   bool        `json:"success"`
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
		Success:   true,
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
			message = CodeMap[exception.ErrCode].Msg
			httpCode = CodeMap[exception.ErrCode].HttpCode
		}
		return &Response{
			HttpCode:  httpCode,
			ErrCode:   exception.ErrCode,
			ErrMsg:    exception.Error(),
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

// PageData 数据分页数据
type PageData struct {
	PageSize   int         `json:"page_size"`   // 一页多少条
	PageIndex  int         `json:"page_index"`  // 当前第几页
	SizeCount  int64       `json:"total"`       // 总共多少条数据
	IndexCount int         `json:"index_count"` // 总共多少页面
	List       interface{} `json:"data"`        // 页面数据
	Success    bool        `json:"success"`
}

func (p *PageData) Format(sizeCount int64, pageSize int, list interface{}) {
	//除尽就不+1
	complement := 1
	if sizeCount%int64(pageSize) == 0 {
		complement = 0
	}
	p.SizeCount = sizeCount
	p.IndexCount = int(sizeCount/int64(pageSize)) + complement
	p.List = list
}

// 分页参数
type QueryPaging struct {
	PageSize  int `form:"pageSize"`  //每页多少条
	PageIndex int `form:"current"` //第几页
}

func (p *QueryPaging) Convert() (limit int, offset int) {
	p.defaultValue()
	limit = p.PageSize
	offset = p.PageSize * (p.PageIndex - 1)
	return
}

func (p *QueryPaging) defaultValue() {
	if p.PageIndex == 0 {
		p.PageIndex = 1
	}
	if p.PageSize == 0 {
		p.PageSize = 10
	}
}
