package TopRequest

import (
	"encoding/json"
	"github.com/XiBao/topsdk"
	"github.com/XiBao/topsdk/dmp"
)

type DmpRptCrowdsGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewDmpRptCrowdsGetRequest() (req *DmpRptCrowdsGetRequest) {
	request := topsdk.Request{MethodName: "taobao.dmp.rpt.crowds.get", Params: make(map[string]interface{}, 3)}
	req = &DmpRptCrowdsGetRequest{
		Request: &request,
	}
	return
}

// 人群效果查询对象
func (req *DmpRptCrowdsGetRequest) SetCrowdRptQuery(query *dmp.CrowdRptQuery) {
	js, _ := json.Marshal(query)
	req.Request.Params["crowd_rpt_query"] = string(js)
}

func (req *DmpRptCrowdsGetRequest) GetCrowdRptQuery() *dmp.CrowdRptQuery {
	js, found := req.Request.Params["crowd_rpt_query"]
	if found {
		query := &dmp.CrowdRptQuery{}
		json.Unmarshal([]byte(js.(string)), query)
		return query
	}
	return nil
}

// 从1开始
func (req *DmpRptCrowdsGetRequest) SetCurrentPage(currentPage uint) {
	req.Request.Params["current_page"] = currentPage
}

func (req *DmpRptCrowdsGetRequest) GetCurrentPage() uint {
	currentPage, found := req.Request.Params["current_page"]
	if found {
		return currentPage.(uint)
	}
	return 0
}

// 默认值为15
func (req *DmpRptCrowdsGetRequest) SetPageSize(pageSize uint) {
	req.Request.Params["page_size"] = pageSize
}

func (req *DmpRptCrowdsGetRequest) GetPageSize() uint {
	pageSize, found := req.Request.Params["page_size"]
	if found {
		return pageSize.(uint)
	}
	return 0
}
