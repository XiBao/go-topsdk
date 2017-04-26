package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type ZuanshiAdvertiserAdzoneRptsTotalGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewZuanshiAdvertiserAdzoneRptsTotalGetRequest() (req *ZuanshiAdvertiserAdzoneRptsTotalGetRequest) {
	request := topsdk.Request{MethodName: "taobao.zuanshi.advertiser.adzone.rpts.total.get", Params: make(map[string]interface{}, 10)}
	req = &ZuanshiAdvertiserAdzoneRptsTotalGetRequest{
		Request: &request,
	}
	return
}

//开始日期
func (req *ZuanshiAdvertiserAdzoneRptsTotalGetRequest) SetStartTime(startTime string) {
	req.Request.Params["start_time"] = startTime
}

func (req *ZuanshiAdvertiserAdzoneRptsTotalGetRequest) GetStartTime() string {
	startTime, found := req.Request.Params["start_time"]
	if found {
		return startTime.(string)
	}
	return ""
}

//结束日期
func (req *ZuanshiAdvertiserAdzoneRptsTotalGetRequest) SetEndTime(endTime string) {
	req.Request.Params["end_time"] = endTime
}

func (req *ZuanshiAdvertiserAdzoneRptsTotalGetRequest) GetEndTime() string {
	endTime, found := req.Request.Params["end_time"]
	if found {
		return endTime.(string)
	}
	return ""
}

// 计划类型。1：全店推广；4单品推广
func (req *ZuanshiAdvertiserAdzoneRptsTotalGetRequest) SetCampaignModel(model uint8) {
	req.Request.Params["campaign_model"] = model
}

func (req *ZuanshiAdvertiserAdzoneRptsTotalGetRequest) GetCampaignModel() uint8 {
	model, found := req.Request.Params["campaign_model"]
	if found {
		return model.(uint8)
	}
	return 0
}

// 效果周期，取值范围：3,7,15。分别表示效果转化周期-3天/7天/15天。
func (req *ZuanshiAdvertiserAdzoneRptsTotalGetRequest) SetEffect(effect uint) {
	req.Request.Params["effect"] = effect
}

func (req *ZuanshiAdvertiserAdzoneRptsTotalGetRequest) GetEffect() uint {
	effect, found := req.Request.Params["effect"]
	if found {
		return effect.(uint)
	}
	return 0
}

// 效果类型。“impression”：展现效果；“click”：点击效果
func (req *ZuanshiAdvertiserAdzoneRptsTotalGetRequest) SetEffectType(effectType string) {
	req.Request.Params["effect_type"] = effectType
}

func (req *ZuanshiAdvertiserAdzoneRptsTotalGetRequest) GetEffectType() string {
	effectType, found := req.Request.Params["effect_type"]
	if found {
		return effectType.(string)
	}
	return ""
}

// 计划id
func (req *ZuanshiAdvertiserAdzoneRptsTotalGetRequest) SetCampaignId(id uint64) {
	req.Request.Params["campaign_id"] = id
}

func (req *ZuanshiAdvertiserAdzoneRptsTotalGetRequest) GetCampaignId() uint64 {
	id, found := req.Request.Params["campaign_id"]
	if found {
		return id.(uint64)
	}
	return 0
}

// 单元id
func (req *ZuanshiAdvertiserAdzoneRptsTotalGetRequest) SetAdgroupId(id uint64) {
	req.Request.Params["adgroup_id"] = id
}

func (req *ZuanshiAdvertiserAdzoneRptsTotalGetRequest) GetAdgroupId() uint64 {
	id, found := req.Request.Params["adgroup_id"]
	if found {
		return id.(uint64)
	}
	return 0
}

// 资源位id
func (req *ZuanshiAdvertiserAdzoneRptsTotalGetRequest) SetAdzoneId(id uint64) {
	req.Request.Params["adzone_id"] = id
}

func (req *ZuanshiAdvertiserAdzoneRptsTotalGetRequest) GetAdzoneId() uint64 {
	id, found := req.Request.Params["adzone_id"]
	if found {
		return id.(uint64)
	}
	return 0
}

// 每页条数，必传；每页最多200
func (req *ZuanshiAdvertiserAdzoneRptsTotalGetRequest) SetPageSize(size uint) {
	req.Request.Params["page_size"] = size
}

func (req *ZuanshiAdvertiserAdzoneRptsTotalGetRequest) GetPageSize() uint {
	size, found := req.Request.Params["page_size"]
	if found {
		return size.(uint)
	}
	return 0
}

// 分页偏移量
func (req *ZuanshiAdvertiserAdzoneRptsTotalGetRequest) SetOffset(offset uint) {
	req.Request.Params["offset"] = offset
}

func (req *ZuanshiAdvertiserAdzoneRptsTotalGetRequest) GetOffset() uint {
	offset, found := req.Request.Params["offset"]
	if found {
		return offset.(uint)
	}
	return 0
}
