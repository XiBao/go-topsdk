package TopRequest

import (
    "github.com/XiBao/topsdk"
)

type SimbaAdgroupDeleteRequest struct {
    Request *topsdk.Request
}
// create new request
func NewSimbaAdgroupDeleteRequest() (req *SimbaAdgroupDeleteRequest) {
    request := topsdk.Request{MethodName:"taobao.simba.adgroup.delete", Params:make(map[string]interface{}, 2)}
    req = &SimbaAdgroupDeleteRequest {
        Request: &request,
    }
    return
}

// 推广组Id
func (req *SimbaAdgroupDeleteRequest) SetAdgroupId(adgroup_id uint64) {
    req.Request.Params["adgroup_id"] = adgroup_id
}

func (req *SimbaAdgroupDeleteRequest) GetAdgroupId() uint64 {
    adgroup_id, found := req.Request.Params["adgroup_id"]
    if found {
        return adgroup_id.(uint64)
    }
    return 0
}

//主人昵称
func (req *SimbaAdgroupDeleteRequest) SetNick(nick string) {
    req.Request.Params["nick"] = nick
}

func (req *SimbaAdgroupDeleteRequest) GetNick() string {
    nick, found := req.Request.Params["nick"]
    if found {
        return nick.(string)
    }
    return ""
}