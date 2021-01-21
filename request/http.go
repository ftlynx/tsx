package request

import (
	"go.etcd.io/etcd/client"
	"io/ioutil"
	"net/http"
	"strings"
	"time"
)

type HttpRequestParam struct {
	Url           string
	Body          string
	Method        string
	User          string
	Password      string
	Authorization string
	Authentication string
}

type HttpResponse struct {
	Code    int         //请求返回的状态码
	Content []byte      //请求返回的内容
	Header  http.Header //请求返回的header
}

func HttpRequest(reqParam HttpRequestParam) (HttpResponse, error) {
	rc := HttpResponse{}
	req, err := http.NewRequest(reqParam.Method, reqParam.Url, strings.NewReader(reqParam.Body))
	if err != nil {
		return rc, err
	}
	req.Header.Set("Content-Type", "application/json")
	if reqParam.Password != "" {
		req.SetBasicAuth(reqParam.User, reqParam.Password)
	}

	if reqParam.Authorization != "" {
		req.Header.Set("Authorization", reqParam.Authorization)
	}

	if reqParam.Authentication != "" {
		req.Header.Set("Authentication", reqParam.Authentication)
	}

	//fmt.Println(req.URL, req.Header.Get("Authorization"))
	client := &http.Client{
		Timeout: time.Duration(15 * time.Second),
	}
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
