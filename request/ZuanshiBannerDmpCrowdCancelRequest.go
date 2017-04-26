package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type ZuanshiBannerDmpCrowdCancelRequest struct {
	Request *topsdk.Request
}

// create new request
func NewZuanshiBannerDmpCrowdCancelRequest() (req *ZuanshiBannerDmpCrowdCancelRequest) {
	request := topsdk.Request{MethodName: "taobao.zuanshi.banner.dmp.crowd.cancel", Params: make(map[string]interface{}, 1)}
	req = &ZuanshiBannerDmpCrowdCancelRequest{
		Request: &request,
	}
	return
}

func (req *ZuanshiBannerDmpCrowdCancelRequest) SetCrowdId(id uint64) {
	req.Request.Params["crowd_id"] = id
}

func (req *ZuanshiBannerDmpCrowdCancelRequest) GetCrowdId() uint64 {
	if crowdId, found := req.Request.Params["crowd_id"]; found {
		return crowdId.(uint64)
	}
	return 0
}
