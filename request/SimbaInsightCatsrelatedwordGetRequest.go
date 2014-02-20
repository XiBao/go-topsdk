package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type SimbaInsightCatsrelatedwordGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewSimbaInsightCatsrelatedwordGetRequest() (req *SimbaInsightCatsrelatedwordGetRequest) {
	request := topsdk.Request{MethodName: "taobao.simba.insight.catsrelatedword.get", Params: make(map[string]interface{}, 3)}
	req = &SimbaInsightCatsrelatedwordGetRequest{
		Request: &request,
	}
	return
}

//主人昵称
func (req *SimbaInsightCatsrelatedwordGetRequest) SetNick(nick string) {
	req.Request.Params["nick"] = nick
}

func (req *SimbaInsightCatsrelatedwordGetRequest) GetNick() string {
	nick, found := req.Request.Params["nick"]
	if found {
		return nick.(string)
	}
	return ""
}

//类目id数组，最大长度200
func (req *SimbaInsightCatsrelatedwordGetRequest) SetWords(words string) {
	req.Request.Params["words"] = words
}

func (req *SimbaInsightCatsrelatedwordGetRequest) GetWords() string {
	words, found := req.Request.Params["words"]
	if found {
		return words.(string)
	}
	return ""
}

//最大返回数量(1-10)
func (req *SimbaInsightCatsrelatedwordGetRequest) SetResultNum(resultNum uint8) {
	req.Request.Params["result_num"] = resultNum
}

func (req *SimbaInsightCatsrelatedwordGetRequest) GetResultNum() uint8 {
	resultNum, found := req.Request.Params["result_num"]
	if found {
		return resultNum.(uint8)
	}
	return 0
}
