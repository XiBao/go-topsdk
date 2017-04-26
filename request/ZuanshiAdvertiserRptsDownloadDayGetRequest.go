package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type ZuanshiAdvertiserRptsDownloadDayGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewZuanshiAdvertiserRptsDownloadDayGetRequest() (req *ZuanshiAdvertiserRptsDownloadDayGetRequest) {
	request := topsdk.Request{MethodName: "taobao.zuanshi.advertiser.rpts.download.day.get", Params: make(map[string]interface{}, 4)}
	req = &ZuanshiAdvertiserRptsDownloadDayGetRequest{
		Request: &request,
	}
	return
}

//开始日期
func (req *ZuanshiAdvertiserRptsDownloadDayGetRequest) SetStartTime(startTime string) {
	req.Request.Params["start_time"] = startTime
}

func (req *ZuanshiAdvertiserRptsDownloadDayGetRequest) GetStartTime() string {
	startTime, found := req.Request.Params["start_time"]
	if found {
		return startTime.(string)
	}
	return ""
}

//结束日期
func (req *ZuanshiAdvertiserRptsDownloadDayGetRequest) SetEndTime(endTime string) {
	req.Request.Params["end_time"] = endTime
}

func (req *ZuanshiAdvertiserRptsDownloadDayGetRequest) GetEndTime() string {
	endTime, found := req.Request.Params["end_time"]
	if found {
		return endTime.(string)
	}
	return ""
}

// 下载那种广告实体的数据取值范围 campaign,adgroup,creative,target,adzone,targetAdzone。一次一个，不能多个。
func (req *ZuanshiAdvertiserRptsDownloadDayGetRequest) SetHierarchy(hierarchy string) {
	req.Request.Params["hierarchy"] = hierarchy
}

func (req *ZuanshiAdvertiserRptsDownloadDayGetRequest) GetHierarchy() string {
	hierarchy, found := req.Request.Params["hierarchy"]
	if found {
		return hierarchy.(string)
	}
	return ""
}

// 计划类型。1：全店推广；4单品推广
func (req *ZuanshiAdvertiserRptsDownloadDayGetRequest) SetCampaignModel(model uint8) {
	req.Request.Params["campaign_model"] = model
}

func (req *ZuanshiAdvertiserRptsDownloadDayGetRequest) GetCampaignModel() uint8 {
	model, found := req.Request.Params["campaign_model"]
	if found {
		return model.(uint8)
	}
	return 0
}

// 效果类型。“impression”：展现效果；“click”：点击效果
func (req *ZuanshiAdvertiserRptsDownloadDayGetRequest) SetEffectType(effectType string) {
	req.Request.Params["effect_type"] = effectType
}

func (req *ZuanshiAdvertiserRptsDownloadDayGetRequest) GetEffectType() string {
	effectType, found := req.Request.Params["effect_type"]
	if found {
		return effectType.(string)
	}
	return ""
}
