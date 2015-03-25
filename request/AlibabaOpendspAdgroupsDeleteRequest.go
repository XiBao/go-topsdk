package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type AlibabaOpendspAdgroupsDeleteRequest struct {
	Request *topsdk.Request
}

// create new request
func NewAlibabaOpendspAdgroupsDeleteRequest() (req *AlibabaOpendspAdgroupsDeleteRequest) {
	request := topsdk.Request{MethodName: "alibaba.opendsp.adgroups.delete", Params: make(map[string]interface{}, 2)}
	req = &AlibabaOpendspAdgroupsDeleteRequest{
		Request: &request,
	}
	return
}

func (req *AlibabaOpendspAdgroupsDeleteRequest) SetDspCustId(dspCustId string) {
	req.Request.Params["dsp_cust_id"] = dspCustId
}

func (req *AlibabaOpendspAdgroupsDeleteRequest) GetDspCustId() string {
	dspCustId, found := req.Request.Params["dsp_cust_id"]
	if found {
		return dspCustId.(string)
	}
	return ""
}

func (req *AlibabaOpendspAdgroupsDeleteRequest) SetAdgroupIdList(ids string) {
	req.Request.Params["adgroup_id_list"] = ids
}

func (req *AlibabaOpendspAdgroupsDeleteRequest) GetAdgroupIdList() string {
	ids, found := req.Request.Params["adgroup_id_list"]
	if found {
		return ids.(string)
	}
	return ""
}
