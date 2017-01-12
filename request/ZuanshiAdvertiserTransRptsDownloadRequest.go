package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type ZuanshiAdvertiserTransRptsDownloadRequest struct {
	Request *topsdk.Request
}

// create new request
func NewZuanshiAdvertiserTransRptsDownloadRequest() (req *ZuanshiAdvertiserTransRptsDownloadRequest) {
	request := topsdk.Request{MethodName: "taobao.zuanshi.advertiser.trans.rpts.download", Params: make(map[string]interface{}, 3)}
	req = &ZuanshiAdvertiserTransRptsDownloadRequest{
		Request: &request,
	}
	return
}

//开始时间
func (req *ZuanshiAdvertiserTransRptsDownloadRequest) SetStartTime(start_time string) {
	req.Request.Params["start_time"] = start_time
}

func (req *ZuanshiAdvertiserTransRptsDownloadRequest) GetStartTime() string {
	start_time, found := req.Request.Params["start_time"]
	if found {
		return start_time.(string)
	}
	return ""
}

//结束时间
func (req *ZuanshiAdvertiserTransRptsDownloadRequest) SetEndTime(end_time string) {
	req.Request.Params["end_time"] = end_time
}

func (req *ZuanshiAdvertiserTransRptsDownloadRequest) GetEndTime() string {
	end_time, found := req.Request.Params["end_time"]
	if found {
		return end_time.(string)
	}
	return ""
}

//计划类型，1代表展示网络，2代表明星店铺，3代表视频网络
func (req *ZuanshiAdvertiserTransRptsDownloadRequest) SetRptType(rptType uint8) {
	req.Request.Params["rpt_type"] = rptType
}

func (req *ZuanshiAdvertiserTransRptsDownloadRequest) GetRptType() uint8 {
	rptType, found := req.Request.Params["rpt_type"]
	if found {
		return rptType.(uint8)
	}
	return 0
}
