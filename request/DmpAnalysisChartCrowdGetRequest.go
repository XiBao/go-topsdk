package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type DmpAnalysisChartCrowdGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewDmpAnalysisChartCrowdGetRequest() (req *DmpAnalysisChartCrowdGetRequest) {
	request := topsdk.Request{MethodName: "taobao.dmp.analysis.chart.crowd.get", Params: make(map[string]interface{}, 2)}
	req = &DmpAnalysisChartCrowdGetRequest{
		Request: &request,
	}
	return
}

// 标签id
func (req *DmpAnalysisChartCrowdGetRequest) SetTagId(tagId uint64) {
	req.Request.Params["tag_id"] = tagId
}

func (req *DmpAnalysisChartCrowdGetRequest) GetTagId() uint64 {
	tagId, found := req.Request.Params["tag_id"]
	if found {
		return tagId.(uint64)
	}
	return 0
}

// 透视的人群id
func (req *DmpAnalysisChartCrowdGetRequest) SetCrowdId(crowdId uint64) {
	req.Request.Params["crowd_id"] = crowdId
}

func (req *DmpAnalysisChartCrowdGetRequest) GetCrowdId() uint64 {
	crowdId, found := req.Request.Params["crowd_id"]
	if found {
		return crowdId.(uint64)
	}
	return 0
}
