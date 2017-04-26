package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type ZuanshiBannerCpcTargetingFindRequest struct {
	Request *topsdk.Request
}

// create new request
func NewZuanshiBannerCpcTargetingFindRequest() (req *ZuanshiBannerCpcTargetingFindRequest) {
	request := topsdk.Request{MethodName: "taobao.zuanshi.banner.cpc.targeting.find", Params: make(map[string]interface{}, 0)}
	req = &ZuanshiBannerCpcTargetingFindRequest{
		Request: &request,
	}
	return
}
