package TopRequest

import (
    "github.com/XiBao/topsdk"
)

type ItemsOnsaleGetRequest struct {
    Request *topsdk.Request
}
// create new request
func NewItemsOnsaleGetRequest() (req *ItemsOnsaleGetRequest) {
    request := topsdk.Request{MethodName:"taobao.items.onsale.get", Params:make(map[string]interface{}, 6)}
    req = &ItemsOnsaleGetRequest {
        Request: &request,
    }
    return
}

// 需返回的字段列表。可选值：Item商品结构体中的以下字段： approve_status,num_iid,title,nick,type,cid,pic_url,num,props,valid_thru,list_time,price,has_discount,has_invoice,has_warranty,has_showcase,modified,delist_time,postage_id,seller_cids,outer_id；字段之间用“,”分隔。不支持其他字段，如果需要获取其他字段数据，调用taobao.item.get。
func (req *ItemsOnsaleGetRequest) SetFields(fields string) {
    req.Request.Params["fields"] = fields
}

func (req *ItemsOnsaleGetRequest) GetFields() string {
    fields, found := req.Request.Params["fields"]
    if found {
        return fields.(string)
    }
    return ""
}

//商品类目ID。ItemCat中的cid字段。可以通过taobao.itemcats.get取到
func (req *ItemsOnsaleGetRequest) SetCid(cid uint64) {
    req.Request.Params["cid"] =cid 
}

func (req *ItemsOnsaleGetRequest) GetCid() uint64 {
    cid, found := req.Request.Params["cid"]
    if found {
        return cid.(uint64)
    }
    return 0
}

//搜索字段。搜索商品的title。
func (req *ItemsOnsaleGetRequest) SetQ(q string) {
    req.Request.Params["q"] = q 
}

func (req *ItemsOnsaleGetRequest) GetQ() string {
    q, found := req.Request.Params["q"]
    if found {
        return q.(string)
    }
    return ""
}

//排序方式。格式为column:asc/desc ，column可选值:list_time(上架时间),delist_time(下架时间),num(商品数量)，modified(最近修改时间);默认上架时间降序(即最新上架排在前面)。如按照上架时间降序排序方式为list_time:desc
func (req *ItemsOnsaleGetRequest) SetOrderBy(orderBy string) {
    req.Request.Params["order_by"] = orderBy 
}

func (req *ItemsOnsaleGetRequest) GetOrderBy() string {
    orderBy, found := req.Request.Params["order_by"]
    if found {
        return orderBy.(string)
    }
    return ""
}

//页码，从1开始
func (req *ItemsOnsaleGetRequest) SetPageNo(page_no uint16) {
    req.Request.Params["page_no"] = page_no
}

func (req *ItemsOnsaleGetRequest) GetPageNo() uint16 {
    page_no, found := req.Request.Params["page_no"]
    if found {
        return page_no.(uint16)
    }
    return 0
}

//每页条数。取值范围:大于零的整数;最大值：200；默认值：40。用此接口获取数据时，当翻页获取的条数（page_no*page_size）超过10万,为了保护后台搜索引擎，接口将报错。所以请大家尽可能的细化自己的搜索条件，例如根据修改时间分段获取商品
func (req *ItemsOnsaleGetRequest) SetPageSize(page_size uint16) {
    req.Request.Params["page_size"] = page_size
}

func (req *ItemsOnsaleGetRequest) GetPageSize() uint16 {
    page_size, found := req.Request.Params["page_size"]
    if found {
        return page_size.(uint16)
    }
    return 0
}
