package TopRequest

import (
	"encoding/json"
	"github.com/XiBao/topsdk"
	"github.com/XiBao/topsdk/opendsp"
)

type AlibabaOpendspAdgroupsUpdateRequest struct {
	Request *topsdk.Request
}

// create new request
func NewAlibabaOpendspAdgroupsUpdateRequest() (req *AlibabaOpendspAdgroupsUpdateRequest) {
	request := topsdk.Request{MethodName: "alibaba.opendsp.adgroups.update", Params: make(map[string]interface{}, 1)}
	req = &AlibabaOpendspAdgroupsUpdateRequest{
		Request: &request,
	}
	return
}

// 推广计划Id
func (req *AlibabaOpendspAdgroupsUpdateRequest) SetAdgroupList(adgroups []*opendsp.Adgroup) {
	js, _ := json.Marshal(adgroups)
	req.Request.Params["adgroup_list"] = string(js)
}

func (req *AlibabaOpendspAdgroupsUpdateRequest) GetAdgroupList() []*opendsp.Adgroup {
	js, found := req.Request.Params["adgroup_list"]
	if found {
		adgroups := []*opendsp.Adgroup{}
		json.Unmarshal([]byte(js.(string)), &adgroups)
		return adgroups
	}
	return nil
}
