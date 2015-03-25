package TopRequest

import (
	"encoding/json"
	"github.com/XiBao/topsdk"
	"github.com/XiBao/topsdk/opendsp"
)

type AlibabaOpendspCreativesUpdateRequest struct {
	Request *topsdk.Request
}

// create new request
func NewAlibabaOpendspCreativesUpdateRequest() (req *AlibabaOpendspCreativesUpdateRequest) {
	request := topsdk.Request{MethodName: "alibaba.opendsp.creatives.update", Params: make(map[string]interface{}, 1)}
	req = &AlibabaOpendspCreativesUpdateRequest{
		Request: &request,
	}
	return
}

func (req *AlibabaOpendspCreativesUpdateRequest) SetCreativeList(creatives []*opendsp.Creative) {
	js, _ := json.Marshal(creatives)
	req.Request.Params["creative_list"] = string(js)
}

func (req *AlibabaOpendspCreativesUpdateRequest) GetCreativeList() []*opendsp.Creative {
	js, found := req.Request.Params["creative_list"]
	if found {
		creatives := []*opendsp.Creative{}
		json.Unmarshal([]byte(js.(string)), &creatives)
		return creatives
	}
	return nil
}
