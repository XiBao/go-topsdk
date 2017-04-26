package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type ZuanshiBannerDmpCrowdApplyRequest struct {
	Request *topsdk.Request
}

// create new request
func NewZuanshiBannerDmpCrowdApplyRequest() (req *ZuanshiBannerDmpCrowdApplyRequest) {
	request := topsdk.Request{MethodName: "taobao.zuanshi.banner.dmp.crowd.apply", Params: make(map[string]interface{}, 1)}
	req = &ZuanshiBannerDmpCrowdApplyRequest{
		Request: &request,
	}
	return
}

func (req *ZuanshiBannerDmpCrowdApplyRequest) SetCrowdId(id uint64) {
	req.Request.Params["crowd_id"] = id
}

func (req *ZuanshiBannerDmpCrowdApplyRequest) GetCrowdId() uint64 {
	if crowdId, found := req.Request.Params["crowd_id"]; found {
		return crowdId.(uint64)
	}
	return 0
}
