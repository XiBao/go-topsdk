package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type DmpDatauploadEndRequest struct {
	Request *topsdk.Request
}

// create new request
func NewDmpDatauploadEndRequest() (req *DmpDatauploadEndRequest) {
	request := topsdk.Request{MethodName: "taobao.dmp.dataupload.end", Params: make(map[string]interface{}, 1)}
	req = &DmpDatauploadEndRequest{
		Request: &request,
	}
	return
}

// 表名
func (req *DmpDatauploadEndRequest) SetTablecode(tablecode string) {
	req.Request.Params["tablecode"] = tablecode
}

func (req *DmpDatauploadEndRequest) GetTablecode() string {
	tablecode, found := req.Request.Params["tablecode"]
	if found {
		return tablecode.(string)
	}
	return ""
}
