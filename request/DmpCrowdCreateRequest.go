package TopRequest

import (
	"encoding/json"
	"github.com/XiBao/topsdk"
	"github.com/XiBao/topsdk/dmp"
)

type DmpCrowdCreateRequest struct {
	Request *topsdk.Request
}

// create new request
func NewDmpCrowdCreateRequest() (req *DmpCrowdCreateRequest) {
	request := topsdk.Request{MethodName: "taobao.dmp.crowd.create", Params: make(map[string]interface{}, 4)}
	req = &DmpCrowdCreateRequest{
		Request: &request,
	}
	return
}

// 过期时间
func (req *DmpCrowdCreateRequest) SetValidDate(validDate string) {
	req.Request.Params["valid_date"] = validDate
}

func (req *DmpCrowdCreateRequest) GetValidDate() string {
	validDate, found := req.Request.Params["valid_date"]
	if found {
		return validDate.(string)
	}
	return ""
}

// 人群名称
func (req *DmpCrowdCreateRequest) SetCrowdName(crowdName string) {
	req.Request.Params["crowd_name"] = crowdName
}

func (req *DmpCrowdCreateRequest) GetCrowdName() string {
	crowdName, found := req.Request.Params["crowd_name"]
	if found {
		return crowdName.(string)
	}
	return ""
}

// 选项
func (req *DmpCrowdCreateRequest) SetSelects(selects *dmp.SelectTagOptionSet) {
	js, _ := json.Marshal(selects)
	req.Request.Params["select_tag_option_set_d_t_o"] = string(js)
}

func (req *DmpCrowdCreateRequest) GetSelects() *dmp.SelectTagOptionSet {
	js, found := req.Request.Params["select_tag_option_set_d_t_o"]
	if found {
		var selects dmp.SelectTagOptionSet
		json.Unmarshal([]byte(js.(string)), &selects)
		return &selects
	}
	return nil
}

// 放大倍数
func (req *DmpCrowdCreateRequest) SetLookalikeMultiple(lookalikeMultiple uint) {
	req.Request.Params["lookalike_multiple"] = lookalikeMultiple
}

func (req *DmpCrowdCreateRequest) GetLookalikeMultiple() uint {
	lookalikeMultiple, found := req.Request.Params["lookalike_multiple"]
	if found {
		return lookalikeMultiple.(uint)
	}
	return 0
}
