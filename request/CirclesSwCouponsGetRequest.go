package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type CirclesSwCouponsGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewCirclesSwCouponsGetRequest() (req *CirclesSwCouponsGetRequest) {
	request := topsdk.Request{MethodName: "taobao.circles.sw.coupons.get", Params: make(map[string]interface{}, 1)}
	req = &CirclesSwCouponsGetRequest{
		Request: &request,
	}
	return
}
