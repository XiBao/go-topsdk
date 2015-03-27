package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type AlibabaOpendspAdgroupCreativesGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewAlibabaOpendspAdgroupCreativesGetRequest() (req *AlibabaOpendspAdgroupCreativesGetRequest) {
	request := topsdk.Request{MethodName: "alibaba.opendsp.adgroup.creatives.get", Params: make(map[string]interface{}, 2)}
	req = &AlibabaOpendspAdgroupCreativesGetRequest{
		Request: &request,
	}
	return
}

// DSP系统中推广主id
func (req *AlibabaOpendspAdgroupCreativesGetRequest) SetDspCustId(dspCustId string) {
	req.Request.Params["dsp_cust_id"] = dspCustId
}

func (req *AlibabaOpendspAdgroupCreativesGetRequest) GetAdgroupList() string {
	dspCustId, found := req.Request.Params["dsp_cust_id"]
	if found {
		return dspCustId.(string)
	}
	return ""
}

// 推广单元ID
func (req *AlibabaOpendspAdgroupCreativesGetRequest) SetAdgroupId(adgroupId uint64) {
	req.Request.Params["adgroup_id"] = adgroupId
}

func (req *AlibabaOpendspAdgroupCreativesGetRequest) GetAdgroupId() uint64 {
	adgroupId, found := req.Request.Params["adgroup_id"]
	if found {
		return adgroupId.(uint64)
	}
	return 0
}
