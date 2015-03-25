package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type DmpCrowdValiddateUpdateRequest struct {
	Request *topsdk.Request
}

// create new request
func NewDmpCrowdValiddateUpdateRequest() (req *DmpCrowdValiddateUpdateRequest) {
	request := topsdk.Request{MethodName: "taobao.dmp.crowd.validdate.update", Params: make(map[string]interface{}, 2)}
	req = &DmpCrowdValiddateUpdateRequest{
		Request: &request,
	}
	return
}

// 人群有效期
func (req *DmpCrowdValiddateUpdateRequest) SetValidDate(validDate string) {
	req.Request.Params["valid_date"] = validDate
}

func (req *DmpCrowdValiddateUpdateRequest) GetValidDate() string {
	validDate, found := req.Request.Params["valid_date"]
	if found {
		return validDate.(string)
	}
	return ""
}

// 人群ID
func (req *DmpCrowdValiddateUpdateRequest) SetCrowdId(id uint64) {
	req.Request.Params["crowd_id"] = id
}

func (req *DmpCrowdValiddateUpdateRequest) GetCrowdId() uint64 {
	crowdId, found := req.Request.Params["crowd_id"]
	if found {
		return crowdId.(uint64)
	}
	return 0
}
