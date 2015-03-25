package TopRequest

import (
	"encoding/json"
	"github.com/XiBao/topsdk"
	"github.com/XiBao/topsdk/opendsp"
)

type AlibabaOpendspAdgroupsListRequest struct {
	Request *topsdk.Request
}

// create new request
func NewAlibabaOpendspAdgroupsListRequest() (req *AlibabaOpendspAdgroupsListRequest) {
	request := topsdk.Request{MethodName: "alibaba.opendsp.adgroups.list", Params: make(map[string]interface{}, 1)}
	req = &AlibabaOpendspAdgroupsListRequest{
		Request: &request,
	}
	return
}

// 推广计划Id
func (req *AlibabaOpendspAdgroupsListRequest) SetQuery(query *opendsp.Adgroup) {
	js, _ := json.Marshal(query)
	req.Request.Params["query"] = string(js)
}

func (req *AlibabaOpendspAdgroupsListRequest) GetQuery() *opendsp.Adgroup {
	js, found := req.Request.Params["query"]
	if found {
		query := &opendsp.Adgroup{}
		json.Unmarshal([]byte(js.(string)), query)
		return query
	}
	return nil
}
