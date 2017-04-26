package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type ZuanshiAdvertiserTargetAdzoneRptsTotalGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewZuanshiAdvertiserTargetAdzoneRptsTotalGetRequest() (req *ZuanshiAdvertiserTargetAdzoneRptsTotalGetRequest) {
	request := topsdk.Request{MethodName: "taobao.zuanshi.advertiser.target.adzone.rpts.total.get", Params: make(map[string]interface{}, 10)}
	req = &ZuanshiAdvertiserTargetAdzoneRptsTotalGetRequest{
		Request: &request,
	}
	return
}

//开始日期
func (req *ZuanshiAdvertiserTargetAdzoneRptsTotalGetRequest) SetStartTime(startTime string) {
	req.Request.Params["start_time"] = startTime
}

func (req *ZuanshiAdvertiserTargetAdzoneRptsTotalGetRequest) GetStartTime() string {
	startTime, found := req.Request.Params["start_time"]
	if found {
		return startTime.(string)
	}
	return ""
}

//结束日期
func (req *ZuanshiAdvertiserTargetAdzoneRptsTotalGetRequest) SetEndTime(endTime string) {
	req.Request.Params["end_time"] = endTime
}

func (req *ZuanshiAdvertiserTargetAdzoneRptsTotalGetRequest) GetEndTime() string {
	endTime, found := req.Request.Params["end_time"]
	if found {
		return endTime.(string)
	}
	return ""
}

// 效果周期，取值范围：3,7,15。分别表示效果转化周期-3天/7天/15天。
func (req *ZuanshiAdvertiserTargetAdzoneRptsTotalGetRequest) SetEffect(effect uint) {
	req.Request.Params["effect"] = effect
}

func (req *ZuanshiAdvertiserTargetAdzoneRptsTotalGetRequest) GetEffect() uint {
	effect, found := req.Request.Params["effect"]
	if found {
		return effect.(uint)
	}
	return 0
}

// 效果类型。“impression”：展现效果；“click”：点击效果
func (req *ZuanshiAdvertiserTargetAdzoneRptsTotalGetRequest) SetEffectType(effectType string) {
	req.Request.Params["effect_type"] = effectType
}

func (req *ZuanshiAdvertiserTargetAdzoneRptsTotalGetRequest) GetEffectType() string {
	effectType, found := req.Request.Params["effect_type"]
	if found {
		return effectType.(string)
	}
	return ""
}

// 计划id
func (req *ZuanshiAdvertiserTargetAdzoneRptsTotalGetRequest) SetCampaignId(id uint64) {
	req.Request.Params["campaign_id"] = id
}

func (req *ZuanshiAdvertiserTargetAdzoneRptsTotalGetRequest) GetCampaignId() uint64 {
	id, found := req.Request.Params["campaign_id"]
	if found {
		return id.(uint64)
	}
	return 0
}

// 单元id
func (req *ZuanshiAdvertiserTargetAdzoneRptsTotalGetRequest) SetAdgroupId(id uint64) {
	req.Request.Params["adgroup_id"] = id
}

func (req *ZuanshiAdvertiserTargetAdzoneRptsTotalGetRequest) GetAdgroupId() uint64 {
	id, found := req.Request.Params["adgroup_id"]
	if found {
		return id.(uint64)
	}
	return 0
}

// 定向id
func (req *ZuanshiAdvertiserTargetAdzoneRptsTotalGetRequest) SetTargetId(id uint64) {
	req.Request.Params["target_id"] = id
}

func (req *ZuanshiAdvertiserTargetAdzoneRptsTotalGetRequest) GetTargetId() uint64 {
	id, found := req.Request.Params["target_id"]
	if found {
		return id.(uint64)
	}
	return 0
}

// 资源位id
func (req *ZuanshiAdvertiserTargetAdzoneRptsTotalGetRequest) SetAdzoneId(id uint64) {
	req.Request.Params["adzone_id"] = id
}

func (req *ZuanshiAdvertiserTargetAdzoneRptsTotalGetRequest) GetAdzoneId() uint64 {
	id, found := req.Request.Params["adzone_id"]
	if found {
		return id.(uint64)
	}
	return 0
}

// 每页条数，必传；每页最多200
func (req *ZuanshiAdvertiserTargetAdzoneRptsTotalGetRequest) SetPageSize(size uint) {
	req.Request.Params["page_size"] = size
}

func (req *ZuanshiAdvertiserTargetAdzoneRptsTotalGetRequest) GetPageSize() uint {
	size, found := req.Request.Params["page_size"]
	if found {
		return size.(uint)
	}
	return 0
}

// 分页偏移量
func (req *ZuanshiAdvertiserTargetAdzoneRptsTotalGetRequest) SetOffset(offset uint) {
	req.Request.Params["offset"] = offset
}

func (req *ZuanshiAdvertiserTargetAdzoneRptsTotalGetRequest) GetOffset() uint {
	offset, found := req.Request.Params["offset"]
	if found {
		return offset.(uint)
	}
	return 0
}
