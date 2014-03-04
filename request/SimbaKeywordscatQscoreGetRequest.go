package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type SimbaKeywordscatQscoreGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewSimbaKeywordscatQscoreGetRequest() (req *SimbaKeywordscatQscoreGetRequest) {
	request := topsdk.Request{MethodName: "taobao.simba.keywordscat.qscore.get", Params: make(map[string]interface{}, 2)}
	req = &SimbaKeywordscatQscoreGetRequest{
		Request: &request,
	}
	return
}

// 推广组id
func (req *SimbaKeywordscatQscoreGetRequest) SetAdgroupId(adgroup_id uint64) {
	req.Request.Params["adgroup_id"] = adgroup_id
}

func (req *SimbaKeywordscatQscoreGetRequest) GetAdgroupId() uint64 {
	adgroup_id, found := req.Request.Params["adgroup_id"]
	if found {
		return adgroup_id.(uint64)
	}
	return 0
}

//主人昵称
func (req *SimbaKeywordscatQscoreGetRequest) SetNick(nick string) {
	req.Request.Params["nick"] = nick
}

func (req *SimbaKeywordscatQscoreGetRequest) GetNick() string {
	nick, found := req.Request.Params["nick"]
	if found {
		return nick.(string)
	}
	return ""
}
