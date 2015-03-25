package TopRequest

import (
	"encoding/json"
	"github.com/XiBao/topsdk"
	"github.com/XiBao/topsdk/opendsp"
)

type AlibabaOpendspReportAdgroupsGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewAlibabaOpendspReportAdgroupsGetRequest() (req *AlibabaOpendspReportAdgroupsGetRequest) {
	request := topsdk.Request{MethodName: "alibaba.opendsp.report.adgroups.get", Params: make(map[string]interface{}, 2)}
	req = &AlibabaOpendspReportAdgroupsGetRequest{
		Request: &request,
	}
	return
}

func (req *AlibabaOpendspReportAdgroupsGetRequest) SetReportName(name string) {
	req.Request.Params["report_name"] = name
}

func (req *AlibabaOpendspReportAdgroupsGetRequest) GetReportName() string {
	name, found := req.Request.Params["report_name"]
	if found {
		return name.(string)
	}
	return ""
}

func (req *AlibabaOpendspReportAdgroupsGetRequest) SetQueryParam(query *opendsp.ReportParams) {
	js, _ := json.Marshal(query)
	req.Request.Params["query_param"] = string(js)
}

func (req *AlibabaOpendspReportAdgroupsGetRequest) GetQueryParam() *opendsp.ReportParams {
	js, found := req.Request.Params["query_param"]
	if found {
		query := &opendsp.ReportParams{}
		json.Unmarshal([]byte(js.(string)), query)
		return query
	}
	return nil
}
