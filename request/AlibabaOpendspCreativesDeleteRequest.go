package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type AlibabaOpendspCreativesDeleteRequest struct {
	Request *topsdk.Request
}

// create new request
func NewAlibabaOpendspCreativesDeleteRequest() (req *AlibabaOpendspCreativesDeleteRequest) {
	request := topsdk.Request{MethodName: "alibaba.opendsp.creatives.delete", Params: make(map[string]interface{}, 2)}
	req = &AlibabaOpendspCreativesDeleteRequest{
		Request: &request,
	}
	return
}

func (req *AlibabaOpendspCreativesDeleteRequest) SetDspCustId(dspCustId string) {
	req.Request.Params["dsp_cust_id"] = dspCustId
}

func (req *AlibabaOpendspCreativesDeleteRequest) GetDspCustId() string {
	dspCustId, found := req.Request.Params["dsp_cust_id"]
	if found {
		return dspCustId.(string)
	}
	return ""
}

func (req *AlibabaOpendspCreativesDeleteRequest) SetCreativeIdList(ids string) {
	req.Request.Params["creative_id_list"] = ids
}

func (req *AlibabaOpendspCreativesDeleteRequest) GetCreativeIdList() string {
	ids, found := req.Request.Params["creative_id_list"]
	if found {
		return ids.(string)
	}
	return ""
}
