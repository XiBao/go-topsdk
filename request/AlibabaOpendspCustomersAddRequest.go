package TopRequest

import (
	"encoding/json"
	"github.com/XiBao/topsdk"
	"github.com/XiBao/topsdk/opendsp"
)

type AlibabaOpendspCustomersAdd struct {
	Request *topsdk.Request
}

// create new request
func NewAlibabaOpendspCustomersAdd() (req *AlibabaOpendspCustomersAdd) {
	request := topsdk.Request{MethodName: "alibaba.opendsp.customers.add", Params: make(map[string]interface{}, 1)}
	req = &AlibabaOpendspCustomersAdd{
		Request: &request,
	}
	return
}

func (req *AlibabaOpendspCustomersAdd) SetCustomerList(customers []*opendsp.Customer) {
	js, _ := json.Marshal(customers)
	req.Request.Params["customer_list"] = string(js)
}

func (req *AlibabaOpendspCustomersAdd) GetCustomerList() []*opendsp.Customer {
	js, found := req.Request.Params["customer_list"]
	if found {
		customers := []*opendsp.Customer{}
		json.Unmarshal([]byte(js.(string)), &customers)
		return customers
	}
	return nil
}
