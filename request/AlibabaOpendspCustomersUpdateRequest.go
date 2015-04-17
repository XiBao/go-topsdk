package TopRequest

import (
	"encoding/json"
	"github.com/XiBao/topsdk"
	"github.com/XiBao/topsdk/opendsp"
)

type AlibabaOpendspCustomersUpdateRequest struct {
	Request *topsdk.Request
}

// create new request
func NewAlibabaOpendspCustomersUpdateRequest() (req *AlibabaOpendspCustomersUpdateRequest) {
	request := topsdk.Request{MethodName: "alibaba.opendsp.customers.update", Params: make(map[string]interface{}, 1)}
	req = &AlibabaOpendspCustomersUpdateRequest{
		Request: &request,
	}
	return
}

func (req *AlibabaOpendspCustomersUpdateRequest) SetCustomerList(customers []*opendsp.Customer) {
	js, _ := json.Marshal(customers)
	req.Request.Params["customer_list"] = string(js)
}

func (req *AlibabaOpendspCustomersUpdateRequest) GetCustomerList() []*opendsp.Customer {
	js, found := req.Request.Params["customer_list"]
	if found {
		customers := []*opendsp.Customer{}
		json.Unmarshal([]byte(js.(string)), &customers)
		return customers
	}
	return nil
}
