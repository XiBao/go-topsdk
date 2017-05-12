package TopRequest

import (
	"encoding/json"
	"github.com/XiBao/topsdk"
	"github.com/XiBao/topsdk/dmp"
)

type DmpAnalysisCoverageRequest struct {
	Request *topsdk.Request
}

// create new request
func NewDmpAnalysisCoverageRequest() (req *DmpAnalysisCoverageRequest) {
	request := topsdk.Request{MethodName: "taobao.dmp.analysis.coverage", Params: make(map[string]interface{}, 1)}
	req = &DmpAnalysisCoverageRequest{
		Request: &request,
	}
	return
}

// 选项
func (req *DmpAnalysisCoverageRequest) SetSelects(selects *dmp.SelectTagOptionSet) {
	js, _ := json.Marshal(selects)
	req.Request.Params["select_tag_option_set_d_t_o"] = string(js)
}

func (req *DmpAnalysisCoverageRequest) GetSelects() *dmp.SelectTagOptionSet {
	js, found := req.Request.Params["select_tag_option_set_d_t_o"]
	if found {
		var selects dmp.SelectTagOptionSet
		json.Unmarshal([]byte(js.(string)), &selects)
		return &selects
	}
	return nil
}
