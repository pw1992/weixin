package auth

import (
	"encoding/json"
	"fmt"
	"github.com/pw1992/weixin/kernel"
	"github.com/pw1992/weixin/kernel/serror"
	"io/ioutil"
)

type AccessToken struct {
	Token       string
	TokenKey    string
	CachePrefix string
	Application *kernel.Application
}

type ResWeixinAccessToken struct {
	Errcode      int    `json:"errcode"`
	Errmsg       string `json:"errmsg"`
	Access_token string `json:"access_token"`
	Expires_in   int    `json:"expires_in"`
}

func NewAccessToken() *AccessToken {
	return &AccessToken{
		Token:       "",
		TokenKey:    "access_token",
		CachePrefix: "weixin.kernel.",
		Application: kernel.NewApplication(),
	}
}

func (a *AccessToken) GetToken() string {
	//从缓存中取
	access_token := a.Application.Cache.Get(a.CachePrefix+a.TokenKey, nil)

	if access_token != nil {
		//return access_token.(string)
	}

	corpid := a.Application.Config.GetString("corpid")
	corpsecret := a.Application.Config.GetString("corpsecret")
	a.Application.HttpClient.Endpoint = fmt.Sprintf("cgi-bin/gettoken?corpid=%s&corpsecret=%s", corpid, corpsecret)
	resp, err := a.Application.HttpClient.Get()
	if err != nil {
		err.Throw()
	}
	var res ResWeixinAccessToken

	body, e := ioutil.ReadAll(resp.Body)
	if e != nil {
		serror.NewError("读取返回值失败", 500, e).Throw()
	}

	json.Unmarshal(body, &res)
	marshal, _ := json.Marshal(res)
	a.Token = res.Access_token

	a.Application.Cache.Set(a.CachePrefix+a.TokenKey, marshal, res.Expires_in)
	return a.Token
}

func (a *AccessToken) Refresh() string {
	return ""
}
