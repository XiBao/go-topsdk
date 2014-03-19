package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type FenxiaoDistributorProcuctStaticGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewFenxiaoDistributorProcuctStaticGetRequest() (req *FenxiaoDistributorProcuctStaticGetRequest) {
	request := topsdk.Request{MethodName: "taobao.fenxiao.distributor.procuct.static.get", Params: make(map[string]interface{}, 2)}
	req = &FenxiaoDistributorProcuctStaticGetRequest{
		Request: &request,
	}
	return
}

// 分销商淘宝店主nick
func (req *FenxiaoDistributorProcuctStaticGetRequest) SetDistributorUserNick(distributorUserNick string) {
	req.Request.Params["distributor_user_nick"] = distributorUserNick
}

func (req *FenxiaoDistributorProcuctStaticGetRequest) GetDistributorUserNick() string {
	distributorUserNick, found := req.Request.Params["distributor_user_nick"]
	if found {
		return distributorUserNick.(string)
	}
	return ""
}

// 供应商商品id，一次可以传多个，每次最多40个。以,(英文)作为分隔符。
func (req *FenxiaoDistributorProcuctStaticGetRequest) SetProductIdArray(productIdArray string) {
	req.Request.Params["product_id_array"] = productIdArray
}

func (req *FenxiaoDistributorProcuctStaticGetRequest) GetProductIdArray() string {
	productIdArray, found := req.Request.Params["product_id_array"]
	if found {
		return productIdArray.(string)
	}
	return ""
}
