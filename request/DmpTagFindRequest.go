package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type DmpTagFindRequest struct {
	Request *topsdk.Request
}

// create new request
func NewDmpTagFindRequest() (req *DmpTagFindRequest) {
	request := topsdk.Request{MethodName: "taobao.dmp.tag.find", Params: make(map[string]interface{}, 1)}
	req = &DmpTagFindRequest{
		Request: &request,
	}
	return
}

// 标签名称模糊查询
func (req *DmpTagFindRequest) SetTagName(tagName string) {
	req.Request.Params["tag_name"] = tagName
}

func (req *DmpTagFindRequest) GetTagName() string {
	if tagName, found := req.Request.Params["tag_name"]; found {
		return tagName.(string)
	}
	return ""
}
