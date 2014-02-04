package TopRequest

import (
    "github.com/XiBao/topsdk"
)

type SimbaCampaignBudgetGetRequest struct {
    Request *topsdk.Request
}
// create new request
func NewSimbaCampaignBudgetGetRequest() (req *SimbaCampaignBudgetGetRequest) {
    request := topsdk.Request{MethodName:"taobao.simba.campaign.budget.get", Params:make(map[string]interface{}, 2)}
    req = &SimbaCampaignBudgetGetRequest {
        Request: &request,
    }
    return
}

// 推广计划Id
func (req *SimbaCampaignBudgetGetRequest) SetCampaignId(campaign_id uint64) {
    req.Request.Params["campaign_id"] = campaign_id
}

func (req *SimbaCampaignBudgetGetRequest) GetCampaignId() uint64 {
    campaign_id, found := req.Request.Params["campaign_id"]
    if found {
        return campaign_id.(uint64)
    }
    return 0
}

//主人昵称
func (req *SimbaCampaignBudgetGetRequest) SetNick(nick string) {
    req.Request.Params["nick"] = nick
}

func (req *SimbaCampaignBudgetGetRequest) GetNick() string {
    nick, found := req.Request.Params["nick"]
    if found {
        return nick.(string)
    }
    return ""
}
