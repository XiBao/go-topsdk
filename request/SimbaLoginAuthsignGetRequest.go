package TopRequest

import (
    "github.com/XiBao/topsdk"
)

type SimbaLoginAuthsignGetRequest struct {
    Request *topsdk.Request
}
// create new request
func NewSimbaLoginAuthsignGetRequest() (req *SimbaLoginAuthsignGetRequest) {
    request := topsdk.Request{MethodName:"taobao.simba.login.authsign.get", Params:make(map[string]interface{}, 1)}
    req = &SimbaLoginAuthsignGetRequest {
        Request: &request,
    }
    return
}

func (req *SimbaLoginAuthsignGetRequest) SetNick(nick string) {
    req.Request.Params["nick"] = nick 
}

func (req *SimbaLoginAuthsignGetRequest) GetNick() string {
    nick, found := req.Request.Params["nick"]
    if found {
        return nick.(string)
    }
    return ""
}
