package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type SimbaRtrptCampaignGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewSimbaRtrptCampaignGetRequest() (req *SimbaRtrptCampaignGetRequest) {
	request := topsdk.Request{MethodName: "taobao.simba.rtrpt.campaign.get", Params: make(map[string]interface{}, 2)}
	req = &SimbaRtrptCampaignGetRequest{
		Request: &request,
	}
	return
}

//昵称
func (req *SimbaRtrptCampaignGetRequest) SetNick(nick string) {
	req.Request.Params["nick"] = nick
}

func (req *SimbaRtrptCampaignGetRequest) GetNick() string {
	nick, found := req.Request.Params["nick"]
	if found {
		return nick.(string)
	}
	return ""
}

//日期，格式yyyy-mm-dd
func (req *SimbaRtrptCampaignGetRequest) SetTheDate(theDate string) {
	req.Request.Params["the_date"] = theDate
}

func (req *SimbaRtrptCampaignGetRequest) GetTheDate() string {
	theDate, found := req.Request.Params["the_date"]
	if found {
		return theDate.(string)
	}
	return ""
}
