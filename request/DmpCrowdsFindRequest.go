package TopRequest

import (
	"encoding/json"
	"github.com/XiBao/topsdk"
	"github.com/XiBao/topsdk/dmp"
)

type DmpCrowdsFindRequest struct {
	Request *topsdk.Request
}

// create new request
func NewDmpCrowdsFindRequest() (req *DmpCrowdsFindRequest) {
	request := topsdk.Request{MethodName: "taobao.dmp.crowds.find", Params: make(map[string]interface{}, 1)}
	req = &DmpCrowdsFindRequest{
		Request: &request,
	}
	return
}

// 人群查询条件
func (req *DmpCrowdsFindRequest) SetCrowdQueryDTO(query *dmp.CrowdQuery) {
	js, _ := json.Marshal(query)
	req.Request.Params["crowd_query_d_t_o"] = string(js)
}

func (req *DmpCrowdsFindRequest) GetCrowdQueryDTO() *dmp.CrowdQuery {
	js, found := req.Request.Params["crowd_query_d_t_o"]
	if found {
		var query dmp.CrowdQuery
		json.Unmarshal([]byte(js.(string)), &query)
		return &query
	}
	return nil
}
