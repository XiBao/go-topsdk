package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type ZuanshiShopItemFindRequest struct {
	Request *topsdk.Request
}

// create new request
func NewZuanshiShopItemFindRequest() (req *ZuanshiShopItemFindRequest) {
	request := topsdk.Request{MethodName: "taobao.zuanshi.shop.item.find", Params: make(map[string]interface{}, 1)}
	req = &ZuanshiShopItemFindRequest{
		Request: &request,
	}
	return
}

// 宝贝名称关键字
func (req *ZuanshiShopItemFindRequest) SetItemName(itemName string) {
	req.Request.Params["item_name"] = itemName
}

func (req *ZuanshiShopItemFindRequest) GetItemName() string {
	if itemName, found := req.Request.Params["item_name"]; found {
		return itemName.(string)
	}
	return ""
}
