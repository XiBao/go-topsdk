package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type SimbaInsightCatstopwordGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewSimbaInsightCatstopwordGetRequest() (req *SimbaInsightCatstopwordGetRequest) {
	request := topsdk.Request{MethodName: "taobao.simba.insight.catstopword.get", Params: make(map[string]interface{}, 3)}
	req = &SimbaInsightCatstopwordGetRequest{
		Request: &request,
	}
	return
}

//主人昵称
func (req *SimbaInsightCatstopwordGetRequest) SetNick(nick string) {
	req.Request.Params["nick"] = nick
}

func (req *SimbaInsightCatstopwordGetRequest) GetNick() string {
	nick, found := req.Request.Params["nick"]
	if found {
		return nick.(string)
	}
	return ""
}

//类目id数组，最大长度200
func (req *SimbaInsightCatstopwordGetRequest) SetCategoryIds(categoryIds string) {
	req.Request.Params["category_ids"] = categoryIds
}

func (req *SimbaInsightCatstopwordGetRequest) GetCategoryIds() string {
	categoryIds, found := req.Request.Params["category_ids"]
	if found {
		return categoryIds.(string)
	}
	return ""
}

//最大返回数量(1-100)
func (req *SimbaInsightCatstopwordGetRequest) SetResultNum(resultNum uint8) {
	req.Request.Params["result_num"] = resultNum
}

func (req *SimbaInsightCatstopwordGetRequest) GetResultNum() uint8 {
	resultNum, found := req.Request.Params["result_num"]
	if found {
		return resultNum.(uint8)
	}
	return 0
}
