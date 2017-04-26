package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type ZuanshiAdvertiserTargetAdzoneRptsDayGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewZuanshiAdvertiserTargetAdzoneRptsDayGetRequest() (req *ZuanshiAdvertiserTargetAdzoneRptsDayGetRequest) {
	request := topsdk.Request{MethodName: "taobao.zuanshi.advertiser.target.adzone.rpts.day.get", Params: make(map[string]interface{}, 8)}
	req = &ZuanshiAdvertiserTargetAdzoneRptsDayGetRequest{
		Request: &request,
	}
	return
}

//开始日期
func (req *ZuanshiAdvertiserTargetAdzoneRptsDayGetRequest) SetStartTime(startTime string) {
	req.Request.Params["start_time"] = startTime
}

func (req *ZuanshiAdvertiserTargetAdzoneRptsDayGetRequest) GetStartTime() string {
	startTime, found := req.Request.Params["start_time"]
	if found {
		return startTime.(string)
	}
	return ""
}

//结束日期
func (req *ZuanshiAdvertiserTargetAdzoneRptsDayGetRequest) SetEndTime(endTime string) {
	req.Request.Params["end_time"] = endTime
}

func (req *ZuanshiAdvertiserTargetAdzoneRptsDayGetRequest) GetEndTime() string {
	endTime, found := req.Request.Params["end_time"]
	if found {
		return endTime.(string)
	}
	return ""
}

// 效果周期，取值范围：3,7,15。分别表示效果转化周期-3天/7天/15天。
func (req *ZuanshiAdvertiserTargetAdzoneRptsDayGetRequest) SetEffect(effect uint) {
	req.Request.Params["effect"] = effect
}

func (req *ZuanshiAdvertiserTargetAdzoneRptsDayGetRequest) GetEffect() uint {
	effect, found := req.Request.Params["effect"]
	if found {
		return effect.(uint)
	}
	return 0
}

// 效果类型。“impression”：展现效果；“click”：点击效果
func (req *ZuanshiAdvertiserTargetAdzoneRptsDayGetRequest) SetEffectType(effectType string) {
	req.Request.Params["effect_type"] = effectType
}

func (req *ZuanshiAdvertiserTargetAdzoneRptsDayGetRequest) GetEffectType() string {
	effectType, found := req.Request.Params["effect_type"]
	if found {
		return effectType.(string)
	}
	return ""
}

// 计划id
func (req *ZuanshiAdvertiserTargetAdzoneRptsDayGetRequest) SetCampaignId(id uint64) {
	req.Request.Params["campaign_id"] = id
}

func (req *ZuanshiAdvertiserTargetAdzoneRptsDayGetRequest) GetCampaignId() uint64 {
	id, found := req.Request.Params["campaign_id"]
	if found {
		return id.(uint64)
	}
	return 0
}

// 单元id
func (req *ZuanshiAdvertiserTargetAdzoneRptsDayGetRequest) SetAdgroupId(id uint64) {
	req.Request.Params["adgroup_id"] = id
}

func (req *ZuanshiAdvertiserTargetAdzoneRptsDayGetRequest) GetAdgroupId() uint64 {
	id, found := req.Request.Params["adgroup_id"]
	if found {
		return id.(uint64)
	}
	return 0
}

// 定向id
func (req *ZuanshiAdvertiserTargetAdzoneRptsDayGetRequest) SetTargetId(id uint64) {
	req.Request.Params["target_id"] = id
}

func (req *ZuanshiAdvertiserTargetAdzoneRptsDayGetRequest) GetTargetId() uint64 {
	id, found := req.Request.Params["target_id"]
	if found {
		return id.(uint64)
	}
	return 0
}

// 资源位id
func (req *ZuanshiAdvertiserTargetAdzoneRptsDayGetRequest) SetAdzoneId(id uint64) {
	req.Request.Params["adzone_id"] = id
}

func (req *ZuanshiAdvertiserTargetAdzoneRptsDayGetRequest) GetAdzoneId() uint64 {
	id, found := req.Request.Params["adzone_id"]
	if found {
		return id.(uint64)
	}
	return 0
}
