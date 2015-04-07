package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type SimbaCampaignAddRequest struct {
	Request *topsdk.Request
}

// create new request
func NewSimbaCampaignAddRequest() (req *SimbaCampaignAddRequest) {
	request := topsdk.Request{MethodName: "taobao.simba.campaign.add", Params: make(map[string]interface{}, 2)}
	req = &SimbaCampaignAddRequest{
		Request: &request,
	}
	return
}

//主人昵称
func (req *SimbaCampaignAddRequest) SetNick(nick string) {
	req.Request.Params["nick"] = nick
}

func (req *SimbaCampaignAddRequest) GetNick() string {
	nick, found := req.Request.Params["nick"]
	if found {
		return nick.(string)
	}
	return ""
}

//推广计划名称，不能多余20个汉字，不能和客户其他推广计划同名。
func (req *SimbaCampaignAddRequest) SetTitle(title string) {
	req.Request.Params["title"] = title
}

func (req *SimbaCampaignAddRequest) GetTitle() string {
	title, found := req.Request.Params["title"]
	if found {
		return title.(string)
	}
	return ""
}
