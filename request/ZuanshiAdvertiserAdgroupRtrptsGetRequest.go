package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type ZuanshiAdvertiserAdgroupRtrptsGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewZuanshiAdvertiserAdgroupRtrptsGetRequest() (req *ZuanshiAdvertiserAdgroupRtrptsGetRequest) {
	request := topsdk.Request{MethodName: "taobao.zuanshi.advertiser.adgroup.rtrpts.get", Params: make(map[string]interface{}, 4)}
	req = &ZuanshiAdvertiserAdgroupRtrptsGetRequest{
		Request: &request,
	}
	return
}

//查询日期
func (req *ZuanshiAdvertiserAdgroupRtrptsGetRequest) SetLogDate(logDate string) {
	req.Request.Params["log_date"] = logDate
}

func (req *ZuanshiAdvertiserAdgroupRtrptsGetRequest) GetLogDate() string {
	logDate, found := req.Request.Params["log_date"]
	if found {
		return logDate.(string)
	}
	return ""
}

// 计划id
func (req *ZuanshiAdvertiserAdgroupRtrptsGetRequest) SetCampaignId(id uint64) {
	req.Request.Params["campaign_id"] = id
}

func (req *ZuanshiAdvertiserAdgroupRtrptsGetRequest) GetCampaignId() uint64 {
	id, found := req.Request.Params["campaign_id"]
	if found {
		return id.(uint64)
	}
	return 0
}

// 推广单元id
func (req *ZuanshiAdvertiserAdgroupRtrptsGetRequest) SetAdgroupId(id uint64) {
	req.Request.Params["adgroup_id"] = id
}

func (req *ZuanshiAdvertiserAdgroupRtrptsGetRequest) GetAdgroupId() uint64 {
	id, found := req.Request.Params["adgroup_id"]
	if found {
		return id.(uint64)
	}
	return 0
}

// 计划类型。1：全店推广；4单品推广
func (req *ZuanshiAdvertiserAdgroupRtrptsGetRequest) SetCampaignModel(model uint8) {
	req.Request.Params["campaign_model"] = model
}

func (req *ZuanshiAdvertiserAdgroupRtrptsGetRequest) GetCampaignModel() uint8 {
	model, found := req.Request.Params["campaign_model"]
	if found {
		return model.(uint8)
	}
	return 0
}
