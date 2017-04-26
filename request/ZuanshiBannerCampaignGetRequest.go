package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type ZuanshiBannerCampaignGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewZuanshiBannerCampaignGetRequest() (req *ZuanshiBannerCampaignGetRequest) {
	request := topsdk.Request{MethodName: "taobao.zuanshi.banner.campaign.get", Params: make(map[string]interface{}, 1)}
	req = &ZuanshiBannerCampaignGetRequest{
		Request: &request,
	}
	return
}

// 计划ID
func (req *ZuanshiBannerCampaignGetRequest) SetId(id uint64) {
	req.Request.Params["id"] = id
}

func (req *ZuanshiBannerCampaignGetRequest) GetId() uint64 {
	if id, found := req.Request.Params["id"]; found {
		return id.(uint64)
	}
	return 0
}
