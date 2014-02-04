package TopRequest

import (
    "github.com/XiBao/topsdk"
)

type SimbaKeywordsvonAddRequest struct {
    Request *topsdk.Request
}
// create new request
func NewSimbaKeywordsvonAddRequest() (req *SimbaKeywordsvonAddRequest) {
    request := topsdk.Request{MethodName:"taobao.simba.keywordsvon.add", Params:make(map[string]interface{}, 4)}
    req = &SimbaKeywordsvonAddRequest {
        Request: &request,
    }
    return
}

// 推广组id
func (req *SimbaKeywordsvonAddRequest) SetAdgroupId(adgroup_id uint64) {
    req.Request.Params["adgroup_id"] = adgroup_id
}

func (req *SimbaKeywordsvonAddRequest) GetAdgroupId() uint64 {
    adgroup_id, found := req.Request.Params["adgroup_id"]
    if found {
        return adgroup_id.(uint64)
    }
    return 0
}

/** 
 * 关键词，出价字符串和匹配方式字符串数组，最多200个;每个字符串：word+  ”^^”+price+”^^”+matchscope,
Price是整数，以“分”为单位，不能小于5，不能大于日限额; 
price为0则设置为使用默认出价；
matchscope只能是1,2,4（1代表精确匹配，2代表子串匹配，4代表广泛匹配）可不传。
关键词不能包含字符”^^”和 ”,”
 **/
func (req *SimbaKeywordsvonAddRequest) SetKeywordPrices(keyword_prices string) {
    req.Request.Params["keyword_prices"] = keyword_prices
}

func (req *SimbaKeywordsvonAddRequest) GetKeywordPrices() string {
    keyword_prices, found := req.Request.Params["keyword_prices"]
    if found {
        return keyword_prices.(string)
    }
    return ""
}

//主人昵称
func (req *SimbaKeywordsvonAddRequest) SetNick(nick string) {
    req.Request.Params["nick"] = nick
}

func (req *SimbaKeywordsvonAddRequest) GetNick() string {
    nick, found := req.Request.Params["nick"]
    if found {
        return nick.(string)
    }
    return ""
}
