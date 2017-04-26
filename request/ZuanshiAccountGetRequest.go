package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type ZuanshiAccountGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewZuanshiAccountGetRequest() (req *ZuanshiAccountGetRequest) {
	request := topsdk.Request{MethodName: "taobao.zuanshi.account.get", Params: make(map[string]interface{}, 0)}
	req = &ZuanshiAccountGetRequest{
		Request: &request,
	}
	return
}
