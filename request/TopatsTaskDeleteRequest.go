package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type TopatsTaskDeleteRequest struct {
	Request *topsdk.Request
}

// create new request
func NewTopatsTaskDeleteRequest() (req *TopatsTaskDeleteRequest) {
	request := topsdk.Request{MethodName: "taobao.topats.task.delete", Params: make(map[string]interface{}, 1)}
	req = &TopatsTaskDeleteRequest{
		Request: &request,
	}
	return
}

// 需要取消的任务ID
func (req *TopatsTaskDeleteRequest) SetTaskId(taskId uint64) {
	req.Request.Params["task_id"] = taskId
}

func (req *TopatsTaskDeleteRequest) GetTaskId() uint64 {
	taskId, found := req.Request.Params["task_id"]
	if found {
		return taskId.(uint64)
	}
	return 0
}
