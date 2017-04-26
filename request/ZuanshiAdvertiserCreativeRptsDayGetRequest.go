package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type ZuanshiAdvertiserCreativeRptsDayGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewZuanshiAdvertiserCreativeRptsDayGetRequest() (req *ZuanshiAdvertiserCreativeRptsDayGetRequest) {
	request := topsdk.Request{MethodName: "taobao.zuanshi.advertiser.creative.rpts.day.get", Params: make(map[string]interface{}, 8)}
	req = &ZuanshiAdvertiserCreativeRptsDayGetRequest{
		Request: &request,
	}
	return
}

//开始日期
func (req *ZuanshiAdvertiserCreativeRptsDayGetRequest) SetStartTime(startTime string) {
	req.Request.Params["start_time"] = startTime
}

func (req *ZuanshiAdvertiserCreativeRptsDayGetRequest) GetStartTime() string {
	startTime, found := req.Request.Params["start_time"]
	if found {
		return startTime.(string)
	}
	return ""
}

//结束日期
func (req *ZuanshiAdvertiserCreativeRptsDayGetRequest) SetEndTime(endTime string) {
	req.Request.Params["end_time"] = endTime
}

func (req *ZuanshiAdvertiserCreativeRptsDayGetRequest) GetEndTime() string {
	endTime, found := req.Request.Params["end_time"]
	if found {
		return endTime.(string)
	}
	return ""
}

// 计划类型。1：全店推广；4单品推广
func (req *ZuanshiAdvertiserCreativeRptsDayGetRequest) SetCampaignModel(model uint8) {
	req.Request.Params["campaign_model"] = model
}

func (req *ZuanshiAdvertiserCreativeRptsDayGetRequest) GetCampaignModel() uint8 {
	model, found := req.Request.Params["campaign_model"]
	if found {
		return model.(uint8)
	}
	return 0
}

// 效果周期，取值范围：3,7,15。分别表示效果转化周期-3天/7天/15天。
func (req *ZuanshiAdvertiserCreativeRptsDayGetRequest) SetEffect(effect uint) {
	req.Request.Params["effect"] = effect
}

func (req *ZuanshiAdvertiserCreativeRptsDayGetRequest) GetEffect() uint {
	effect, found := req.Request.Params["effect"]
	if found {
		return effect.(uint)
	}
	return 0
}

// 效果类型。“impression”：展现效果；“click”：点击效果
func (req *ZuanshiAdvertiserCreativeRptsDayGetRequest) SetEffectType(effectType string) {
	req.Request.Params["effect_type"] = effectType
}

func (req *ZuanshiAdvertiserCreativeRptsDayGetRequest) GetEffectType() string {
	effectType, found := req.Request.Params["effect_type"]
	if found {
		return effectType.(string)
	}
	return ""
}

// 计划id
func (req *ZuanshiAdvertiserCreativeRptsDayGetRequest) SetCampaignId(id uint64) {
	req.Request.Params["campaign_id"] = id
}

func (req *ZuanshiAdvertiserCreativeRptsDayGetRequest) GetCampaignId() uint64 {
	id, found := req.Request.Params["campaign_id"]
	if found {
		return id.(uint64)
	}
	return 0
}

// 单元id
func (req *ZuanshiAdvertiserCreativeRptsDayGetRequest) SetAdgroupId(id uint64) {
	req.Request.Params["adgroup_id"] = id
}

func (req *ZuanshiAdvertiserCreativeRptsDayGetRequest) GetAdgroupId() uint64 {
	id, found := req.Request.Params["adgroup_id"]
	if found {
		return id.(uint64)
	}
	return 0
}

// 创意id
func (req *ZuanshiAdvertiserCreativeRptsDayGetRequest) SetCreativeId(id uint64) {
	req.Request.Params["creative_id"] = id
}

func (req *ZuanshiAdvertiserCreativeRptsDayGetRequest) GetCreativeId() uint64 {
	id, found := req.Request.Params["creative_id"]
	if found {
		return id.(uint64)
	}
	return 0
}
