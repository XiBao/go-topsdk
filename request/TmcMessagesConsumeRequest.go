package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type TmcMessagesConsumeRequest struct {
	Request *topsdk.Request
}

// create new request
func NewTmcMessagesConsumeRequest() (req *TmcMessagesConsumeRequest) {
	request := topsdk.Request{MethodName: "taobao.tmc.messages.consume", Params: make(map[string]interface{}, 2)}
	req = &TmcMessagesConsumeRequest{
		Request: &request,
	}
	return
}

// 用户分组名称，不传表示消费默认分组，如果应用没有设置用户分组，传入分组名称将会返回错误
func (req *TmcMessagesConsumeRequest) SetGroupName(fields string) {
	req.Request.Params["group_name"] = fields
}

func (req *TmcMessagesConsumeRequest) GetGroupName() string {
	fields, found := req.Request.Params["fields"]
	if found {
		return fields.(string)
	}
	return ""
}

// 每次批量消费消息的条数
func (req *TmcMessagesConsumeRequest) SetQuantity(quantity uint16) {
	req.Request.Params["quantity"] = quantity
}

func (req *TmcMessagesConsumeRequest) GetQuantity() uint16 {
	fields, found := req.Request.Params["quantity"]
	if found {
		return fields.(uint16)
	}
	return 0
}
