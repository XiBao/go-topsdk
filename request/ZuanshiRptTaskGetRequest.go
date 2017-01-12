package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type ZuanshiRptTaskGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewZuanshiRptTaskGetRequest() (req *ZuanshiRptTaskGetRequest) {
	request := topsdk.Request{MethodName: "taobao.zuanshi.rpt.task.get", Params: make(map[string]interface{}, 1)}
	req = &ZuanshiRptTaskGetRequest{
		Request: &request,
	}
	return
}

// 推广单元ID
func (req *ZuanshiRptTaskGetRequest) SetTaskId(taskId uint64) {
	req.Request.Params["task_id"] = taskId
}

func (req *ZuanshiRptTaskGetRequest) GetTaskId() uint64 {
	taskId, found := req.Request.Params["task_id"]
	if found {
		return taskId.(uint64)
	}
	return 0
}
