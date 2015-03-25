package TopRequest

import (
	"encoding/json"
	"github.com/XiBao/topsdk"
	"github.com/XiBao/topsdk/opendsp"
)

type AlibabaOpendspReportCampaignsGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewAlibabaOpendspReportCampaignsGetRequest() (req *AlibabaOpendspReportCampaignsGetRequest) {
	request := topsdk.Request{MethodName: "alibaba.opendsp.report.campaigns.get", Params: make(map[string]interface{}, 2)}
	req = &AlibabaOpendspReportCampaignsGetRequest{
		Request: &request,
	}
	return
}

func (req *AlibabaOpendspReportCampaignsGetRequest) SetReportName(name string) {
	req.Request.Params["report_name"] = name
}

func (req *AlibabaOpendspReportCampaignsGetRequest) GetReportName() string {
	name, found := req.Request.Params["report_name"]
	if found {
		return name.(string)
	}
	return ""
}

func (req *AlibabaOpendspReportCampaignsGetRequest) SetQueryParam(query *opendsp.ReportParams) {
	js, _ := json.Marshal(query)
	req.Request.Params["query_param"] = string(js)
}

func (req *AlibabaOpendspReportCampaignsGetRequest) GetQueryParam() *opendsp.ReportParams {
	js, found := req.Request.Params["query_param"]
	if found {
		query := &opendsp.ReportParams{}
		json.Unmarshal([]byte(js.(string)), query)
		return query
	}
	return nil
}
