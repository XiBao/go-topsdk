package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type DmpCrowdNameGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewDmpCrowdNameGetRequest() (req *DmpCrowdNameGetRequest) {
	request := topsdk.Request{MethodName: "taobao.dmp.crowd.name.get", Params: make(map[string]interface{}, 1)}
	req = &DmpCrowdNameGetRequest{
		Request: &request,
	}
	return
}

// 人群名称
func (req *DmpCrowdNameGetRequest) SetCrowdName(crowd_name string) {
	req.Request.Params["crowd_name"] = crowd_name
}

func (req *DmpCrowdNameGetRequest) GetCrowdName() string {
	crowd_name, found := req.Request.Params["crowd_name"]
	if found {
		return crowd_name.(string)
	}
	return ""
}
