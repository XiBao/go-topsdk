package TopRequest

import (
    "github.com/XiBao/topsdk"
)

type ItemsListGetRequest struct {
    Request *topsdk.Request
}
// create new request
func NewItemsListGetRequest() (req *ItemsListGetRequest) {
    request := topsdk.Request{MethodName:"taobao.items.list.get", Params:make(map[string]interface{}, 2)}
    req = &ItemsListGetRequest {
        Request: &request,
    }
    return
}

// 需返回的商品结构字段列表。可选值为Item中的以下字段：num_iid,title,nick,pic_url,cid,price,type,delist_time,post_fee;多个字段用“,”分隔。如：num_iid,title。新增字段score(卖家信用等级数),volume(最近成交量),location（卖家地址，可以分别获取location.city,location.state）,num_iid商品数字id。此接口返回的post_fee是快递费用，volume对应搜索商品列表页的最近成交量
func (req *ItemsListGetRequest) SetFields(fields string) {
    req.Request.Params["fields"] = fields
}

func (req *ItemsListGetRequest) GetFields() string {
    fields, found := req.Request.Params["fields"]
    if found {
        return fields.(string)
    }
    return ""
}

func (req *ItemsListGetRequest) SetNumIids(numiids string) {
    req.Request.Params["num_iids"] = numiids
}

func (req *ItemsListGetRequest) GetNumIids() string{
    numiids, found := req.Request.Params["num_iids"]
    if found {
        return numiids.(string)
    }
    return ""
}
