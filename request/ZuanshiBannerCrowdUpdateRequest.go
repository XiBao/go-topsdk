package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type ZuanshiBannerCrowdUpdateRequest struct {
	Request *topsdk.Request
}

// create new request
func NewZuanshiBannerCrowdUpdateRequest() (req *ZuanshiBannerCrowdUpdateRequest) {
	request := topsdk.Request{MethodName: "taobao.zuanshi.banner.crowd.update", Params: make(map[string]interface{}, 3)}
	req = &ZuanshiBannerCrowdUpdateRequest{
		Request: &request,
	}
	return
}

// 计划ID
func (req *ZuanshiBannerCrowdUpdateRequest) SetCampaignId(campaignId uint64) {
	req.Request.Params["campaign_id"] = campaignId
}

func (req *ZuanshiBannerCrowdUpdateRequest) GetCampaignId() uint64 {
	if campaignId, found := req.Request.Params["campaign_id"]; found {
		return campaignId.(uint64)
	}
	return 0
}

// 单元ID
func (req *ZuanshiBannerCrowdUpdateRequest) SetAdgroupId(adgroupId uint64) {
	req.Request.Params["adgroup_id"] = adgroupId
}

func (req *ZuanshiBannerCrowdUpdateRequest) GetAdgroupId() uint64 {
	if adgroupId, found := req.Request.Params["adgroup_id"]; found {
		return adgroupId.(uint64)
	}
	return 0
}

// 绑定定向
func (req *ZuanshiBannerCrowdUpdateRequest) SetCrowds(crowds string) {
	req.Request.Params["crowds"] = crowds
}

func (req *ZuanshiBannerCrowdUpdateRequest) GetCrowds() string {
	if crowds, found := req.Request.Params["crowds"]; found {
		return crowds.(string)
	}
	return ""
}
