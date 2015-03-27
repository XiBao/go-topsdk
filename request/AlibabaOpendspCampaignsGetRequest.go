package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type AlibabaOpendspCampaignsGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewAlibabaOpendspCampaignsGetRequest() (req *AlibabaOpendspCampaignsGetRequest) {
	request := topsdk.Request{MethodName: "alibaba.opendsp.campaigns.get", Params: make(map[string]interface{}, 2)}
	req = &AlibabaOpendspCampaignsGetRequest{
		Request: &request,
	}
	return
}

func (req *AlibabaOpendspCampaignsGetRequest) SetDspCustId(dspCustId string) {
	req.Request.Params["dsp_cust_id"] = dspCustId
}

func (req *AlibabaOpendspCampaignsGetRequest) GetDspCustId() string {
	dspCustId, found := req.Request.Params["dsp_cust_id"]
	if found {
		return dspCustId.(string)
	}
	return ""
}

func (req *AlibabaOpendspCampaignsGetRequest) SetCampaignIdList(ids string) {
	req.Request.Params["campaign_id_list"] = ids
}

func (req *AlibabaOpendspCampaignsGetRequest) GetCampaignIdList() string {
	ids, found := req.Request.Params["campaign_id_list"]
	if found {
		return ids.(string)
	}
	return ""
}
