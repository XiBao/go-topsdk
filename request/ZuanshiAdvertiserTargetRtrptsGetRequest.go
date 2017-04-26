package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type ZuanshiAdvertiserTargetRtrptsGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewZuanshiAdvertiserTargetRtrptsGetRequest() (req *ZuanshiAdvertiserTargetRtrptsGetRequest) {
	request := topsdk.Request{MethodName: "taobao.zuanshi.advertiser.target.rtrpts.get", Params: make(map[string]interface{}, 5)}
	req = &ZuanshiAdvertiserTargetRtrptsGetRequest{
		Request: &request,
	}
	return
}

//查询日期
func (req *ZuanshiAdvertiserTargetRtrptsGetRequest) SetLogDate(logDate string) {
	req.Request.Params["log_date"] = logDate
}

func (req *ZuanshiAdvertiserTargetRtrptsGetRequest) GetLogDate() string {
	logDate, found := req.Request.Params["log_date"]
	if found {
		return logDate.(string)
	}
	return ""
}

// 计划id
func (req *ZuanshiAdvertiserTargetRtrptsGetRequest) SetCampaignId(id uint64) {
	req.Request.Params["campaign_id"] = id
}

func (req *ZuanshiAdvertiserTargetRtrptsGetRequest) GetCampaignId() uint64 {
	id, found := req.Request.Params["campaign_id"]
	if found {
		return id.(uint64)
	}
	return 0
}

// 推广单元id
func (req *ZuanshiAdvertiserTargetRtrptsGetRequest) SetAdgroupId(id uint64) {
	req.Request.Params["adgroup_id"] = id
}

func (req *ZuanshiAdvertiserTargetRtrptsGetRequest) GetAdgroupId() uint64 {
	id, found := req.Request.Params["adgroup_id"]
	if found {
		return id.(uint64)
	}
	return 0
}

// 定向标签id
func (req *ZuanshiAdvertiserTargetRtrptsGetRequest) SetTargetId(id uint64) {
	req.Request.Params["target_id"] = id
}

func (req *ZuanshiAdvertiserTargetRtrptsGetRequest) GetTargetId() uint64 {
	id, found := req.Request.Params["target_id"]
	if found {
		return id.(uint64)
	}
	return 0
}

// 计划类型。1：全店推广；4单品推广
func (req *ZuanshiAdvertiserTargetRtrptsGetRequest) SetCampaignModel(model uint8) {
	req.Request.Params["campaign_model"] = model
}

func (req *ZuanshiAdvertiserTargetRtrptsGetRequest) GetCampaignModel() uint8 {
	model, found := req.Request.Params["campaign_model"]
	if found {
		return model.(uint8)
	}
	return 0
}
