package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type FenxiaoDistributorArchivesGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewFenxiaoDistributorArchivesGetRequest() (req *FenxiaoDistributorArchivesGetRequest) {
	request := topsdk.Request{MethodName: "taobao.fenxiao.distributor.archives.get", Params: make(map[string]interface{}, 1)}
	req = &FenxiaoDistributorArchivesGetRequest{
		Request: &request,
	}
	return
}

// 分销商淘宝店主nick
func (req *FenxiaoDistributorArchivesGetRequest) SetDistributorUserNick(distributorUserNick string) {
	req.Request.Params["distributor_user_nick"] = distributorUserNick
}

func (req *FenxiaoDistributorArchivesGetRequest) GetDistributorUserNick() string {
	distributorUserNick, found := req.Request.Params["distributor_user_nick"]
	if found {
		return distributorUserNick.(string)
	}
	return ""
}
