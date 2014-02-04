package TopRequest

import (
    "github.com/XiBao/topsdk"
)

type SimbaAdgroupOnlineitemsvonGetRequest struct {
    Request *topsdk.Request
}
// create new request
func NewSimbaAdgroupOnlineitemsvonGetRequest() (req *SimbaAdgroupOnlineitemsvonGetRequest) {
    request := topsdk.Request{MethodName:"taobao.simba.adgroup.onlineitemsvon.get", Params:make(map[string]interface{}, 5)}
    req = &SimbaAdgroupOnlineitemsvonGetRequest {
        Request: &request,
    }
    return
}

// 主人昵称
func (req *SimbaAdgroupOnlineitemsvonGetRequest) SetNick(nick string) {
    req.Request.Params["nick"] = nick 
}

func (req *SimbaAdgroupOnlineitemsvonGetRequest) GetNick() string {
    nick, found := req.Request.Params["nick"]
    if found {
        return nick.(string)
    }
    return ""
}

//排序字段，starts：按开始时间排序bidCount:按销量排序
func (req *SimbaAdgroupOnlineitemsvonGetRequest) SetOrderField(orderField string) {
    req.Request.Params["order_field"] = orderField 
}

func (req *SimbaAdgroupOnlineitemsvonGetRequest) GetOrderField() string {
    field, found := req.Request.Params["order_field"]
    if found {
        return field.(string)
    }
    return ""
}

//排序，true:降序， false:升序
func (req *SimbaAdgroupOnlineitemsvonGetRequest) SetOrderBy(orderBy bool) {
    req.Request.Params["order_by"] = orderBy 
}

func (req *SimbaAdgroupOnlineitemsvonGetRequest) GetOrderBy() bool {
    orderBy, found := req.Request.Params["order_by"]
    if found {
        return orderBy.(bool)
    }
    return false
}

//页码，从1开始,最大50。最大只能获取1W个宝贝
func (req *SimbaAdgroupOnlineitemsvonGetRequest) SetPageNo(page_no uint16) {
    req.Request.Params["page_no"] = page_no
}

func (req *SimbaAdgroupOnlineitemsvonGetRequest) GetPageNo() uint16 {
    page_no, found := req.Request.Params["page_no"]
    if found {
        return page_no.(uint16)
    }
    return 0
}

//页尺寸，最大200
func (req *SimbaAdgroupOnlineitemsvonGetRequest) SetPageSize(page_size uint16) {
    req.Request.Params["page_size"] = page_size
}

func (req *SimbaAdgroupOnlineitemsvonGetRequest) GetPageSize() uint16 {
    page_size, found := req.Request.Params["page_size"]
    if found {
        return page_size.(uint16)
    }
    return 0
}
