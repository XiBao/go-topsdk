package TopRequest

import (
    "github.com/XiBao/topsdk"
)

type SimbaAdgroupOnlineitemsGetRequest struct {
    Request *topsdk.Request
}
// create new request
func NewSimbaAdgroupOnlineitemsGetRequest() (req *SimbaAdgroupOnlineitemsGetRequest) {
    request := topsdk.Request{MethodName:"taobao.simba.adgroup.onlineitems.get", Params:make(map[string]interface{}, 5)}
    req = &SimbaAdgroupOnlineitemsGetRequest {
        Request: &request,
    }
    return
}

// 主人昵称
func (req *SimbaAdgroupOnlineitemsGetRequest) SetNick(nick string) {
    req.Request.Params["nick"] = nick 
}

func (req *SimbaAdgroupOnlineitemsGetRequest) GetNick() string {
    nick, found := req.Request.Params["nick"]
    if found {
        return nick.(string)
    }
    return ""
}

//排序字段，starts：按开始时间排序bidCount:按销量排序
func (req *SimbaAdgroupOnlineitemsGetRequest) SetOrderField(orderField string) {
    req.Request.Params["order_field"] = orderField 
}

func (req *SimbaAdgroupOnlineitemsGetRequest) GetOrderField() string {
    field, found := req.Request.Params["order_field"]
    if found {
        return field.(string)
    }
    return ""
}

//排序，true:降序， false:升序
func (req *SimbaAdgroupOnlineitemsGetRequest) SetOrderBy(orderBy bool) {
    req.Request.Params["order_by"] = orderBy 
}

func (req *SimbaAdgroupOnlineitemsGetRequest) GetOrderBy() bool {
    orderBy, found := req.Request.Params["order_by"]
    if found {
        return orderBy.(bool)
    }
    return false
}

//页码，从1开始,最大50。最大只能获取1W个宝贝
func (req *SimbaAdgroupOnlineitemsGetRequest) SetPageNo(page_no uint16) {
    req.Request.Params["page_no"] = page_no
}

func (req *SimbaAdgroupOnlineitemsGetRequest) GetPageNo() uint16 {
    page_no, found := req.Request.Params["page_no"]
    if found {
        return page_no.(uint16)
    }
    return 0
}

//页尺寸，最大200
func (req *SimbaAdgroupOnlineitemsGetRequest) SetPageSize(page_size uint16) {
    req.Request.Params["page_size"] = page_size
}

func (req *SimbaAdgroupOnlineitemsGetRequest) GetPageSize() uint16 {
    page_size, found := req.Request.Params["page_size"]
    if found {
        return page_size.(uint16)
    }
    return 0
}
