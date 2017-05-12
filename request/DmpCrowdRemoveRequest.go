package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type DmpCrowdRemoveRequest struct {
	Request *topsdk.Request
}

// create new request
func NewDmpCrowdRemoveRequest() (req *DmpCrowdRemoveRequest) {
	request := topsdk.Request{MethodName: "taobao.dmp.crowd.remove", Params: make(map[string]interface{}, 1)}
	req = &DmpCrowdRemoveRequest{
		Request: &request,
	}
	return
}

// 过期时间
func (req *DmpCrowdRemoveRequest) SetCrowdId(id uint64) {
	req.Request.Params["crowd_id"] = id
}

func (req *DmpCrowdRemoveRequest) GetCrowdId() uint64 {
	id, found := req.Request.Params["crowd_id"]
	if found {
		return id.(uint64)
	}
	return 0
}
