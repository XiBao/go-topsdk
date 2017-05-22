package TopRequest

import (
	"github.com/XiBao/topsdk"
	"strconv"
	"strings"
)

type ZuanshiBannerAdgroupFindRequest struct {
	Request *topsdk.Request
}

// create new request
func NewZuanshiBannerAdgroupFindRequest() (req *ZuanshiBannerAdgroupFindRequest) {
	request := topsdk.Request{MethodName: "taobao.zuanshi.banner.adgroup.find", Params: make(map[string]interface{}, 6)}
	req = &ZuanshiBannerAdgroupFindRequest{
		Request: &request,
	}
	return
}

// 计划ID
func (req *ZuanshiBannerAdgroupFindRequest) SetCampaignId(id uint64) {
	req.Request.Params["campaign_id"] = id
}

func (req *ZuanshiBannerAdgroupFindRequest) GetCampaignId() uint64 {
	if id, found := req.Request.Params["campaign_id"]; found {
		return id.(uint64)
	}
	return 0
}

// 计划ID
func (req *ZuanshiBannerAdgroupFindRequest) SetAdgroupIdList(ids []uint64) {
	var idsStr []string
	for _, id := range ids {
		idsStr = append(idsStr, strconv.FormatUint(id, 10))
	}
	req.Request.Params["adgroup_id_list"] = strings.Join(idsStr, ",")
}

func (req *ZuanshiBannerAdgroupFindRequest) GetAdgroupIdList() []uint64 {
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

// 状态，1：正常投放，0：暂停投放，9：结束投放
func (req *ZuanshiBannerAdgroupFindRequest) SetStatus(status uint8) {
	req.Request.Params["status"] = status
}

func (req *ZuanshiBannerAdgroupFindRequest) GetStatus() uint8 {
	if status, found := req.Request.Params["status"]; found {
		return status.(uint8)
	}
	return 0
}

// 分页大小, 默认值：15 最小值：1 最大长度：200
func (req *ZuanshiBannerAdgroupFindRequest) SetPageSize(pageSize uint) {
	if pageSize == 0 {
		pageSize = 15
	}
	req.Request.Params["page_size"] = pageSize
}

func (req *ZuanshiBannerAdgroupFindRequest) GetPageSize() uint {
	if pageSize, found := req.Request.Params["page_size"]; found {
		return pageSize.(uint)
	}
	return 15
}

// 当前页码, 默认值：1
func (req *ZuanshiBannerAdgroupFindRequest) SetPageNum(pageNum uint) {
	if pageNum == 0 {
		pageNum = 1
	}
	req.Request.Params["page_num"] = pageNum
}

func (req *ZuanshiBannerAdgroupFindRequest) GetPageNum() uint {
	if pageNum, found := req.Request.Params["page_num"]; found {
		return pageNum.(uint)
	}
	return 1
}

// 单元名称
func (req *ZuanshiBannerAdgroupFindRequest) SetName(name string) {
	req.Request.Params["name"] = name
}

func (req *ZuanshiBannerAdgroupFindRequest) GetName() string {
	if name, found := req.Request.Params["name"]; found {
		return name.(string)
	}
	return ""
}
