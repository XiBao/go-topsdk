package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type AlibabaOpendspAdgroupsGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewAlibabaOpendspAdgroupsGetRequest() (req *AlibabaOpendspAdgroupsGetRequest) {
	request := topsdk.Request{MethodName: "alibaba.opendsp.adgroups.get", Params: make(map[string]interface{}, 2)}
	req = &AlibabaOpendspAdgroupsGetRequest{
		Request: &request,
	}
	return
}

func (req *AlibabaOpendspAdgroupsGetRequest) SetDspCustId(dspCustId string) {
	req.Request.Params["dsp_cust_id"] = dspCustId
}

func (req *AlibabaOpendspAdgroupsGetRequest) GetDspCustId() string {
	dspCustId, found := req.Request.Params["dsp_cust_id"]
	if found {
		return dspCustId.(string)
	}
	return ""
}

func (req *AlibabaOpendspAdgroupsGetRequest) SetAdgroupIdList(ids string) {
	req.Request.Params["adgroup_id_list"] = ids
}

func (req *AlibabaOpendspAdgroupsGetRequest) GetAdgroupIdList() string {
	ids, found := req.Request.Params["adgroup_id_list"]
	if found {
		return ids.(string)
	}
	return ""
}
