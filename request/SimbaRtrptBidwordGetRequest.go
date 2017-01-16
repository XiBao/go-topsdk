package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type SimbaRtrptBidwordGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewSimbaRtrptBidwordGetRequest() (req *SimbaRtrptBidwordGetRequest) {
	request := topsdk.Request{MethodName: "taobao.simba.rtrpt.bidword.get", Params: make(map[string]interface{}, 4)}
	req = &SimbaRtrptBidwordGetRequest{
		Request: &request,
	}
	return
}

//昵称
func (req *SimbaRtrptBidwordGetRequest) SetNick(nick string) {
	req.Request.Params["nick"] = nick
}

func (req *SimbaRtrptBidwordGetRequest) GetNick() string {
	nick, found := req.Request.Params["nick"]
	if found {
		return nick.(string)
	}
	return ""
}

//日期，格式yyyy-mm-dd
func (req *SimbaRtrptBidwordGetRequest) SetTheDate(theDate string) {
	req.Request.Params["the_date"] = theDate
}

func (req *SimbaRtrptBidwordGetRequest) GetTheDate() string {
	theDate, found := req.Request.Params["the_date"]
	if found {
		return theDate.(string)
	}
	return ""
}

func (req *SimbaRtrptBidwordGetRequest) SetCampaignId(campaignId uint64) {
	req.Request.Params["campaign_id"] = campaignId
}

func (req *SimbaRtrptBidwordGetRequest) GetCampaignId() uint64 {
	campaignId, found := req.Request.Params["campaign_id"]
	if found {
		return campaignId.(uint64)
	}
	return 0
}

func (req *SimbaRtrptBidwordGetRequest) SetAdgroupId(adgroupId uint64) {
	req.Request.Params["adgroup_id"] = adgroupId
}

func (req *SimbaRtrptBidwordGetRequest) GetAdgroupId() uint64 {
	adgroupId, found := req.Request.Params["adgroup_id"]
	if found {
		return adgroupId.(uint64)
	}
	return 0
}
