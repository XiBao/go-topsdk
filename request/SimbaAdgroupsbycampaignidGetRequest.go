package TopRequest

import (
    "github.com/XiBao/topsdk"
)

type SimbaAdgroupsbycampaignidGetRequest struct {
    Request *topsdk.Request
}
// create new request
func NewSimbaAdgroupsbycampaignidGetRequest() (req *SimbaAdgroupsbycampaignidGetRequest) {
    request := topsdk.Request{MethodName:"taobao.simba.adgroupsbycampaignid.get", Params:make(map[string]interface{}, 5)}
    req = &SimbaAdgroupsbycampaignidGetRequest {
        Request: &request,
    }
    return
}

// 推广计划Id
func (req *SimbaAdgroupsbycampaignidGetRequest) SetCampaignId(campaign_id uint64) {
    req.Request.Params["campaign_id"] = campaign_id
}

func (req *SimbaAdgroupsbycampaignidGetRequest) GetCampaignId() uint64 {
    campaign_id, found := req.Request.Params["campaign_id"]
    if found {
        return campaign_id.(uint64)
    }
    return 0
}

//主人昵称
func (req *SimbaAdgroupsbycampaignidGetRequest) SetNick(nick string) {
    req.Request.Params["nick"] = nick
}

func (req *SimbaAdgroupsbycampaignidGetRequest) GetNick() string {
    nick, found := req.Request.Params["nick"]
    if found {
        return nick.(string)
    }
    return ""
}

//页码，从1开始
func (req *SimbaAdgroupsbycampaignidGetRequest) SetPageNo(page_no uint16) {
    req.Request.Params["page_no"] = page_no
}

func (req *SimbaAdgroupsbycampaignidGetRequest) GetPageNo() uint16 {
    page_no, found := req.Request.Params["page_no"]
    if found {
        return page_no.(uint16)
    }
    return 0
}

// 页尺寸，最大200，如果入参adgroup_ids有传入值，则page_size和page_no值不起作用。如果adgrpup_ids为空而campaign_id有值，此时page_size和page_no值才是返回的页数据大小和页码
func (req *SimbaAdgroupsbycampaignidGetRequest) SetPageSize(page_size uint16) {
    req.Request.Params["page_size"] = page_size
}

func (req *SimbaAdgroupsbycampaignidGetRequest) GetPageSize() uint16 {
    page_size, found := req.Request.Params["page_size"]
    if found {
        return page_size.(uint16)
    }
    return 0
}
