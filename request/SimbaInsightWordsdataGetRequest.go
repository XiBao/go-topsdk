package TopRequest

import (
    "github.com/XiBao/topsdk"
)

type SimbaInsightWordsdataGetRequest struct {
    Request *topsdk.Request
}
// create new request
func NewSimbaInsightWordsdataGetRequest() (req *SimbaInsightWordsdataGetRequest) {
    request := topsdk.Request{MethodName:"taobao.simba.insight.wordsdata.get", Params:make(map[string]interface{}, 4)}
    req = &SimbaInsightWordsdataGetRequest {
        Request: &request,
    }
    return
}

// 开始时间
func (req *SimbaInsightWordsdataGetRequest) SetStartDate(startDate string) {
    req.Request.Params["start_date"] = startDate
}

func (req *SimbaInsightWordsdataGetRequest) GetStartDate() string {
    date, found := req.Request.Params["start_date"]
    if found {
        return date.(string)
    }
    return ""
}

// 结束时间
func (req *SimbaInsightWordsdataGetRequest) SetEndDate(endDate string) {
    req.Request.Params["end_date"] = endDate
}

func (req *SimbaInsightWordsdataGetRequest) GetEndDate() string {
    date, found := req.Request.Params["end_date"]
    if found {
        return date.(string)
    }
    return ""
}

//查询词组
func (req *SimbaInsightWordsdataGetRequest) SetWords(words string) {
    req.Request.Params["bidword_list"] = words
}

func (req *SimbaInsightWordsdataGetRequest) GetWords() string {
    words, found := req.Request.Params["bidword_list"]
    if found {
        return words.(string)
    }
    return ""
}
