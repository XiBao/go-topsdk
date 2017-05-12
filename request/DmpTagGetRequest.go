package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type DmpTagGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewDmpTagGetRequest() (req *DmpTagGetRequest) {
	request := topsdk.Request{MethodName: "taobao.dmp.tag.get", Params: make(map[string]interface{}, 1)}
	req = &DmpTagGetRequest{
		Request: &request,
	}
	return
}

// 标签名称模糊查询
func (req *DmpTagGetRequest) SetTagId(id uint64) {
	req.Request.Params["tag_id"] = id
}

func (req *DmpTagGetRequest) GetTagId() uint64 {
	id, found := req.Request.Params["tag_id"]
	if found {
		return id.(uint64)
	}
	return 0
}
