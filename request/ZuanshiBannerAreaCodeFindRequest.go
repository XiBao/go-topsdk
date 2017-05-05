package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type ZuanshiBannerAreaCodeFindRequest struct {
	Request *topsdk.Request
}

// create new request
func NewZuanshiBannerAreaCodeFindRequest() (req *ZuanshiBannerAreaCodeFindRequest) {
	request := topsdk.Request{MethodName: "taobao.zuanshi.banner.area.code.find", Params: make(map[string]interface{}, 0)}
	req = &ZuanshiBannerAreaCodeFindRequest{
		Request: &request,
	}
	return
}
