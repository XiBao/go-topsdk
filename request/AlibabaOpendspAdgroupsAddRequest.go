package TopRequest

import (
	"encoding/json"
	"github.com/XiBao/topsdk"
	"github.com/XiBao/topsdk/opendsp"
)

type AlibabaOpendspAdgroupsAddRequest struct {
	Request *topsdk.Request
}

// create new request
func NewAlibabaOpendspAdgroupsAddRequest() (req *AlibabaOpendspAdgroupsAddRequest) {
	request := topsdk.Request{MethodName: "alibaba.opendsp.adgroups.add", Params: make(map[string]interface{}, 1)}
	req = &AlibabaOpendspAdgroupsAddRequest{
		Request: &request,
	}
	return
}

// 推广计划Id
func (req *AlibabaOpendspAdgroupsAddRequest) SetAdgroupList(adgroups []*opendsp.Adgroup) {
	js, _ := json.Marshal(adgroups)
	req.Request.Params["adgroup_list"] = string(js)
}

func (req *AlibabaOpendspAdgroupsAddRequest) GetAdgroupList() []*opendsp.Adgroup {
	js, found := req.Request.Params["adgroup_list"]
	if found {
		adgroups := []*opendsp.Adgroup{}
		json.Unmarshal([]byte(js.(string)), &adgroups)
		return adgroups
	}
	return nil
}
