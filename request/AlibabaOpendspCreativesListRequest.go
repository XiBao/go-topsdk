package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type AlibabaOpendspCreativesListRequest struct {
	Request *topsdk.Request
}

// create new request
func NewAlibabaOpendspCreativesListRequest() (req *AlibabaOpendspCreativesListRequest) {
	request := topsdk.Request{MethodName: "alibaba.opendsp.creatives.list", Params: make(map[string]interface{}, 0)}
	req = &AlibabaOpendspCreativesListRequest{
		Request: &request,
	}
	return
}
