package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type WirelessShareTpwdCreateRequest struct {
	Request *topsdk.Request
}

// create new request
func NewWirelessShareTpwdCreateRequest() (req *WirelessShareTpwdCreateRequest) {
	request := topsdk.Request{MethodName: "taobao.wireless.share.tpwd.create", Params: make(map[string]interface{}, 1)}
	req = &WirelessShareTpwdCreateRequest{
		Request: &request,
	}
	return
}

// 口令参数
func (req *WirelessShareTpwdCreateRequest) SetTpwdParam(tpwdParam string) {
	req.Request.Params["tpwd_param"] = tpwdParam
}

func (req *WirelessShareTpwdCreateRequest) GetTpwdParam() string {
	tpwdParam, found := req.Request.Params["tpwd_param"]
	if found {
		return tpwdParam.(string)
	}
	return ""
}
