package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type ZuanshiAdvertiserAdzoneRtrptsGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewZuanshiAdvertiserAdzoneRtrptsGetRequest() (req *ZuanshiAdvertiserAdzoneRtrptsGetRequest) {
	request := topsdk.Request{MethodName: "taobao.zuanshi.advertiser.adzone.rtrpts.get", Params: make(map[string]interface{}, 5)}
	req = &ZuanshiAdvertiserAdzoneRtrptsGetRequest{
		Request: &request,
	}
	return
}

//查询日期
func (req *ZuanshiAdvertiserAdzoneRtrptsGetRequest) SetLogDate(logDate string) {
	req.Request.Params["log_date"] = logDate
}

func (req *ZuanshiAdvertiserAdzoneRtrptsGetRequest) GetLogDate() string {
	logDate, found := req.Request.Params["log_date"]
	if found {
		return logDate.(string)
	}
	return ""
}

// 计划id
func (req *ZuanshiAdvertiserAdzoneRtrptsGetRequest) SetCampaignId(id uint64) {
	req.Request.Params["campaign_id"] = id
}

func (req *ZuanshiAdvertiserAdzoneRtrptsGetRequest) GetCampaignId() uint64 {
	id, found := req.Request.Params["campaign_id"]
	if found {
		return id.(uint64)
	}
	return 0
}

// 推广单元id
func (req *ZuanshiAdvertiserAdzoneRtrptsGetRequest) SetAdgroupId(id uint64) {
	req.Request.Params["adgroup_id"] = id
}

func (req *ZuanshiAdvertiserAdzoneRtrptsGetRequest) GetAdgroupId() uint64 {
	id, found := req.Request.Params["adgroup_id"]
	if found {
		return id.(uint64)
	}
	return 0
}

// 定向标签id
func (req *ZuanshiAdvertiserAdzoneRtrptsGetRequest) SetAdzoneId(id uint64) {
	req.Request.Params["adzone_id"] = id
}

func (req *ZuanshiAdvertiserAdzoneRtrptsGetRequest) GetAdzoneId() uint64 {
	id, found := req.Request.Params["adzone_id"]
	if found {
		return id.(uint64)
	}
	return 0
}

// 计划类型。1：全店推广；4单品推广
func (req *ZuanshiAdvertiserAdzoneRtrptsGetRequest) SetCampaignModel(model uint8) {
	req.Request.Params["campaign_model"] = model
}

func (req *ZuanshiAdvertiserAdzoneRtrptsGetRequest) GetCampaignModel() uint8 {
	model, found := req.Request.Params["campaign_model"]
	if found {
		return model.(uint8)
	}
	return 0
}
