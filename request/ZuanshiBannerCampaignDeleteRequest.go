package TopRequest

import (
	"github.com/XiBao/topsdk"
	"strconv"
	"strings"
)

type ZuanshiBannerCampaignDeleteRequest struct {
	Request *topsdk.Request
}

// create new request
func NewZuanshiBannerCampaignDeleteRequest() (req *ZuanshiBannerCampaignDeleteRequest) {
	request := topsdk.Request{MethodName: "taobao.zuanshi.banner.campaign.delete", Params: make(map[string]interface{}, 1)}
	req = &ZuanshiBannerCampaignDeleteRequest{
		Request: &request,
	}
	return
}

// 计划ID
func (req *ZuanshiBannerCampaignDeleteRequest) SetCampaignIdList(ids []uint64) {
	var idsStr []string
	for _, id := range ids {
		idsStr = append(idsStr, strconv.FormatUint(id, 10))
	}
	req.Request.Params["campaign_id_list"] = strings.Join(idsStr, ",")
}

func (req *ZuanshiBannerCampaignDeleteRequest) GetCampaignIdList() []uint64 {
	var ids []uint64
	if idStr, found := req.Request.Params["campaign_id_list"]; found {
		idsStr := strings.Split(idStr.(string), ",")
		for _, id := range idsStr {
			if aId, err := strconv.ParseUint(id, 10, 64); err == nil && aId > 0 {
				ids = append(ids, aId)
			}
		}
	}
	return ids
}
