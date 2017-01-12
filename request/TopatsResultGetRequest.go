package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type TopatsResultGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewTopatsResultGetRequest() (req *TopatsResultGetRequest) {
	request := topsdk.Request{MethodName: "taobao.topats.result.get", Params: make(map[string]interface{}, 1)}
	req = &TopatsResultGetRequest{
		Request: &request,
	}
	return
}

// 任务ID，创建任务时返回的task_id
func (req *TopatsResultGetRequest) SetTaskId(taskId uint64) {
	req.Request.Params["task_id"] = taskId
}

func (req *TopatsResultGetRequest) GetTaskId() uint64 {
	taskId, found := req.Request.Params["task_id"]
	if found {
		return taskId.(uint64)
	}
	return 0
}
