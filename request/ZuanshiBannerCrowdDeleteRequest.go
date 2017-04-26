package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type ZuanshiBannerCrowdDeleteRequest struct {
	Request *topsdk.Request
}

// create new request
func NewZuanshiBannerCrowdDeleteRequest() (req *ZuanshiBannerCrowdDeleteRequest) {
	request := topsdk.Request{MethodName: "taobao.zuanshi.banner.crowd.delete", Params: make(map[string]interface{}, 3)}
	req = &ZuanshiBannerCrowdDeleteRequest{
		Request: &request,
	}
	return
}

// 计划ID
func (req *ZuanshiBannerCrowdDeleteRequest) SetCampaignId(campaignId uint64) {
	req.Request.Params["campaign_id"] = campaignId
}

func (req *ZuanshiBannerCrowdDeleteRequest) GetCampaignId() uint64 {
	if campaignId, found := req.Request.Params["campaign_id"]; found {
		return campaignId.(uint64)
	}
	return 0
}

// 单元ID
func (req *ZuanshiBannerCrowdDeleteRequest) SetAdgroupId(adgroupId uint64) {
	req.Request.Params["adgroup_id"] = adgroupId
}

func (req *ZuanshiBannerCrowdDeleteRequest) GetAdgroupId() uint64 {
	if adgroupId, found := req.Request.Params["adgroup_id"]; found {
		return adgroupId.(uint64)
	}
	return 0
}

// 绑定定向
func (req *ZuanshiBannerCrowdDeleteRequest) SetCrowds(crowds string) {
	req.Request.Params["crowds"] = crowds
}

func (req *ZuanshiBannerCrowdDeleteRequest) GetCrowds() string {
	if crowds, found := req.Request.Params["crowds"]; found {
		return crowds.(string)
	}
	return ""
}
