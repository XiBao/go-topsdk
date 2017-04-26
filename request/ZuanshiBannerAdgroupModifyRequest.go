package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type ZuanshiBannerAdgroupModifyRequest struct {
	Request *topsdk.Request
}

// create new request
func NewZuanshiBannerAdgroupModifyRequest() (req *ZuanshiBannerAdgroupModifyRequest) {
	request := topsdk.Request{MethodName: "taobao.zuanshi.banner.adgroup.modify", Params: make(map[string]interface{}, 5)}
	req = &ZuanshiBannerAdgroupModifyRequest{
		Request: &request,
	}
	return
}

// 单元ID
func (req *ZuanshiBannerAdgroupModifyRequest) SetId(id uint64) {
	req.Request.Params["id"] = id
}

func (req *ZuanshiBannerAdgroupModifyRequest) GetId() uint64 {
	if id, found := req.Request.Params["id"]; found {
		return id.(uint64)
	}
	return 0
}

// 计划ID
func (req *ZuanshiBannerAdgroupModifyRequest) SetCampaignId(campaignId uint64) {
	req.Request.Params["campaign_id"] = campaignId
}

func (req *ZuanshiBannerAdgroupModifyRequest) GetCampaignId() uint64 {
	if campaignId, found := req.Request.Params["campaign_id"]; found {
		return campaignId.(uint64)
	}
	return 0
}

// 智能出价，0：不使用,1：优化进店，2：优化成交，cpc不能选择2优化成交
func (req *ZuanshiBannerAdgroupModifyRequest) SetIntelligentBid(intelligentBid uint8) {
	req.Request.Params["intelligent_bid"] = intelligentBid
}

func (req *ZuanshiBannerAdgroupModifyRequest) GetIntelligentBid() uint8 {
	if intelligentBid, found := req.Request.Params["intelligent_bid"]; found {
		return intelligentBid.(uint8)
	}
	return 0
}

// 创意优选，1：开启，0关闭，其他值默认开启，cpc不能修改这个字段
func (req *ZuanshiBannerAdgroupModifyRequest) SetAdboardFilter(adboardFilter uint8) {
	req.Request.Params["adboard_filter"] = adboardFilter
}

func (req *ZuanshiBannerAdgroupModifyRequest) GetAdboardFilter() uint8 {
	if adboardFilter, found := req.Request.Params["adboard_filter"]; found {
		return adboardFilter.(uint8)
	}
	return 0
}

// 单元名称
func (req *ZuanshiBannerAdgroupModifyRequest) SetName(name string) {
	req.Request.Params["name"] = name
}

func (req *ZuanshiBannerAdgroupModifyRequest) GetName() string {
	if name, found := req.Request.Params["name"]; found {
		return name.(string)
	}
	return ""
}
