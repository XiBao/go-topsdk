package TopRequest

import (
	"encoding/json"
	"github.com/XiBao/topsdk"
	"github.com/XiBao/topsdk/opendsp"
)

type AlibabaOpendspTemplatesAddRequest struct {
	Request *topsdk.Request
}

// create new request
func NewAlibabaOpendspTemplatesAddRequest() (req *AlibabaOpendspTemplatesAddRequest) {
	request := topsdk.Request{MethodName: "alibaba.opendsp.templates.add", Params: make(map[string]interface{}, 1)}
	req = &AlibabaOpendspTemplatesAddRequest{
		Request: &request,
	}
	return
}

func (req *AlibabaOpendspTemplatesAddRequest) SetTemplateList(templates []*opendsp.Template) {
	js, _ := json.Marshal(templates)
	req.Request.Params["template_list"] = string(js)
}

func (req *AlibabaOpendspTemplatesAddRequest) GetTemplateList() []*opendsp.Template {
	js, found := req.Request.Params["template_list"]
	if found {
		templates := []*opendsp.Template{}
		json.Unmarshal([]byte(js.(string)), &templates)
		return templates
	}
	return nil
}
