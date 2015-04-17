package TopRequest

import (
	"encoding/json"
	"github.com/XiBao/topsdk"
	"github.com/XiBao/topsdk/opendsp"
)

type AlibabaOpendspCustomersAddRequest struct {
	Request *topsdk.Request
}

// create new request
func NewAlibabaOpendspCustomersAddRequest() (req *AlibabaOpendspCustomersAddRequest) {
	request := topsdk.Request{MethodName: "alibaba.opendsp.customers.add", Params: make(map[string]interface{}, 1)}
	req = &AlibabaOpendspCustomersAddRequest{
		Request: &request,
	}
	return
}

func (req *AlibabaOpendspCustomersAddRequest) SetCustomerList(customers []*opendsp.Customer) {
	js, _ := json.Marshal(customers)
	req.Request.Params["customer_list"] = string(js)
}

func (req *AlibabaOpendspCustomersAddRequest) GetCustomerList() []*opendsp.Customer {
	js, found := req.Request.Params["customer_list"]
	if found {
		customers := []*opendsp.Customer{}
		json.Unmarshal([]byte(js.(string)), &customers)
		return customers
	}
	return nil
}
