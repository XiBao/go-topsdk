package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type FenxiaoDistributorsGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewFenxiaoDistributorsGetRequest() (req *FenxiaoDistributorsGetRequest) {
	request := topsdk.Request{MethodName: "taobao.fenxiao.distributors.get", Params: make(map[string]interface{}, 1)}
	req = &FenxiaoDistributorsGetRequest{
		Request: &request,
	}
	return
}

// 分销商用户名列表。多个之间以“,”分隔;最多支持50个分销商用户名。
func (req *FenxiaoDistributorsGetRequest) SetNicks(nicks string) {
	req.Request.Params["nicks"] = nicks
}

func (req *FenxiaoDistributorsGetRequest) GetNicks() string {
	nicks, found := req.Request.Params["nicks"]
	if found {
		return nicks.(string)
	}
	return ""
}
