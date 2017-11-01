package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type TbkAdzoneCreateRequest struct {
	Request *topsdk.Request
}

// create new request
func NewTbkAdzoneCreateRequest() (req *TbkAdzoneCreateRequest) {
	request := topsdk.Request{MethodName: "taobao.tbk.adzone.create", Params: make(map[string]interface{}, 2)}
	req = &TbkAdzoneCreateRequest{
		Request: &request,
	}
	return
}

// 网站ID
func (req *TbkAdzoneCreateRequest) SetSiteId(siteId uint64) {
	req.Request.Params["site_id"] = siteId
}

func (req *TbkAdzoneCreateRequest) GetSiteId() uint64 {
	siteId, found := req.Request.Params["site_id"]
	if found {
		return siteId.(uint64)
	}
	return 0
}

// 广告位名称，最大长度64字符
func (req *TbkAdzoneCreateRequest) SetAdzoneName(adzoneName string) {
	req.Request.Params["adzone_name"] = adzoneName
}

func (req *TbkAdzoneCreateRequest) GetAdzoneName() string {
	adzoneName, found := req.Request.Params["num_iids"]
	if found {
		return adzoneName.(string)
	}
	return ""
}
