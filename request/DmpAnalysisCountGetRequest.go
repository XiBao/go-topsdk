package TopRequest

import (
	"encoding/json"
	"github.com/XiBao/topsdk"
	"github.com/XiBao/topsdk/dmp"
)

type DmpAnalysisCountGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewDmpAnalysisCountGetRequest() (req *DmpAnalysisCountGetRequest) {
	request := topsdk.Request{MethodName: "taobao.dmp.analysis.count.get", Params: make(map[string]interface{}, 2)}
	req = &DmpAnalysisCountGetRequest{
		Request: &request,
	}
	return
}

// 放到倍数
func (req *DmpAnalysisCountGetRequest) SetLookalikeMultiple(lookalikeMultiple uint) {
	req.Request.Params["lookalike_multiple"] = lookalikeMultiple
}

func (req *DmpAnalysisCountGetRequest) GetLookalikeMultiple() uint {
	lookalikeMultiple, found := req.Request.Params["lookalike_multiple"]
	if found {
		return lookalikeMultiple.(uint)
	}
	return 0
}

// 选项
func (req *DmpAnalysisCountGetRequest) SetSelects(selects []*dmp.DmpSelectTagOption) {
	js, _ := json.Marshal(selects)
	req.Request.Params["selects"] = string(js)
}

func (req *DmpAnalysisCountGetRequest) GetSelects() []*dmp.DmpSelectTagOption {
	js, found := req.Request.Params["selects"]
	if found {
		selects := []*dmp.DmpSelectTagOption{}
		json.Unmarshal([]byte(js.(string)), &selects)
		return selects
	}
	return nil
}
