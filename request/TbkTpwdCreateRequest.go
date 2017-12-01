package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type TbkTpwdCreateRequest struct {
	Request *topsdk.Request
}

// create new request
func NewTbkTpwdCreateRequest() (req *TbkTpwdCreateRequest) {
	request := topsdk.Request{MethodName: "taobao.tbk.tpwd.create", Params: make(map[string]interface{}, 5)}
	req = &TbkTpwdCreateRequest{
		Request: &request,
	}
	return
}

// 口令跳转目标页
func (req *TbkTpwdCreateRequest) SetUrl(url string) {
	req.Request.Params["url"] = url
}

func (req *TbkTpwdCreateRequest) GetUrl() string {
	url, found := req.Request.Params["url"]
	if found {
		return url.(string)
	}
	return ""
}

// 口令弹框内容
func (req *TbkTpwdCreateRequest) SetText(text string) {
	req.Request.Params["text"] = text
}

func (req *TbkTpwdCreateRequest) GetText() string {
	text, found := req.Request.Params["text"]
	if found {
		return text.(string)
	}
	return ""
}

// 口令弹框logoURL
func (req *TbkTpwdCreateRequest) SetLogo(logo string) {
	req.Request.Params["logo"] = logo
}

func (req *TbkTpwdCreateRequest) GetLogo() string {
	logo, found := req.Request.Params["logo"]
	if found {
		return logo.(string)
	}
	return ""
}

// 扩展字段JSON格式
func (req *TbkTpwdCreateRequest) SetExt(ext string) {
	req.Request.Params["ext"] = ext
}

func (req *TbkTpwdCreateRequest) GetExt() string {
	ext, found := req.Request.Params["ext"]
	if found {
		return ext.(string)
	}
	return ""
}

// 生成口令的淘宝用户ID
func (req *TbkTpwdCreateRequest) SetUserId(userId uint64) {
	req.Request.Params["user_id"] = userId
}

func (req *TbkTpwdCreateRequest) GetUserId() uint64 {
	userId, found := req.Request.Params["user_id"]
	if found {
		return userId.(uint64)
	}
	return 0
}
