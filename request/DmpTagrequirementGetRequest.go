package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type DmpTagrequirementGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewDmpTagrequirementGetRequest() (req *DmpTagrequirementGetRequest) {
	request := topsdk.Request{MethodName: "taobao.dmp.tagrequirement.get", Params: make(map[string]interface{}, 3)}
	req = &DmpTagrequirementGetRequest{
		Request: &request,
	}
	return
}

// 标签需求id
func (req *DmpTagrequirementGetRequest) SetId(id uint64) {
	req.Request.Params["id"] = id
}

func (req *DmpTagrequirementGetRequest) GetId() uint64 {
	id, found := req.Request.Params["id"]
	if found {
		return id.(uint64)
	}
	return 0
}

// 需求状态
func (req *DmpTagrequirementGetRequest) SetStatus(status uint8) {
	req.Request.Params["status"] = status
}

func (req *DmpTagrequirementGetRequest) GetStatus() uint8 {
	status, found := req.Request.Params["status"]
	if found {
		return status.(uint8)
	}
	return 0
}

// 需求类型
func (req *DmpTagrequirementGetRequest) SetType(atype uint) {
	req.Request.Params["type"] = atype
}

func (req *DmpTagrequirementGetRequest) GetType() uint {
	atype, found := req.Request.Params["type"]
	if found {
		return atype.(uint)
	}
	return 0
}
