package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type DmpDatauploadStartRequest struct {
	Request *topsdk.Request
}

// create new request
func NewDmpDatauploadStartRequest() (req *DmpDatauploadStartRequest) {
	request := topsdk.Request{MethodName: "taobao.dmp.dataupload.start", Params: make(map[string]interface{}, 1)}
	req = &DmpDatauploadStartRequest{
		Request: &request,
	}
	return
}

// 表名
func (req *DmpDatauploadStartRequest) SetTablecode(tablecode string) {
	req.Request.Params["tablecode"] = tablecode
}

func (req *DmpDatauploadStartRequest) GetTablecode() string {
	tablecode, found := req.Request.Params["tablecode"]
	if found {
		return tablecode.(string)
	}
	return ""
}
