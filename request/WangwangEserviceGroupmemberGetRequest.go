package TopRequest

import (
    "github.com/XiBao/topsdk"
)

type WangwangEserviceGroupmemberGetRequest struct {
    Request *topsdk.Request
}
// create new request
func NewWangwangEserviceGroupmemberGetRequest() (req *WangwangEserviceGroupmemberGetRequest) {
    request := topsdk.Request{MethodName:"taobao.wangwang.eservice.groupmember.get", Format:topsdk.XML_FORMAT, Params:make(map[string]interface{}, 1)}
    req = &WangwangEserviceGroupmemberGetRequest {
        Request: &request,
    }
    return
}

// 被查询用户组管理员ID：cntaobao+淘宝nick，例如cntaobaotest
func (req *WangwangEserviceGroupmemberGetRequest) SetManagerId(managerId string) {
    req.Request.Params["manager_id"] = managerId 
}

func (req *WangwangEserviceGroupmemberGetRequest) GetManagerId() string {
    managerId, found := req.Request.Params["manager_id"]
    if found {
        return managerId.(string)
    }
    return ""
}
