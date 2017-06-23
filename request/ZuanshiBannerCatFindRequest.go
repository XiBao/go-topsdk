package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type ZuanshiBannerCatFindRequest struct {
	Request *topsdk.Request
}

// create new request
func NewZuanshiBannerCatFindRequest() (req *ZuanshiBannerCatFindRequest) {
	request := topsdk.Request{MethodName: "taobao.zuanshi.banner.cat.find", Params: make(map[string]interface{}, 2)}
	req = &ZuanshiBannerCatFindRequest{
		Request: &request,
	}
	return
}

// 计划类型，banner_cpm:2，banner_cpc:8
func (req *ZuanshiBannerCatFindRequest) SetCampaignType(campaignType uint8) {
	req.Request.Params["campaign_type"] = campaignType
}

func (req *ZuanshiBannerCatFindRequest) GetCampaignType() uint8 {
	campaignType, found := req.Request.Params["campaign_type"]
	if found {
		return campaignType.(uint8)
	}
	return 0
}

// 类目名称
func (req *ZuanshiBannerCatFindRequest) SetInterestName(interestName string) {
	req.Request.Params["interest_name"] = interestName
}

func (req *ZuanshiBannerCatFindRequest) GetInterestName() string {
	interestName, found := req.Request.Params["interest_name"]
	if found {
		return interestName.(string)
	}
	return ""
}
