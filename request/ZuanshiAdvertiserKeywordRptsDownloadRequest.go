package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type ZuanshiAdvertiserKeywordRptsDownloadRequest struct {
	Request *topsdk.Request
}

// create new request
func NewZuanshiAdvertiserKeywordRptsDownloadRequest() (req *ZuanshiAdvertiserKeywordRptsDownloadRequest) {
	request := topsdk.Request{MethodName: "taobao.zuanshi.advertiser.keyword.rpts.download", Params: make(map[string]interface{}, 2)}
	req = &ZuanshiAdvertiserKeywordRptsDownloadRequest{
		Request: &request,
	}
	return
}

//开始时间
func (req *ZuanshiAdvertiserKeywordRptsDownloadRequest) SetStartTime(start_time string) {
	req.Request.Params["start_time"] = start_time
}

func (req *ZuanshiAdvertiserKeywordRptsDownloadRequest) GetStartTime() string {
	start_time, found := req.Request.Params["start_time"]
	if found {
		return start_time.(string)
	}
	return ""
}

//结束时间
func (req *ZuanshiAdvertiserKeywordRptsDownloadRequest) SetEndTime(end_time string) {
	req.Request.Params["end_time"] = end_time
}

func (req *ZuanshiAdvertiserKeywordRptsDownloadRequest) GetEndTime() string {
	end_time, found := req.Request.Params["end_time"]
	if found {
		return end_time.(string)
	}
	return ""
}
