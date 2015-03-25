package TopRequest

import (
	"encoding/json"
	"github.com/XiBao/topsdk"
	"github.com/XiBao/topsdk/dmp"
)

type DmpAnalysisChartGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewDmpAnalysisChartGetRequest() (req *DmpAnalysisChartGetRequest) {
	request := topsdk.Request{MethodName: "taobao.dmp.analysis.chart.get", Params: make(map[string]interface{}, 2)}
	req = &DmpAnalysisChartGetRequest{
		Request: &request,
	}
	return
}

// 标签ID
func (req *DmpAnalysisChartGetRequest) SetTagId(id uint64) {
	req.Request.Params["tag_id"] = id
}

func (req *DmpAnalysisChartGetRequest) GetTagId() uint64 {
	id, found := req.Request.Params["tag_id"]
	if found {
		return id.(uint64)
	}
	return 0
}

// 选项
func (req *DmpAnalysisChartGetRequest) SetSelects(selects []*dmp.DmpSelectTagOption) {
	js, _ := json.Marshal(selects)
	req.Request.Params["selects"] = string(js)
}

func (req *DmpAnalysisChartGetRequest) GetSelects() []*dmp.DmpSelectTagOption {
	js, found := req.Request.Params["selects"]
	if found {
		selects := []*dmp.DmpSelectTagOption{}
		json.Unmarshal([]byte(js.(string)), &selects)
		return selects
	}
	return nil
}
