package TopRequest

import (
    "github.com/XiBao/topsdk"
)

type ItemsGetRequest struct {
    Request *topsdk.Request
}
// create new request
func NewItemsGetRequest() (req *ItemsGetRequest) {
    request := topsdk.Request{MethodName:"taobao.items.get", Params:make(map[string]interface{}, 7)}
    req = &ItemsGetRequest {
        Request: &request,
    }
    return
}

// 需返回的商品结构字段列表。可选值为Item中的以下字段：num_iid,title,nick,pic_url,cid,price,type,delist_time,post_fee;多个字段用“,”分隔。如：num_iid,title。新增字段score(卖家信用等级数),volume(最近成交量),location（卖家地址，可以分别获取location.city,location.state）,num_iid商品数字id。此接口返回的post_fee是快递费用，volume对应搜索商品列表页的最近成交量。
func (req *ItemsGetRequest) SetFields(fields string) {
    req.Request.Params["fields"] = fields
}

func (req *ItemsGetRequest) GetFields() string {
    fields, found := req.Request.Params["fields"]
    if found {
        return fields.(string)
    }
    return ""
}

//商品类目ID。ItemCat中的cid字段。可以通过taobao.itemcats.get取到
func (req *ItemsGetRequest) SetCid(cid uint64) {
    req.Request.Params["cid"] =cid 
}

func (req *ItemsGetRequest) GetCid() uint64 {
    cid, found := req.Request.Params["cid"]
    if found {
        return cid.(uint64)
    }
    return 0
}

//搜索字段。搜索商品的title。
func (req *ItemsGetRequest) SetQ(q string) {
    req.Request.Params["q"] = q 
}

func (req *ItemsGetRequest) GetQ() string {
    q, found := req.Request.Params["q"]
    if found {
        return q.(string)
    }
    return ""
}

//卖家昵称列表。多个之间以“,”分隔;最多支持5个卖家昵称。如：nick1,nick2,nick3。
func (req *ItemsGetRequest) SetNicks(nicks string) {
    req.Request.Params["nicks"] = nicks
}

func (req *ItemsGetRequest) GetNicks() string {
    nicks, found := req.Request.Params["nicks"]
    if found {
        return nicks.(string)
    }
    return ""
}

//排序方式。格式为column:asc/desc,column可选值为: price（价格）, delist_time（下架时间）, seller_credit（卖家信用）,popularity(人气)。如按价格升序排列表示为：price:asc。新增排序字段：volume（最近成交量）
func (req *ItemsGetRequest) SetOrderBy(orderBy string) {
    req.Request.Params["order_by"] = orderBy 
}

func (req *ItemsGetRequest) GetOrderBy() string {
    orderBy, found := req.Request.Params["order_by"]
    if found {
        return orderBy.(string)
    }
    return ""
}

//页码，从1开始
func (req *ItemsGetRequest) SetPageNo(page_no uint16) {
    req.Request.Params["page_no"] = page_no
}

func (req *ItemsGetRequest) GetPageNo() uint16 {
    page_no, found := req.Request.Params["page_no"]
    if found {
        return page_no.(uint16)
    }
    return 0
}

//页码。取值范围:大于零的整数。默认值为1,即默认返回第一页数据。用此接口获取数据时，当翻页获取的条数（page_no*page_size）超过10240,为了保护后台搜索引擎，接口将报错。所以请大家尽可能的细化自己的搜索条件，例如根据修改时间分段获取商品
func (req *ItemsGetRequest) SetPageSize(page_size uint16) {
    req.Request.Params["page_size"] = page_size
}

func (req *ItemsGetRequest) GetPageSize() uint16 {
    page_size, found := req.Request.Params["page_size"]
    if found {
        return page_size.(uint16)
    }
    return 0
}
