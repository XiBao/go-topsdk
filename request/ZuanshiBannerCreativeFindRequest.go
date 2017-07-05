package TopRequest

import (
	"github.com/XiBao/topsdk"
	"strconv"
	"strings"
)

type ZuanshiBannerCreativeFindRequest struct {
	Request *topsdk.Request
}

// create new request
func NewZuanshiBannerCreativeFindRequest() (req *ZuanshiBannerCreativeFindRequest) {
	request := topsdk.Request{MethodName: "taobao.zuanshi.banner.creative.find", Params: make(map[string]interface{}, 9)}
	req = &ZuanshiBannerCreativeFindRequest{
		Request: &request,
	}
	return
}

// 计划ID
func (req *ZuanshiBannerCreativeFindRequest) SetCampaignId(id uint64) {
	req.Request.Params["campaign_id"] = id
}

func (req *ZuanshiBannerCreativeFindRequest) GetCampaignId() uint64 {
	if id, found := req.Request.Params["campaign_id"]; found {
		return id.(uint64)
	}
	return 0
}

// 单元ID
func (req *ZuanshiBannerCreativeFindRequest) SetAdgroupId(id uint64) {
	req.Request.Params["adgroup_id"] = id
}

func (req *ZuanshiBannerCreativeFindRequest) GetAdgroupId() uint64 {
	if id, found := req.Request.Params["adgroup_id"]; found {
		return id.(uint64)
	}
	return 0
}

// 尺寸列表，200x300
func (req *ZuanshiBannerCreativeFindRequest) SetSizeList(sizes []string) {
	req.Request.Params["size_list"] = strings.Join(sizes, ",")
}

func (req *ZuanshiBannerCreativeFindRequest) GetSizeList() []string {
	var sizes []string
	if sizeStr, found := req.Request.Params["size_list"]; found {
		sizes = strings.Split(sizeStr.(string), ",")
	}
	return sizes
}

// 图片：2，flash：3，视频：4，文字链：5，Flash不遮盖：9，模板：10，
func (req *ZuanshiBannerCreativeFindRequest) SetFormatList(formatList []uint) {
	var formatsStr []string
	for _, f := range formatList {
		formatsStr = append(formatsStr, strconv.FormatUint(uint64(f), 10))
	}
	req.Request.Params["format_list"] = strings.Join(formatsStr, ",")
}

func (req *ZuanshiBannerCreativeFindRequest) GetFormatList() []uint {
	var formatList []uint
	if flStr, found := req.Request.Params["format_list"]; found {
		fs := strings.Split(flStr.(string), ",")
		for _, f := range fs {
			if aF, err := strconv.ParseUint(f, 10, 64); err == nil && aF > 0 {
				formatList = append(formatList, uint(aF))
			}
		}
	}
	return formatList
}

// 待审核：-4,-1,0，审核通过：1，审核拒绝：-2,-5,-9，
func (req *ZuanshiBannerCreativeFindRequest) SetAuditStatusList(auditStatusList []int) {
	var asList []string
	for _, a := range auditStatusList {
		asList = append(asList, strconv.FormatInt(int64(a), 10))
	}
	req.Request.Params["audit_status_list"] = strings.Join(asList, ",")
}

func (req *ZuanshiBannerCreativeFindRequest) GetAuditStatusList() []int {
	var auditStatusList []int
	if asListStr, found := req.Request.Params["audit_status_list"]; found {
		asList := strings.Split(asListStr.(string), ",")
		for _, as := range asList {
			if aAs, err := strconv.ParseInt(as, 10, 64); err == nil && aAs > 0 {
				auditStatusList = append(auditStatusList, int(aAs))
			}
		}
	}
	return auditStatusList
}

// 在线状态：1：正常，-1：回收站
func (req *ZuanshiBannerCreativeFindRequest) SetOnlineStatus(onlineStatus int8) {
	req.Request.Params["online_status"] = onlineStatus
}

func (req *ZuanshiBannerCreativeFindRequest) GetOnlineStatus() int8 {
	if onlineStatus, found := req.Request.Params["online_status"]; found {
		return onlineStatus.(int8)
	}
	return 0
}

// 创意等级,1：一级，2：二级，3：三级，4：四级，10：十级，99：未分级
func (req *ZuanshiBannerCreativeFindRequest) SetCreativeLevel(level uint) {
	req.Request.Params["creative_level"] = level
}

func (req *ZuanshiBannerCreativeFindRequest) GetCreativeLevel() uint {
	if level, found := req.Request.Params["creative_level"]; found {
		return level.(uint)
	}
	return 0
}

// 分页大小, 默认值：15 最小值：1 最大长度：200
func (req *ZuanshiBannerCreativeFindRequest) SetPageSize(pageSize uint) {
	if pageSize == 0 {
		pageSize = 15
	}
	req.Request.Params["page_size"] = pageSize
}

func (req *ZuanshiBannerCreativeFindRequest) GetPageSize() uint {
	if pageSize, found := req.Request.Params["page_size"]; found {
		return pageSize.(uint)
	}
	return 15
}

// 当前页码, 默认值：1
func (req *ZuanshiBannerCreativeFindRequest) SetPageNum(pageNum uint) {
	if pageNum == 0 {
		pageNum = 1
	}
	req.Request.Params["page_num"] = pageNum
}

func (req *ZuanshiBannerCreativeFindRequest) GetPageNum() uint {
	if pageNum, found := req.Request.Params["page_num"]; found {
		return pageNum.(uint)
	}
	return 1
}

// 创意名称
func (req *ZuanshiBannerCreativeFindRequest) SetName(name string) {
	req.Request.Params["name"] = name
}

func (req *ZuanshiBannerCreativeFindRequest) GetName() string {
	if name, found := req.Request.Params["name"]; found {
		return name.(string)
	}
	return ""
}
