package TopRequest

import (
    "github.com/XiBao/topsdk"
)

type SimbaCampaignsGetRequest struct {
    Request *topsdk.Request
}
// create new request
func NewSimbaCampaignsGetRequest() (req *SimbaCampaignsGetRequest) {
    request := topsdk.Request{MethodName:"taobao.simba.campaigns.get", Params:make(map[string]interface{}, 1)}
    req = &SimbaCampaignsGetRequest {
        Request: &request,
    }
    return
}

//主人昵称
func (req *SimbaCampaignsGetRequest) SetNick(nick string) {
    req.Request.Params["nick"] = nick
}

func (req *SimbaCampaignsGetRequest) GetNick() string {
    nick, found := req.Request.Params["nick"]
    if found {
        return nick.(string)
    }
    return ""
}
