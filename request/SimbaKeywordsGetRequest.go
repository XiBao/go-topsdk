package TopRequest

import (
    "github.com/XiBao/topsdk"
)

type SimbaKeywordsGetRequest struct {
    Request *topsdk.Request
}
// create new request
func NewSimbaKeywordsGetRequest() (req *SimbaKeywordsGetRequest) {
    request := topsdk.Request{MethodName:"taobao.simba.keywords.get", Params:make(map[string]interface{}, 3)}
    req = &SimbaKeywordsGetRequest {
        Request: &request,
    }
    return
}

// 推广组id
func (req *SimbaKeywordsGetRequest) SetAdgroupId(adgroup_id uint64) {
    req.Request.Params["adgroup_id"] = adgroup_id
}

func (req *SimbaKeywordsGetRequest) GetAdgroupId() uint64 {
    adgroup_id, found := req.Request.Params["adgroup_id"]
    if found {
        return adgroup_id.(uint64)
    }
    return 0
}

//关键词Id数组，最多200个
func (req *SimbaKeywordsGetRequest) SetKeywordIds(keyword_ids string) {
    req.Request.Params["keyword_ids"] = keyword_ids
}

func (req *SimbaKeywordsGetRequest) GetKeywordIds() string {
    keyword_ids, found := req.Request.Params["keyword_ids"]
    if found {
        return keyword_ids.(string)
    }
    return ""
}

//主人昵称
func (req *SimbaKeywordsGetRequest) SetNick(nick string) {
    req.Request.Params["nick"] = nick
}

func (req *SimbaKeywordsGetRequest) GetNick() string {
    nick, found := req.Request.Params["nick"]
    if found {
        return nick.(string)
    }
    return ""
}
