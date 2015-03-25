package TopRequest

import (
	"encoding/json"
	"github.com/XiBao/topsdk"
	"github.com/XiBao/topsdk/opendsp"
)

type AlibabaOpendspCustomersUpdate struct {
	Request *topsdk.Request
}

// create new request
func NewAlibabaOpendspCustomersUpdate() (req *AlibabaOpendspCustomersUpdate) {
	request := topsdk.Request{MethodName: "alibaba.opendsp.customers.update", Params: make(map[string]interface{}, 1)}
	req = &AlibabaOpendspCustomersUpdate{
		Request: &request,
	}
	return
}

func (req *AlibabaOpendspCustomersUpdate) SetCustomerList(customers []*opendsp.Customer) {
	js, _ := json.Marshal(customers)
	req.Request.Params["customer_list"] = string(js)
}

func (req *AlibabaOpendspCustomersUpdate) GetCustomerList() []*opendsp.Customer {
	js, found := req.Request.Params["customer_list"]
	if found {
		customers := []*opendsp.Customer{}
		json.Unmarshal([]byte(js.(string)), &customers)
		return customers
	}
	return nil
}
