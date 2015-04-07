package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type SimbaAccountBalanceGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewSimbaAccountBalanceGetRequest() (req *SimbaAccountBalanceGetRequest) {
	request := topsdk.Request{MethodName: "taobao.simba.adgroups.get", Params: make(map[string]interface{}, 1)}
	req = &SimbaAccountBalanceGetRequest{
		Request: &request,
	}
	return
}

// 主人昵称
func (req *SimbaAccountBalanceGetRequest) SetNick(nick string) {
	req.Request.Params["nick"] = nick
}

func (req *SimbaAccountBalanceGetRequest) GetNick() string {
	nick, found := req.Request.Params["nick"]
	if found {
		return nick.(string)
	}
	return ""
}
