package TopRequest

import (
	"github.com/XiBao/topsdk"
	"strconv"
	"strings"
)

type ZuanshiBannerAdgroupStatusRequest struct {
	Request *topsdk.Request
}

// create new request
func NewZuanshiBannerAdgroupStatusRequest() (req *ZuanshiBannerAdgroupStatusRequest) {
	request := topsdk.Request{MethodName: "taobao.zuanshi.banner.adgroup.status", Params: make(map[string]interface{}, 3)}
	req = &ZuanshiBannerAdgroupStatusRequest{
		Request: &request,
	}
	return
}

// 计划ID
func (req *ZuanshiBannerAdgroupStatusRequest) SetCampaignId(campaignId uint64) {
	req.Request.Params["campaign_id"] = campaignId
}

func (req *ZuanshiBannerAdgroupStatusRequest) GetCampaignId() uint64 {
	if campaignId, found := req.Request.Params["campaign_id"]; found {
		return campaignId.(uint64)
	}
	return 0
}

// 单元ID
func (req *ZuanshiBannerAdgroupStatusRequest) SetAdgroupIdList(ids []uint64) {
	var idsStr []string
	for _, id := range ids {
		idsStr = append(idsStr, strconv.FormatUint(id, 10))
	}
	req.Request.Params["adgroup_id_list"] = strings.Join(idsStr, ",")
}

func (req *ZuanshiBannerAdgroupStatusRequest) GetAdgroupIdList() []uint64 {
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

// 计划状态，0:暂停,1:投放,9:投放结束
func (req *ZuanshiBannerAdgroupStatusRequest) SetStatus(status uint8) {
	req.Request.Params["status"] = status
}

func (req *ZuanshiBannerAdgroupStatusRequest) GetStatus() uint8 {
	if status, found := req.Request.Params["status"]; found {
		return status.(uint8)
	}
	return 0
}
