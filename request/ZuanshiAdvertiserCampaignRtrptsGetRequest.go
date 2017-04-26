package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type ZuanshiAdvertiserCampaignRtrptsGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewZuanshiAdvertiserCampaignRtrptsGetRequest() (req *ZuanshiAdvertiserCampaignRtrptsGetRequest) {
	request := topsdk.Request{MethodName: "taobao.zuanshi.advertiser.campaign.rtrpts.get", Params: make(map[string]interface{}, 3)}
	req = &ZuanshiAdvertiserCampaignRtrptsGetRequest{
		Request: &request,
	}
	return
}

//查询日期
func (req *ZuanshiAdvertiserCampaignRtrptsGetRequest) SetLogDate(logDate string) {
	req.Request.Params["log_date"] = logDate
}

func (req *ZuanshiAdvertiserCampaignRtrptsGetRequest) GetLogDate() string {
	logDate, found := req.Request.Params["log_date"]
	if found {
		return logDate.(string)
	}
	return ""
}

// 计划id
func (req *ZuanshiAdvertiserCampaignRtrptsGetRequest) SetCampaignId(id uint64) {
	req.Request.Params["campaign_id"] = id
}

func (req *ZuanshiAdvertiserCampaignRtrptsGetRequest) GetCampaignId() uint64 {
	id, found := req.Request.Params["campaign_id"]
	if found {
		return id.(uint64)
	}
	return 0
}

// 计划类型。1：全店推广；4单品推广
func (req *ZuanshiAdvertiserCampaignRtrptsGetRequest) SetCampaignModel(model uint8) {
	req.Request.Params["campaign_model"] = model
}

func (req *ZuanshiAdvertiserCampaignRtrptsGetRequest) GetCampaignModel() uint8 {
	model, found := req.Request.Params["campaign_model"]
	if found {
		return model.(uint8)
	}
	return 0
}
