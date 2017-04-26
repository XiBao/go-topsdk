package TopRequest

import (
	"github.com/XiBao/topsdk"
	"strconv"
	"strings"
)

type ZuanshiBannerCrowdFindRequest struct {
	Request *topsdk.Request
}

// create new request
func NewZuanshiBannerCrowdFindRequest() (req *ZuanshiBannerCrowdFindRequest) {
	request := topsdk.Request{MethodName: "taobao.zuanshi.banner.crowd.find", Params: make(map[string]interface{}, 6)}
	req = &ZuanshiBannerCrowdFindRequest{
		Request: &request,
	}
	return
}

// 计划ID
func (req *ZuanshiBannerCrowdFindRequest) SetCampaignId(campaignId uint64) {
	req.Request.Params["campaign_id"] = campaignId
}

func (req *ZuanshiBannerCrowdFindRequest) GetCampaignId() uint64 {
	if campaignId, found := req.Request.Params["campaign_id"]; found {
		return campaignId.(uint64)
	}
	return 0
}

// 单元ID
func (req *ZuanshiBannerCrowdFindRequest) SetAdgroupId(adgroupId uint64) {
	req.Request.Params["adgroup_id"] = adgroupId
}

func (req *ZuanshiBannerCrowdFindRequest) GetAdgroupId() uint64 {
	if adgroupId, found := req.Request.Params["adgroup_id"]; found {
		return adgroupId.(uint64)
	}
	return 0
}

// 定向人群ID
func (req *ZuanshiBannerCrowdFindRequest) SetTargetId(targetId uint64) {
	req.Request.Params["target_id"] = targetId
}

func (req *ZuanshiBannerCrowdFindRequest) GetTargetId() uint64 {
	if targetId, found := req.Request.Params["adgroup_id"]; found {
		return targetId.(uint64)
	}
	return 0
}

// 指定多个定向类型
func (req *ZuanshiBannerCrowdFindRequest) SetTargetTypes(ids []uint) {
	var idsStr []string
	for _, id := range ids {
		idsStr = append(idsStr, strconv.FormatUint(uint64(id), 10))
	}
	req.Request.Params["target_types"] = idsStr
}

func (req *ZuanshiBannerCrowdFindRequest) GetTargetTypes() []uint {
	var ids []uint
	if idStr, found := req.Request.Params["target_types"]; found {
		idsStr := strings.Split(idStr.(string), ",")
		for _, id := range idsStr {
			if aId, err := strconv.ParseUint(id, 10, 64); err == nil && aId > 0 {
				ids = append(ids, uint(aId))
			}
		}
	}
	return ids
}

// 分页大小, 默认值：15 最小值：1 最大长度：200
func (req *ZuanshiBannerCrowdFindRequest) SetPageSize(pageSize uint) {
	if pageSize == 0 {
		pageSize = 15
	}
	req.Request.Params["page_size"] = pageSize
}

func (req *ZuanshiBannerCrowdFindRequest) GetPageSize() uint {
	if pageSize, found := req.Request.Params["page_size"]; found {
		return pageSize.(uint)
	}
	return 15
}

// 当前页码, 默认值：1
func (req *ZuanshiBannerCrowdFindRequest) SetPageNum(pageNum uint) {
	if pageNum == 0 {
		pageNum = 1
	}
	req.Request.Params["page_num"] = pageNum
}

func (req *ZuanshiBannerCrowdFindRequest) GetPageNum() uint {
	if pageNum, found := req.Request.Params["page_num"]; found {
		return pageNum.(uint)
	}
	return 1
}
