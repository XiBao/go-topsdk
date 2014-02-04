package TopRequest

import (
    "github.com/XiBao/topsdk"
)

type ShopGetRequest struct {
    Request *topsdk.Request
}
// create new request
func NewShopGetRequest() (req *ShopGetRequest) {
    request := topsdk.Request{MethodName:"taobao.shop.get", Params:make(map[string]interface{}, 2)}
    req = &ShopGetRequest {
        Request: &request,
    }
    return
}

// 需返回的字段列表。可选值：Shop 结构中的所有字段；多个字段之间用逗号(,)分隔
func (req *ShopGetRequest) SetFields(fields string) {
    req.Request.Params["fields"] = fields
}

func (req *ShopGetRequest) GetFields() string {
    fields, found := req.Request.Params["fields"]
    if found {
        return fields.(string)
    }
    return ""
}

func (req *ShopGetRequest) SetNick(nick string) {
    req.Request.Params["nick"] = nick 
}

func (req *ShopGetRequest) GetNick() string {
    nick, found := req.Request.Params["nick"]
    if found {
        return nick.(string)
    }
    return ""
}
