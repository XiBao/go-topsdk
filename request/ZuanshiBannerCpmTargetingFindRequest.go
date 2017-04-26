package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type ZuanshiBannerCpmTargetingFindRequest struct {
	Request *topsdk.Request
}

// create new request
func NewZuanshiBannerCpmTargetingFindRequest() (req *ZuanshiBannerCpmTargetingFindRequest) {
	request := topsdk.Request{MethodName: "taobao.zuanshi.banner.cpm.targeting.find", Params: make(map[string]interface{}, 0)}
	req = &ZuanshiBannerCpmTargetingFindRequest{
		Request: &request,
	}
	return
}
