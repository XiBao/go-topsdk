package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type ZuanshiBannerAdgroupAdzoneBindRequest struct {
	Request *topsdk.Request
}

// create new request
func NewZuanshiBannerAdgroupAdzoneBindRequest() (req *ZuanshiBannerAdgroupAdzoneBindRequest) {
	request := topsdk.Request{MethodName: "taobao.zuanshi.banner.adgroup.adzone.bind", Params: make(map[string]interface{}, 3)}
	req = &ZuanshiBannerAdgroupAdzoneBindRequest{
		Request: &request,
	}
	return
}

//计划ID
func (req *ZuanshiBannerAdgroupAdzoneBindRequest) SetCampaignId(id uint64) {
	req.Request.Params["campaign_id"] = id
}

func (req *ZuanshiBannerAdgroupAdzoneBindRequest) GetCampaignId() uint64 {
	id, found := req.Request.Params["campaign_id"]
	if found {
		return id.(uint64)
	}
	return 0
}

//单元ID
func (req *ZuanshiBannerAdgroupAdzoneBindRequest) SetAdgroupId(id uint64) {
	req.Request.Params["adgroup_id"] = id
}

func (req *ZuanshiBannerAdgroupAdzoneBindRequest) GetAdgroupId() uint64 {
	id, found := req.Request.Params["adgroup_id"]
	if found {
		return id.(uint64)
	}
	return 0
}

// 广告位列表
func (req *ZuanshiBannerAdgroupAdzoneBindRequest) SetAdzoneBidDtoList(adzoneBidDtoList string) {
	req.Request.Params["adzone_bid_dto_list"] = adzoneBidDtoList
}

func (req *ZuanshiBannerAdgroupAdzoneBindRequest) GetAdzoneBidDtoList() string {
	if adzoneBidDtoList, found := req.Request.Params["adzone_bid_dto_list"]; found {
		return adzoneBidDtoList.(string)
	}
	return ""
}
