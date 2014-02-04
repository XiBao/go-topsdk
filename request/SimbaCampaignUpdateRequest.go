package TopRequest

import (
    "github.com/XiBao/topsdk"
)

type SimbaCampaignUpdateRequest struct {
    Request *topsdk.Request
}
// create new request
func NewSimbaCampaignUpdateRequest() (req *SimbaCampaignUpdateRequest) {
    request := topsdk.Request{MethodName:"taobao.simba.campaign.update", Params:make(map[string]interface{}, 4)}
    req = &SimbaCampaignUpdateRequest {
        Request: &request,
    }
    return
}

// 推广计划Id
func (req *SimbaCampaignUpdateRequest) SetCampaignId(campaign_id uint64) {
    req.Request.Params["campaign_id"] = campaign_id
}

func (req *SimbaCampaignUpdateRequest) GetCampaignId() uint64 {
    campaign_id, found := req.Request.Params["campaign_id"]
    if found {
        return campaign_id.(uint64)
    }
    return 0
}

//用户设置的上下限状态；offline-下线；online-上线；
func (req *SimbaCampaignUpdateRequest) SetOnlineStatus(online_status string) {
    req.Request.Params["online_status"] = online_status
}

func (req *SimbaCampaignUpdateRequest) GetOnlineStatus() string {
    online_status, found := req.Request.Params["online_status"]
    if found {
        return online_status.(string)
    }
    return ""
}

//推广计划名称，不能多余20个字符，不能和客户其他推广计划同名。
func (req *SimbaCampaignUpdateRequest) SetTitle(title string) {
    req.Request.Params["title"] = title
}

func (req *SimbaCampaignUpdateRequest) GetTitle() string {
    title, found := req.Request.Params["title"]
    if found {
        return title.(string)
    }
    return ""
}

//主人昵称
func (req *SimbaCampaignUpdateRequest) SetNick(nick string) {
    req.Request.Params["nick"] = nick
}

func (req *SimbaCampaignUpdateRequest) GetNick() string {
    nick, found := req.Request.Params["nick"]
    if found {
        return nick.(string)
    }
    return ""
}
