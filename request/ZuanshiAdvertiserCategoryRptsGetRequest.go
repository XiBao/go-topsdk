package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type ZuanshiAdvertiserCategoryRptsGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewZuanshiAdvertiserCategoryRptsGetRequest() (req *ZuanshiAdvertiserCategoryRptsGetRequest) {
	request := topsdk.Request{MethodName: "taobao.zuanshi.advertiser.category.rpts.get", Params: make(map[string]interface{}, 2)}
	req = &ZuanshiAdvertiserCategoryRptsGetRequest{
		Request: &request,
	}
	return
}

//开始时间
func (req *ZuanshiAdvertiserCategoryRptsGetRequest) SetStartTime(start_time string) {
	req.Request.Params["start_time"] = start_time
}

func (req *ZuanshiAdvertiserCategoryRptsGetRequest) GetStartTime() string {
	start_time, found := req.Request.Params["start_time"]
	if found {
		return start_time.(string)
	}
	return ""
}

//结束时间
func (req *ZuanshiAdvertiserCategoryRptsGetRequest) SetEndTime(end_time string) {
	req.Request.Params["end_time"] = end_time
}

func (req *ZuanshiAdvertiserCategoryRptsGetRequest) GetEndTime() string {
	end_time, found := req.Request.Params["end_time"]
	if found {
		return end_time.(string)
	}
	return ""
}
