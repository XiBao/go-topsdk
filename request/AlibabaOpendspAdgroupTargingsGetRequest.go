package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type AlibabaOpendspAdgroupTargetingsGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewAlibabaOpendspAdgroupTargetingsGetRequest() (req *AlibabaOpendspAdgroupTargetingsGetRequest) {
	request := topsdk.Request{MethodName: "alibaba.opendsp.adgroup.targetings.get", Params: make(map[string]interface{}, 2)}
	req = &AlibabaOpendspAdgroupTargetingsGetRequest{
		Request: &request,
	}
	return
}

func (req *AlibabaOpendspAdgroupTargetingsGetRequest) SetDspCustId(dspCustId string) {
	req.Request.Params["dsp_cust_id"] = dspCustId
}

func (req *AlibabaOpendspAdgroupTargetingsGetRequest) GetDspCustId() string {
	dspCustId, found := req.Request.Params["dsp_cust_id"]
	if found {
		return dspCustId.(string)
	}
	return ""
}

func (req *AlibabaOpendspAdgroupTargetingsGetRequest) SetAdgroupIdList(ids string) {
	req.Request.Params["adgroup_id_list"] = ids
}

func (req *AlibabaOpendspAdgroupTargetingsGetRequest) GetAdgroupIdList() string {
	ids, found := req.Request.Params["adgroup_id_list"]
	if found {
		return ids.(string)
	}
	return ""
}
