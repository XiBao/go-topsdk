package TopRequest

import (
	"github.com/XiBao/topsdk"
)

type SimbaCreativesGetRequest struct {
	Request *topsdk.Request
}

// create new request
func NewSimbaCreativesGetRequest() (req *SimbaCreativesGetRequest) {
	request := topsdk.Request{MethodName: "taobao.simba.creatives.get", Params: make(map[string]interface{}, 3)}
	req = &SimbaCreativesGetRequest{
		Request: &request,
	}
	return
}

//主人昵称
func (req *SimbaCreativesGetRequest) SetNick(nick string) {
	req.Request.Params["nick"] = nick
}

func (req *SimbaCreativesGetRequest) GetNick() string {
	nick, found := req.Request.Params["nick"]
	if found {
		return nick.(string)
	}
	return ""
}

func (req *SimbaCreativesGetRequest) SetAdgroupId(adgroupId uint64) {
	req.Request.Params["adgroup_id"] = adgroupId
}

func (req *SimbaCreativesGetRequest) GetAdgroupId() uint64 {
	adgroupId, found := req.Request.Params["adgroup_id"]
	if found {
		return adgroupId.(uint64)
	}
	return 0
}

func (req *SimbaCreativesGetRequest) SetCreativeIds(creativeIds string) {
	req.Request.Params["creative_ids"] = creativeIds
}

func (req *SimbaCreativesGetRequest) GetCreativeIds() string {
	creativeIds, found := req.Request.Params["creative_ids"]
	if found {
		return creativeIds.(string)
	}
	return ""
}
