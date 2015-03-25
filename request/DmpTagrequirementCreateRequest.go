package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type DmpTagrequirementCreateRequest struct {
	Request *topsdk.Request
}

// create new request
func NewDmpTagrequirementCreateRequest() (req *DmpTagrequirementCreateRequest) {
	request := topsdk.Request{MethodName: "taobao.dmp.tagrequirement.create", Params: make(map[string]interface{}, 4)}
	req = &DmpTagrequirementCreateRequest{
		Request: &request,
	}
	return
}

// 扩展属性
func (req *DmpTagrequirementCreateRequest) SetAttributes(attributes string) {
	req.Request.Params["attributes"] = attributes
}

func (req *DmpTagrequirementCreateRequest) GetAttributes() string {
	attributes, found := req.Request.Params["attributes"]
	if found {
		return attributes.(string)
	}
	return ""
}

// 内容
func (req *DmpTagrequirementCreateRequest) SetContent(content string) {
	req.Request.Params["content"] = content
}

func (req *DmpTagrequirementCreateRequest) GetContent() string {
	content, found := req.Request.Params["content"]
	if found {
		return content.(string)
	}
	return ""
}

// 标题
func (req *DmpTagrequirementCreateRequest) SetTitle(title string) {
	req.Request.Params["title"] = title
}

func (req *DmpTagrequirementCreateRequest) GetTitle() string {
	title, found := req.Request.Params["title"]
	if found {
		return title.(string)
	}
	return ""
}

// 类型
func (req *DmpTagrequirementCreateRequest) SetType(atype uint) {
	req.Request.Params["type"] = atype
}

func (req *DmpTagrequirementCreateRequest) GetType() uint {
	atype, found := req.Request.Params["type"]
	if found {
		return atype.(uint)
	}
	return 0
}
