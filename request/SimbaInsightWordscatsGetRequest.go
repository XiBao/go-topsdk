package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type SimbaInsightWordscatsGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewSimbaInsightWordscatsGetRequest() (req *SimbaInsightWordscatsGetRequest) {
	request := topsdk.Request{MethodName: "taobao.simba.insight.wordscats.get", Params: make(map[string]interface{}, 4)}
	req = &SimbaInsightWordscatsGetRequest{
		Request: &request,
	}
	return
}

//结果过滤。PV：返回展现量；CLICK：返回点击量；AVGCPC：返回平均出价；COMPETITION ：返回竞争宝贝数;CTR 点击率。filter可由,组合
func (req *SimbaInsightWordscatsGetRequest) SetFilter(filter string) {
	req.Request.Params["filter"] = filter
}

func (req *SimbaInsightWordscatsGetRequest) GetFilter() string {
	filter, found := req.Request.Params["filter"]
	if found {
		return filter.(string)
	}
	return ""
}

//主人昵称
func (req *SimbaInsightWordscatsGetRequest) SetNick(nick string) {
	req.Request.Params["nick"] = nick
}

func (req *SimbaInsightWordscatsGetRequest) GetNick() string {
	nick, found := req.Request.Params["nick"]
	if found {
		return nick.(string)
	}
	return ""
}

//查询词和类目数组，最大长度200，每一个数组元素都是词+类目，以”^^”分割如下：
func (req *SimbaInsightWordscatsGetRequest) SetWordCategories(wordCategories []string) {
	req.Request.Params["word_categories"] = wordCategories
}

func (req *SimbaInsightWordscatsGetRequest) GetWordCategories() []string {
	wordCategories, found := req.Request.Params["word_categories"]
	if found {
		return wordCategories.([]string)
	}
	return []string{}
}
