package TopRequest

import (
    "github.com/XiBao/topsdk"
)

type SimbaAdgroupsItemExistRequest struct {
    Request *topsdk.Request
}
// create new request
func NewSimbaAdgroupsItemExistRequest() (req *SimbaAdgroupsItemExistRequest) {
    request := topsdk.Request{MethodName:"taobao.simba.adgroups.item.exist", Params:make(map[string]interface{}, 3)}
    req = &SimbaAdgroupsItemExistRequest {
        Request: &request,
    }
    return
}

// 推广计划Id
func (req *SimbaAdgroupsItemExistRequest) SetCampaignId(campaign_id uint64) {
    req.Request.Params["campaign_id"] = campaign_id
}

func (req *SimbaAdgroupsItemExistRequest) GetCampaignId() uint64 {
    campaign_id, found := req.Request.Params["campaign_id"]
    if found {
        return campaign_id.(uint64)
    }
    return 0
}

// 商品Id
func (req *SimbaAdgroupsItemExistRequest) SetItemId(item_id uint64) {
    req.Request.Params["item_id"] = item_id
}

func (req *SimbaAdgroupsItemExistRequest) GetItemId() uint64 {
    item_id, found := req.Request.Params["item_id"]
    if found {
        return item_id.(uint64)
    }
    return 0
}

//主人昵称
func (req *SimbaAdgroupsItemExistRequest) SetNick(nick string) {
    req.Request.Params["nick"] = nick
}

func (req *SimbaAdgroupsItemExistRequest) GetNick() string {
    nick, found := req.Request.Params["nick"]
    if found {
        return nick.(string)
    }
    return ""
}
