package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type DmpTagsGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewDmpTagsGetRequest() (req *DmpTagsGetRequest) {
	request := topsdk.Request{MethodName: "taobao.dmp.tags.get", Params: make(map[string]interface{}, 1)}
	req = &DmpTagsGetRequest{
		Request: &request,
	}
	return
}

// 标签名称模糊查询
func (req *DmpTagsGetRequest) SetTagName(tagName string) {
	req.Request.Params["tag_name"] = tagName
}

func (req *DmpTagsGetRequest) GetTagName() string {
	if tagName, found := req.Request.Params["tag_name"]; found {
		return tagName.(string)
	}
	return ""
}
