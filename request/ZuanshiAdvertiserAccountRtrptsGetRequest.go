package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type ZuanshiAdvertiserAccountRtrptsGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewZuanshiAdvertiserAccountRtrptsGetRequest() (req *ZuanshiAdvertiserAccountRtrptsGetRequest) {
	request := topsdk.Request{MethodName: "taobao.zuanshi.advertiser.account.rtrpts.get", Params: make(map[string]interface{}, 2)}
	req = &ZuanshiAdvertiserAccountRtrptsGetRequest{
		Request: &request,
	}
	return
}

//查询日期
func (req *ZuanshiAdvertiserAccountRtrptsGetRequest) SetLogDate(logDate string) {
	req.Request.Params["log_date"] = logDate
}

func (req *ZuanshiAdvertiserAccountRtrptsGetRequest) GetLogDate() string {
	logDate, found := req.Request.Params["log_date"]
	if found {
		return logDate.(string)
	}
	return ""
}

// 计划类型。1：全店推广；4单品推广
func (req *ZuanshiAdvertiserAccountRtrptsGetRequest) SetCampaignModel(model uint8) {
	req.Request.Params["campaign_model"] = model
}

func (req *ZuanshiAdvertiserAccountRtrptsGetRequest) GetCampaignModel() uint8 {
	model, found := req.Request.Params["campaign_model"]
	if found {
		return model.(uint8)
	}
	return 0
}
