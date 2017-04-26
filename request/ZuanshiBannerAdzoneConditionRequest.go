package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type ZuanshiBannerAdzoneConditionRequest struct {
	Request *topsdk.Request
}

// create new request
func NewZuanshiBannerAdzoneConditionRequest() (req *ZuanshiBannerAdzoneConditionRequest) {
	request := topsdk.Request{MethodName: "taobao.zuanshi.banner.adzone.condition", Params: make(map[string]interface{}, 0)}
	req = &ZuanshiBannerAdzoneConditionRequest{
		Request: &request,
	}
	return
}
