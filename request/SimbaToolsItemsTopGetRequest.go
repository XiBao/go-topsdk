package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type SimbaToolsItemsTopGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewSimbaToolsItemsTopGetRequest() (req *SimbaToolsItemsTopGetRequest) {
	request := topsdk.Request{MethodName: "taobao.simba.tools.items.top.get", Params: make(map[string]interface{}, 3)}
	req = &SimbaToolsItemsTopGetRequest{
		Request: &request,
	}
	return
}

//昵称
func (req *SimbaToolsItemsTopGetRequest) SetNick(nick string) {
	req.Request.Params["nick"] = nick
}

func (req *SimbaToolsItemsTopGetRequest) GetNick() string {
	nick, found := req.Request.Params["nick"]
	if found {
		return nick.(string)
	}
	return ""
}

//关键词
func (req *SimbaToolsItemsTopGetRequest) SetKeyword(kw string) {
	req.Request.Params["keyword"] = kw
}

func (req *SimbaToolsItemsTopGetRequest) GetKeyword() string {
	if kw, found := req.Request.Params["keyword"]; found {
		return kw.(string)
	}
	return ""
}

//输入的必须是一个符合ipv4或者ipv6格式的IP地址
func (req *SimbaToolsItemsTopGetRequest) SetIp(ip string) {
	req.Request.Params["ip"] = ip
}

func (req *SimbaToolsItemsTopGetRequest) GetIp() string {
	if ip, found := req.Request.Params["ip"]; found {
		return ip.(string)
	}
	return ""
}
