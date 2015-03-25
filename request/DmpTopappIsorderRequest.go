package TopRequest

import (
	"encoding/json"
	"github.com/XiBao/topsdk"
	"github.com/XiBao/topsdk/dmp"
)

type DmpTopappIsorderRequest struct {
	Request *topsdk.Request
}

// create new request
func NewDmpTopappIsorderRequest() (req *DmpTopappIsorderRequest) {
	request := topsdk.Request{MethodName: "taobao.dmp.topapp.isorder", Params: make(map[string]interface{}, 1)}
	req = &DmpTopappIsorderRequest{
		Request: &request,
	}
	return
}

// query对象
func (req *DmpTopappIsorderRequest) SetCrowdRptQuery(query *dmp.TopQuery) {
	js, _ := json.Marshal(query)
	req.Request.Params["top_query"] = string(js)
}

func (req *DmpTopappIsorderRequest) GetCrowdRptQuery() *dmp.TopQuery {
	js, found := req.Request.Params["top_query"]
	if found {
		query := &dmp.TopQuery{}
		json.Unmarshal([]byte(js.(string)), query)
		return query
	}
	return nil
}
