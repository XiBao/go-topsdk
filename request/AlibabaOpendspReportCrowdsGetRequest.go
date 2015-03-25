package TopRequest

import (
	"encoding/json"
	"github.com/XiBao/topsdk"
	"github.com/XiBao/topsdk/opendsp"
)

type AlibabaOpendspReportCrowdsGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewAlibabaOpendspReportCrowdsGetRequest() (req *AlibabaOpendspReportCrowdsGetRequest) {
	request := topsdk.Request{MethodName: "alibaba.opendsp.report.crowds.get", Params: make(map[string]interface{}, 2)}
	req = &AlibabaOpendspReportCrowdsGetRequest{
		Request: &request,
	}
	return
}

func (req *AlibabaOpendspReportCrowdsGetRequest) SetReportName(name string) {
	req.Request.Params["report_name"] = name
}

func (req *AlibabaOpendspReportCrowdsGetRequest) GetReportName() string {
	name, found := req.Request.Params["report_name"]
	if found {
		return name.(string)
	}
	return ""
}

func (req *AlibabaOpendspReportCrowdsGetRequest) SetQueryParam(query *opendsp.ReportParams) {
	js, _ := json.Marshal(query)
	req.Request.Params["query_param"] = string(js)
}

func (req *AlibabaOpendspReportCrowdsGetRequest) GetQueryParam() *opendsp.ReportParams {
	js, found := req.Request.Params["query_param"]
	if found {
		query := &opendsp.ReportParams{}
		json.Unmarshal([]byte(js.(string)), query)
		return query
	}
	return nil
}
