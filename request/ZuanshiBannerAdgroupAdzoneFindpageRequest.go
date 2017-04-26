package TopRequest

import (
	"github.com/XiBao/topsdk"
	"strconv"
	"strings"
)

type ZuanshiBannerAdgroupAdzoneFindpageRequest struct {
	Request *topsdk.Request
}

// create new request
func NewZuanshiBannerAdgroupAdzoneFindpageRequest() (req *ZuanshiBannerAdgroupAdzoneFindpageRequest) {
	request := topsdk.Request{MethodName: "taobao.zuanshi.banner.adgroup.adzone.findpage", Params: make(map[string]interface{}, 9)}
	req = &ZuanshiBannerAdgroupAdzoneFindpageRequest{
		Request: &request,
	}
	return
}

// 广告位id列表
func (req *ZuanshiBannerAdgroupAdzoneFindpageRequest) SetAdzoneIdList(ids []uint64) {
	var idsStr []string
	for _, id := range ids {
		idsStr = append(idsStr, strconv.FormatUint(id, 10))
	}
	req.Request.Params["adzone_id_list"] = strings.Join(idsStr, ",")
}

func (req *ZuanshiBannerAdgroupAdzoneFindpageRequest) GetAdzoneIdList() []uint64 {
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

// 计划id
func (req *ZuanshiBannerAdgroupAdzoneFindpageRequest) SetCampaignId(id uint64) {
	req.Request.Params["campaign_id"] = id
}

func (req *ZuanshiBannerAdgroupAdzoneFindpageRequest) GetCampaignId() uint64 {
	if id, found := req.Request.Params["campaign_id"]; found {
		return id.(uint64)
	}
	return 0
}

// 推广单元id
func (req *ZuanshiBannerAdgroupAdzoneFindpageRequest) SetAdgroupId(id uint64) {
	req.Request.Params["adgroup_id"] = id
}

func (req *ZuanshiBannerAdgroupAdzoneFindpageRequest) GetAdgroupId() uint64 {
	if id, found := req.Request.Params["adgroup_id"]; found {
		return id.(uint64)
	}
	return 0
}

// 分页大小, 默认值：15 最小值：1 最大长度：200
func (req *ZuanshiBannerAdgroupAdzoneFindpageRequest) SetPageSize(pageSize uint) {
	if pageSize == 0 {
		pageSize = 15
	}
	req.Request.Params["page_size"] = pageSize
}

func (req *ZuanshiBannerAdgroupAdzoneFindpageRequest) GetPageSize() uint {
	if pageSize, found := req.Request.Params["page_size"]; found {
		return pageSize.(uint)
	}
	return 15
}

// 当前页码, 默认值：1
func (req *ZuanshiBannerAdgroupAdzoneFindpageRequest) SetPageNum(pageNum uint) {
	if pageNum == 0 {
		pageNum = 1
	}
	req.Request.Params["page_num"] = pageNum
}

func (req *ZuanshiBannerAdgroupAdzoneFindpageRequest) GetPageNum() uint {
	if pageNum, found := req.Request.Params["page_num"]; found {
		return pageNum.(uint)
	}
	return 1
}

// 广告位名称
func (req *ZuanshiBannerAdgroupAdzoneFindpageRequest) SetAdzoneName(name string) {
	req.Request.Params["adzone_name"] = name
}

func (req *ZuanshiBannerAdgroupAdzoneFindpageRequest) GetAdzoneName() string {
	if name, found := req.Request.Params["adzone_name"]; found {
		return name.(string)
	}
	return ""
}

// 广告位尺寸
func (req *ZuanshiBannerAdgroupAdzoneFindpageRequest) SetAdzoneSize(size string) {
	req.Request.Params["adzone_size"] = size
}

func (req *ZuanshiBannerAdgroupAdzoneFindpageRequest) GetAdzoneSize() string {
	if size, found := req.Request.Params["adzone_size"]; found {
		return size.(string)
	}
	return ""
}
