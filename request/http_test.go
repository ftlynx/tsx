package request

//自定义一个参数
//func customOption(token string) Option {
//	return func(req *http.Request) {
//		req.Header.Set("token", token)
//	}
//}
//
//func TestHttpSendForm(t *testing.T) {
//	body, _ := FormAuth("xxx", "xxx")
//	param := tsx.HttpRequestParam{
//		Url:            "http://xxx/login",
//		Body:           body,
//		Method:         http.MethodPost,
//	}
//	rc, err := HttpSend(DefaultClient(), param, ContentTypeByForm())
//	fmt.Println(rc.Code, string(rc.Content), err)
//}
//
//func TestHttpSendToken(t *testing.T) {
//	param := tsx.HttpRequestParam{
//		Url:            "http://xxx/login",
//		Body:           "",
//		Method:         http.MethodPost,
//	}
//	rc, err := HttpSend(DefaultClient(), param, ContentTypeByJSON(), customOption("abc"))
//	fmt.Println(rc.Code, string(rc.Content), err)
//}
