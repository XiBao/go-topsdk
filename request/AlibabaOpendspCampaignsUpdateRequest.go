package TopRequest

import (
	"encoding/json"
	"github.com/XiBao/topsdk"
	"github.com/XiBao/topsdk/opendsp"
)

type AlibabaOpendspCampaignsUpdateRequest struct {
	Request *topsdk.Request
}

// create new request
func NewAlibabaOpendspCampaignsUpdateRequest() (req *AlibabaOpendspCampaignsUpdateRequest) {
	request := topsdk.Request{MethodName: "alibaba.opendsp.campaigns.update", Params: make(map[string]interface{}, 1)}
	req = &AlibabaOpendspCampaignsUpdateRequest{
		Request: &request,
	}
	return
}

// 推广计划Id
func (req *AlibabaOpendspCampaignsUpdateRequest) SetCampaignList(campaigns []*opendsp.Campaign) {
	js, _ := json.Marshal(campaigns)
	req.Request.Params["campaign_list"] = string(js)
}

func (req *AlibabaOpendspCampaignsUpdateRequest) GetCampaignList() []*opendsp.Campaign {
	js, found := req.Request.Params["campaign_list"]
	if found {
		campaigns := []*opendsp.Campaign{}
		json.Unmarshal([]byte(js.(string)), &campaigns)
		return campaigns
	}
	return nil
}
