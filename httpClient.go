package weixin

import (
	"fmt"
	"github.com/pw1992/weixin/kernel/serror"
	"io"
	"io/ioutil"
	"net/http"
)

type HttpClient struct {
	BaseUri     string //企业微信前缀
	Method      string //POST  GET
	Endpoint    string //请求地址的后缀
	ContentType string
}

func NewHttpClient() *HttpClient {
	return &HttpClient{
		BaseUri:     "https://qyapi.weixin.qq.com/",
		Method:      "",
		Endpoint:    "",
		ContentType: "application/json",
	}
}

func (client *HttpClient) Get() (*http.Response, *serror.Error) {
	resp, err := http.Get(fmt.Sprintf("%s%s", client.BaseUri, client.Endpoint))
	if err != nil {
		return nil, serror.NewError("get weixin api fail", 500, err)
	}
	return resp, nil
}

func (client *HttpClient) Post(body io.Reader) ([]byte, *serror.Error) {
	resp, err := http.Post(fmt.Sprintf("%s%s", client.BaseUri, client.Endpoint), client.ContentType, body)
	if err != nil {
		return nil, serror.NewError("get weixin api fail", 500, err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, serror.NewError("read weixin api fail", 500, err)
	}
	return data, nil
}
