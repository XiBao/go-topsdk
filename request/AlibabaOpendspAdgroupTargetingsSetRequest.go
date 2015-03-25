package TopRequest

import (
	"encoding/json"
	"github.com/XiBao/topsdk"
	"github.com/XiBao/topsdk/opendsp"
)

type AlibabaOpendspAdgroupTargetingsSetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewAlibabaOpendspAdgroupTargetingsSetRequest() (req *AlibabaOpendspAdgroupTargetingsSetRequest) {
	request := topsdk.Request{MethodName: "alibaba.opendsp.adgroup.targetings.set", Params: make(map[string]interface{}, 3)}
	req = &AlibabaOpendspAdgroupTargetingsSetRequest{
		Request: &request,
	}
	return
}

func (req *AlibabaOpendspAdgroupTargetingsSetRequest) SetDspCustId(dspCustId string) {
	req.Request.Params["dsp_cust_id"] = dspCustId
}

func (req *AlibabaOpendspAdgroupTargetingsSetRequest) GetDspCustId() string {
	dspCustId, found := req.Request.Params["dsp_cust_id"]
	if found {
		return dspCustId.(string)
	}
	return ""
}

func (req *AlibabaOpendspAdgroupTargetingsSetRequest) SetAdgroupId(id uint64) {
	req.Request.Params["adgroup_id"] = id
}

func (req *AlibabaOpendspAdgroupTargetingsSetRequest) GetAdgroupId() uint64 {
	id, found := req.Request.Params["adgroup_id"]
	if found {
		return id.(uint64)
	}
	return 0
}

func (req *AlibabaOpendspAdgroupTargetingsSetRequest) SetTargetingList(adgroupTargetings []*opendsp.AdgroupTargeting) {
	js, _ := json.Marshal(adgroupTargetings)
	req.Request.Params["targeting_list"] = string(js)
}

func (req *AlibabaOpendspAdgroupTargetingsSetRequest) GetTargetingList() []*opendsp.AdgroupTargeting {
	js, found := req.Request.Params["targeting_list"]
	if found {
		adgroupTargetings := []*opendsp.AdgroupTargeting{}
		json.Unmarshal([]byte(js.(string)), &adgroupTargetings)
		return adgroupTargetings
	}
	return nil
}
