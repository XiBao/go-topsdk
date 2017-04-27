package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type ZuanshiShopItemFindRequest struct {
	Request *topsdk.Request
}

// create new request
func NewZuanshiShopItemFindRequest() (req *ZuanshiShopItemFindRequest) {
	request := topsdk.Request{MethodName: "taobao.zuanshi.shop.item.find", Params: make(map[string]interface{}, 3)}
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

// 分页大小, 默认值：15 最小值：1 最大长度：200
func (req *ZuanshiShopItemFindRequest) SetPageSize(pageSize uint) {
	if pageSize == 0 {
		pageSize = 15
	}
	req.Request.Params["page_size"] = pageSize
}

func (req *ZuanshiShopItemFindRequest) GetPageSize() uint {
	if pageSize, found := req.Request.Params["page_size"]; found {
		return pageSize.(uint)
	}
	return 15
}

// 当前页码, 默认值：1
func (req *ZuanshiShopItemFindRequest) SetPageNum(pageNum uint) {
	if pageNum == 0 {
		pageNum = 1
	}
	req.Request.Params["page_num"] = pageNum
}

func (req *ZuanshiShopItemFindRequest) GetPageNum() uint {
	if pageNum, found := req.Request.Params["page_num"]; found {
		return pageNum.(uint)
	}
	return 1
}
