package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type AlibabaOpendspCustomersListRequest struct {
	Request *topsdk.Request
}

// create new request
func NewAlibabaOpendspCustomersListRequest() (req *AlibabaOpendspCustomersListRequest) {
	request := topsdk.Request{MethodName: "alibaba.opendsp.customers.list", Params: make(map[string]interface{}, 2)}
	req = &AlibabaOpendspCustomersListRequest{
		Request: &request,
	}
	return
}

func (req *AlibabaOpendspCustomersListRequest) SetDspCustId(dspCustId string) {
	req.Request.Params["dsp_cust_id"] = dspCustId
}

func (req *AlibabaOpendspCustomersListRequest) GetDspCustId() string {
	dspCustId, found := req.Request.Params["dsp_cust_id"]
	if found {
		return dspCustId.(string)
	}
	return ""
}

func (req *AlibabaOpendspCustomersListRequest) SetStatus(status uint8) {
	req.Request.Params["status"] = status
}

func (req *AlibabaOpendspCustomersListRequest) GetStatus() uint8 {
	status, found := req.Request.Params["status"]
	if found {
		return status.(uint8)
	}
	return 0
}
