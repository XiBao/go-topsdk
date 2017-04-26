package TopRequest

import (
	"github.com/XiBao/topsdk"
	"strconv"
	"strings"
)

type ZuanshiBannerAdgroupAdzoneUnbindRequest struct {
	Request *topsdk.Request
}

// create new request
func NewZuanshiBannerAdgroupAdzoneUnbindRequest() (req *ZuanshiBannerAdgroupAdzoneUnbindRequest) {
	request := topsdk.Request{MethodName: "taobao.zuanshi.banner.adgroup.adzone.unbind", Params: make(map[string]interface{}, 3)}
	req = &ZuanshiBannerAdgroupAdzoneUnbindRequest{
		Request: &request,
	}
	return
}

//计划ID
func (req *ZuanshiBannerAdgroupAdzoneUnbindRequest) SetCampaignId(id uint64) {
	req.Request.Params["campaign_id"] = id
}

func (req *ZuanshiBannerAdgroupAdzoneUnbindRequest) GetCampaignId() uint64 {
	id, found := req.Request.Params["campaign_id"]
	if found {
		return id.(uint64)
	}
	return 0
}

//单元ID
func (req *ZuanshiBannerAdgroupAdzoneUnbindRequest) SetAdgroupId(id uint64) {
	req.Request.Params["adgroup_id"] = id
}

func (req *ZuanshiBannerAdgroupAdzoneUnbindRequest) GetAdgroupId() uint64 {
	id, found := req.Request.Params["adgroup_id"]
	if found {
		return id.(uint64)
	}
	return 0
}

// 广告位列表
func (req *ZuanshiBannerAdgroupAdzoneUnbindRequest) SetAdzoneIdList(ids []uint64) {
	var idsStr []string
	for _, id := range ids {
		idsStr = append(idsStr, strconv.FormatUint(id, 10))
	}
	req.Request.Params["adzone_id_list"] = strings.Join(idsStr, ",")
}

func (req *ZuanshiBannerAdgroupAdzoneUnbindRequest) GetAdzoneIdList() []uint64 {
	var ids []uint64
	if idStr, found := req.Request.Params["adzone_id_list"]; found {
		idsStr := strings.Split(idStr.(string), ",")
		for _, id := range idsStr {
			if aId, err := strconv.ParseUint(id, 10, 64); err == nil && aId > 0 {
				ids = append(ids, aId)
			}
		}
	}
	return ids
}
