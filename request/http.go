package request

import (
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type HttpRequestParam struct {
	Url    string
	Body   string
	Method string
}

type HttpResponse struct {
	Code    int         //请求返回的状态码
	Content []byte      //请求返回的内容
	Header  http.Header //请求返回的header
}

//自定义参数
type Option func(*http.Request)

// json请求类型
func ContentTypeByJSON() Option {
	return func(req *http.Request) {
		req.Header.Set("Content-Type", "application/json")
	}
}

// form请求类型
func ContentTypeByForm() Option {
	return func(req *http.Request) {
		req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	}
}

// 自定义请求类型
func SetContentType(contentType string) Option {
	return func(req *http.Request) {
		req.Header.Set("Content-Type", contentType)
	}
}

// http basic 认证
func BasicAuth(username string, password string) Option {
	return func(req *http.Request) {
		req.SetBasicAuth(username, password)
	}
}

// 标准http的认证信息应该使用Authorization header
func Authorization(token string) Option {
	return func(req *http.Request) {
		req.Header.Set("Authorization", token)
	}
}

// 自定义认证信息的header
func SetAuthorization(key string, token string) Option {
	return func(req *http.Request) {
		req.Header.Set(key, token)
	}
}

//Form认证，自己定义用户的key为username, 密码的key为password
func FormAuth(username string, password string) (string, error) {
	r := http.Request{}
	if err := r.ParseForm(); err != nil {
		return "", err
	}
	r.Form.Add("userName", username)
	r.Form.Add("password", password)
	return r.Form.Encode(), nil
}

//默认client参数
func DefaultClient() *http.Client {
	return &http.Client{
		Timeout: 15 * time.Second,
	}
}

//需要注意options的顺序
func (param *HttpRequestParam)HttpSend(client *http.Client, options ...Option) (HttpResponse, error) {
	rc := HttpResponse{}
	req, err := http.NewRequest(param.Method, param.Url, strings.NewReader(param.Body))
	if err != nil {
		return rc, err
	}

	// 自定义配置
	for _, option := range options {
		option(req)
	}

	//client := &http.Client{
	//	Timeout: time.Duration(15 * time.Second),
	//}
	resp, err := client.Do(req)
	if err != nil {
		return rc, err
	}
	defer resp.Body.Close()
	rc.Content, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return rc, err
	}
	rc.Code = resp.StatusCode
	rc.Header = resp.Header

	//fmt.Printf("req_method:%s req_url:%s req_body:%s resp_code:%d resp_content:%s", req.Method, req.URL.String(), string(reqParam.Body), rc.Code, string(rc.Content))
	return rc, nil
}

//自定义结果处理
type ResultParse func(HttpResponse) error

/*
样例
type TestResponse struct {
	Code int
	Msg string
}
func MyResponse(t *TestResponse) ResultParse{
	return func(resp HttpResponse) error {
		if resp.Code != http.StatusOK {
			return fmt.Error("error")
		}
		return json.Unmarshal(resp.Content, t)
	}
}
*/

//需要注意options的顺序，通过自定义ResultParse函数参数自行处理解析的结果
func (param *HttpRequestParam)HttpResultParse(client *http.Client, resultParse ResultParse, options ...Option) error {
	resp, err := param.HttpSend(client, options...)
	if err != nil {
		return err
	}
	return resultParse(resp)
}
