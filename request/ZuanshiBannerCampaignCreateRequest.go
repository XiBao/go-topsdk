package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type ZuanshiBannerCampaignCreateRequest struct {
	Request *topsdk.Request
}

// create new request
func NewZuanshiBannerCampaignCreateRequest() (req *ZuanshiBannerCampaignCreateRequest) {
	request := topsdk.Request{MethodName: "taobao.zuanshi.banner.campaign.create", Params: make(map[string]interface{}, 9)}
	req = &ZuanshiBannerCampaignCreateRequest{
		Request: &request,
	}
	return
}

// 周一到周五，数组24个元素，true：投放，false：不投放
func (req *ZuanshiBannerCampaignCreateRequest) SetWorkday(workday string) {
	req.Request.Params["workday"] = workday
}

func (req *ZuanshiBannerCampaignCreateRequest) GetWorkday() string {
	if workday, found := req.Request.Params["workday"]; found {
		return workday.(string)
	}
	return ""
}

// 周六和周日，数组24个元素，true：投放，false：不投放
func (req *ZuanshiBannerCampaignCreateRequest) SetWeekend(weekend string) {
	req.Request.Params["weekend"] = weekend
}

func (req *ZuanshiBannerCampaignCreateRequest) GetWeekend() string {
	if weekend, found := req.Request.Params["weekend"]; found {
		return weekend.(string)
	}
	return ""
}

// 2：cpm，8：cpc
func (req *ZuanshiBannerCampaignCreateRequest) SetType(aType uint8) {
	req.Request.Params["type"] = aType
}

func (req *ZuanshiBannerCampaignCreateRequest) GetType() uint8 {
	if aType, found := req.Request.Params["type"]; found {
		return aType.(uint8)
	}
	return 2
}

// 计划名称
func (req *ZuanshiBannerCampaignCreateRequest) SetName(name string) {
	req.Request.Params["name"] = name
}

func (req *ZuanshiBannerCampaignCreateRequest) GetName() string {
	if name, found := req.Request.Params["name"]; found {
		return name.(string)
	}
	return ""
}

// 投放地区
func (req *ZuanshiBannerCampaignCreateRequest) SetAreaIdList(areaIdList string) {
	req.Request.Params["area_id_list"] = areaIdList
}

func (req *ZuanshiBannerCampaignCreateRequest) GetAreaIdList() string {
	if areaIdList, found := req.Request.Params["area_id_list"]; found {
		return areaIdList.(string)
	}
	return ""
}

// 1：尽快，2：均匀
func (req *ZuanshiBannerCampaignCreateRequest) SetSpeedType(speedType uint8) {
	req.Request.Params["speed_type"] = speedType
}

func (req *ZuanshiBannerCampaignCreateRequest) GetSpeedType() uint8 {
	if speedType, found := req.Request.Params["speed_type"]; found {
		return speedType.(uint8)
	}
	return 0
}

// 日预算，单位分
func (req *ZuanshiBannerCampaignCreateRequest) SetDayBudget(dayBudget uint) {
	req.Request.Params["day_budget"] = dayBudget
}

func (req *ZuanshiBannerCampaignCreateRequest) GetDayBudget() uint {
	if dayBudget, found := req.Request.Params["day_budget"]; found {
		return dayBudget.(uint)
	}
	return 0
}

// 投放开始时间
func (req *ZuanshiBannerCampaignCreateRequest) SetStartTime(startTime string) {
	req.Request.Params["start_time"] = startTime
}

func (req *ZuanshiBannerCampaignCreateRequest) GetStartTime() string {
	if startTime, found := req.Request.Params["start_time"]; found {
		return startTime.(string)
	}
	return ""
}

// 投放结束时间
func (req *ZuanshiBannerCampaignCreateRequest) SetEndTime(endTime string) {
	req.Request.Params["end_time"] = endTime
}

func (req *ZuanshiBannerCampaignCreateRequest) GetEndTime() string {
	if endTime, found := req.Request.Params["end_time"]; found {
		return endTime.(string)
	}
	return ""
}
