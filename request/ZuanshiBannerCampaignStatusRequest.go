package TopRequest

import (
	"github.com/XiBao/topsdk"
	"strconv"
	"strings"
)

type ZuanshiBannerCampaignStatusRequest struct {
	Request *topsdk.Request
}

// create new request
func NewZuanshiBannerCampaignStatusRequest() (req *ZuanshiBannerCampaignStatusRequest) {
	request := topsdk.Request{MethodName: "taobao.zuanshi.banner.campaign.status", Params: make(map[string]interface{}, 2)}
	req = &ZuanshiBannerCampaignStatusRequest{
		Request: &request,
	}
	return
}

// 计划ID
func (req *ZuanshiBannerCampaignStatusRequest) SetCampaignIdList(ids []uint64) {
	var idsStr []string
	for _, id := range ids {
		idsStr = append(idsStr, strconv.FormatUint(id, 10))
	}
	req.Request.Params["campaign_id_list"] = idsStr
}

func (req *ZuanshiBannerCampaignStatusRequest) GetCampaignIdList() []uint64 {
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

// 计划状态，0:暂停,1:投放,9:投放结束
func (req *ZuanshiBannerCampaignStatusRequest) SetStatus(status uint8) {
	req.Request.Params["status"] = status
}

func (req *ZuanshiBannerCampaignStatusRequest) GetStatus() uint8 {
	if status, found := req.Request.Params["status"]; found {
		return status.(uint8)
	}
	return 0
}
