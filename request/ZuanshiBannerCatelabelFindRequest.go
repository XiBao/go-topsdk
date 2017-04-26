package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type ZuanshiBannerCatelabelFindRequest struct {
	Request *topsdk.Request
}

// create new request
func NewZuanshiBannerCatelabelFindRequest() (req *ZuanshiBannerCatelabelFindRequest) {
	request := topsdk.Request{MethodName: "taobao.zuanshi.banner.catelabel.find", Params: make(map[string]interface{}, 0)}
	req = &ZuanshiBannerCatelabelFindRequest{
		Request: &request,
	}
	return
}
