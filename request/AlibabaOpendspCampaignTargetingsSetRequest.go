package TopRequest

import (
	"encoding/json"
	"github.com/XiBao/topsdk"
	"github.com/XiBao/topsdk/opendsp"
)

type AlibabaOpendspCampaignTargetingsSetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewAlibabaOpendspCampaignTargetingsSetRequest() (req *AlibabaOpendspCampaignTargetingsSetRequest) {
	request := topsdk.Request{MethodName: "alibaba.opendsp.campaign.targetings.set", Params: make(map[string]interface{}, 3)}
	req = &AlibabaOpendspCampaignTargetingsSetRequest{
		Request: &request,
	}
	return
}

func (req *AlibabaOpendspCampaignTargetingsSetRequest) SetDspCustId(dspCustId string) {
	req.Request.Params["dsp_cust_id"] = dspCustId
}

func (req *AlibabaOpendspCampaignTargetingsSetRequest) GetDspCustId() string {
	dspCustId, found := req.Request.Params["dsp_cust_id"]
	if found {
		return dspCustId.(string)
	}
	return ""
}

func (req *AlibabaOpendspCampaignTargetingsSetRequest) SetCampaignId(id uint64) {
	req.Request.Params["campaign_id"] = id
}

func (req *AlibabaOpendspCampaignTargetingsSetRequest) GetCampaignId() uint64 {
	id, found := req.Request.Params["campaign_id"]
	if found {
		return id.(uint64)
	}
	return 0
}

func (req *AlibabaOpendspCampaignTargetingsSetRequest) SetTargetingList(campaignTargetings []*opendsp.CampaignTargeting) {
	js, _ := json.Marshal(campaignTargetings)
	req.Request.Params["targeting_list"] = string(js)
}

func (req *AlibabaOpendspCampaignTargetingsSetRequest) GetTargetingList() []*opendsp.CampaignTargeting {
	js, found := req.Request.Params["targeting_list"]
	if found {
		campaignTargetings := []*opendsp.CampaignTargeting{}
		json.Unmarshal([]byte(js.(string)), &campaignTargetings)
		return campaignTargetings
	}
	return nil
}
