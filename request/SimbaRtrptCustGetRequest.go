package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type SimbaRtrptCustGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewSimbaRtrptCustGetRequest() (req *SimbaRtrptCustGetRequest) {
	request := topsdk.Request{MethodName: "taobao.simba.rtrpt.cust.get", Params: make(map[string]interface{}, 2)}
	req = &SimbaRtrptCustGetRequest{
		Request: &request,
	}
	return
}

//昵称
func (req *SimbaRtrptCustGetRequest) SetNick(nick string) {
	req.Request.Params["nick"] = nick
}

func (req *SimbaRtrptCustGetRequest) GetNick() string {
	nick, found := req.Request.Params["nick"]
	if found {
		return nick.(string)
	}
	return ""
}

//日期，格式yyyy-mm-dd
func (req *SimbaRtrptCustGetRequest) SetTheDate(theDate string) {
	req.Request.Params["the_date"] = theDate
}

func (req *SimbaRtrptCustGetRequest) GetTheDate() string {
	theDate, found := req.Request.Params["the_date"]
	if found {
		return theDate.(string)
	}
	return ""
}
