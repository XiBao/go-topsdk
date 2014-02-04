package TopRequest

import (
    "github.com/XiBao/topsdk"
)

type SimbaCampaignPlatformGetRequest struct {
    Request *topsdk.Request
}
// create new request
func NewSimbaCampaignPlatformGetRequest() (req *SimbaCampaignPlatformGetRequest) {
    request := topsdk.Request{MethodName:"taobao.simba.campaign.platform.get", Params:make(map[string]interface{}, 2)}
    req = &SimbaCampaignPlatformGetRequest {
        Request: &request,
    }
    return
}

// 推广计划Id
func (req *SimbaCampaignPlatformGetRequest) SetCampaignId(campaign_id uint64) {
    req.Request.Params["campaign_id"] = campaign_id
}

func (req *SimbaCampaignPlatformGetRequest) GetCampaignId() uint64 {
    campaign_id, found := req.Request.Params["campaign_id"]
    if found {
        return campaign_id.(uint64)
    }
    return 0
}

//主人昵称
func (req *SimbaCampaignPlatformGetRequest) SetNick(nick string) {
    req.Request.Params["nick"] = nick
}

func (req *SimbaCampaignPlatformGetRequest) GetNick() string {
    nick, found := req.Request.Params["nick"]
    if found {
        return nick.(string)
    }
    return ""
}
