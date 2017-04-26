package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type ZuanshiBannerAccountBudgetGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewZuanshiBannerAccountBudgetGetRequest() (req *ZuanshiBannerAccountBudgetGetRequest) {
	request := topsdk.Request{MethodName: "taobao.zuanshi.banner.account.budget.get", Params: make(map[string]interface{}, 0)}
	req = &ZuanshiBannerAccountBudgetGetRequest{
		Request: &request,
	}
	return
}
