package TopRequest

import (
    "github.com/XiBao/topsdk"
)

type SimbaKeywordsQscoreGetRequest struct {
    Request *topsdk.Request
}
// create new request
func NewSimbaKeywordsQscoreGetRequest() (req *SimbaKeywordsQscoreGetRequest) {
    request := topsdk.Request{MethodName:"taobao.simba.keywords.qscore.get", Params:make(map[string]interface{}, 2)}
    req = &SimbaKeywordsQscoreGetRequest {
        Request: &request,
    }
    return
}

// 推广组id
func (req *SimbaKeywordsQscoreGetRequest) SetAdgroupId(adgroup_id uint64) {
    req.Request.Params["adgroup_id"] = adgroup_id
}

func (req *SimbaKeywordsQscoreGetRequest) GetAdgroupId() uint64 {
    adgroup_id, found := req.Request.Params["adgroup_id"]
    if found {
        return adgroup_id.(uint64)
    }
    return 0
}

//主人昵称
func (req *SimbaKeywordsQscoreGetRequest) SetNick(nick string) {
    req.Request.Params["nick"] = nick
}

func (req *SimbaKeywordsQscoreGetRequest) GetNick() string {
    nick, found := req.Request.Params["nick"]
    if found {
        return nick.(string)
    }
    return ""
}
