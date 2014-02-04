package TopRequest

import (
    "github.com/XiBao/topsdk"
)

type SimbaInsightWordsbaseGetRequest struct {
    Request *topsdk.Request
}
// create new request
func NewSimbaInsightWordsbaseGetRequest() (req *SimbaInsightWordsbaseGetRequest) {
    request := topsdk.Request{MethodName:"taobao.simba.insight.wordsbase.get", Params:make(map[string]interface{}, 4)}
    req = &SimbaInsightWordsbaseGetRequest {
        Request: &request,
    }
    return
}

// 时间格式(DAY: 最近一天； WEEK：最近一周。MONTH：最近一个月。3MONTH：最近三个月)
func (req *SimbaInsightWordsbaseGetRequest) SetTime(time string) {
    req.Request.Params["time"] = time
}

func (req *SimbaInsightWordsbaseGetRequest) GetTime() string {
    time, found := req.Request.Params["time"]
    if found {
        return time.(string)
    }
    return ""
}

//结果过滤。PV：返回展现量；CLICK：返回点击量；AVGCPC：返回平均出价；COMPETITION ：返回竞争宝贝数;CTR 点击率。filter可由,组合
func (req *SimbaInsightWordsbaseGetRequest) SetFilter(filter string) {
    req.Request.Params["filter"] = filter
}

func (req *SimbaInsightWordsbaseGetRequest) GetFilter() string {
    filter, found := req.Request.Params["filter"]
    if found {
        return filter.(string)
    }
    return ""
}

//主人昵称
func (req *SimbaInsightWordsbaseGetRequest) SetNick(nick string) {
    req.Request.Params["nick"] = nick
}

func (req *SimbaInsightWordsbaseGetRequest) GetNick() string {
    nick, found := req.Request.Params["nick"]
    if found {
        return nick.(string)
    }
    return ""
}

//查询词组，最大长度170
func (req *SimbaInsightWordsbaseGetRequest) SetWords(words string) {
    req.Request.Params["words"] = words
}

func (req *SimbaInsightWordsbaseGetRequest) GetWords() string {
    words, found := req.Request.Params["words"]
    if found {
        return words.(string)
    }
    return ""
}