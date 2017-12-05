package TopRequest

import (
	"encoding/json"
	"github.com/XiBao/topsdk"
)

type TbkSpreadGetRequest struct {
	Request *topsdk.Request
}

type TbkSpreadRequest struct {
	Url string `json:"url"`
}

// create new request
func NewTbkSpreadGetRequest() (req *TbkSpreadGetRequest) {
	request := topsdk.Request{MethodName: "taobao.tbk.spread.get", Params: make(map[string]interface{}, 1)}
	req = &TbkSpreadGetRequest{
		Request: &request,
	}
	return
}

// 请求列表，内部包含多个url, 最大列表长度：20
func (req *TbkSpreadGetRequest) SetRequests(urls []string) {
	var requests []TbkSpreadRequest
	for _, url := range urls {
		requests = append(requests, TbkSpreadRequest{Url: url})
	}
	js, _ := json.Marshal(&requests)
	req.Request.Params["requests"] = string(js)
}

func (req *TbkSpreadGetRequest) GetRequests() []string {
	js, found := req.Request.Params["requests"]
	if found {
		var (
			requests []TbkSpreadRequest
			urls     []string
		)
		json.Unmarshal([]byte(js.(string)), &requests)
		for _, request := range requests {
			urls = append(urls, request.Url)
		}
		return urls
	}
	return nil
}
