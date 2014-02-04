package TopRequest

import (
    "github.com/XiBao/topsdk"
)

type SimbaAdgroupUpdateRequest struct {
    Request *topsdk.Request
}
// create new request
func NewSimbaAdgroupUpdateRequest() (req *SimbaAdgroupUpdateRequest) {
    request := topsdk.Request{MethodName:"taobao.simba.adgroup.update", Params:make(map[string]interface{}, 6)}
    req = &SimbaAdgroupUpdateRequest {
        Request: &request,
    }
    return
}

// 推广宝贝Id
func (req *SimbaAdgroupUpdateRequest) SetAdgroupId(adgroup_id uint64) {
    req.Request.Params["adgroup_id"] = adgroup_id
}

func (req *SimbaAdgroupUpdateRequest) GetAdgroupId() uint64 {
    adgroup_id, found := req.Request.Params["adgroup_id"]
    if found {
        return adgroup_id.(uint64)
    }
    return 0
}

//推广组默认出价，单位为分，不能小于5 不能大于日最高限额
func (req *SimbaAdgroupUpdateRequest) SetDefaultPrice(default_price uint32) {
    req.Request.Params["default_price"] = default_price
}

func (req *SimbaAdgroupUpdateRequest) GetDefaultPrice() uint32 {
    default_price, found := req.Request.Params["default_price"]
    if found {
        return default_price.(uint32)
    }
    return 0
}

//主人昵称
func (req *SimbaAdgroupUpdateRequest) SetNick(nick string) {
    req.Request.Params["nick"] = nick
}

func (req *SimbaAdgroupUpdateRequest) GetNick() string {
    nick, found := req.Request.Params["nick"]
    if found {
        return nick.(string)
    }
    return ""
}

func (req *SimbaAdgroupUpdateRequest) SetNonsearchMaxPrice(nonsearch_max_price uint32) {
    req.Request.Params["nonsearch_max_price"] = nonsearch_max_price
}

func (req *SimbaAdgroupUpdateRequest) GetNonsearchMaxPrice() uint32 {
    nonsearch_max_price, found := req.Request.Params["nonsearch_max_price"]
    if found {
        return nonsearch_max_price.(uint32)
    }
    return 0
}

func (req *SimbaAdgroupUpdateRequest) SetOnlineStatus(online_status string) {
    req.Request.Params["online_status"] = online_status
}

func (req *SimbaAdgroupUpdateRequest) GetOnlineStatus() string {
    online_status, found := req.Request.Params["online_status"]
    if found {
        return online_status.(string)
    }
    return ""
}

func (req *SimbaAdgroupUpdateRequest) SetUseNonsearchDefaultPrice(use_nonsearch_default_price bool) {
    req.Request.Params["use_nonsearch_default_price"] = use_nonsearch_default_price
}

func (req *SimbaAdgroupUpdateRequest) GetUseNonsearchDefaultPrice() bool {
    use_nonsearch_default_price, found := req.Request.Params["use_nonsearch_default_price"]
    if found {
        return use_nonsearch_default_price.(bool)
    }
    return false
}
