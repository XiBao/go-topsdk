package TopRequest

import (
	"encoding/json"
	"github.com/XiBao/topsdk"
	"github.com/XiBao/topsdk/dmp"
)

type DmpTableUpdateRequest struct {
	Request *topsdk.Request
}

// create new request
func NewDmpTableUpdateRequest() (req *DmpTableUpdateRequest) {
	request := topsdk.Request{MethodName: "taobao.dmp.table.update", Params: make(map[string]interface{}, 1)}
	req = &DmpTableUpdateRequest{
		Request: &request,
	}
	return
}

// 选项
func (req *DmpTableUpdateRequest) SetTable(table []*dmp.Table) {
	js, _ := json.Marshal(table)
	req.Request.Params["table"] = string(js)
}

func (req *DmpTableUpdateRequest) GetSelects() []*dmp.Table {
	js, found := req.Request.Params["table"]
	if found {
		table := []*dmp.Table{}
		json.Unmarshal([]byte(js.(string)), &table)
		return table
	}
	return nil
}
