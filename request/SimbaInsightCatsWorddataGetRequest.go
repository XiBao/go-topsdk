package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type SimbaInsightCatsWorddataGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewSimbaInsightCatsWorddataGetRequest() (req *SimbaInsightCatsWorddataGetRequest) {
	request := topsdk.Request{MethodName: "taobao.simba.insight.catsworddata.get", Params: make(map[string]interface{}, 3)}
	req = &SimbaInsightCatsWorddataGetRequest{
		Request: &request,
	}
	return
}

func (req *SimbaInsightCatsWorddataGetRequest) SetCid(cid uint64) {
    req.Request.Params["cat_id"] = cid
}

func (req *SimbaInsightCatsWorddataGetRequest) GetCid() uint64 {
    date, found := req.Request.Params["cat_id"]
    if found {
        return date.(uint64)
    }
    return 0
}

// 开始时间
func (req *SimbaInsightCatsWorddataGetRequest) SetStartDate(startDate string) {
    req.Request.Params["start_date"] = startDate
}

func (req *SimbaInsightCatsWorddataGetRequest) GetStartDate() string {
    date, found := req.Request.Params["start_date"]
    if found {
        return date.(string)
    }
    return ""
}

// 结束时间
func (req *SimbaInsightCatsWorddataGetRequest) SetEndDate(endDate string) {
    req.Request.Params["end_date"] = endDate
}

func (req *SimbaInsightCatsWorddataGetRequest) GetEndDate() string {
    date, found := req.Request.Params["end_date"]
    if found {
        return date.(string)
    }
    return ""
}

//查询词组
func (req *SimbaInsightCatsWorddataGetRequest) SetWords(words string) {
    req.Request.Params["bidword_list"] = words
}

func (req *SimbaInsightCatsWorddataGetRequest) GetWords() string {
    words, found := req.Request.Params["bidword_list"]
    if found {
        return words.(string)
    }
    return ""
}
