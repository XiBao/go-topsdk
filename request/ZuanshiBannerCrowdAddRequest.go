package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type ZuanshiBannerCrowdAddRequest struct {
	Request *topsdk.Request
}

// create new request
func NewZuanshiBannerCrowdAddRequest() (req *ZuanshiBannerCrowdAddRequest) {
	request := topsdk.Request{MethodName: "taobao.zuanshi.banner.crowd.add", Params: make(map[string]interface{}, 3)}
	req = &ZuanshiBannerCrowdAddRequest{
		Request: &request,
	}
	return
}

// 计划ID
func (req *ZuanshiBannerCrowdAddRequest) SetCampaignId(campaignId uint64) {
	req.Request.Params["campaign_id"] = campaignId
}

func (req *ZuanshiBannerCrowdAddRequest) GetCampaignId() uint64 {
	if campaignId, found := req.Request.Params["campaign_id"]; found {
		return campaignId.(uint64)
	}
	return 0
}

// 单元ID
func (req *ZuanshiBannerCrowdAddRequest) SetAdgroupId(adgroupId uint64) {
	req.Request.Params["adgroup_id"] = adgroupId
}

func (req *ZuanshiBannerCrowdAddRequest) GetAdgroupId() uint64 {
	if adgroupId, found := req.Request.Params["adgroup_id"]; found {
		return adgroupId.(uint64)
	}
	return 0
}

// 绑定定向
func (req *ZuanshiBannerCrowdAddRequest) SetCrowds(crowds string) {
	req.Request.Params["crowds"] = crowds
}

func (req *ZuanshiBannerCrowdAddRequest) GetCrowds() string {
	if crowds, found := req.Request.Params["crowds"]; found {
		return crowds.(string)
	}
	return ""
}
