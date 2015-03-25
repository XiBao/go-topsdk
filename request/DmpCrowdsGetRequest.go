package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type DmpCrowdsGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewDmpCrowdsGetRequest() (req *DmpCrowdsGetRequest) {
	request := topsdk.Request{MethodName: "taobao.dmp.crowds.get", Params: make(map[string]interface{}, 0)}
	req = &DmpCrowdsGetRequest{
		Request: &request,
	}
	return
}
