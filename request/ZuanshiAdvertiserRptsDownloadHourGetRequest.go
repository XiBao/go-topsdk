package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type ZuanshiAdvertiserRptsDownloadHourGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewZuanshiAdvertiserRptsDownloadHourGetRequest() (req *ZuanshiAdvertiserRptsDownloadHourGetRequest) {
	request := topsdk.Request{MethodName: "taobao.zuanshi.advertiser.rpts.download.hour.get", Params: make(map[string]interface{}, 3)}
	req = &ZuanshiAdvertiserRptsDownloadHourGetRequest{
		Request: &request,
	}
	return
}

//查询日期
func (req *ZuanshiAdvertiserRptsDownloadHourGetRequest) SetLogDate(logDate string) {
	req.Request.Params["log_date"] = logDate
}

func (req *ZuanshiAdvertiserRptsDownloadHourGetRequest) GetLogDate() string {
	logDate, found := req.Request.Params["log_date"]
	if found {
		return logDate.(string)
	}
	return ""
}

// 下载那种广告实体的数据取值范围 campaign,adgroup,creative,target,adzone,targetAdzone。一次一个，不能多个。
func (req *ZuanshiAdvertiserRptsDownloadHourGetRequest) SetHierarchy(hierarchy string) {
	req.Request.Params["hierarchy"] = hierarchy
}

func (req *ZuanshiAdvertiserRptsDownloadHourGetRequest) GetHierarchy() string {
	hierarchy, found := req.Request.Params["hierarchy"]
	if found {
		return hierarchy.(string)
	}
	return ""
}

// 计划类型。1：全店推广；4单品推广
func (req *ZuanshiAdvertiserRptsDownloadHourGetRequest) SetCampaignModel(model uint8) {
	req.Request.Params["campaign_model"] = model
}

func (req *ZuanshiAdvertiserRptsDownloadHourGetRequest) GetCampaignModel() uint8 {
	model, found := req.Request.Params["campaign_model"]
	if found {
		return model.(uint8)
	}
	return 0
}
