package TopRequest

import (
    "github.com/XiBao/topsdk"
)

type WangwangEserviceChatpeersGetRequest struct {
    Request *topsdk.Request
}
// create new request
func NewWangwangEserviceChatpeersGetRequest() (req *WangwangEserviceChatpeersGetRequest) {
    request := topsdk.Request{MethodName:"taobao.wangwang.eservice.chatpeers.get", Format:topsdk.XML_FORMAT, Params:make(map[string]interface{}, 3)}
    req = &WangwangEserviceChatpeersGetRequest {
        Request: &request,
    }
    return
}

// 聊天用户ID：cntaobao+淘宝nick，例如cntaobaotest
func (req *WangwangEserviceChatpeersGetRequest) SetChatId(chatId string) {
    req.Request.Params["chat_id"] = chatId
}

func (req *WangwangEserviceChatpeersGetRequest) GetChatId() string {
    chatId, found := req.Request.Params["chat_id"]
    if found {
        return chatId.(string)
    }
    return ""
}

//查询起始日期。如2010-02-01，与当前日期间隔小于1个月。
func (req *WangwangEserviceChatpeersGetRequest) SetStartDate(startDate string) {
    req.Request.Params["start_date"] = startDate
}

func (req *WangwangEserviceChatpeersGetRequest) GetStartDate() string {
    startDate, found := req.Request.Params["start_date"]
    if found {
        return startDate.(string)
    }
    return ""
}

//查询结束日期。如2010-03-24，与起始日期跨度不能超过7天
func (req *WangwangEserviceChatpeersGetRequest) SetEndDate(endDate string) {
    req.Request.Params["end_date"] = endDate
}

func (req *WangwangEserviceChatpeersGetRequest) GetEndDate() string {
    endDate, found := req.Request.Params["end_date"]
    if found {
        return endDate.(string)
    }
    return ""
}

//字符集
func (req *WangwangEserviceChatpeersGetRequest) SetCharset(charset string) {
    req.Request.Params["charset"] = charset
}

func (req *WangwangEserviceChatpeersGetRequest) GetCharset() string {
    charset, found := req.Request.Params["charset"]
    if found {
        return charset.(string)
    }
    return ""
}
