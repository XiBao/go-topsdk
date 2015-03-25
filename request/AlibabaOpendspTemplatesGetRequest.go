package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type AlibabaOpendspTemplatesGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewAlibabaOpendspTemplatesGetRequest() (req *AlibabaOpendspTemplatesGetRequest) {
	request := topsdk.Request{MethodName: "alibaba.opendsp.templates.get", Params: make(map[string]interface{}, 1)}
	req = &AlibabaOpendspTemplatesGetRequest{
		Request: &request,
	}
	return
}

func (req *AlibabaOpendspTemplatesGetRequest) SetTemplateIdList(ids string) {
	req.Request.Params["template_id_list"] = ids
}

func (req *AlibabaOpendspTemplatesGetRequest) GetTemplateIdList() string {
	ids, found := req.Request.Params["template_id_list"]
	if found {
		return ids.(string)
	}
	return ""
}
