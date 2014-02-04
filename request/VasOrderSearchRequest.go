package TopRequest

import (
    "github.com/XiBao/topsdk"
)

type VasOrderSearchRequest struct {
    Request *topsdk.Request
}
// create new request
func NewVasOrderSearchRequest() (req *VasOrderSearchRequest) {
    request := topsdk.Request{MethodName:"taobao.vas.order.search", Params:make(map[string]interface{}, 10)}
    req = &VasOrderSearchRequest {
        Request: &request,
    }
    return
}

// 应用收费代码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得该应用的收费代码
func (req *VasOrderSearchRequest) SetArticleCode(articleCode string) {
    req.Request.Params["article_code"] = articleCode
}

func (req *VasOrderSearchRequest) GetArticleCode() string {
    articleCode, found := req.Request.Params["article_code"]
    if found {
        return articleCode.(string)
    }
    return ""
}

// 订单号
func (req *VasOrderSearchRequest) SetBizOrderId(bizOrderId uint64) {
    req.Request.Params["biz_order_id"] = bizOrderId 
}

func (req *VasOrderSearchRequest) GetBizOrderId() uint64 {
    bizOrderId, found := req.Request.Params["biz_order_id"]
    if found {
        return bizOrderId.(uint64)
    }
    return 0
}

//订单类型，1=新订 2=续订 3=升级 4=后台赠送 5=后台自动续订 6=订单审核后生成订购关系（暂时用不到） 空=全部
func (req *VasOrderSearchRequest) SetBizType(bizType uint8) {
    req.Request.Params["biz_type"] = bizType 
}

func (req *VasOrderSearchRequest) GetBizType() uint8 {
    bizType, found := req.Request.Params["biz_type"]
    if found {
        return bizType.(uint8)
    }
    return 0
}

//订单创建时间（订购时间）结束值
func (req *VasOrderSearchRequest) SetEndCreated(endCreated string) {
    req.Request.Params["end_created"] = endCreated
}

func (req *VasOrderSearchRequest) GetEndCreated() string {
    endCreated, found := req.Request.Params["end_created"]
    if found {
        return endCreated.(string)
    }
    return ""
}

//收费项目代码，从合作伙伴后台（my.open.taobao.com）-收费管理-收费项目列表 能够获得收费项目代码
func (req *VasOrderSearchRequest) SetItemCode(itemCode string) {
    req.Request.Params["item_code"] = itemCode
}

func (req *VasOrderSearchRequest) GetItemCode() string {
    itemCode, found := req.Request.Params["item_code"]
    if found {
        return itemCode.(string)
    }
    return ""
}

//淘宝会员名
func (req *VasOrderSearchRequest) SetNick(nick string) {
    req.Request.Params["nick"] = nick 
}

func (req *VasOrderSearchRequest) GetNick() string {
    nick, found := req.Request.Params["nick"]
    if found {
        return nick.(string)
    }
    return ""
}

// 子订单号
func (req *VasOrderSearchRequest) SetOrderId(orderId uint64) {
    req.Request.Params["order_id"] = orderId 
}

func (req *VasOrderSearchRequest) GetOrderId() uint64 {
    orderId, found := req.Request.Params["order_id"]
    if found {
        return orderId.(uint64)
    }
    return 0
}

//页码，从1开始
func (req *VasOrderSearchRequest) SetPageNo(page_no uint16) {
    req.Request.Params["page_no"] = page_no
}

func (req *VasOrderSearchRequest) GetPageNo() uint16 {
    page_no, found := req.Request.Params["page_no"]
    if found {
        return page_no.(uint16)
    }
    return 0
}

// 每页大小
func (req *VasOrderSearchRequest) SetPageSize(page_size uint16) {
    req.Request.Params["page_size"] = page_size
}

func (req *VasOrderSearchRequest) GetPageSize() uint16 {
    page_size, found := req.Request.Params["page_size"]
    if found {
        return page_size.(uint16)
    }
    return 0
}

//订单创建时间（订购时间）起始值（当start_created和end_created都不填写时，默认返回最近90天的数据）
func (req *VasOrderSearchRequest) SetStartCreated(startCreated string) {
    req.Request.Params["start_created"] = startCreated
}

func (req *VasOrderSearchRequest) GetStartCreated() string {
    startCreated, found := req.Request.Params["start_created"]
    if found {
        return startCreated.(string)
    }
    return ""
}
