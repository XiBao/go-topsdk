package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type ZuanshiAdvertiserCreativeRtrptsGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewZuanshiAdvertiserCreativeRtrptsGetRequest() (req *ZuanshiAdvertiserCreativeRtrptsGetRequest) {
	request := topsdk.Request{MethodName: "taobao.zuanshi.advertiser.creative.rtrpts.get", Params: make(map[string]interface{}, 5)}
	req = &ZuanshiAdvertiserCreativeRtrptsGetRequest{
		Request: &request,
	}
	return
}

//查询日期
func (req *ZuanshiAdvertiserCreativeRtrptsGetRequest) SetLogDate(logDate string) {
	req.Request.Params["log_date"] = logDate
}

func (req *ZuanshiAdvertiserCreativeRtrptsGetRequest) GetLogDate() string {
	logDate, found := req.Request.Params["log_date"]
	if found {
		return logDate.(string)
	}
	return ""
}

// 计划id
func (req *ZuanshiAdvertiserCreativeRtrptsGetRequest) SetCampaignId(id uint64) {
	req.Request.Params["campaign_id"] = id
}

func (req *ZuanshiAdvertiserCreativeRtrptsGetRequest) GetCampaignId() uint64 {
	id, found := req.Request.Params["campaign_id"]
	if found {
		return id.(uint64)
	}
	return 0
}

// 推广单元id
func (req *ZuanshiAdvertiserCreativeRtrptsGetRequest) SetAdgroupId(id uint64) {
	req.Request.Params["adgroup_id"] = id
}

func (req *ZuanshiAdvertiserCreativeRtrptsGetRequest) GetAdgroupId() uint64 {
	id, found := req.Request.Params["adgroup_id"]
	if found {
		return id.(uint64)
	}
	return 0
}

// 创意id
func (req *ZuanshiAdvertiserCreativeRtrptsGetRequest) SetCreativeId(id uint64) {
	req.Request.Params["creative_id"] = id
}

func (req *ZuanshiAdvertiserCreativeRtrptsGetRequest) GetCreativeId() uint64 {
	id, found := req.Request.Params["creative_id"]
	if found {
		return id.(uint64)
	}
	return 0
}

// 计划类型。1：全店推广；4单品推广
func (req *ZuanshiAdvertiserCreativeRtrptsGetRequest) SetCampaignModel(model uint8) {
	req.Request.Params["campaign_model"] = model
}

func (req *ZuanshiAdvertiserCreativeRtrptsGetRequest) GetCampaignModel() uint8 {
	model, found := req.Request.Params["campaign_model"]
	if found {
		return model.(uint8)
	}
	return 0
}
