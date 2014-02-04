package TopRequest

import (
    "github.com/XiBao/topsdk"
)

type WangwangEserviceChatlogGetRequest struct {
    Request *topsdk.Request
}
// create new request
func NewWangwangEserviceChatlogGetRequest() (req *WangwangEserviceChatlogGetRequest) {
    request := topsdk.Request{MethodName:"taobao.wangwang.eservice.chatlog.get", Format:topsdk.XML_FORMAT, Params:make(map[string]interface{}, 4)}
    req = &WangwangEserviceChatlogGetRequest {
        Request: &request,
    }
    return
}

// 聊天消息被查询用户ID：cntaobao+淘宝nick，例如cntaobaotest
func (req *WangwangEserviceChatlogGetRequest) SetFromId(fromId string) {
    req.Request.Params["from_id"] = fromId
}

func (req *WangwangEserviceChatlogGetRequest) GetFromId() string {
    fromId, found := req.Request.Params["from_id"]
    if found {
        return fromId.(string)
    }
    return ""
}

// 聊天消息被查询用户ID：cntaobao+淘宝nick，例如cntaobaotest
func (req *WangwangEserviceChatlogGetRequest) SetToId(toId string) {
    req.Request.Params["to_id"] = toId
}

func (req *WangwangEserviceChatlogGetRequest) GetToId() string {
    toId, found := req.Request.Params["to_id"]
    if found {
        return toId.(string)
    }
    return ""
}

//聊天消息起始时间，如2010-02-01
func (req *WangwangEserviceChatlogGetRequest) SetStartDate(startDate string) {
    req.Request.Params["start_date"] = startDate
}

func (req *WangwangEserviceChatlogGetRequest) GetStartDate() string {
    startDate, found := req.Request.Params["start_date"]
    if found {
        return startDate.(string)
    }
    return ""
}

//聊天消息终止时间，如2010-03-24
func (req *WangwangEserviceChatlogGetRequest) SetEndDate(endDate string) {
    req.Request.Params["end_date"] = endDate
}

func (req *WangwangEserviceChatlogGetRequest) GetEndDate() string {
    endDate, found := req.Request.Params["end_date"]
    if found {
        return endDate.(string)
    }
    return ""
}
