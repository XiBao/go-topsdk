package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type AlibabaOpendspCampaignsDeleteRequest struct {
	Request *topsdk.Request
}

// create new request
func NewAlibabaOpendspCampaignsDeleteRequest() (req *AlibabaOpendspCampaignsDeleteRequest) {
	request := topsdk.Request{MethodName: "alibaba.opendsp.campaigns.delete", Params: make(map[string]interface{}, 2)}
	req = &AlibabaOpendspCampaignsDeleteRequest{
		Request: &request,
	}
	return
}

func (req *AlibabaOpendspCampaignsDeleteRequest) SetDspCustId(dspCustId string) {
	req.Request.Params["dsp_cust_id"] = dspCustId
}

func (req *AlibabaOpendspCampaignsDeleteRequest) GetDspCustId() string {
	dspCustId, found := req.Request.Params["dsp_cust_id"]
	if found {
		return dspCustId.(string)
	}
	return ""
}

func (req *AlibabaOpendspCampaignsDeleteRequest) SetCampaignIdList(ids string) {
	req.Request.Params["campaign_id_list"] = ids
}

func (req *AlibabaOpendspCampaignsDeleteRequest) GetCampaignIdList() string {
	ids, found := req.Request.Params["campaign_id_list"]
	if found {
		return ids.(string)
	}
	return ""
}
