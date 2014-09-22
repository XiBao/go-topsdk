package TopRequest

import (
    "github.com/XiBao/topsdk"
)

type SimbaCampaignAreaGetRequest struct {
    Request *topsdk.Request
}
// create new request
func NewSimbaCampaignAreaGetRequest() (req *SimbaCampaignAreaGetRequest) {
    request := topsdk.Request{MethodName:"taobao.simba.campaign.area.get", Params:make(map[string]interface{}, 2)}
    req = &SimbaCampaignAreaGetRequest {
        Request: &request,
    }
    return
}

// 推广计划Id
func (req *SimbaCampaignAreaGetRequest) SetCampaignId(campaign_id uint64) {
    req.Request.Params["campaign_id"] = campaign_id
}

func (req *SimbaCampaignAreaGetRequest) GetCampaignId() uint64 {
    campaign_id, found := req.Request.Params["campaign_id"]
    if found {
        return campaign_id.(uint64)
    }
    return 0
}

//主人昵称
func (req *SimbaCampaignAreaGetRequest) SetNick(nick string) {
    req.Request.Params["nick"] = nick
}

func (req *SimbaCampaignAreaGetRequest) GetNick() string {
    nick, found := req.Request.Params["nick"]
    if found {
        return nick.(string)
    }
    return ""
}
