package TopRequest

import (
	"encoding/json"
	"github.com/XiBao/topsdk"
	"github.com/XiBao/topsdk/opendsp"
)

type AlibabaOpendspCampaignsListRequest struct {
	Request *topsdk.Request
}

// create new request
func NewAlibabaOpendspCampaignsListRequest() (req *AlibabaOpendspCampaignsListRequest) {
	request := topsdk.Request{MethodName: "alibaba.opendsp.campaigns.list", Params: make(map[string]interface{}, 1)}
	req = &AlibabaOpendspCampaignsListRequest{
		Request: &request,
	}
	return
}

// 推广计划Id
func (req *AlibabaOpendspCampaignsListRequest) SetQuery(query *opendsp.Campaign) {
	js, _ := json.Marshal(query)
	req.Request.Params["query"] = string(js)
}

func (req *AlibabaOpendspCampaignsListRequest) GetQuery() *opendsp.Campaign {
	js, found := req.Request.Params["query"]
	if found {
		query := &opendsp.Campaign{}
		json.Unmarshal([]byte(js.(string)), query)
		return query
	}
	return nil
}
