package TopRequest

import (
	"github.com/XiBao/topsdk"
	"strconv"
	"strings"
)

type ZuanshiBannerAdgroupDeleteRequest struct {
	Request *topsdk.Request
}

// create new request
func NewZuanshiBannerAdgroupDeleteRequest() (req *ZuanshiBannerAdgroupDeleteRequest) {
	request := topsdk.Request{MethodName: "taobao.zuanshi.banner.adgroup.delete", Params: make(map[string]interface{}, 2)}
	req = &ZuanshiBannerAdgroupDeleteRequest{
		Request: &request,
	}
	return
}

// 单元ID
func (req *ZuanshiBannerAdgroupDeleteRequest) SetAdgroupIdList(ids []uint64) {
	var idsStr []string
	for _, id := range ids {
		idsStr = append(idsStr, strconv.FormatUint(id, 10))
	}
	req.Request.Params["adgroup_id_list"] = idsStr
}

func (req *ZuanshiBannerAdgroupDeleteRequest) GetAdgroupIdList() []uint64 {
	var ids []uint64
	if idStr, found := req.Request.Params["adgroup_id_list"]; found {
		idsStr := strings.Split(idStr.(string), ",")
		for _, id := range idsStr {
			if aId, err := strconv.ParseUint(id, 10, 64); err == nil && aId > 0 {
				ids = append(ids, aId)
			}
		}
	}
	return ids
}

// 计划ID
func (req *ZuanshiBannerAdgroupDeleteRequest) SetCampaignId(campaignId uint64) {
	req.Request.Params["campaign_id"] = campaignId
}

func (req *ZuanshiBannerAdgroupDeleteRequest) GetCampaignId() uint64 {
	if campaignId, found := req.Request.Params["campaign_id"]; found {
		return campaignId.(uint64)
	}
	return 0
}
