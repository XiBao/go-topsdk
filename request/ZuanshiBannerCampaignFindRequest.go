package TopRequest

import (
	"github.com/XiBao/topsdk"
	"strconv"
	"strings"
)

type ZuanshiBannerCampaignFindRequest struct {
	Request *topsdk.Request
}

// create new request
func NewZuanshiBannerCampaignFindRequest() (req *ZuanshiBannerCampaignFindRequest) {
	request := topsdk.Request{MethodName: "taobao.zuanshi.banner.campaign.find", Params: make(map[string]interface{}, 6)}
	req = &ZuanshiBannerCampaignFindRequest{
		Request: &request,
	}
	return
}

// 计划ID
func (req *ZuanshiBannerCampaignFindRequest) SetCampaignIdList(ids []uint64) {
	var idsStr []string
	for _, id := range ids {
		idsStr = append(idsStr, strconv.FormatUint(id, 10))
	}
	req.Request.Params["campaign_id_list"] = strings.Join(idsStr, ",")
}

func (req *ZuanshiBannerCampaignFindRequest) GetCampaignIdList() []uint64 {
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
func (req *ZuanshiBannerCampaignFindRequest) SetStatusList(statusList []uint8) {
	var statusStr []string
	for _, status := range statusList {
		statusStr = append(statusStr, strconv.FormatUint(uint64(status), 10))
	}
	req.Request.Params["status_list"] = statusStr
}

func (req *ZuanshiBannerCampaignFindRequest) GetStatusList() []uint8 {
	var statusList []uint8
	if statusStr, found := req.Request.Params["status_list"]; found {
		statusArr := strings.Split(statusStr.(string), ",")
		for _, status := range statusArr {
			if status, err := strconv.ParseUint(status, 10, 64); err == nil && status > 0 {
				statusList = append(statusList, uint8(status))
			}
		}
	}
	return statusList
}

// 分页大小, 默认值：15 最小值：1 最大长度：200
func (req *ZuanshiBannerCampaignFindRequest) SetPageSize(pageSize uint) {
	if pageSize == 0 {
		pageSize = 15
	}
	req.Request.Params["page_size"] = pageSize
}

func (req *ZuanshiBannerCampaignFindRequest) GetPageSize() uint {
	if pageSize, found := req.Request.Params["page_size"]; found {
		return pageSize.(uint)
	}
	return 15
}

// 当前页码, 默认值：1
func (req *ZuanshiBannerCampaignFindRequest) SetPageNum(pageNum uint) {
	if pageNum == 0 {
		pageNum = 1
	}
	req.Request.Params["page_num"] = pageNum
}

func (req *ZuanshiBannerCampaignFindRequest) GetPageNum() uint {
	if pageNum, found := req.Request.Params["page_num"]; found {
		return pageNum.(uint)
	}
	return 1
}

// 2:cpm,8:cpc
func (req *ZuanshiBannerCampaignFindRequest) SetType(aType uint8) {
	req.Request.Params["type"] = aType
}

func (req *ZuanshiBannerCampaignFindRequest) GetType() uint8 {
	if aType, found := req.Request.Params["type"]; found {
		return aType.(uint8)
	}
	return 0
}

// 计划名称
func (req *ZuanshiBannerCampaignFindRequest) SetName(name string) {
	req.Request.Params["name"] = name
}

func (req *ZuanshiBannerCampaignFindRequest) GetName() string {
	if name, found := req.Request.Params["name"]; found {
		return name.(string)
	}
	return ""
}
