package media

import (
	"encoding/json"
	"fmt"
	"github.com/gabriel-vasile/mimetype"
	"github.com/pw1992/weixin"
	"github.com/pw1992/weixin/kernel/auth"
	"github.com/pw1992/weixin/kernel/contracts"
	"github.com/pw1992/weixin/kernel/serror"
	"io/ioutil"
	"os"
	"time"
)

type Media struct {
	AccessTokener contracts.AccessTokener
	HttpClient    *weixin.HttpClient
}

type ResMedia struct {
	Errcode    int
	Errmsg     string
	Type       string
	Media_id   string
	Created_at string
}

func NewMedia() *Media {
	return &Media{
		AccessTokener: auth.NewAccessToken(),
		HttpClient:    weixin.NewHttpClient(),
	}
}

// 获取临时素材 see@https://developer.work.weixin.qq.com/document/path/90254
func (m *Media) Get(mediaId string) (string, *serror.Error) {
	token := m.AccessTokener.GetToken()
	m.HttpClient.Endpoint = fmt.Sprintf("cgi-bin/media/get?access_token=%s&media_id=%s", token, mediaId)
	resp, e := m.HttpClient.Get()
	if e != nil {
		return "", e
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", serror.NewError("读取企业微信返回值失败", 500, err)
	}

	//获取文件类型
	detect := mimetype.Detect(body)
	ext := detect.Extension()

	if len(ext) == 0 {
		return "", serror.NewError("文件类型未获取", 500, ext)
	}

	localpath := fmt.Sprintf("%s%v%s", os.TempDir()+"/", time.Now().Unix(), ext)

	ioutil.WriteFile(localpath, body, os.ModePerm)
	return localpath, nil
}

// 上传临时素材  see@https://developer.work.weixin.qq.com/document/path/90253
func (m *Media) UploadImage(path string) (*ResMedia, *serror.Error) {
	return m.upload(path, "image")
}

// 上传视频
func (m *Media) UploadVideo(path string) (*ResMedia, *serror.Error) {
	return m.upload(path, "video")
}

// 上传语音
func (m *Media) UploadVoice(path string) (*ResMedia, *serror.Error) {
	return m.upload(path, "voice")
}

// 上传普通文件
func (m *Media) UploadFile(path string) (*ResMedia, *serror.Error) {
	return m.upload(path, "file")
}

func (m *Media) upload(path, mediaType string) (*ResMedia, *serror.Error) {
	token := m.AccessTokener.GetToken()
	m.HttpClient.Endpoint = fmt.Sprintf("cgi-bin/media/upload?access_token=%s&type=%s", token, mediaType)

	m.HttpClient.ContentType = "application/octet-stream"

	openfile, err := os.Open(path)
	if err != nil {
		return nil, serror.NewError("上传图片打开文件失败:", 500, err)
	}
	resData, e := m.HttpClient.Post(openfile)
	if e != nil {
		return nil, e
	}

	var res *ResMedia
	json.Unmarshal(resData, &res)
	return res, nil
}
