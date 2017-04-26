package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type ZuanshiBannerDmpFindRequest struct {
	Request *topsdk.Request
}

// create new request
func NewZuanshiBannerDmpFindRequest() (req *ZuanshiBannerDmpFindRequest) {
	request := topsdk.Request{MethodName: "taobao.zuanshi.banner.dmp.find", Params: make(map[string]interface{}, 0)}
	req = &ZuanshiBannerDmpFindRequest{
		Request: &request,
	}
	return
}
