package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type TmcMessagesConfirmRequest struct {
	Request *topsdk.Request
}

// create new request
func NewTmcMessagesConfirmRequest() (req *TmcMessagesConfirmRequest) {
	request := topsdk.Request{MethodName: "taobao.tmc.messages.confirm", Params: make(map[string]interface{}, 2)}
	req = &TmcMessagesConfirmRequest{
		Request: &request,
	}
	return
}

// 分组名称，不传代表默认分组
func (req *TmcMessagesConfirmRequest) SetGroupName(fields string) {
	req.Request.Params["group_name"] = fields
}

func (req *TmcMessagesConfirmRequest) GetGroupName() string {
	fields, found := req.Request.Params["fields"]
	if found {
		return fields.(string)
	}
	return ""
}

// 每次批量消费消息的条数
func (req *TmcMessagesConfirmRequest) SetsMessageIds(sMessageIds string) {
	req.Request.Params["s_message_ids"] = sMessageIds
}

// 处理成功的消息ID列表 最大 200个ID
func (req *TmcMessagesConfirmRequest) GetsMessageIds() string {
	fields, found := req.Request.Params["s_message_ids"]
	if found {
		return fields.(string)
	}
	return ""
}
