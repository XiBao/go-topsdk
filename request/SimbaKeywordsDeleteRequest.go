package TopRequest

import (
    "github.com/XiBao/topsdk"
)

type SimbaKeywordsDeleteRequest struct {
    Request *topsdk.Request
}
// create new request
func NewSimbaKeywordsDeleteRequest() (req *SimbaKeywordsDeleteRequest) {
    request := topsdk.Request{MethodName:"taobao.simba.keywords.delete", Params:make(map[string]interface{}, 3)}
    req = &SimbaKeywordsDeleteRequest {
        Request: &request,
    }
    return
}

// 推广计划Id
func (req *SimbaKeywordsDeleteRequest) SetCampaignId(campaign_id uint64) {
    req.Request.Params["campaign_id"] = campaign_id
}

func (req *SimbaKeywordsDeleteRequest) GetCampaignId() uint64 {
    campaign_id, found := req.Request.Params["campaign_id"]
    if found {
        return campaign_id.(uint64)
    }
    return 0
}

//关键词Id数组，最多200个
func (req *SimbaKeywordsDeleteRequest) SetKeywordIds(keyword_ids string) {
    req.Request.Params["keyword_ids"] = keyword_ids
}

func (req *SimbaKeywordsDeleteRequest) GetKeywordIds() string {
    keyword_ids, found := req.Request.Params["keyword_ids"]
    if found {
        return keyword_ids.(string)
    }
    return ""
}

//主人昵称
func (req *SimbaKeywordsDeleteRequest) SetNick(nick string) {
    req.Request.Params["nick"] = nick
}

func (req *SimbaKeywordsDeleteRequest) GetNick() string {
    nick, found := req.Request.Params["nick"]
    if found {
        return nick.(string)
    }
    return ""
}
