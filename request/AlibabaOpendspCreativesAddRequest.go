package TopRequest

import (
	"encoding/json"
	"github.com/XiBao/topsdk"
	"github.com/XiBao/topsdk/opendsp"
)

type AlibabaOpendspCreativesAddRequest struct {
	Request *topsdk.Request
}

// create new request
func NewAlibabaOpendspCreativesAddRequest() (req *AlibabaOpendspCreativesAddRequest) {
	request := topsdk.Request{MethodName: "alibaba.opendsp.creatives.add", Params: make(map[string]interface{}, 1)}
	req = &AlibabaOpendspCreativesAddRequest{
		Request: &request,
	}
	return
}

func (req *AlibabaOpendspCreativesAddRequest) SetCreativeList(creatives []*opendsp.Creative) {
	js, _ := json.Marshal(creatives)
	req.Request.Params["creative_list"] = string(js)
}

func (req *AlibabaOpendspCreativesAddRequest) GetCreativeList() []*opendsp.Creative {
	js, found := req.Request.Params["creative_list"]
	if found {
		creatives := []*opendsp.Creative{}
		json.Unmarshal([]byte(js.(string)), &creatives)
		return creatives
	}
	return nil
}
