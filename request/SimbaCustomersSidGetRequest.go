package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type SimbaCustomersSidGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewSimbaCustomersSidGetRequest() (req *SimbaCustomersSidGetRequest) {
	request := topsdk.Request{MethodName: "taobao.simba.customers.sid.get", Params: make(map[string]interface{}, 0)}
	req = &SimbaCustomersSidGetRequest{
		Request: &request,
	}
	return
}
