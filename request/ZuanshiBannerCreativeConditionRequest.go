package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type ZuanshiBannerCreativeConditionRequest struct {
	Request *topsdk.Request
}

// create new request
func NewZuanshiBannerCreativeConditionRequest() (req *ZuanshiBannerCreativeConditionRequest) {
	request := topsdk.Request{MethodName: "taobao.zuanshi.banner.creative.condition", Params: make(map[string]interface{}, 0)}
	req = &ZuanshiBannerCreativeConditionRequest{
		Request: &request,
	}
	return
}
