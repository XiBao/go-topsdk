package TopRequest

import (
    "github.com/XiBao/topsdk"
)

type ItemUpdateRequest struct {
    Request *topsdk.Request
}
// create new request
func NewItemUpdateRequest() (req *ItemUpdateRequest) {
    request := topsdk.Request{MethodName:"taobao.item.update", Params:make(map[string]interface{}, 2)}
    req = &ItemUpdateRequest {
        Request: &request,
    }
    return
}

func (req *ItemUpdateRequest) SetNumIid(numiid uint64) {
    req.Request.Params["num_iid"] = numiid
}

func (req *ItemUpdateRequest) GetNumIid() uint64 {
    numiid, found := req.Request.Params["num_iid"]
    if found {
        return numiid.(uint64)
    }
    return 0
}

func (req *ItemUpdateRequest) SetTitle(title string) {
    req.Request.Params["title"] = title
}

func (req *ItemUpdateRequest) GetTitle() string {
    title, found := req.Request.Params["title"]
    if found {
        return title.(string)
    }
    return ""
}
