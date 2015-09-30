package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type TaeItemsListRequest struct {
	Request *topsdk.Request
}

// create new request
func NewTaeItemsListRequest() (req *TaeItemsListRequest) {
	request := topsdk.Request{MethodName: "taobao.tae.items.list", Params: make(map[string]interface{}, 3)}
	req = &TaeItemsListRequest{
		Request: &request,
	}
	return
}

// 返回值中需要的字段. 可选值 title,nick,pic_url,location,cid,price,post_fee,promoted_service,ju,shop_name字段间用 (,) 逗号分隔
func (req *TaeItemsListRequest) SetFields(fields string) {
	req.Request.Params["fields"] = fields
}

func (req *TaeItemsListRequest) GetFields() string {
	fields, found := req.Request.Params["fields"]
	if found {
		return fields.(string)
	}
	return ""
}

// 商品ID，英文逗号(,)分隔，最多50个。优先级低于open_iid，open_iids存在的话，则忽略本参数
func (req *TaeItemsListRequest) SetNumIids(numIids string) {
	req.Request.Params["num_iids"] = numIids
}

func (req *TaeItemsListRequest) GetNumIids() string {
	fields, found := req.Request.Params["num_iids"]
	if found {
		return fields.(string)
	}
	return ""
}

// 商品混淆ID，英文逗号(,)分隔，最多50个。优先级高于open_iid，本参数存在的话，则忽略num_iids
func (req *TaeItemsListRequest) SetOpenIids(numIids string) {
	req.Request.Params["open_iids"] = numIids
}

func (req *TaeItemsListRequest) GetOpenIids() string {
	fields, found := req.Request.Params["open_iids"]
	if found {
		return fields.(string)
	}
	return ""
}
