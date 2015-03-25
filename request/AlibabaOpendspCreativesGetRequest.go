package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type AlibabaOpendspCreativesGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewAlibabaOpendspCreativesGetRequest() (req *AlibabaOpendspCreativesGetRequest) {
	request := topsdk.Request{MethodName: "alibaba.opendsp.creatives.get", Params: make(map[string]interface{}, 2)}
	req = &AlibabaOpendspCreativesGetRequest{
		Request: &request,
	}
	return
}

func (req *AlibabaOpendspCreativesGetRequest) SetDspCustId(dspCustId string) {
	req.Request.Params["dsp_cust_id"] = dspCustId
}

func (req *AlibabaOpendspCreativesGetRequest) GetDspCustId() string {
	dspCustId, found := req.Request.Params["dsp_cust_id"]
	if found {
		return dspCustId.(string)
	}
	return ""
}

func (req *AlibabaOpendspCreativesGetRequest) SetCreativeIdList(ids string) {
	req.Request.Params["creative_id_list"] = ids
}

func (req *AlibabaOpendspCreativesGetRequest) GetCreativeIdList() string {
	ids, found := req.Request.Params["creative_id_list"]
	if found {
		return ids.(string)
	}
	return ""
}
