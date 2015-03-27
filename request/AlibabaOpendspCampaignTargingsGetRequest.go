package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type AlibabaOpendspCampaignTargetingsGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewAlibabaOpendspCampaignTargetingsGetRequest() (req *AlibabaOpendspCampaignTargetingsGetRequest) {
	request := topsdk.Request{MethodName: "alibaba.opendsp.campaigns.targetings.get", Params: make(map[string]interface{}, 2)}
	req = &AlibabaOpendspCampaignTargetingsGetRequest{
		Request: &request,
	}
	return
}

func (req *AlibabaOpendspCampaignTargetingsGetRequest) SetDspCustId(dspCustId string) {
	req.Request.Params["dsp_cust_id"] = dspCustId
}

func (req *AlibabaOpendspCampaignTargetingsGetRequest) GetDspCustId() string {
	dspCustId, found := req.Request.Params["dsp_cust_id"]
	if found {
		return dspCustId.(string)
	}
	return ""
}

func (req *AlibabaOpendspCampaignTargetingsGetRequest) SetCampaignIdList(ids string) {
	req.Request.Params["campaign_id_list"] = ids
}

func (req *AlibabaOpendspCampaignTargetingsGetRequest) GetCampaignIdList() string {
	ids, found := req.Request.Params["campaign_id_list"]
	if found {
		return ids.(string)
	}
	return ""
}
