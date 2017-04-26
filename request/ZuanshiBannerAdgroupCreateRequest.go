package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type ZuanshiBannerAdgroupCreateRequest struct {
	Request *topsdk.Request
}

// create new request
func NewZuanshiBannerAdgroupCreateRequest() (req *ZuanshiBannerAdgroupCreateRequest) {
	request := topsdk.Request{MethodName: "taobao.zuanshi.banner.adgroup.create", Params: make(map[string]interface{}, 5)}
	req = &ZuanshiBannerAdgroupCreateRequest{
		Request: &request,
	}
	return
}

// 计划ID
func (req *ZuanshiBannerAdgroupCreateRequest) SetCampaignId(campaignId uint64) {
	req.Request.Params["campaign_id"] = campaignId
}

func (req *ZuanshiBannerAdgroupCreateRequest) GetCampaignId() uint64 {
	if campaignId, found := req.Request.Params["campaign_id"]; found {
		return campaignId.(uint64)
	}
	return 0
}

// 智能出价，0：不使用,1：优化进店，2：优化成交，cpc不能选择2优化成交
func (req *ZuanshiBannerAdgroupCreateRequest) SetIntelligentBid(intelligentBid uint8) {
	req.Request.Params["intelligent_bid"] = intelligentBid
}

func (req *ZuanshiBannerAdgroupCreateRequest) GetIntelligentBid() uint8 {
	if intelligentBid, found := req.Request.Params["intelligent_bid"]; found {
		return intelligentBid.(uint8)
	}
	return 0
}

// 单元名称
func (req *ZuanshiBannerAdgroupCreateRequest) SetName(name string) {
	req.Request.Params["name"] = name
}

func (req *ZuanshiBannerAdgroupCreateRequest) GetName() string {
	if name, found := req.Request.Params["name"]; found {
		return name.(string)
	}
	return ""
}

// 绑定定向
func (req *ZuanshiBannerAdgroupCreateRequest) SetCrowds(crowds string) {
	req.Request.Params["crowds"] = crowds
}

func (req *ZuanshiBannerAdgroupCreateRequest) GetCrowds() string {
	if crowds, found := req.Request.Params["crowds"]; found {
		return crowds.(string)
	}
	return ""
}

// 资源位列表
func (req *ZuanshiBannerAdgroupCreateRequest) SetAdzoneBidList(adzoneBidList string) {
	req.Request.Params["adzone_bid_list"] = adzoneBidList
}

func (req *ZuanshiBannerAdgroupCreateRequest) GetAdzoneBidList() string {
	if adzoneBidList, found := req.Request.Params["adzone_bid_list"]; found {
		return adzoneBidList.(string)
	}
	return ""
}
