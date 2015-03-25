package TopRequest

import (
	"encoding/json"
	"github.com/XiBao/topsdk"
	"github.com/XiBao/topsdk/dmp"
)

type DmpCrowdAddRequest struct {
	Request *topsdk.Request
}

// create new request
func NewDmpCrowdAddRequest() (req *DmpCrowdAddRequest) {
	request := topsdk.Request{MethodName: "taobao.dmp.crowd.add", Params: make(map[string]interface{}, 4)}
	req = &DmpCrowdAddRequest{
		Request: &request,
	}
	return
}

// 过期时间
func (req *DmpCrowdAddRequest) SetValidDate(validDate string) {
	req.Request.Params["valid_date"] = validDate
}

func (req *DmpCrowdAddRequest) GetValidDate() string {
	validDate, found := req.Request.Params["valid_date"]
	if found {
		return validDate.(string)
	}
	return ""
}

// 人群名称
func (req *DmpCrowdAddRequest) SetCrowdName(crowdName string) {
	req.Request.Params["crowd_name"] = crowdName
}

func (req *DmpCrowdAddRequest) GetCrowdName() string {
	crowdName, found := req.Request.Params["crowd_name"]
	if found {
		return crowdName.(string)
	}
	return ""
}

// 选项
func (req *DmpCrowdAddRequest) SetSelects(selects []*dmp.DmpSelectTagOption) {
	js, _ := json.Marshal(selects)
	req.Request.Params["selects"] = string(js)
}

func (req *DmpCrowdAddRequest) GetSelects() []*dmp.DmpSelectTagOption {
	js, found := req.Request.Params["selects"]
	if found {
		selects := []*dmp.DmpSelectTagOption{}
		json.Unmarshal([]byte(js.(string)), &selects)
		return selects
	}
	return nil
}

// 放大倍数
func (req *DmpCrowdAddRequest) SetLookalikeMultiple(lookalikeMultiple uint) {
	req.Request.Params["lookalike_multiple"] = lookalikeMultiple
}

func (req *DmpCrowdAddRequest) GetLookalikeMultiple() uint {
	lookalikeMultiple, found := req.Request.Params["lookalike_multiple"]
	if found {
		return lookalikeMultiple.(uint)
	}
	return 0
}
