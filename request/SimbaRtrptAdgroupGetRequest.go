package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type SimbaRtrptAdgroupGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewSimbaRtrptAdgroupGetRequest() (req *SimbaRtrptAdgroupGetRequest) {
	request := topsdk.Request{MethodName: "taobao.simba.rtrpt.adgroup.get", Params: make(map[string]interface{}, 2)}
	req = &SimbaRtrptAdgroupGetRequest{
		Request: &request,
	}
	return
}

//昵称
func (req *SimbaRtrptAdgroupGetRequest) SetNick(nick string) {
	req.Request.Params["nick"] = nick
}

func (req *SimbaRtrptAdgroupGetRequest) GetNick() string {
	nick, found := req.Request.Params["nick"]
	if found {
		return nick.(string)
	}
	return ""
}

//日期，格式yyyy-mm-dd
func (req *SimbaRtrptAdgroupGetRequest) SetTheDate(theDate string) {
	req.Request.Params["the_date"] = theDate
}

func (req *SimbaRtrptAdgroupGetRequest) GetTheDate() string {
	theDate, found := req.Request.Params["the_date"]
	if found {
		return theDate.(string)
	}
	return ""
}

func (req *SimbaRtrptAdgroupGetRequest) SetCampaignId(campaignId uint64) {
	req.Request.Params["campaign_id"] = campaignId
}

func (req *SimbaRtrptAdgroupGetRequest) GetCampaignId() uint64 {
	campaignId, found := req.Request.Params["campaign_id"]
	if found {
		return campaignId.(uint64)
	}
	return 0
}

func (req *SimbaRtrptAdgroupGetRequest) SetPageSize(pageSize uint) {
	req.Request.Params["page_size"] = pageSize
}

func (req *SimbaRtrptAdgroupGetRequest) GetPageSize() uint {
	pageSize, found := req.Request.Params["page_size"]
	if found {
		return pageSize.(uint)
	}
	return 0
}

func (req *SimbaRtrptAdgroupGetRequest) SetPageNo(pageNo uint) {
	req.Request.Params["page_number"] = pageNo
}

func (req *SimbaRtrptAdgroupGetRequest) GetPageNo() uint {
	pageNo, found := req.Request.Params["page_number"]
	if found {
		return pageNo.(uint)
	}
	return 0
}
