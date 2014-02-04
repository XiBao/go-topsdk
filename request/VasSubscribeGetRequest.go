package TopRequest

import (
    "github.com/XiBao/topsdk"
)

type VasSubscribeGetRequest struct {
    Request *topsdk.Request
}
// create new request
func NewVasSubscribeGetRequest() (req *VasSubscribeGetRequest) {
    request := topsdk.Request{MethodName:"taobao.vas.subscribe.get", Params:make(map[string]interface{}, 10)}
    req = &VasSubscribeGetRequest {
        Request: &request,
    }
    return
}

// 应用收费代码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得该应用的收费代码
func (req *VasSubscribeGetRequest) SetArticleCode(articleCode string) {
    req.Request.Params["article_code"] = articleCode
}

func (req *VasSubscribeGetRequest) GetArticleCode() string {
    articleCode, found := req.Request.Params["article_code"]
    if found {
        return articleCode.(string)
    }
    return ""
}

//淘宝会员名
func (req *VasSubscribeGetRequest) SetNick(nick string) {
    req.Request.Params["nick"] = nick
}

func (req *VasSubscribeGetRequest) GetNick() string {
    nick, found := req.Request.Params["nick"]
    if found {
        return nick.(string)
    }
    return ""
}
