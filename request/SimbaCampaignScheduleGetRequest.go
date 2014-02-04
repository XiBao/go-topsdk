package TopRequest

import (
    "github.com/XiBao/topsdk"
)

type SimbaCampaignScheduleGetRequest struct {
    Request *topsdk.Request
}
// create new request
func NewSimbaCampaignScheduleGetRequest() (req *SimbaCampaignScheduleGetRequest) {
    request := topsdk.Request{MethodName:"taobao.simba.campaign.schedule.get", Params:make(map[string]interface{}, 2)}
    req = &SimbaCampaignScheduleGetRequest {
        Request: &request,
    }
    return
}

// 推广计划Id
func (req *SimbaCampaignScheduleGetRequest) SetCampaignId(campaign_id uint64) {
    req.Request.Params["campaign_id"] = campaign_id
}

func (req *SimbaCampaignScheduleGetRequest) GetCampaignId() uint64 {
    campaign_id, found := req.Request.Params["campaign_id"]
    if found {
        return campaign_id.(uint64)
    }
    return 0
}

//主人昵称
func (req *SimbaCampaignScheduleGetRequest) SetNick(nick string) {
    req.Request.Params["nick"] = nick
}

func (req *SimbaCampaignScheduleGetRequest) GetNick() string {
    nick, found := req.Request.Params["nick"]
    if found {
        return nick.(string)
    }
    return ""
}
