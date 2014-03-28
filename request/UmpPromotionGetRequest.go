package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type UmpPromotionGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewUmpPromotionGetRequest() (req *UmpPromotionGetRequest) {
	request := topsdk.Request{MethodName: "taobao.ump.promotion.get", Params: make(map[string]interface{}, 1)}
	req = &UmpPromotionGetRequest{
		Request: &request,
	}
	return
}

// 商品id
func (req *UmpPromotionGetRequest) SetItemId(itemId uint64) {
	req.Request.Params["item_id"] = itemId
}

func (req *UmpPromotionGetRequest) GetItemId() uint64 {
	itemId, found := req.Request.Params["item_id"]
	if found {
		return itemId.(uint64)
	}
	return 0
}
