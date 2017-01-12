package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type TopatsSimbaCampkeywordbaseGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewTopatsSimbaCampkeywordbaseGetRequest() (req *TopatsSimbaCampkeywordbaseGetRequest) {
	request := topsdk.Request{MethodName: "taobao.topats.simba.campkeywordbase.get", Params: make(map[string]interface{}, 5)}
	req = &TopatsSimbaCampkeywordbaseGetRequest{
		Request: &request,
	}
	return
}

//查询推广计划ID
func (req *TopatsSimbaCampkeywordbaseGetRequest) SetCampaignId(campaignId uint64) {
	req.Request.Params["campaign_id"] = campaignId
}

func (req *TopatsSimbaCampkeywordbaseGetRequest) GetCampaignId() uint64 {
	campaignId, found := req.Request.Params["campaign_id"]
	if found {
		return campaignId.(uint64)
	}
	return 0
}

func (req *TopatsSimbaCampkeywordbaseGetRequest) SetNick(nick string) {
	req.Request.Params["nick"] = nick
}

func (req *TopatsSimbaCampkeywordbaseGetRequest) GetNick() string {
	nick, found := req.Request.Params["nick"]
	if found {
		return nick.(string)
	}
	return ""
}

//时间参数（昨天：DAY、 前一周：7DAY 、前15天：15DAY 、前30天：30DAY 、前90天：90DAY）单选
func (req *TopatsSimbaCampkeywordbaseGetRequest) SetTimeSlot(timeSlot string) {
	req.Request.Params["time_slot"] = timeSlot
}

func (req *TopatsSimbaCampkeywordbaseGetRequest) GetTimeSlot() string {
	timeSlot, found := req.Request.Params["time_slot"]
	if found {
		return timeSlot.(string)
	}
	return ""
}

// 报表类型。可选值：SEARCH（搜索）、CAT（类目出价）、 NOSEARCH（定向投放），可多选，如：SEARCH,CAT
func (req *TopatsSimbaCampkeywordbaseGetRequest) SetSearchType(searchType string) {
	req.Request.Params["search_type"] = searchType
}

func (req *TopatsSimbaCampkeywordbaseGetRequest) GetSearchType() string {
	searchType, found := req.Request.Params["search_type"]
	if found {
		return searchType.(string)
	}
	return ""
}

// 数据来源。可选值：1（站内）、2（站外）、SUMMARY（汇总），其中SUMMARY必须单选，其它值可多选，如：1,2
func (req *TopatsSimbaCampkeywordbaseGetRequest) SetSource(source string) {
	req.Request.Params["source"] = source
}

func (req *TopatsSimbaCampkeywordbaseGetRequest) GetSource() string {
	source, found := req.Request.Params["source"]
	if found {
		return source.(string)
	}
	return ""
}
