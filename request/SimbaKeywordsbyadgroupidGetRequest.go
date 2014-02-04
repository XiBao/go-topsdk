package TopRequest

import (
    "github.com/XiBao/topsdk"
)

type SimbaKeywordsbyadgroupidGetRequest struct {
    Request *topsdk.Request
}
// create new request
func NewSimbaKeywordsbyadgroupidGetRequest() (req *SimbaKeywordsbyadgroupidGetRequest) {
    request := topsdk.Request{MethodName:"taobao.simba.keywordsbyadgroupid.get", Params:make(map[string]interface{}, 3)}
    req = &SimbaKeywordsbyadgroupidGetRequest {
        Request: &request,
    }
    return
}

// 推广组id
func (req *SimbaKeywordsbyadgroupidGetRequest) SetAdgroupId(adgroup_id uint64) {
    req.Request.Params["adgroup_id"] = adgroup_id
}

func (req *SimbaKeywordsbyadgroupidGetRequest) GetAdgroupId() uint64 {
    adgroup_id, found := req.Request.Params["adgroup_id"]
    if found {
        return adgroup_id.(uint64)
    }
    return 0
}

//主人昵称
func (req *SimbaKeywordsbyadgroupidGetRequest) SetNick(nick string) {
    req.Request.Params["nick"] = nick
}

func (req *SimbaKeywordsbyadgroupidGetRequest) GetNick() string {
    nick, found := req.Request.Params["nick"]
    if found {
        return nick.(string)
    }
    return ""
}
