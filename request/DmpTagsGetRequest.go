package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type DmpTagsGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewDmpTagsGetRequest() (req *DmpTagsGetRequest) {
	request := topsdk.Request{MethodName: "taobao.dmp.tags.get", Params: make(map[string]interface{}, 0)}
	req = &DmpTagsGetRequest{
		Request: &request,
	}
	return
}
