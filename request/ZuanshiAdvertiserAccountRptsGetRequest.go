package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type ZuanshiAdvertiserAccountRptsGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewZuanshiAdvertiserAccountRptsGetRequest() (req *ZuanshiAdvertiserAccountRptsGetRequest) {
	request := topsdk.Request{MethodName: "taobao.zuanshi.advertiser.account.rpts.get", Params: make(map[string]interface{}, 3)}
	req = &ZuanshiAdvertiserAccountRptsGetRequest{
		Request: &request,
	}
	return
}

//开始时间
func (req *ZuanshiAdvertiserAccountRptsGetRequest) SetStartTime(start_time string) {
	req.Request.Params["start_time"] = start_time
}

func (req *ZuanshiAdvertiserAccountRptsGetRequest) GetStartTime() string {
	start_time, found := req.Request.Params["start_time"]
	if found {
		return start_time.(string)
	}
	return ""
}

//结束时间
func (req *ZuanshiAdvertiserAccountRptsGetRequest) SetEndTime(end_time string) {
	req.Request.Params["end_time"] = end_time
}

func (req *ZuanshiAdvertiserAccountRptsGetRequest) GetEndTime() string {
	end_time, found := req.Request.Params["end_time"]
	if found {
		return end_time.(string)
	}
	return ""
}

//计划类型，1代表展示网络，2代表明星店铺，3代表视频网络
func (req *ZuanshiAdvertiserAccountRptsGetRequest) SetRptType(rptType uint8) {
	req.Request.Params["rpt_type"] = rptType
}

func (req *ZuanshiAdvertiserAccountRptsGetRequest) GetRptType() uint8 {
	rptType, found := req.Request.Params["rpt_type"]
	if found {
		return rptType.(uint8)
	}
	return 0
}
