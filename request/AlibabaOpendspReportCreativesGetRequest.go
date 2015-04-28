package TopRequest

import (
	"encoding/json"
	"github.com/XiBao/topsdk"
	"github.com/XiBao/topsdk/opendsp"
)

type AlibabaOpendspReportCreativesGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewAlibabaOpendspReportCreativesGetRequest() (req *AlibabaOpendspReportCreativesGetRequest) {
	request := topsdk.Request{MethodName: "alibaba.opendsp.report.creatives.get", Params: make(map[string]interface{}, 2)}
	req = &AlibabaOpendspReportCreativesGetRequest{
		Request: &request,
	}
	return
}

func (req *AlibabaOpendspReportCreativesGetRequest) SetReportName(name string) {
	req.Request.Params["report_name"] = name
}

func (req *AlibabaOpendspReportCreativesGetRequest) GetReportName() string {
	name, found := req.Request.Params["report_name"]
	if found {
		return name.(string)
	}
	return ""
}

func (req *AlibabaOpendspReportCreativesGetRequest) SetQueryParam(query *opendsp.CreativeReportParams) {
	js, _ := json.Marshal(query)
	req.Request.Params["query_param"] = string(js)
}

func (req *AlibabaOpendspReportCreativesGetRequest) GetQueryParam() *opendsp.CreativeReportParams {
	js, found := req.Request.Params["query_param"]
	if found {
		query := &opendsp.CreativeReportParams{}
		json.Unmarshal([]byte(js.(string)), query)
		return query
	}
	return nil
}
