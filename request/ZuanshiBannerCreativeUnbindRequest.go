package TopRequest

import (
	"github.com/XiBao/topsdk"
	"strconv"
	"strings"
)

type ZuanshiBannerCreativeUnbindRequest struct {
	Request *topsdk.Request
}

// create new request
func NewZuanshiBannerCreativeUnbindRequest() (req *ZuanshiBannerCreativeUnbindRequest) {
	request := topsdk.Request{MethodName: "taobao.zuanshi.banner.creative.unbind", Params: make(map[string]interface{}, 3)}
	req = &ZuanshiBannerCreativeUnbindRequest{
		Request: &request,
	}
	return
}

//计划ID
func (req *ZuanshiBannerCreativeUnbindRequest) SetCampaignId(id uint64) {
	req.Request.Params["campaign_id"] = id
}

func (req *ZuanshiBannerCreativeUnbindRequest) GetCampaignId() uint64 {
	id, found := req.Request.Params["campaign_id"]
	if found {
		return id.(uint64)
	}
	return 0
}

//单元ID
func (req *ZuanshiBannerCreativeUnbindRequest) SetAdgroupId(id uint64) {
	req.Request.Params["adgroup_id"] = id
}

func (req *ZuanshiBannerCreativeUnbindRequest) GetAdgroupId() uint64 {
	id, found := req.Request.Params["adgroup_id"]
	if found {
		return id.(uint64)
	}
	return 0
}

// 创意ID列表
func (req *ZuanshiBannerCreativeUnbindRequest) SetCreativeIdList(ids []uint64) {
	var idsStr []string
	for _, id := range ids {
		idsStr = append(idsStr, strconv.FormatUint(id, 10))
	}
	req.Request.Params["creative_id_list"] = strings.Join(idsStr, ",")
}

func (req *ZuanshiBannerCreativeUnbindRequest) GetCreativeIdList() []uint64 {
	var ids []uint64
	if idStr, found := req.Request.Params["creative_id_list"]; found {
		idsStr := strings.Split(idStr.(string), ",")
		for _, id := range idsStr {
			if aId, err := strconv.ParseUint(id, 10, 64); err == nil && aId > 0 {
				ids = append(ids, aId)
			}
		}
	}
	return ids
}
