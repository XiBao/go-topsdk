package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type ZuanshiBannerAdgroupGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewZuanshiBannerAdgroupGetRequest() (req *ZuanshiBannerAdgroupGetRequest) {
	request := topsdk.Request{MethodName: "taobao.zuanshi.banner.adgroup.get", Params: make(map[string]interface{}, 1)}
	req = &ZuanshiBannerAdgroupGetRequest{
		Request: &request,
	}
	return
}

// 单元ID
func (req *ZuanshiBannerAdgroupGetRequest) SetId(id uint64) {
	req.Request.Params["id"] = id
}

func (req *ZuanshiBannerAdgroupGetRequest) GetId() uint64 {
	if id, found := req.Request.Params["id"]; found {
		return id.(uint64)
	}
	return 0
}
