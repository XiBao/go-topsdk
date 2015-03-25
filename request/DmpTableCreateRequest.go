package TopRequest

import (
	"encoding/json"
	"github.com/XiBao/topsdk"
	"github.com/XiBao/topsdk/dmp"
)

type DmpTableCreateRequest struct {
	Request *topsdk.Request
}

// create new request
func NewDmpTableCreateRequest() (req *DmpTableCreateRequest) {
	request := topsdk.Request{MethodName: "taobao.dmp.table.create", Params: make(map[string]interface{}, 1)}
	req = &DmpTableCreateRequest{
		Request: &request,
	}
	return
}

// 选项
func (req *DmpTableCreateRequest) SetTable(table []*dmp.Table) {
	js, _ := json.Marshal(table)
	req.Request.Params["table"] = string(js)
}

func (req *DmpTableCreateRequest) GetSelects() []*dmp.Table {
	js, found := req.Request.Params["table"]
	if found {
		table := []*dmp.Table{}
		json.Unmarshal([]byte(js.(string)), &table)
		return table
	}
	return nil
}
